package files

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"load-generator/internal/data"
	
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// ImageGenerator handles image file generation
type ImageGenerator struct{}

// NewImageGenerator creates a new image generator
func NewImageGenerator() *ImageGenerator {
	return &ImageGenerator{}
}

// Generate creates image files (PNG and JPG) with the given records
func (g *ImageGenerator) Generate(outputDir, country, sensitiveType string, records []*data.Record, fileIndex int) error {
	// Generate PNG
	if err := g.generatePNG(outputDir, country, sensitiveType, records, fileIndex); err != nil {
		return err
	}
	
	// Generate JPG
	if err := g.generateJPG(outputDir, country, sensitiveType, records, fileIndex); err != nil {
		return err
	}
	
	return nil
}

func (g *ImageGenerator) generatePNG(outputDir, country, sensitiveType string, records []*data.Record, fileIndex int) error {
	dir := filepath.Join(outputDir, "png", country, sensitiveType)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	filename := filepath.Join(dir, fmt.Sprintf("data_%d.png", fileIndex))
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	img := g.createImage(records, sensitiveType)
	if err := png.Encode(file, img); err != nil {
		return fmt.Errorf("failed to encode PNG: %w", err)
	}

	return nil
}

func (g *ImageGenerator) generateJPG(outputDir, country, sensitiveType string, records []*data.Record, fileIndex int) error {
	dir := filepath.Join(outputDir, "jpg", country, sensitiveType)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	filename := filepath.Join(dir, fmt.Sprintf("data_%d.jpg", fileIndex))
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	img := g.createImage(records, sensitiveType)
	if err := jpeg.Encode(file, img, &jpeg.Options{Quality: 90}); err != nil {
		return fmt.Errorf("failed to encode JPG: %w", err)
	}

	return nil
}

func (g *ImageGenerator) createImage(records []*data.Record, sensitiveType string) *image.RGBA {
	// Create a decent sized image
	width := 800
	recordsToShow := len(records)
	if recordsToShow > 10 {
		recordsToShow = 10 // Show only first 10 records in image
	}
	height := 100 + (recordsToShow * 180) // Header + records
	
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	
	// Fill background with white
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	
	// Draw header background
	headerRect := image.Rect(0, 0, width, 80)
	draw.Draw(img, headerRect, &image.Uniform{color.RGBA{25, 118, 210, 255}}, image.Point{}, draw.Src)
	
	// Draw text
	col := color.RGBA{255, 255, 255, 255}
	g.drawText(img, 20, 30, "UK Sensitive Test Data", col)
	g.drawText(img, 20, 55, fmt.Sprintf("Type: %s | Records: %d", sensitiveType, len(records)), col)
	
	// Draw records
	y := 100
	textCol := color.RGBA{0, 0, 0, 255}
	for i := 0; i < recordsToShow; i++ {
		record := records[i]
		g.drawText(img, 20, y, fmt.Sprintf("Record #%d: %s", i+1, record.FullName), textCol)
		g.drawText(img, 20, y+15, fmt.Sprintf("NHS: %s | NI: %s", record.NHSNumber, record.NationalInsurance), textCol)
		g.drawText(img, 20, y+30, fmt.Sprintf("DOB: %s | Email: %s", record.DateOfBirth, record.Email), textCol)
		g.drawText(img, 20, y+45, fmt.Sprintf("Address: %s", record.FullAddress), textCol)
		g.drawText(img, 20, y+60, fmt.Sprintf("Card: %s (%s) CVV: %s Exp: %s", 
			record.CardNumber, record.CardType, record.CVV, record.ExpiryDate), textCol)
		g.drawText(img, 20, y+75, fmt.Sprintf("Bank: %s Sort: %s IBAN: %s", 
			record.BankAccountNumber, record.SortCode, record.IBAN), textCol)
		
		// Draw separator line
		lineY := y + 90
		for x := 10; x < width-10; x++ {
			img.Set(x, lineY, color.RGBA{200, 200, 200, 255})
		}
		
		y += 180
	}
	
	return img
}

func (g *ImageGenerator) drawText(img *image.RGBA, x, y int, text string, col color.Color) {
	point := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}
	
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(text)
}

