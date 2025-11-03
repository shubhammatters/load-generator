# Quick Start Guide

## Get Started in 3 Steps

### 1. Build the Application

```bash
cd load-generator
go build -o load-generator
```

If you encounter permission issues with Go module cache:
```bash
GOPATH=$(pwd)/.go GOMODCACHE=$(pwd)/.gomodcache go build -o load-generator
```

### 2. Configure (Optional)

The default `config.json` is already set up for a quick test. Edit it to customize:

```json
{
  "files_per_type": 10,
  "records_per_file": 100,
  "sensitive_types": ["pii", "pci", "financial"],
  "file_types": ["pdf", "xlsx", "html", "csv", "json", "zip", "png", "jpg"],
  "worker_count": 10,
  "output_dir": "output"
}
```

### 3. Run

```bash
./load-generator
```

That's it! Your files will be generated in the `output/` directory.

---

## Common Scenarios

### Small Test Run (Quick Validation)

**Config:** 6 files, small data set
```json
{
  "files_per_type": 1,
  "records_per_file": 10,
  "sensitive_types": ["pii"],
  "file_types": ["csv", "json", "html"],
  "worker_count": 5,
  "output_dir": "output"
}
```

**Result:** 3 files in seconds

---

### Medium Test (Development/Testing)

**Config:** 900 files, ~50MB
```json
{
  "files_per_type": 10,
  "records_per_file": 100,
  "sensitive_types": ["pii", "pci", "financial"],
  "file_types": ["pdf", "xlsx", "html", "csv", "json", "zip", "png", "jpg"],
  "worker_count": 10,
  "output_dir": "output"
}
```

**Result:** ~240 files, ~25MB in 1-2 seconds

---

### Large Test (Load Testing - ~1GB)

**Config:** 2,400 files, ~1GB
```json
{
  "files_per_type": 100,
  "records_per_file": 100,
  "sensitive_types": ["pii", "pci", "financial"],
  "file_types": ["pdf", "xlsx", "html", "csv", "json", "zip", "png", "jpg"],
  "worker_count": 20,
  "output_dir": "output"
}
```

**Result:** 2,400 files, ~1GB in 5-10 seconds

---

### Massive Dataset (100GB+)

**Config:** 24,000 files, 100+ GB
```json
{
  "files_per_type": 1000,
  "records_per_file": 1000,
  "sensitive_types": ["pii", "pci", "financial"],
  "file_types": ["pdf", "xlsx", "html", "csv", "json", "zip", "png", "jpg"],
  "worker_count": 20,
  "output_dir": "output"
}
```

**Result:** 24,000 files, 100+ GB in 2-5 minutes (depending on hardware)

**Requirements:**
- 8+ CPU cores
- 16+ GB RAM
- SSD storage
- 200+ GB free disk space

---

## Understanding the Output

### Directory Structure

```
output/
├── csv/
│   ├── pii/
│   │   ├── data_1.csv
│   │   ├── data_2.csv
│   │   └── ...
│   ├── pci/
│   │   └── ...
│   └── financial/
│       └── ...
├── json/
│   └── ...
├── pdf/
│   └── ...
└── ... (other formats)
```

### File Naming

Files are named sequentially:
- `data_1.csv`
- `data_2.csv`
- `data_3.csv`
- etc.

---

## Tips for Best Performance

### 1. Adjust Worker Count
- **Recommended:** Number of CPU cores × 2
- **Example:** 8-core CPU = `"worker_count": 16`

### 2. Use SSD Storage
- Dramatically faster than HDD
- Essential for large datasets

### 3. Batch Generation
Generate data in batches if targeting very large volumes:
```bash
# Generate 100GB in 10 batches of 10GB each
for i in {1..10}; do
  sed -i '' "s/\"output_dir\": \".*\"/\"output_dir\": \"output_batch_$i\"/" config.json
  ./load-generator
done
```

### 4. Monitor System Resources
```bash
# Run in one terminal
./load-generator

# Monitor in another
htop  # or top
```

---

## Specific File Type Generation

Want only specific formats? Customize `file_types`:

### CSV and JSON only (fastest)
```json
{
  "file_types": ["csv", "json"]
}
```

### Documents only
```json
{
  "file_types": ["pdf", "xlsx", "html"]
}
```

### Images only
```json
{
  "file_types": ["png", "jpg"]
}
```

### Archives only
```json
{
  "file_types": ["zip"]
}
```

---

## Specific Data Type Generation

Want only specific sensitive data? Customize `sensitive_types`:

### PII only
```json
{
  "sensitive_types": ["pii"]
}
```

### Financial data only
```json
{
  "sensitive_types": ["financial"]
}
```

### PCI data only
```json
{
  "sensitive_types": ["pci"]
}
```

---

## Viewing Generated Data

### CSV Files
```bash
# View in terminal
head output/csv/pii/data_1.csv

# Open in Excel/LibreOffice
open output/csv/pii/data_1.csv
```

### JSON Files
```bash
# Pretty print
cat output/json/pii/data_1.json | jq

# View first record
cat output/json/pii/data_1.json | jq '.records[0]'
```

### HTML Files
```bash
# Open in browser
open output/html/pii/data_1.html
```

### PDF Files
```bash
# Open with default PDF viewer
open output/pdf/pii/data_1.pdf
```

### Excel Files
```bash
# Open in Excel/LibreOffice
open output/xlsx/pii/data_1.xlsx
```

---

## Cleaning Up

### Remove all generated files
```bash
rm -rf output/
```

### Remove specific type
```bash
rm -rf output/pdf/
```

### Remove specific sensitive type
```bash
rm -rf output/*/pii/
```

---

## Troubleshooting

### "Permission denied" when running
```bash
chmod +x load-generator
./load-generator
```

### Go module permission errors
```bash
# Use local cache
GOPATH=$(pwd)/.go GOMODCACHE=$(pwd)/.gomodcache go build -o load-generator
```

### Out of memory
Reduce `worker_count` and `records_per_file` in config.json

### Slow generation
- Use SSD storage
- Increase `worker_count`
- Close other applications
- Use simpler file formats (CSV, JSON)

### Disk full
- Check free space: `df -h`
- Reduce `files_per_type`
- Generate fewer file types

---

## Next Steps

1. ✅ Generate your first dataset
2. ✅ Open and inspect some files
3. ✅ Use the data in your testing environment
4. ✅ Adjust config for your needs
5. ✅ Integrate into your CI/CD pipeline (optional)

---

## Questions?

Check the main [README.md](README.md) for detailed documentation.

