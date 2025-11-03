package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config represents the application configuration
type Config struct {
	FilesPerType      int      `json:"files_per_type"`
	RecordsPerFile    int      `json:"records_per_file"`
	FileSizeTargetMB  int      `json:"file_size_target_mb"`
	Countries         []string `json:"countries"`           // Countries to generate data for (uk, india, us)
	SensitiveTypes    []string `json:"sensitive_types"`
	FileTypes         []string `json:"file_types"`
	WorkerCount       int      `json:"worker_count"`        // Number of concurrent workers
	OutputDir         string   `json:"output_dir"`          // Output directory
}

// DefaultConfig returns default configuration values
func DefaultConfig() *Config {
	return &Config{
		FilesPerType:     10,
		RecordsPerFile:   100,
		FileSizeTargetMB: 10,
		Countries:        []string{"uk", "india", "us"},
		SensitiveTypes:   []string{"pii", "pci", "financial"},
		FileTypes:        []string{"pdf", "xlsx", "html", "csv", "json", "zip", "png", "jpg"},
		WorkerCount:      10,
		OutputDir:        "output",
	}
}

// Load loads configuration from a JSON file
func Load(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	config := DefaultConfig()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}

	return config, nil
}

// LoadOrDefault attempts to load config from file, falls back to defaults
func LoadOrDefault(filename string) *Config {
	config, err := Load(filename)
	if err != nil {
		fmt.Printf("Using default config (failed to load %s: %v)\n", filename, err)
		return DefaultConfig()
	}
	return config
}

// Save saves the configuration to a JSON file
func (c *Config) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(c); err != nil {
		return fmt.Errorf("failed to encode config: %w", err)
	}

	return nil
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.FilesPerType <= 0 {
		return fmt.Errorf("files_per_type must be positive")
	}
	if c.RecordsPerFile <= 0 {
		return fmt.Errorf("records_per_file must be positive")
	}
	if c.FileSizeTargetMB < 0 {
		return fmt.Errorf("file_size_target_mb cannot be negative")
	}
	if len(c.Countries) == 0 {
		return fmt.Errorf("countries cannot be empty")
	}
	// Validate country names
	validCountries := map[string]bool{"uk": true, "india": true, "us": true}
	for _, country := range c.Countries {
		if !validCountries[country] {
			return fmt.Errorf("invalid country: %s (must be uk, india, or us)", country)
		}
	}
	if len(c.SensitiveTypes) == 0 {
		return fmt.Errorf("sensitive_types cannot be empty")
	}
	if len(c.FileTypes) == 0 {
		return fmt.Errorf("file_types cannot be empty")
	}
	if c.WorkerCount <= 0 {
		c.WorkerCount = 10 // Set default
	}
	if c.OutputDir == "" {
		c.OutputDir = "output"
	}
	return nil
}

