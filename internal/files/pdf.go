package files

import (
	"fmt"
	"os"
	"path/filepath"

	"load-generator/internal/data"
	
	"github.com/jung-kurt/gofpdf"
)

// PDFGenerator handles PDF file generation
type PDFGenerator struct{}

// NewPDFGenerator creates a new PDF generator
func NewPDFGenerator() *PDFGenerator {
	return &PDFGenerator{}
}

// Generate creates a PDF file with the given records
func (g *PDFGenerator) Generate(outputDir, country, sensitiveType string, records []*data.Record, fileIndex int) error {
	// Create directory structure
	dir := filepath.Join(outputDir, "pdf", country, sensitiveType)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(10, 10, 10)
	
	// Add first page with title
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 20)
	pdf.SetTextColor(25, 118, 210)
	pdf.Cell(0, 10, "UK Sensitive Test Data")
	pdf.Ln(12)
	
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 0)
	pdf.Cell(0, 8, fmt.Sprintf("Data Type: %s", sensitiveType))
	pdf.Ln(6)
	pdf.Cell(0, 8, fmt.Sprintf("Total Records: %d", len(records)))
	pdf.Ln(10)
	
	pdf.SetFont("Arial", "I", 10)
	pdf.SetTextColor(200, 0, 0)
	pdf.MultiCell(0, 5, "WARNING: This document contains sensitive test data including Personal Identifiable Information (PII), Payment Card Industry (PCI) data, and Financial information. For TESTING PURPOSES ONLY.", "", "L", false)
	pdf.Ln(10)

	// Add records
	for i, record := range records {
		// Check if we need a new page
		if pdf.GetY() > 250 {
			pdf.AddPage()
		}
		
		// Record header
		pdf.SetFont("Arial", "B", 14)
		pdf.SetTextColor(25, 118, 210)
		pdf.Cell(0, 8, fmt.Sprintf("Record #%d", i+1))
		pdf.Ln(8)
		
		// Personal Information
		pdf.SetFont("Arial", "B", 11)
		pdf.SetTextColor(0, 0, 0)
		pdf.Cell(0, 6, "Personal Information")
		pdf.Ln(6)
		
		pdf.SetFont("Arial", "", 10)
		g.addField(pdf, "Name", record.FullName)
		g.addField(pdf, "Date of Birth", record.DateOfBirth)
		g.addField(pdf, "Email", record.Email)
		pdf.Ln(3)
		
		// Identification
		pdf.SetFont("Arial", "B", 11)
		pdf.Cell(0, 6, "Identification")
		pdf.Ln(6)
		
		pdf.SetFont("Arial", "", 10)
		pdf.SetTextColor(211, 47, 47) // Red for sensitive data
		g.addField(pdf, "NHS Number", record.NHSNumber)
		g.addField(pdf, "National Insurance", record.NationalInsurance)
		g.addField(pdf, "Driving License", getDrivingLicensePDF(record))
		g.addField(pdf, "Passport Number", record.PassportNumber)
		pdf.Ln(3)
		
		// Contact Information
		pdf.SetFont("Arial", "B", 11)
		pdf.SetTextColor(0, 0, 0)
		pdf.Cell(0, 6, "Contact Information")
		pdf.Ln(6)
		
		pdf.SetFont("Arial", "", 10)
		g.addField(pdf, "Address", record.FullAddress)
		g.addField(pdf, "Postcode", record.UKPostcode)
		g.addField(pdf, "Phone", record.PhoneNumber)
		g.addField(pdf, "Mobile", record.MobileNumber)
		pdf.Ln(3)
		
		// Payment Information
		pdf.SetFont("Arial", "B", 11)
		pdf.SetTextColor(0, 0, 0)
		pdf.Cell(0, 6, "Payment Information (PCI)")
		pdf.Ln(6)
		
		pdf.SetFont("Arial", "", 10)
		pdf.SetTextColor(211, 47, 47) // Red for sensitive data
		g.addField(pdf, "Card Number", record.CardNumber)
		g.addField(pdf, "Card Type", record.CardType)
		g.addField(pdf, "CVV", record.CVV)
		g.addField(pdf, "Expiry Date", record.ExpiryDate)
		pdf.Ln(3)
		
		// Banking Information
		pdf.SetFont("Arial", "B", 11)
		pdf.SetTextColor(0, 0, 0)
		pdf.Cell(0, 6, "Banking Information")
		pdf.Ln(6)
		
		pdf.SetFont("Arial", "", 10)
		pdf.SetTextColor(211, 47, 47) // Red for sensitive data
		g.addField(pdf, "Account Number", record.BankAccountNumber)
		g.addField(pdf, "Sort Code", record.SortCode)
		g.addField(pdf, "IBAN", record.IBAN)
		pdf.Ln(3)
		
		// Tax Information
		pdf.SetFont("Arial", "B", 11)
		pdf.SetTextColor(0, 0, 0)
		pdf.Cell(0, 6, "Tax & Financial Information")
		pdf.Ln(6)
		
		pdf.SetFont("Arial", "", 10)
		pdf.SetTextColor(211, 47, 47) // Red for sensitive data
		g.addField(pdf, "Tax Reference", record.TaxReference)
		pdf.SetTextColor(0, 0, 0)
		g.addField(pdf, "Annual Income", record.AnnualIncome)
		
		// Separator
		pdf.Ln(5)
		pdf.SetDrawColor(200, 200, 200)
		pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
		pdf.Ln(5)
	}

	// Save PDF
	filename := filepath.Join(dir, fmt.Sprintf("data_%d.pdf", fileIndex))
	if err := pdf.OutputFileAndClose(filename); err != nil {
		return fmt.Errorf("failed to save PDF: %w", err)
	}

	return nil
}

func (g *PDFGenerator) addField(pdf *gofpdf.Fpdf, label, value string) {
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(50, 5, label+":")
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 5, value)
	pdf.Ln(5)
}


func getDrivingLicensePDF(record *data.Record) string {
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
