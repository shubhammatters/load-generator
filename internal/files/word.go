package files

import (
	"fmt"
	"os"
	"path/filepath"

	"load-generator/internal/data"
	
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

// WordGenerator handles Word document generation
type WordGenerator struct{}

// NewWordGenerator creates a new Word generator
func NewWordGenerator() *WordGenerator {
	return &WordGenerator{}
}

// Generate creates a Word document with the given records
func (g *WordGenerator) Generate(outputDir, country, sensitiveType string, records []*data.Record, fileIndex int) error {
	// Create directory structure
	dir := filepath.Join(outputDir, "docx", country, sensitiveType)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create document
	doc := document.New()
	
	// Add title
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText("UK Sensitive Test Data")
	run.Properties().SetSize(20)
	run.Properties().SetBold(true)
	run.Properties().Color().SetThemeColor(wml.ST_ThemeColorAccent1)
	para.Properties().SetAlignment(wml.ST_JcCenter)
	
	// Add subtitle
	para = doc.AddParagraph()
	run = para.AddRun()
	run.AddText(fmt.Sprintf("Data Type: %s | Total Records: %d", sensitiveType, len(records)))
	run.Properties().SetSize(12)
	para.Properties().SetAlignment(wml.ST_JcCenter)
	para.Properties().Spacing().SetAfter(measurement.Distance(12 * measurement.Point))
	
	// Add warning
	para = doc.AddParagraph()
	run = para.AddRun()
	run.AddText("⚠ WARNING: ")
	run.Properties().SetBold(true)
	run = para.AddRun()
	run.AddText("This document contains sensitive test data including Personal Identifiable Information (PII), Payment Card Industry (PCI) data, and Financial information. For TESTING PURPOSES ONLY.")
	run.Properties().SetItalic(true)
	para.Properties().Spacing().SetAfter(measurement.Distance(12 * measurement.Point))
	
	// Add records
	for i, record := range records {
		// Record header
		para = doc.AddParagraph()
		run = para.AddRun()
		run.AddText(fmt.Sprintf("Record #%d", i+1))
		run.Properties().SetSize(14)
		run.Properties().SetBold(true)
		para.Properties().Spacing().SetBefore(measurement.Distance(6 * measurement.Point))
		para.Properties().Spacing().SetAfter(measurement.Distance(6 * measurement.Point))
		
		// Create table for this record
		table := doc.AddTable()
		table.Properties().SetWidthPercent(100)
		
		// Add data rows
		g.addTableRow(table, "Name", record.FullName, false)
		g.addTableRow(table, "Date of Birth", record.DateOfBirth, false)
		g.addTableRow(table, "Email", record.Email, false)
		g.addTableRow(table, "NHS Number", record.NHSNumber, true)
		g.addTableRow(table, "National Insurance", record.NationalInsurance, true)
		g.addTableRow(table, "Postcode", record.UKPostcode, false)
		g.addTableRow(table, "Address", record.FullAddress, false)
		g.addTableRow(table, "Phone", record.PhoneNumber, false)
		g.addTableRow(table, "Mobile", record.MobileNumber, false)
		g.addTableRow(table, "Driving License", getDrivingLicenseWord(record), true)
		g.addTableRow(table, "Passport Number", record.PassportNumber, true)
		g.addTableRow(table, "Card Number", record.CardNumber, true)
		g.addTableRow(table, "Card Type", record.CardType, false)
		g.addTableRow(table, "CVV", record.CVV, true)
		g.addTableRow(table, "Expiry Date", record.ExpiryDate, false)
		g.addTableRow(table, "Bank Account", record.BankAccountNumber, true)
		g.addTableRow(table, "Sort Code", record.SortCode, true)
		g.addTableRow(table, "IBAN", record.IBAN, true)
		g.addTableRow(table, "Tax Reference", record.TaxReference, true)
		g.addTableRow(table, "Annual Income", record.AnnualIncome, false)
		
		// Add spacing after table
		para = doc.AddParagraph()
		para.Properties().Spacing().SetAfter(measurement.Distance(6 * measurement.Point))
		
		// Limit to 50 records per document to avoid huge files
		if i >= 49 {
			break
		}
	}

	// Save document
	filename := filepath.Join(dir, fmt.Sprintf("data_%d.docx", fileIndex))
	if err := doc.SaveToFile(filename); err != nil {
		return fmt.Errorf("failed to save Word document: %w", err)
	}

	return nil
}

func (g *WordGenerator) addTableRow(table document.Table, label, value string, sensitive bool) {
	row := table.AddRow()
	
	// Label cell
	cellLabel := row.AddCell()
	cellLabel.Properties().SetWidthPercent(30)
	paraLabel := cellLabel.AddParagraph()
	runLabel := paraLabel.AddRun()
	runLabel.AddText(label)
	runLabel.Properties().SetBold(true)
	
	// Value cell
	cellValue := row.AddCell()
	cellValue.Properties().SetWidthPercent(70)
	paraValue := cellValue.AddParagraph()
	runValue := paraValue.AddRun()
	runValue.AddText(value)
	// Simplified - removed color setting due to API differences
}


func getDrivingLicenseWord(record *data.Record) string {
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
