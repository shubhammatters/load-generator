package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"load-generator/internal/config"
	"load-generator/internal/data"
	"load-generator/internal/files"
)

// FileGenerator interface for all file generators
type FileGenerator interface {
	Generate(outputDir, country, sensitiveType string, records []*data.Record, fileIndex int) error
}

// GeneratorJob represents a job for file generation
type GeneratorJob struct {
	FileType      string
	Country       string
	SensitiveType string
	FileIndex     int
}

func main() {
	fmt.Println("═══════════════════════════════════════════════")
	fmt.Println("   Multi-Country Sensitive Data Load Generator")
	fmt.Println("═══════════════════════════════════════════════")
	fmt.Println()

	// Load configuration
	cfg := config.LoadOrDefault("config.json")
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	// If config.json doesn't exist, create it
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		fmt.Println("Creating default config.json...")
		if err := cfg.Save("config.json"); err != nil {
			log.Printf("Warning: Failed to save config.json: %v", err)
		}
	}

	fmt.Printf("Configuration:\n")
	fmt.Printf("  Files per type: %d\n", cfg.FilesPerType)
	fmt.Printf("  Records per file: %d\n", cfg.RecordsPerFile)
	fmt.Printf("  Countries: %v\n", cfg.Countries)
	fmt.Printf("  Sensitive types: %v\n", cfg.SensitiveTypes)
	fmt.Printf("  File types: %v\n", cfg.FileTypes)
	fmt.Printf("  Worker count: %d\n", cfg.WorkerCount)
	fmt.Printf("  Output directory: %s\n", cfg.OutputDir)
	fmt.Println()

	// Create output directory
	if err := os.MkdirAll(cfg.OutputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Initialize generators
	generators := map[string]FileGenerator{
		"csv":  files.NewCSVGenerator(),
		"json": files.NewJSONGenerator(),
		"html": files.NewHTMLGenerator(),
		"png":  files.NewImageGenerator(),
		"jpg":  files.NewImageGenerator(),
		"zip":  files.NewZIPGenerator(),
		"pdf":  files.NewPDFGenerator(),
		"xlsx": files.NewExcelGenerator(),
		"docx": files.NewWordGenerator(),
		"pptx": files.NewPowerPointGenerator(),
	}

	// Calculate total jobs (now including countries dimension)
	totalJobs := len(cfg.FileTypes) * len(cfg.Countries) * len(cfg.SensitiveTypes) * cfg.FilesPerType
	fmt.Printf("Total files to generate: %d\n", totalJobs)
	fmt.Println("Starting generation...\n")

	startTime := time.Now()

	// Create job queue
	jobs := make(chan GeneratorJob, totalJobs)
	var completedJobs int32
	var failedJobs int32

	// Create worker pool
	var wg sync.WaitGroup
	for i := 0; i < cfg.WorkerCount; i++ {
		wg.Add(1)
		go worker(i+1, jobs, &wg, generators, cfg, &completedJobs, &failedJobs)
	}

	// Progress reporter
	stopProgress := make(chan bool)
	go progressReporter(&completedJobs, totalJobs, stopProgress)

	// Enqueue jobs (now with country dimension)
	for _, fileType := range cfg.FileTypes {
		for _, country := range cfg.Countries {
			for _, sensitiveType := range cfg.SensitiveTypes {
				for fileIndex := 1; fileIndex <= cfg.FilesPerType; fileIndex++ {
					jobs <- GeneratorJob{
						FileType:      fileType,
						Country:       country,
						SensitiveType: sensitiveType,
						FileIndex:     fileIndex,
					}
				}
			}
		}
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()
	stopProgress <- true

	elapsed := time.Since(startTime)

	// Print summary
	fmt.Println("\n═══════════════════════════════════════════════")
	fmt.Println("   Generation Complete!")
	fmt.Println("═══════════════════════════════════════════════")
	fmt.Printf("Total files generated: %d\n", completedJobs)
	fmt.Printf("Failed: %d\n", failedJobs)
	fmt.Printf("Time taken: %s\n", elapsed.Round(time.Second))
	fmt.Printf("Average: %.2f files/second\n", float64(completedJobs)/elapsed.Seconds())
	fmt.Printf("\nOutput directory: %s\n", cfg.OutputDir)

	// Calculate directory size
	go func() {
		size, err := calculateDirSize(cfg.OutputDir)
		if err == nil {
			fmt.Printf("Total size: %s\n", formatBytes(size))
		}
	}()

	time.Sleep(100 * time.Millisecond) // Give goroutine time to complete
	fmt.Println("═══════════════════════════════════════════════")
}

func worker(id int, jobs <-chan GeneratorJob, wg *sync.WaitGroup, generators map[string]FileGenerator, cfg *config.Config, completedJobs, failedJobs *int32) {
	defer wg.Done()

	// Create data generator with worker-specific seed
	dataGen := data.NewGenerator(time.Now().UnixNano() + int64(id))

	for job := range jobs {
		// Convert country string to data.Country type
		country := data.Country(job.Country)
		
		// Generate records for the specific country
		records := dataGen.GenerateRecords(cfg.RecordsPerFile, country)

		// Get appropriate generator
		generator, exists := generators[job.FileType]
		if !exists {
			log.Printf("Worker %d: Unknown file type: %s", id, job.FileType)
			atomic.AddInt32(failedJobs, 1)
			continue
		}

		// Generate file with country
		err := generator.Generate(cfg.OutputDir, job.Country, job.SensitiveType, records, job.FileIndex)
		if err != nil {
			log.Printf("Worker %d: Failed to generate %s/%s/%s file %d: %v", id, job.FileType, job.Country, job.SensitiveType, job.FileIndex, err)
			atomic.AddInt32(failedJobs, 1)
		} else {
			atomic.AddInt32(completedJobs, 1)
		}
	}
}

func progressReporter(completedJobs *int32, totalJobs int, stop chan bool) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			completed := atomic.LoadInt32(completedJobs)
			percentage := float64(completed) / float64(totalJobs) * 100
			fmt.Printf("\rProgress: %d/%d (%.1f%%) ", completed, totalJobs, percentage)
		case <-stop:
			completed := atomic.LoadInt32(completedJobs)
			percentage := float64(completed) / float64(totalJobs) * 100
			fmt.Printf("\rProgress: %d/%d (%.1f%%) ", completed, totalJobs, percentage)
			return
		}
	}
}

func calculateDirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

