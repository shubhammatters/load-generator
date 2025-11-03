package files

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"load-generator/internal/data"
)

// JSONGenerator handles JSON file generation
type JSONGenerator struct{}

// NewJSONGenerator creates a new JSON generator
func NewJSONGenerator() *JSONGenerator {
	return &JSONGenerator{}
}

// Generate creates a JSON file with the given records
func (g *JSONGenerator) Generate(outputDir, country, sensitiveType string, records []*data.Record, fileIndex int) error {
	// Create directory structure with country
	dir := filepath.Join(outputDir, "json", country, sensitiveType)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create file
	filename := filepath.Join(dir, fmt.Sprintf("data_%d.json", fileIndex))
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write JSON with indentation for readability
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	
	// Wrap records in an object for better structure
	output := map[string]interface{}{
		"records": records,
		"count":   len(records),
		"country": country,
		"type":    sensitiveType,
	}
	
	if err := encoder.Encode(output); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}

