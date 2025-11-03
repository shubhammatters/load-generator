package files

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"load-generator/internal/data"
)

// CSVGenerator handles CSV file generation
type CSVGenerator struct{}

// NewCSVGenerator creates a new CSV generator
func NewCSVGenerator() *CSVGenerator {
	return &CSVGenerator{}
}

// Generate creates a CSV file with the given records
func (g *CSVGenerator) Generate(outputDir, country, sensitiveType string, records []*data.Record, fileIndex int) error {
	// Create directory structure with country
	dir := filepath.Join(outputDir, "csv", country, sensitiveType)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create file
	filename := filepath.Join(dir, fmt.Sprintf("data_%d.csv", fileIndex))
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers based on country
	headers := g.getHeaders(country)
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("failed to write headers: %w", err)
	}

	// Write records
	for _, record := range records {
		row := g.getRowData(record, country)
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write record: %w", err)
		}
	}

	return nil
}

func (g *CSVGenerator) getHeaders(country string) []string {
	commonHeaders := []string{
		"Country", "FirstName", "LastName", "FullName", "DateOfBirth", "Email",
		"FullAddress", "PhoneNumber", "MobileNumber", "PassportNumber",
		"CardNumber", "CardType", "CVV", "ExpiryDate",
		"BankAccountNumber", "TaxReference", "AnnualIncome",
	}
	
	switch country {
	case "uk":
		return append(commonHeaders, "NHSNumber", "NationalInsurance", "UKPostcode", "UKDrivingLicense", "SortCode", "IBAN")
	case "india":
		return append(commonHeaders, "AadhaarNumber", "PANNumber", "IndianPostcode", "IndianState", "IndianDrivingLicense", "IFSCCode")
	case "us":
		return append(commonHeaders, "SSN", "GreenCard", "USZipCode", "USState", "USDrivingLicense", "BankRoutingNumber")
	default:
		return commonHeaders
	}
}

func (g *CSVGenerator) getRowData(record *data.Record, country string) []string {
	commonData := []string{
		string(record.Country), record.FirstName, record.LastName, record.FullName, record.DateOfBirth, record.Email,
		record.FullAddress, record.PhoneNumber, record.MobileNumber, record.PassportNumber,
		record.CardNumber, record.CardType, record.CVV, record.ExpiryDate,
		record.BankAccountNumber, record.TaxReference, record.AnnualIncome,
	}
	
	switch country {
	case "uk":
		return append(commonData, record.NHSNumber, record.NationalInsurance, record.UKPostcode, record.UKDrivingLicense, record.SortCode, record.IBAN)
	case "india":
		return append(commonData, record.AadhaarNumber, record.PANNumber, record.IndianPostcode, record.IndianState, record.IndianDrivingLicense, record.IFSCCode)
	case "us":
		return append(commonData, record.SSN, record.GreenCard, record.USZipCode, record.USState, record.USDrivingLicense, record.BankRoutingNumber)
	default:
		return commonData
	}
}

