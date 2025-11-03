# Multi-Country Data Support

## Overview

The load generator now supports generating authentic data for three countries:
- **UK (United Kingdom)**
- **India**
- **US (United States)**

## Country-Specific Identifiers

### UK 🇬🇧

| Data Type | Format | Example |
|-----------|--------|---------|
| NHS Number | XXX XXX XXXX | 123 456 7890 |
| National Insurance | AB123456C | MP347116B |
| Postcode | SW1A 1AA | BT49 3KK |
| Driving License | MORGA657054SM9IJ | Wilso40522EI0FL |
| Sort Code | XX-XX-XX | 24-93-89 |
| IBAN | GB29 NWBK 6016 1331 9268 19 | GB22 0249 9374... |

### India 🇮🇳

| Data Type | Format | Example |
|-----------|--------|---------|
| Aadhaar | XXXX XXXX XXXX | 8656 3331 5502 |
| PAN | ABCDE1234F | BUZPC4880D |
| PIN Code | 6 digits | 250062 |
| Driving License | MH01 20160012345 | MH49 20226602742 |
| IFSC Code | ABCD0123456 | HDFC0865894 |
| Mobile | +91 XXXXX XXXXX | +91 89198 67432 |

### US 🇺🇸

| Data Type | Format | Example |
|-----------|--------|---------|
| SSN | XXX-XX-XXXX | 735-80-5959 |
| Green Card | ABC1234567890 | UKY7587894857 |
| ZIP Code | XXXXX or XXXXX-XXXX | 24576-2055 |
| Driving License | A1234567 or 12345678 | Q6564541 |
| Routing Number | 9 digits | 995655328 |
| Phone | (XXX) XXX-XXXX | (703) 409-0348 |

## Output Structure

Files are organized by: **FileType → Country → SensitiveType**

```
output/
├── csv/
│   ├── uk/
│   │   ├── pii/
│   │   ├── pci/
│   │   └── financial/
│   ├── india/
│   │   ├── pii/
│   │   ├── pci/
│   │   └── financial/
│   └── us/
│       ├── pii/
│       ├── pci/
│       └── financial/
└── ... (other file types)
```

## Configuration

Select which countries to generate data for in `config.json`:

```json
{
  "countries": ["uk", "india", "us"]
}
```

Options:
- Generate all three: `["uk", "india", "us"]`
- UK only: `["uk"]`
- India only: `["india"]`
- US only: `["us"]`
- Any combination: `["uk", "us"]`, `["india", "us"]`, etc.

## Sample Data

### UK Sample
```csv
Country,FirstName,LastName,FullName,Email,NHSNumber,NationalInsurance,UKPostcode
uk,George,Wilson,George Wilson,GeorgeWilson@outlook.com,546 701 3970,MP347116B,BT49 3KK
```

### India Sample
```csv
Country,FirstName,LastName,FullName,Email,AadhaarNumber,PANNumber,IndianPostcode
india,Kavya,Chopra,Kavya Chopra,Kavya_Chopra@yahoo.in,8656 3331 5502,BUZPC4880D,250062
```

### US Sample
```csv
Country,FirstName,LastName,FullName,Email,SSN,GreenCard,USZipCode
us,Susan,Brown,Susan Brown,SusanBrown@hotmail.com,735-80-5959,UKY7587894857,24576-2055
```

## Currency Support

Annual income is generated with appropriate currency symbols:
- UK: £ (Pound Sterling)
- India: ₹ (Indian Rupee)
- US: $ (US Dollar)

## Names

Each country uses culturally appropriate names:
- **UK**: James, Emma, Smith, Johnson
- **India**: Rahul, Priya, Sharma, Patel
- **US**: Robert, Jennifer, Brown, Garcia

## Addresses

Addresses include country-specific formats:
- **UK**: Street names, cities like London, Manchester
- **India**: Nagar/Colony, cities like Mumbai, Delhi
- **US**: Avenue/Boulevard, cities like New York, Los Angeles

## Email Domains

Email addresses use appropriate TLDs:
- **UK**: .co.uk
- **India**: .in
- **US**: .com
