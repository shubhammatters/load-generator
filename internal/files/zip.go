package files

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"load-generator/internal/data"
)

// ZIPGenerator handles ZIP file generation
type ZIPGenerator struct{}

// NewZIPGenerator creates a new ZIP generator
func NewZIPGenerator() *ZIPGenerator {
	return &ZIPGenerator{}
}

// Generate creates a ZIP file containing CSV data with the given records
func (g *ZIPGenerator) Generate(outputDir, country, sensitiveType string, records []*data.Record, fileIndex int) error {
	// Create directory structure
	dir := filepath.Join(outputDir, "zip", country, sensitiveType)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create ZIP file
	filename := filepath.Join(dir, fmt.Sprintf("data_%d.zip", fileIndex))
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	// Add main data CSV file
	if err := g.addCSVToZip(zipWriter, "data.csv", records); err != nil {
		return err
	}

	// Add individual record files (text format)
	for i, record := range records {
		if i >= 100 { // Limit to 100 individual files to avoid bloat
			break
		}
		if err := g.addRecordToZip(zipWriter, fmt.Sprintf("records/record_%d.txt", i+1), record); err != nil {
			return err
		}
	}

	// Add README
	if err := g.addReadmeToZip(zipWriter, sensitiveType, len(records)); err != nil {
		return err
	}

	return nil
}

func (g *ZIPGenerator) addCSVToZip(zipWriter *zip.Writer, filename string, records []*data.Record) error {
	writer, err := zipWriter.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file in zip: %w", err)
	}

	csvWriter := csv.NewWriter(writer)
	defer csvWriter.Flush()

	// Write headers
	headers := []string{
		"FirstName", "LastName", "FullName", "DateOfBirth", "Email",
		"NHSNumber", "NationalInsurance", "UKPostcode", "FullAddress",
		"PhoneNumber", "MobileNumber", "DrivingLicense", "PassportNumber",
		"CardNumber", "CardType", "CVV", "ExpiryDate",
		"BankAccountNumber", "SortCode", "IBAN", "TaxReference", "AnnualIncome",
	}
	if err := csvWriter.Write(headers); err != nil {
		return fmt.Errorf("failed to write headers: %w", err)
	}

	// Write records
	for _, record := range records {
		row := []string{
			record.FirstName, record.LastName, record.FullName, record.DateOfBirth, record.Email,
			record.NHSNumber, record.NationalInsurance, record.UKPostcode, record.FullAddress,
			record.PhoneNumber, record.MobileNumber, getDrivingLicenseZIP(record), record.PassportNumber,
			record.CardNumber, record.CardType, record.CVV, record.ExpiryDate,
			record.BankAccountNumber, record.SortCode, record.IBAN, record.TaxReference, record.AnnualIncome,
		}
		if err := csvWriter.Write(row); err != nil {
			return fmt.Errorf("failed to write record: %w", err)
		}
	}

	return nil
}

func (g *ZIPGenerator) addRecordToZip(zipWriter *zip.Writer, filename string, record *data.Record) error {
	writer, err := zipWriter.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file in zip: %w", err)
	}

	content := fmt.Sprintf(`UK Sensitive Data Record
========================

Personal Information:
---------------------
Name: %s
Date of Birth: %s
Email: %s

Identification:
--------------
NHS Number: %s
National Insurance: %s
Driving License: %s
Passport Number: %s

Contact Information:
-------------------
Address: %s
Postcode: %s
Phone: %s
Mobile: %s

Payment Information:
-------------------
Card Number: %s
Card Type: %s
CVV: %s
Expiry Date: %s

Banking Information:
-------------------
Account Number: %s
Sort Code: %s
IBAN: %s

Tax Information:
---------------
Tax Reference: %s
Annual Income: %s
`,
		record.FullName, record.DateOfBirth, record.Email,
		record.NHSNumber, record.NationalInsurance, getDrivingLicenseZIP(record), record.PassportNumber,
		record.FullAddress, record.UKPostcode, record.PhoneNumber, record.MobileNumber,
		record.CardNumber, record.CardType, record.CVV, record.ExpiryDate,
		record.BankAccountNumber, record.SortCode, record.IBAN,
		record.TaxReference, record.AnnualIncome,
	)

	if _, err := writer.Write([]byte(content)); err != nil {
		return fmt.Errorf("failed to write content: %w", err)
	}

	return nil
}

func (g *ZIPGenerator) addReadmeToZip(zipWriter *zip.Writer, sensitiveType string, recordCount int) error {
	writer, err := zipWriter.Create("README.txt")
	if err != nil {
		return fmt.Errorf("failed to create README in zip: %w", err)
	}

	content := fmt.Sprintf(`UK Sensitive Test Data Archive
==============================

Sensitive Data Type: %s
Total Records: %d

Contents:
---------
- data.csv: All records in CSV format
- records/: Individual record files in text format

WARNING:
--------
This file contains sensitive test data including:
- Personal Identifiable Information (PII)
- Payment Card Industry (PCI) data
- Financial information

This data is for TESTING PURPOSES ONLY.
Do not use in production environments.

Generated by UK Load Generator
`, sensitiveType, recordCount)

	if _, err := writer.Write([]byte(content)); err != nil {
		return fmt.Errorf("failed to write README: %w", err)
	}

	return nil
}


func getDrivingLicenseZIP(record *data.Record) string {
	switch record.Country {
	case data.UK:
		return record.UKDrivingLicense
	case data.India:
		return record.IndianDrivingLicense
	case data.US:
		return record.USDrivingLicense
	default:
		return ""
	}
}
