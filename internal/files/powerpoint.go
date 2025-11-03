package files

import (
	"fmt"
	"os"
	"path/filepath"

	"load-generator/internal/data"
	
	"github.com/unidoc/unioffice/presentation"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/dml"
)

// PowerPointGenerator handles PowerPoint presentation generation
type PowerPointGenerator struct{}

// NewPowerPointGenerator creates a new PowerPoint generator
func NewPowerPointGenerator() *PowerPointGenerator {
	return &PowerPointGenerator{}
}

// Generate creates a PowerPoint presentation with the given records
func (g *PowerPointGenerator) Generate(outputDir, country, sensitiveType string, records []*data.Record, fileIndex int) error {
	// Create directory structure
	dir := filepath.Join(outputDir, "pptx", country, sensitiveType)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create presentation
	ppt := presentation.New()
	
	// Add title slide
	titleSlide := ppt.AddSlide()
	
	// Add title text box
	titleBox := titleSlide.AddTextBox()
	titleBox.Properties().SetPosition(measurement.Distance(1*measurement.Inch), measurement.Distance(2*measurement.Inch))
	titleBox.Properties().SetSize(measurement.Distance(8*measurement.Inch), measurement.Distance(1.5*measurement.Inch))
	
	titlePara := titleBox.AddParagraph()
	titlePara.Properties().SetAlign(dml.ST_TextAlignTypeCtr)
	titleRun := titlePara.AddRun()
	titleRun.SetText("UK Sensitive Test Data")
	titleRun.Properties().SetSize(44)
	titleRun.Properties().SetBold(true)
	
	// Add subtitle
	subtitleBox := titleSlide.AddTextBox()
	subtitleBox.Properties().SetPosition(measurement.Distance(1*measurement.Inch), measurement.Distance(3.8*measurement.Inch))
	subtitleBox.Properties().SetSize(measurement.Distance(8*measurement.Inch), measurement.Distance(1*measurement.Inch))
	
	subtitlePara := subtitleBox.AddParagraph()
	subtitlePara.Properties().SetAlign(dml.ST_TextAlignTypeCtr)
	subtitleRun := subtitlePara.AddRun()
	subtitleRun.SetText(fmt.Sprintf("Data Type: %s | Records: %d", sensitiveType, len(records)))
	subtitleRun.Properties().SetSize(20)
	
	// Add warning slide
	warningSlide := ppt.AddSlide()
	warningBox := warningSlide.AddTextBox()
	warningBox.Properties().SetPosition(measurement.Distance(0.5*measurement.Inch), measurement.Distance(0.5*measurement.Inch))
	warningBox.Properties().SetSize(measurement.Distance(9*measurement.Inch), measurement.Distance(6.5*measurement.Inch))
	
	warningTitle := warningBox.AddParagraph()
	warningTitle.Properties().SetAlign(dml.ST_TextAlignTypeCtr)
	warningTitleRun := warningTitle.AddRun()
	warningTitleRun.SetText("⚠ WARNING ⚠")
	warningTitleRun.Properties().SetSize(36)
	warningTitleRun.Properties().SetBold(true)
	
	warningBox.AddParagraph() // Empty line
	
	warningText := warningBox.AddParagraph()
	warningText.Properties().SetAlign(dml.ST_TextAlignTypeCtr)
	warningRun := warningText.AddRun()
	warningRun.SetText("This presentation contains sensitive test data including:")
	warningRun.Properties().SetSize(20)
	
	warningBox.AddParagraph() // Empty line
	
	for _, item := range []string{
		"• Personal Identifiable Information (PII)",
		"• Payment Card Industry (PCI) data",
		"• Financial information",
	} {
		itemPara := warningBox.AddParagraph()
		itemPara.Properties().SetAlign(dml.ST_TextAlignTypeCtr)
		itemRun := itemPara.AddRun()
		itemRun.SetText(item)
		itemRun.Properties().SetSize(18)
	}
	
	warningBox.AddParagraph() // Empty line
	warningBox.AddParagraph() // Empty line
	
	disclaimerPara := warningBox.AddParagraph()
	disclaimerPara.Properties().SetAlign(dml.ST_TextAlignTypeCtr)
	disclaimerRun := disclaimerPara.AddRun()
	disclaimerRun.SetText("For TESTING PURPOSES ONLY")
	disclaimerRun.Properties().SetSize(24)
	disclaimerRun.Properties().SetBold(true)
	
	// Add record slides (limit to 20 to avoid huge files)
	recordsToShow := len(records)
	if recordsToShow > 20 {
		recordsToShow = 20
	}
	
	for i := 0; i < recordsToShow; i++ {
		record := records[i]
		g.addRecordSlide(ppt, record, i+1)
	}

	// Save presentation
	filename := filepath.Join(dir, fmt.Sprintf("data_%d.pptx", fileIndex))
	if err := ppt.SaveToFile(filename); err != nil {
		return fmt.Errorf("failed to save PowerPoint presentation: %w", err)
	}

	return nil
}

func (g *PowerPointGenerator) addRecordSlide(ppt *presentation.Presentation, record *data.Record, index int) {
	slide := ppt.AddSlide()
	
	// Add title
	titleBox := slide.AddTextBox()
	titleBox.Properties().SetPosition(measurement.Distance(0.5*measurement.Inch), measurement.Distance(0.3*measurement.Inch))
	titleBox.Properties().SetSize(measurement.Distance(9*measurement.Inch), measurement.Distance(0.5*measurement.Inch))
	
	titlePara := titleBox.AddParagraph()
	titleRun := titlePara.AddRun()
	titleRun.SetText(fmt.Sprintf("Record #%d: %s", index, record.FullName))
	titleRun.Properties().SetSize(28)
	titleRun.Properties().SetBold(true)
	
	// Add content
	contentBox := slide.AddTextBox()
	contentBox.Properties().SetPosition(measurement.Distance(0.5*measurement.Inch), measurement.Distance(1*measurement.Inch))
	contentBox.Properties().SetSize(measurement.Distance(9*measurement.Inch), measurement.Distance(5.5*measurement.Inch))
	
	fields := []struct {
		label string
		value string
	}{
		{"Date of Birth", record.DateOfBirth},
		{"Email", record.Email},
		{"NHS Number", record.NHSNumber},
		{"National Insurance", record.NationalInsurance},
		{"Address", record.FullAddress},
		{"Postcode", record.UKPostcode},
		{"Phone", record.PhoneNumber},
		{"Mobile", record.MobileNumber},
		{"Card Number", record.CardNumber},
		{"Card Type", record.CardType},
		{"CVV", record.CVV},
		{"Expiry", record.ExpiryDate},
		{"Bank Account", record.BankAccountNumber},
		{"Sort Code", record.SortCode},
		{"Tax Reference", record.TaxReference},
		{"Income", record.AnnualIncome},
	}
	
	for _, field := range fields {
		para := contentBox.AddParagraph()
		
		// Label
		labelRun := para.AddRun()
		labelRun.SetText(field.label + ": ")
		labelRun.Properties().SetSize(14)
		labelRun.Properties().SetBold(true)
		
		// Value
		valueRun := para.AddRun()
		valueRun.SetText(field.value)
		valueRun.Properties().SetSize(14)
	}
}

