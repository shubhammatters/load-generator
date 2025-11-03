package files

import (
	"fmt"
	"html"
	"os"
	"path/filepath"

	"load-generator/internal/data"
)

// HTMLGenerator handles HTML file generation
type HTMLGenerator struct{}

// NewHTMLGenerator creates a new HTML generator
func NewHTMLGenerator() *HTMLGenerator {
	return &HTMLGenerator{}
}

// Generate creates an HTML file with the given records
func (g *HTMLGenerator) Generate(outputDir, country, sensitiveType string, records []*data.Record, fileIndex int) error {
	// Create directory structure
	dir := filepath.Join(outputDir, "html", country, sensitiveType)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create file
	filename := filepath.Join(dir, fmt.Sprintf("data_%d.html", fileIndex))
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write HTML header
	fmt.Fprintf(file, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UK Sensitive Data - %s</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; background-color: #f5f5f5; }
        h1 { color: #333; }
        .record { 
            background: white; 
            padding: 15px; 
            margin: 10px 0; 
            border-radius: 5px; 
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
        }
        .field { margin: 5px 0; }
        .label { font-weight: bold; color: #555; }
        .value { color: #333; }
        .sensitive { color: #d32f2f; }
        .header { background: #1976d2; color: white; padding: 20px; border-radius: 5px; margin-bottom: 20px; }
    </style>
</head>
<body>
    <div class="header">
        <h1>UK Sensitive Test Data</h1>
        <p>Data Type: %s | Record Count: %d</p>
    </div>
`, html.EscapeString(sensitiveType), html.EscapeString(sensitiveType), len(records))

	// Write records
	for i, record := range records {
		fmt.Fprintf(file, `    <div class="record">
        <h3>Record #%d</h3>
        <div class="field"><span class="label">Name:</span> <span class="value">%s</span></div>
        <div class="field"><span class="label">Date of Birth:</span> <span class="value">%s</span></div>
        <div class="field"><span class="label">Email:</span> <span class="value">%s</span></div>
        <div class="field"><span class="label">NHS Number:</span> <span class="value sensitive">%s</span></div>
        <div class="field"><span class="label">National Insurance:</span> <span class="value sensitive">%s</span></div>
        <div class="field"><span class="label">Postcode:</span> <span class="value">%s</span></div>
        <div class="field"><span class="label">Address:</span> <span class="value">%s</span></div>
        <div class="field"><span class="label">Phone:</span> <span class="value">%s</span></div>
        <div class="field"><span class="label">Mobile:</span> <span class="value">%s</span></div>
        <div class="field"><span class="label">Driving License:</span> <span class="value sensitive">%s</span></div>
        <div class="field"><span class="label">Passport:</span> <span class="value sensitive">%s</span></div>
        <div class="field"><span class="label">Card Number:</span> <span class="value sensitive">%s</span></div>
        <div class="field"><span class="label">Card Type:</span> <span class="value">%s</span></div>
        <div class="field"><span class="label">CVV:</span> <span class="value sensitive">%s</span></div>
        <div class="field"><span class="label">Expiry Date:</span> <span class="value">%s</span></div>
        <div class="field"><span class="label">Bank Account:</span> <span class="value sensitive">%s</span></div>
        <div class="field"><span class="label">Sort Code:</span> <span class="value sensitive">%s</span></div>
        <div class="field"><span class="label">IBAN:</span> <span class="value sensitive">%s</span></div>
        <div class="field"><span class="label">Tax Reference:</span> <span class="value sensitive">%s</span></div>
        <div class="field"><span class="label">Annual Income:</span> <span class="value">%s</span></div>
    </div>
`,
			i+1,
			html.EscapeString(record.FullName),
			html.EscapeString(record.DateOfBirth),
			html.EscapeString(record.Email),
			html.EscapeString(record.NHSNumber),
			html.EscapeString(record.NationalInsurance),
			html.EscapeString(record.UKPostcode),
			html.EscapeString(record.FullAddress),
			html.EscapeString(record.PhoneNumber),
			html.EscapeString(record.MobileNumber),
			html.EscapeString(getDrivingLicenseHTML(record)),
			html.EscapeString(record.PassportNumber),
			html.EscapeString(record.CardNumber),
			html.EscapeString(record.CardType),
			html.EscapeString(record.CVV),
			html.EscapeString(record.ExpiryDate),
			html.EscapeString(record.BankAccountNumber),
			html.EscapeString(record.SortCode),
			html.EscapeString(record.IBAN),
			html.EscapeString(record.TaxReference),
			html.EscapeString(record.AnnualIncome),
		)
	}

	// Write HTML footer
	fmt.Fprintf(file, `</body>
</html>`)

	return nil
}


func getDrivingLicenseHTML(record *data.Record) string {
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
