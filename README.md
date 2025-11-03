# Multi-Country Sensitive Data Load Generator

A high-performance Go application that generates various file types containing country-specific sensitive data for testing purposes. This tool creates realistic test data including PII (Personal Identifiable Information), PCI (Payment Card Industry), and Financial data for **UK, India, and US**.

## Features

- **Multi-Country Support**: Generate data for UK, India, and US with country-specific formats
- **Multiple File Formats**: PDF, Excel (XLSX), HTML, CSV, JSON, ZIP, PNG, JPG
  - Note: Word (DOCX) and PowerPoint (PPTX) require a commercial license for the unioffice library
- **Country-Specific Data**: 
  - **UK**: NHS numbers, National Insurance, UK postcodes, sort codes
  - **India**: Aadhaar numbers, PAN cards, Indian postcodes, IFSC codes
  - **US**: SSN, Green Cards, US ZIP codes, routing numbers
- **High Performance**: Concurrent generation using worker pools
- **Configurable**: JSON-based configuration for flexible data generation
- **Organized Output**: Files organized by format → country → sensitivity type

## Generated Data Types

### UK-Specific Data
**PII:**
- NHS numbers (XXX XXX XXXX)
- National Insurance numbers (AB123456C)
- UK postcodes
- UK driving license numbers
- UK addresses and phone numbers

**Financial:**
- Bank account numbers (8 digits)
- Sort codes (XX-XX-XX)
- IBAN (GB format)
- UTR (Unique Taxpayer Reference)

### India-Specific Data
**PII:**
- Aadhaar numbers (XXXX XXXX XXXX)
- PAN numbers (ABCDE1234F)
- Indian PIN codes (6 digits)
- Indian driving license numbers
- Indian addresses and phone numbers (+91)
- Indian states

**Financial:**
- Indian bank account numbers (11 digits)
- IFSC codes (ABCD0123456)
- PAN as tax reference

### US-Specific Data
**PII:**
- SSN (XXX-XX-XXXX)
- Green Card numbers
- US ZIP codes (XXXXX or XXXXX-XXXX)
- US driving license numbers
- US addresses and phone numbers
- US states

**Financial:**
- US bank account numbers (12 digits)
- Routing numbers (9 digits)
- SSN as tax reference

### Common Data (All Countries)
**PII:**
- Full names (country-appropriate)
- Dates of birth
- Email addresses (country TLDs)
- Passport numbers

**PCI:**
- Credit/debit card numbers
- Card types (Visa, Mastercard, Amex, Maestro)
- CVV codes
- Expiry dates

**Financial:**
- Annual income (with appropriate currency symbols: £, ₹, $)

## Installation

### Prerequisites
- Go 1.19 or higher

### Build

```bash
# Clone or download the repository
cd load-generator

# Build the binary
go build -o load-generator

# Make it executable (Unix/Linux/Mac)
chmod +x load-generator
```

## Configuration

The application uses a `config.json` file for configuration. If the file doesn't exist, it will be created with default values on first run.

### config.json Structure

```json
{
  "files_per_type": 10,
  "records_per_file": 100,
  "file_size_target_mb": 10,
  "countries": ["uk", "india", "us"],
  "sensitive_types": ["pii", "pci", "financial"],
  "file_types": ["pdf", "xlsx", "html", "csv", "json", "zip", "png", "jpg"],
  "worker_count": 10,
  "output_dir": "output"
}
```

### Configuration Options

| Option | Description | Default |
|--------|-------------|---------|
| `files_per_type` | Number of files to generate per type, country, and sensitivity combination | 10 |
| `records_per_file` | Number of data records in each file | 100 |
| `file_size_target_mb` | Target file size in MB (informational) | 10 |
| `countries` | Countries to generate data for | `["uk", "india", "us"]` |
| `sensitive_types` | Types of sensitive data to generate | `["pii", "pci", "financial"]` |
| `file_types` | File formats to generate | `["pdf", "xlsx", "html", ...]` |
| `worker_count` | Number of concurrent workers | 10 |
| `output_dir` | Output directory for generated files | `"output"` |

## Usage

### Basic Usage

```bash
./load-generator
```

The application will:
1. Load configuration from `config.json`
2. Create the output directory structure
3. Generate files concurrently using worker pools
4. Display progress and statistics

### Output Structure

Files are organized in a structured hierarchy: **format → country → sensitivity type**

```
output/
├── pdf/
│   ├── uk/
│   │   ├── pii/
│   │   │   ├── data_1.pdf
│   │   │   ├── data_2.pdf
│   │   │   └── ...
│   │   ├── pci/
│   │   │   └── ...
│   │   └── financial/
│   │       └── ...
│   ├── india/
│   │   ├── pii/
│   │   ├── pci/
│   │   └── financial/
│   └── us/
│       ├── pii/
│       ├── pci/
│       └── financial/
├── csv/
│   ├── uk/
│   ├── india/
│   └── us/
├── xlsx/
│   └── ...
└── ... (other file types)
```

### Example: Generate 100GB+ of Data

To generate approximately 100GB+ of data, adjust your `config.json`:

```json
{
  "files_per_type": 50,
  "records_per_file": 1000,
  "countries": ["uk", "india", "us"],
  "file_types": ["pdf", "xlsx", "html", "csv", "json", "zip"],
  "sensitive_types": ["pii", "pci", "financial"],
  "worker_count": 20
}
```

This configuration will create:
- 6 file types × 3 countries × 3 sensitive types × 50 files = **2,700 files**
- Each file with 1,000 records
- Using 20 concurrent workers for faster generation
- Approximately 120GB+ of diverse data

### Example: Quick Test Run

For a quick test with minimal data:

```json
{
  "files_per_type": 1,
  "records_per_file": 10,
  "countries": ["uk", "india"],
  "file_types": ["csv", "json", "html"],
  "sensitive_types": ["pii"],
  "worker_count": 5
}
```

**Result:** 3 file types × 2 countries × 1 sensitive type × 1 file = **6 files** in seconds

## Performance

The generator is optimized for performance:

- **Concurrent Processing**: Uses goroutine worker pools
- **Buffered I/O**: Efficient file writing
- **Batch Generation**: Generates data in batches

Typical performance on modern hardware:
- ~50-100 files/second for simple formats (CSV, JSON, HTML)
- ~10-20 files/second for complex formats (PDF, DOCX, XLSX)

## Output Examples

### CSV Format
```csv
FirstName,LastName,FullName,DateOfBirth,Email,NHSNumber,...
James,Smith,James Smith,15/03/1985,james.smith@gmail.com,123 456 7890,...
```

### JSON Format
```json
{
  "records": [
    {
      "FirstName": "James",
      "LastName": "Smith",
      "FullName": "James Smith",
      "NHSNumber": "123 456 7890",
      ...
    }
  ],
  "count": 100,
  "type": "pii"
}
```

### PDF/Word/Excel/PowerPoint
Formatted documents with tables, styling, and organized sections for different data types.

## Security & Legal Notice

⚠️ **WARNING**: This tool generates realistic-looking sensitive data for TESTING PURPOSES ONLY.

- **Do NOT use this data in production environments**
- **Do NOT use this data for fraudulent purposes**
- All generated data is completely fictional
- No real personal information is used or generated
- Generated data should be handled according to your organization's test data policies

## Dependencies

The project uses the following Go libraries:

- `github.com/jung-kurt/gofpdf` - PDF generation (free, open source)
- `github.com/xuri/excelize/v2` - Excel file generation (free, open source)
- `github.com/unidoc/unioffice` - Word and PowerPoint generation (requires commercial license)
- `golang.org/x/image` - Image text rendering (free, open source)

Dependencies are automatically downloaded when you run `go build`.

### Note on Word and PowerPoint Generation

The `unioffice` library used for DOCX and PPTX generation requires a commercial license. The default configuration excludes these file types. If you have a license, you can:
1. Add `"docx"` and `"pptx"` to the `file_types` array in `config.json`
2. Set your license key in the code or environment variable as per unioffice documentation

The other file formats (PDF, Excel, CSV, JSON, HTML, ZIP, Images) work without any licensing restrictions.

## Building from Source

```bash
# Initialize module (if starting fresh)
go mod init load-generator

# Download dependencies
go mod tidy

# Build
go build -o load-generator

# Run
./load-generator
```

## Troubleshooting

### "Permission denied" error
Make sure the binary is executable:
```bash
chmod +x load-generator
```

### Out of memory errors
Reduce the `worker_count` and `records_per_file` values in config.json.

### Slow generation
- Increase `worker_count` (but not beyond your CPU core count × 2)
- Use faster storage (SSD recommended)
- Generate simpler file types first (CSV, JSON) for quick testing

## System Requirements

- **Minimum**: 2 CPU cores, 4GB RAM
- **Recommended**: 4+ CPU cores, 8GB+ RAM, SSD storage
- **For 100GB+ generation**: 8+ CPU cores, 16GB+ RAM, 200GB+ free disk space

## License

This tool is provided as-is for testing purposes.

## Contributing

Feel free to contribute improvements, bug fixes, or new features.

## Support

For issues or questions, please check the documentation or create an issue in the repository.

