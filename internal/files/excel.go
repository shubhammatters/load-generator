package files

import (
	"fmt"
	"os"
	"path/filepath"

	"load-generator/internal/data"
	
	"github.com/xuri/excelize/v2"
)

// ExcelGenerator handles Excel file generation
type ExcelGenerator struct{}

// NewExcelGenerator creates a new Excel generator
func NewExcelGenerator() *ExcelGenerator {
	return &ExcelGenerator{}
}

// Generate creates an Excel file with the given records
func (g *ExcelGenerator) Generate(outputDir, country, sensitiveType string, records []*data.Record, fileIndex int) error {
	// Create directory structure
	dir := filepath.Join(outputDir, "xlsx", country, sensitiveType)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create Excel file
	f := excelize.NewFile()
	defer f.Close()
	
	sheetName := "Sensitive Data"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf("failed to create sheet: %w", err)
	}
	f.SetActiveSheet(index)
	
	// Delete default Sheet1
	f.DeleteSheet("Sheet1")
	
	// Create styles
	headerStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Color: "FFFFFF",
			Size:  12,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"1976D2"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create header style: %w", err)
	}
	
	sensitiveStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Color: "D32F2F",
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create sensitive style: %w", err)
	}

	// Write headers
	headers := []string{
		"First Name", "Last Name", "Full Name", "Date of Birth", "Email",
		"NHS Number", "National Insurance", "UK Postcode", "Full Address",
		"Phone Number", "Mobile Number", "Driving License", "Passport Number",
		"Card Number", "Card Type", "CVV", "Expiry Date",
		"Bank Account Number", "Sort Code", "IBAN", "Tax Reference", "Annual Income",
	}
	
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
		f.SetCellStyle(sheetName, cell, cell, headerStyle)
	}
	
	// Set column widths
	columnWidths := map[string]float64{
		"A": 12, "B": 12, "C": 20, "D": 12, "E": 25,
		"F": 15, "G": 18, "H": 12, "I": 35,
		"J": 15, "K": 15, "L": 20, "M": 15,
		"N": 20, "O": 15, "P": 8, "Q": 12,
		"R": 18, "S": 12, "T": 28, "U": 15, "V": 15,
	}
	
	for col, width := range columnWidths {
		f.SetColWidth(sheetName, col, col, width)
	}
	
	// Write records
	for i, record := range records {
		row := i + 2 // Start from row 2 (after header)
		
		drivingLicense := getDrivingLicense(record)
		values := []interface{}{
			record.FirstName, record.LastName, record.FullName, record.DateOfBirth, record.Email,
			record.NHSNumber, record.NationalInsurance, record.UKPostcode, record.FullAddress,
			record.PhoneNumber, record.MobileNumber, drivingLicense, record.PassportNumber,
			record.CardNumber, record.CardType, record.CVV, record.ExpiryDate,
			record.BankAccountNumber, record.SortCode, record.IBAN, record.TaxReference, record.AnnualIncome,
		}
		
		for j, value := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			f.SetCellValue(sheetName, cell, value)
			
			// Apply sensitive data style to specific columns
			sensitiveColumns := []int{6, 7, 12, 13, 14, 16, 18, 19, 20, 21} // NHS, NI, License, Passport, Card, CVV, Bank, Sort, IBAN, Tax
			for _, sensCol := range sensitiveColumns {
				if j+1 == sensCol {
					f.SetCellStyle(sheetName, cell, cell, sensitiveStyle)
					break
				}
			}
		}
	}
	
	// Freeze first row
	f.SetPanes(sheetName, &excelize.Panes{
		Freeze:      true,
		XSplit:      0,
		YSplit:      1,
		TopLeftCell: "A2",
		ActivePane:  "bottomLeft",
	})
	
	// Add info sheet
	if err := g.addInfoSheet(f, sensitiveType, len(records)); err != nil {
		return err
	}

	// Save file
	filename := filepath.Join(dir, fmt.Sprintf("data_%d.xlsx", fileIndex))
	if err := f.SaveAs(filename); err != nil {
		return fmt.Errorf("failed to save Excel file: %w", err)
	}

	return nil
}

func (g *ExcelGenerator) addInfoSheet(f *excelize.File, sensitiveType string, recordCount int) error {
	infoSheet := "Information"
	_, err := f.NewSheet(infoSheet)
	if err != nil {
		return fmt.Errorf("failed to create info sheet: %w", err)
	}
	// Set as first sheet by setting it active
	f.SetActiveSheet(0)
	
	// Title style
	titleStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Size:  16,
			Color: "1976D2",
		},
	})
	
	// Warning style
	warningStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Color: "D32F2F",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"FFEBEE"},
			Pattern: 1,
		},
	})
	
	// Content
	f.SetCellValue(infoSheet, "A1", "UK Sensitive Test Data")
	f.SetCellStyle(infoSheet, "A1", "A1", titleStyle)
	
	f.SetCellValue(infoSheet, "A3", "Data Type:")
	f.SetCellValue(infoSheet, "B3", sensitiveType)
	
	f.SetCellValue(infoSheet, "A4", "Total Records:")
	f.SetCellValue(infoSheet, "B4", recordCount)
	
	f.SetCellValue(infoSheet, "A6", "WARNING")
	f.SetCellStyle(infoSheet, "A6", "C6", warningStyle)
	
	f.SetCellValue(infoSheet, "A7", "This workbook contains sensitive test data including:")
	f.SetCellValue(infoSheet, "A8", "• Personal Identifiable Information (PII)")
	f.SetCellValue(infoSheet, "A9", "• Payment Card Industry (PCI) data")
	f.SetCellValue(infoSheet, "A10", "• Financial information")
	f.SetCellValue(infoSheet, "A12", "This data is for TESTING PURPOSES ONLY.")
	f.SetCellValue(infoSheet, "A13", "Do not use in production environments.")
	
	f.SetColWidth(infoSheet, "A", "C", 40)
	
	return nil
}

func getDrivingLicense(record *data.Record) string {
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
