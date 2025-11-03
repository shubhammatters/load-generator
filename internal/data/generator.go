package data

import (
	"fmt"
	"math/rand"
	"time"
)

// SensitiveType represents the type of sensitive data
type SensitiveType string

const (
	PII       SensitiveType = "pii"
	PCI       SensitiveType = "pci"
	Financial SensitiveType = "financial"
)

// Country represents the country for data generation
type Country string

const (
	UK    Country = "uk"
	India Country = "india"
	US    Country = "us"
)

// Record represents a complete record with sensitive data
type Record struct {
	Country Country
	
	// Common PII Data
	FirstName   string
	LastName    string
	FullName    string
	DateOfBirth string
	Email       string
	FullAddress string
	PhoneNumber string
	MobileNumber string
	PassportNumber string
	
	// UK-specific PII
	NHSNumber         string
	NationalInsurance string
	UKPostcode        string
	UKDrivingLicense  string
	
	// India-specific PII
	AadhaarNumber     string
	PANNumber         string
	IndianPostcode    string
	IndianState       string
	IndianDrivingLicense string
	
	// US-specific PII
	SSN              string
	GreenCard        string
	USZipCode        string
	USState          string
	USDrivingLicense string
	
	// PCI Data (common)
	CardNumber string
	CardType   string
	CVV        string
	ExpiryDate string
	
	// Financial Data (country-specific)
	BankAccountNumber string
	BankRoutingNumber string // US
	SortCode          string // UK
	IFSCCode          string // India
	IBAN              string
	TaxReference      string // UTR (UK), PAN (India), EIN/SSN (US)
	AnnualIncome      string
}

var (
	// UK names
	ukFirstNames = []string{
		"James", "John", "Robert", "Michael", "William", "David", "Richard", "Joseph", "Thomas", "Charles",
		"Mary", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara", "Susan", "Jessica", "Sarah", "Karen",
		"Emma", "Olivia", "Ava", "Isabella", "Sophia", "Mia", "Charlotte", "Amelia", "Harper", "Evelyn",
		"Oliver", "George", "Harry", "Jack", "Jacob", "Noah", "Charlie", "Muhammad", "Oscar", "Henry",
	}
	
	ukLastNames = []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez",
		"Wilson", "Anderson", "Taylor", "Thomas", "Moore", "Jackson", "Martin", "Lee", "Thompson", "White",
		"Harris", "Clark", "Lewis", "Robinson", "Walker", "Young", "Hall", "Allen", "King", "Wright",
		"Patel", "Khan", "Murphy", "O'Brien", "Kelly", "Walsh", "O'Connor", "McLaughlin", "MacDonald",
	}
	
	ukStreetNames = []string{
		"High Street", "Station Road", "Main Street", "Church Lane", "Park Road", "London Road",
		"Manor Road", "Church Street", "Mill Lane", "Victoria Road", "Green Lane", "The Avenue",
		"Queens Road", "New Road", "Grange Road", "Kings Road", "Kingsway", "Windsor Road",
		"School Lane", "The Crescent", "Springfield Road", "The Grove", "York Road", "Broadway",
	}
	
	ukCities = []string{
		"London", "Birmingham", "Manchester", "Leeds", "Glasgow", "Sheffield", "Liverpool", "Edinburgh",
		"Bristol", "Cardiff", "Leicester", "Coventry", "Bradford", "Belfast", "Nottingham", "Newcastle",
		"Brighton", "Hull", "Plymouth", "Stoke", "Wolverhampton", "Derby", "Southampton", "Portsmouth",
	}
	
	// Indian names
	indianFirstNames = []string{
		"Rahul", "Priya", "Amit", "Anjali", "Raj", "Neha", "Vikram", "Pooja", "Arjun", "Kavya",
		"Rohan", "Shreya", "Aditya", "Divya", "Karan", "Riya", "Varun", "Sanya", "Aryan", "Ishita",
		"Siddharth", "Ananya", "Abhishek", "Aarti", "Kunal", "Megha", "Manish", "Nisha", "Suresh", "Sunita",
		"Rajesh", "Lakshmi", "Manoj", "Deepika", "Sanjay", "Rekha", "Vijay", "Meera", "Ashok", "Gayatri",
	}
	
	indianLastNames = []string{
		"Sharma", "Patel", "Kumar", "Singh", "Gupta", "Reddy", "Verma", "Joshi", "Nair", "Menon",
		"Chopra", "Kapoor", "Mehta", "Shah", "Iyer", "Rao", "Desai", "Kulkarni", "Agarwal", "Banerjee",
		"Chatterjee", "Das", "Ghosh", "Malhotra", "Pillai", "Shetty", "Bhat", "Naik", "Pandey", "Mishra",
	}
	
	indianCities = []string{
		"Mumbai", "Delhi", "Bangalore", "Hyderabad", "Chennai", "Kolkata", "Pune", "Ahmedabad",
		"Surat", "Jaipur", "Lucknow", "Kanpur", "Nagpur", "Indore", "Bhopal", "Visakhapatnam",
		"Patna", "Vadodara", "Ghaziabad", "Ludhiana", "Agra", "Nashik", "Faridabad", "Meerut",
	}
	
	indianStates = []string{
		"Maharashtra", "Karnataka", "Tamil Nadu", "Uttar Pradesh", "Gujarat", "West Bengal",
		"Rajasthan", "Kerala", "Madhya Pradesh", "Punjab", "Haryana", "Bihar", "Telangana",
	}
	
	// US names
	usFirstNames = []string{
		"James", "Mary", "John", "Patricia", "Robert", "Jennifer", "Michael", "Linda", "William", "Elizabeth",
		"David", "Barbara", "Richard", "Susan", "Joseph", "Jessica", "Thomas", "Sarah", "Charles", "Karen",
		"Christopher", "Nancy", "Daniel", "Lisa", "Matthew", "Betty", "Anthony", "Margaret", "Mark", "Sandra",
		"Donald", "Ashley", "Steven", "Kimberly", "Paul", "Emily", "Andrew", "Donna", "Joshua", "Michelle",
	}
	
	usLastNames = []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez",
		"Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin",
		"Lee", "Perez", "Thompson", "White", "Harris", "Sanchez", "Clark", "Ramirez", "Lewis", "Robinson",
	}
	
	usStreetNames = []string{
		"Main Street", "Oak Avenue", "Maple Drive", "Cedar Lane", "Pine Street", "Elm Avenue",
		"Washington Street", "Park Avenue", "Lake Road", "Hill Street", "Sunset Boulevard", "River Road",
		"First Avenue", "Second Street", "Broadway", "Madison Avenue", "Lincoln Way", "Jefferson Street",
	}
	
	usCities = []string{
		"New York", "Los Angeles", "Chicago", "Houston", "Phoenix", "Philadelphia", "San Antonio", "San Diego",
		"Dallas", "San Jose", "Austin", "Jacksonville", "Fort Worth", "Columbus", "Charlotte", "San Francisco",
		"Indianapolis", "Seattle", "Denver", "Washington", "Boston", "Nashville", "Baltimore", "Portland",
	}
	
	usStates = []string{
		"California", "Texas", "Florida", "New York", "Pennsylvania", "Illinois", "Ohio", "Georgia",
		"North Carolina", "Michigan", "New Jersey", "Virginia", "Washington", "Arizona", "Massachusetts",
	}
	
	cardTypes = []string{"Visa", "Mastercard", "American Express", "Maestro"}
)

// NewGenerator creates a new data generator with a seeded random
func NewGenerator(seed int64) *Generator {
	return &Generator{
		rand: rand.New(rand.NewSource(seed)),
	}
}

// Generator generates UK-specific sensitive data
type Generator struct {
	rand *rand.Rand
}

// GenerateRecord creates a complete record with country-specific sensitive data
func (g *Generator) GenerateRecord(country Country) *Record {
	record := &Record{
		Country:     country,
		DateOfBirth: g.generateDOB(),
		CardNumber:  g.generateCardNumber(),
		CardType:    g.randomChoice(cardTypes),
		CVV:         g.generateCVV(),
		ExpiryDate:  g.generateExpiryDate(),
	}
	
	switch country {
	case UK:
		g.fillUKData(record)
	case India:
		g.fillIndianData(record)
	case US:
		g.fillUSData(record)
	}
	
	return record
}

func (g *Generator) fillUKData(record *Record) {
	firstName := g.randomChoice(ukFirstNames)
	lastName := g.randomChoice(ukLastNames)
	
	record.FirstName = firstName
	record.LastName = lastName
	record.FullName = fmt.Sprintf("%s %s", firstName, lastName)
	record.Email = g.generateEmail(firstName, lastName, "co.uk")
	record.NHSNumber = g.generateNHSNumber()
	record.NationalInsurance = g.generateNINumber()
	record.UKPostcode = g.generateUKPostcode()
	record.FullAddress = g.generateUKAddress()
	record.PhoneNumber = g.generateUKLandline()
	record.MobileNumber = g.generateUKMobile()
	record.UKDrivingLicense = g.generateUKDrivingLicense(lastName)
	record.PassportNumber = g.generatePassportNumber()
	record.BankAccountNumber = g.generateUKBankAccount()
	record.SortCode = g.generateSortCode()
	record.IBAN = g.generateUKIBAN()
	record.TaxReference = g.generateUTR()
	record.AnnualIncome = g.generateIncome("£")
}

func (g *Generator) fillIndianData(record *Record) {
	firstName := g.randomChoice(indianFirstNames)
	lastName := g.randomChoice(indianLastNames)
	
	record.FirstName = firstName
	record.LastName = lastName
	record.FullName = fmt.Sprintf("%s %s", firstName, lastName)
	record.Email = g.generateEmail(firstName, lastName, "in")
	record.AadhaarNumber = g.generateAadhaar()
	record.PANNumber = g.generatePAN(lastName)
	record.IndianPostcode = g.generateIndianPostcode()
	record.IndianState = g.randomChoice(indianStates)
	record.FullAddress = g.generateIndianAddress()
	record.PhoneNumber = g.generateIndianLandline()
	record.MobileNumber = g.generateIndianMobile()
	record.IndianDrivingLicense = g.generateIndianDrivingLicense(record.IndianState)
	record.PassportNumber = g.generatePassportNumber()
	record.BankAccountNumber = g.generateIndianBankAccount()
	record.IFSCCode = g.generateIFSC()
	record.IBAN = ""  // India doesn't commonly use IBAN
	record.TaxReference = record.PANNumber  // PAN is the tax reference in India
	record.AnnualIncome = g.generateIncome("₹")
}

func (g *Generator) fillUSData(record *Record) {
	firstName := g.randomChoice(usFirstNames)
	lastName := g.randomChoice(usLastNames)
	
	record.FirstName = firstName
	record.LastName = lastName
	record.FullName = fmt.Sprintf("%s %s", firstName, lastName)
	record.Email = g.generateEmail(firstName, lastName, "com")
	record.SSN = g.generateSSN()
	record.GreenCard = g.generateGreenCard()
	record.USZipCode = g.generateUSZipCode()
	record.USState = g.randomChoice(usStates)
	record.FullAddress = g.generateUSAddress()
	record.PhoneNumber = g.generateUSPhone()
	record.MobileNumber = g.generateUSMobile()
	record.USDrivingLicense = g.generateUSDrivingLicense(record.USState)
	record.PassportNumber = g.generatePassportNumber()
	record.BankAccountNumber = g.generateUSBankAccount()
	record.BankRoutingNumber = g.generateUSRoutingNumber()
	record.IBAN = ""  // US doesn't typically use IBAN
	record.TaxReference = record.SSN  // SSN is the tax reference in US
	record.AnnualIncome = g.generateIncome("$")
}

// generateNHSNumber generates a UK NHS number (XXX XXX XXXX format)
func (g *Generator) generateNHSNumber() string {
	return fmt.Sprintf("%03d %03d %04d", 
		g.rand.Intn(1000), 
		g.rand.Intn(1000), 
		g.rand.Intn(10000))
}

// generateNINumber generates a UK National Insurance number (AB123456C format)
func (g *Generator) generateNINumber() string {
	// First letter cannot be D, F, I, Q, U, V
	validFirst := "ABCEGHJKLMNOPRSTWXYZ"
	// Second letter cannot be D, F, I, O, Q, U, V
	validSecond := "ABCEGHJKLMNPRSTWXYZ"
	// Last letter is always A, B, C, or D
	validLast := "ABCD"
	
	return fmt.Sprintf("%c%c%06d%c",
		validFirst[g.rand.Intn(len(validFirst))],
		validSecond[g.rand.Intn(len(validSecond))],
		g.rand.Intn(1000000),
		validLast[g.rand.Intn(len(validLast))],
	)
}

// UK-specific generators

// generateUKPostcode generates a realistic UK postcode
func (g *Generator) generateUKPostcode() string {
	area := []string{"SW", "SE", "NW", "NE", "E", "W", "N", "S", "EC", "WC", "M", "B", "LS", "G", "EH", "CF", "BT"}
	selectedArea := area[g.rand.Intn(len(area))]
	district := g.rand.Intn(99) + 1
	sector := g.rand.Intn(9) + 1
	unit := fmt.Sprintf("%c%c", 'A'+byte(g.rand.Intn(26)), 'A'+byte(g.rand.Intn(26)))
	
	return fmt.Sprintf("%s%d %d%s", selectedArea, district, sector, unit)
}

// generateUKAddress generates a full UK address
func (g *Generator) generateUKAddress() string {
	houseNum := g.rand.Intn(200) + 1
	street := g.randomChoice(ukStreetNames)
	city := g.randomChoice(ukCities)
	postcode := g.generateUKPostcode()
	
	return fmt.Sprintf("%d %s, %s, %s", houseNum, street, city, postcode)
}

// generateUKLandline generates a UK landline number
func (g *Generator) generateUKLandline() string {
	areaCodes := []string{"020", "0121", "0131", "0141", "0151", "0161", "0117", "01632"}
	areaCode := areaCodes[g.rand.Intn(len(areaCodes))]
	
	return fmt.Sprintf("%s %d%04d", areaCode, g.rand.Intn(1000), g.rand.Intn(10000))
}

// generateUKMobile generates a UK mobile number
func (g *Generator) generateUKMobile() string {
	return fmt.Sprintf("07%03d %06d", g.rand.Intn(1000), g.rand.Intn(1000000))
}

// generateUKDrivingLicense generates a UK driving license number
func (g *Generator) generateUKDrivingLicense(lastName string) string {
	// Format: MORGA657054SM9IJ (surname-based + numbers + initials + numbers)
	surname := lastName
	if len(surname) > 5 {
		surname = surname[:5]
	}
	for len(surname) < 5 {
		surname += "9"
	}
	
	return fmt.Sprintf("%s%d%02d%02d%c%c%d%c%c",
		surname,
		g.rand.Intn(10),
		g.rand.Intn(13)+1,  // Month
		g.rand.Intn(32)+1,  // Day
		'A'+byte(g.rand.Intn(26)),
		'A'+byte(g.rand.Intn(26)),
		g.rand.Intn(10),
		'A'+byte(g.rand.Intn(26)),
		'A'+byte(g.rand.Intn(26)),
	)
}

// generateUKBankAccount generates an 8-digit UK bank account number
func (g *Generator) generateUKBankAccount() string {
	return fmt.Sprintf("%08d", g.rand.Intn(100000000))
}

// generateUKIBAN generates a UK IBAN
func (g *Generator) generateUKIBAN() string {
	return fmt.Sprintf("GB%02d %s %s %04d %04d %04d %02d",
		g.rand.Intn(100),
		fmt.Sprintf("%04d", g.rand.Intn(10000)),
		fmt.Sprintf("%04d", g.rand.Intn(10000)),
		g.rand.Intn(10000),
		g.rand.Intn(10000),
		g.rand.Intn(10000),
		g.rand.Intn(100),
	)
}

// generatePassportNumber generates a UK passport number
func (g *Generator) generatePassportNumber() string {
	return fmt.Sprintf("%09d", 100000000+g.rand.Intn(900000000))
}

// generateCardNumber generates a valid-looking credit card number
func (g *Generator) generateCardNumber() string {
	// Visa starts with 4, Mastercard with 5
	prefix := []int{4, 5}
	firstDigit := prefix[g.rand.Intn(len(prefix))]
	
	return fmt.Sprintf("%d%03d %04d %04d %04d",
		firstDigit,
		g.rand.Intn(1000),
		g.rand.Intn(10000),
		g.rand.Intn(10000),
		g.rand.Intn(10000),
	)
}

// generateCVV generates a 3-digit CVV
func (g *Generator) generateCVV() string {
	return fmt.Sprintf("%03d", g.rand.Intn(1000))
}

// generateExpiryDate generates a card expiry date in MM/YY format
func (g *Generator) generateExpiryDate() string {
	month := g.rand.Intn(12) + 1
	year := time.Now().Year() + g.rand.Intn(5) + 1
	return fmt.Sprintf("%02d/%02d", month, year%100)
}

// generateSortCode generates a UK sort code (XX-XX-XX format)
func (g *Generator) generateSortCode() string {
	return fmt.Sprintf("%02d-%02d-%02d",
		g.rand.Intn(100),
		g.rand.Intn(100),
		g.rand.Intn(100),
	)
}

// generateUTR generates a Unique Taxpayer Reference
func (g *Generator) generateUTR() string {
	return fmt.Sprintf("%010d", g.rand.Intn(10000000000))
}

// India-specific generators

// generateAadhaar generates an Indian Aadhaar number (XXXX XXXX XXXX format)
func (g *Generator) generateAadhaar() string {
	return fmt.Sprintf("%04d %04d %04d",
		g.rand.Intn(10000),
		g.rand.Intn(10000),
		g.rand.Intn(10000))
}

// generatePAN generates an Indian PAN number (ABCDE1234F format)
func (g *Generator) generatePAN(lastName string) string {
	// First 3 letters: random uppercase
	// 4th letter: Person type (P for individual)
	// 5th letter: First letter of last name
	// Next 4 digits: random numbers
	// Last letter: check digit (random for test data)
	firstLetter := byte('A' + g.rand.Intn(26))
	lastNameInitial := byte('A')
	if len(lastName) > 0 {
		lastNameInitial = byte(lastName[0])
		if lastNameInitial >= 'a' && lastNameInitial <= 'z' {
			lastNameInitial = lastNameInitial - 32 // Convert to uppercase
		}
	}
	
	return fmt.Sprintf("%c%c%c%c%c%04d%c",
		firstLetter,
		'A'+byte(g.rand.Intn(26)),
		'A'+byte(g.rand.Intn(26)),
		'P', // Person
		lastNameInitial,
		g.rand.Intn(10000),
		'A'+byte(g.rand.Intn(26)),
	)
}

// generateIndianPostcode generates an Indian PIN code (6 digits)
func (g *Generator) generateIndianPostcode() string {
	// First digit: 1-8 (regional distribution)
	// Remaining 5: random
	return fmt.Sprintf("%d%05d", g.rand.Intn(8)+1, g.rand.Intn(100000))
}

// generateIndianAddress generates a full Indian address
func (g *Generator) generateIndianAddress() string {
	houseNum := g.rand.Intn(500) + 1
	streetTypes := []string{"Road", "Street", "Lane", "Nagar", "Colony", "Enclave"}
	streetType := streetTypes[g.rand.Intn(len(streetTypes))]
	area := []string{"MG", "Nehru", "Gandhi", "Station", "Market", "Park", "Lake", "Main"}
	areaName := area[g.rand.Intn(len(area))]
	city := g.randomChoice(indianCities)
	state := g.randomChoice(indianStates)
	pincode := g.generateIndianPostcode()
	
	return fmt.Sprintf("%d, %s %s, %s, %s - %s", houseNum, areaName, streetType, city, state, pincode)
}

// generateIndianLandline generates an Indian landline number
func (g *Generator) generateIndianLandline() string {
	// Format: 0XX-XXXXXXXX (STD code + number)
	stdCode := g.rand.Intn(89) + 11 // 11 to 99
	number := g.rand.Intn(90000000) + 10000000
	return fmt.Sprintf("0%d-%d", stdCode, number)
}

// generateIndianMobile generates an Indian mobile number
func (g *Generator) generateIndianMobile() string {
	// Format: +91 XXXXX XXXXX (starts with 6,7,8,9)
	firstDigit := []int{6, 7, 8, 9}
	return fmt.Sprintf("+91 %d%04d %05d",
		firstDigit[g.rand.Intn(len(firstDigit))],
		g.rand.Intn(10000),
		g.rand.Intn(100000))
}

// generateIndianDrivingLicense generates an Indian driving license number
func (g *Generator) generateIndianDrivingLicense(state string) string {
	// Format: MH01 20160012345 (State code + RTO + year + number)
	stateCode := "MH" // Default to Maharashtra
	if len(state) >= 2 {
		stateCode = state[:2]
	}
	rto := fmt.Sprintf("%02d", g.rand.Intn(99)+1)
	year := 2010 + g.rand.Intn(15)
	number := fmt.Sprintf("%07d", g.rand.Intn(10000000))
	
	return fmt.Sprintf("%s%s %d%s", stateCode, rto, year, number)
}

// generateIndianBankAccount generates an Indian bank account number
func (g *Generator) generateIndianBankAccount() string {
	// Indian account numbers are typically 9-18 digits
	return fmt.Sprintf("%011d", g.rand.Intn(100000000000))
}

// generateIFSC generates an Indian IFSC code
func (g *Generator) generateIFSC() string {
	// Format: ABCD0123456 (4 letters + 0 + 6 digits)
	banks := []string{"SBIN", "HDFC", "ICIC", "AXIS", "PUNB", "UBIN", "BARB"}
	bank := banks[g.rand.Intn(len(banks))]
	branch := fmt.Sprintf("%06d", g.rand.Intn(1000000))
	
	return fmt.Sprintf("%s0%s", bank, branch)
}

// US-specific generators

// generateSSN generates a US Social Security Number (XXX-XX-XXXX format)
func (g *Generator) generateSSN() string {
	// Avoiding invalid SSN patterns
	area := g.rand.Intn(899) + 1    // 001-899 (avoiding 000, 666, 900-999)
	if area == 666 {
		area = 667
	}
	group := g.rand.Intn(99) + 1     // 01-99 (avoiding 00)
	serial := g.rand.Intn(9999) + 1  // 0001-9999 (avoiding 0000)
	
	return fmt.Sprintf("%03d-%02d-%04d", area, group, serial)
}

// generateGreenCard generates a US Green Card number
func (g *Generator) generateGreenCard() string {
	// Format: ABC1234567890 (3 letters + 10 digits)
	return fmt.Sprintf("%c%c%c%010d",
		'A'+byte(g.rand.Intn(26)),
		'A'+byte(g.rand.Intn(26)),
		'A'+byte(g.rand.Intn(26)),
		g.rand.Intn(10000000000),
	)
}

// generateUSZipCode generates a US ZIP code
func (g *Generator) generateUSZipCode() string {
	// Format: XXXXX or XXXXX-XXXX
	if g.rand.Intn(2) == 0 {
		return fmt.Sprintf("%05d", g.rand.Intn(100000))
	}
	return fmt.Sprintf("%05d-%04d", g.rand.Intn(100000), g.rand.Intn(10000))
}

// generateUSAddress generates a full US address
func (g *Generator) generateUSAddress() string {
	houseNum := g.rand.Intn(9999) + 1
	street := g.randomChoice(usStreetNames)
	city := g.randomChoice(usCities)
	state := g.randomChoice(usStates)
	zipcode := g.generateUSZipCode()
	
	return fmt.Sprintf("%d %s, %s, %s %s", houseNum, street, city, state, zipcode)
}

// generateUSPhone generates a US phone number
func (g *Generator) generateUSPhone() string {
	// Format: (XXX) XXX-XXXX
	// Avoiding area codes starting with 0 or 1
	areaCode := g.rand.Intn(800) + 200 // 200-999
	exchange := g.rand.Intn(800) + 200 // 200-999
	subscriber := g.rand.Intn(10000)
	
	return fmt.Sprintf("(%03d) %03d-%04d", areaCode, exchange, subscriber)
}

// generateUSMobile generates a US mobile number
func (g *Generator) generateUSMobile() string {
	// Same format as landline in the US
	return g.generateUSPhone()
}

// generateUSDrivingLicense generates a US driving license number
func (g *Generator) generateUSDrivingLicense(state string) string {
	// Format varies by state, using generic format
	// Example: A1234567 or 12345678
	if g.rand.Intn(2) == 0 {
		// Letter + numbers
		return fmt.Sprintf("%c%07d",
			'A'+byte(g.rand.Intn(26)),
			g.rand.Intn(10000000))
	}
	// Just numbers
	return fmt.Sprintf("%08d", g.rand.Intn(100000000))
}

// generateUSBankAccount generates a US bank account number
func (g *Generator) generateUSBankAccount() string {
	// US account numbers are typically 8-17 digits
	return fmt.Sprintf("%012d", g.rand.Intn(1000000000000))
}

// generateUSRoutingNumber generates a US bank routing number (ABA number)
func (g *Generator) generateUSRoutingNumber() string {
	// Format: 9 digits
	return fmt.Sprintf("%09d", g.rand.Intn(1000000000))
}

// Common helper methods

// generateIncome generates an annual income with currency symbol
func (g *Generator) generateIncome(currency string) string {
	income := (g.rand.Intn(150) + 20) * 1000 // 20k to 170k
	return fmt.Sprintf("%s%d", currency, income)
}

// generateDOB generates a date of birth
func (g *Generator) generateDOB() string {
	year := 1950 + g.rand.Intn(55) // 1950-2004
	month := g.rand.Intn(12) + 1
	day := g.rand.Intn(28) + 1
	return fmt.Sprintf("%02d/%02d/%04d", day, month, year)
}

// generateEmail generates an email address with country-specific domain
func (g *Generator) generateEmail(firstName, lastName, tld string) string {
	providers := []string{"gmail", "yahoo", "hotmail", "outlook"}
	provider := providers[g.rand.Intn(len(providers))]
	
	patterns := []string{
		fmt.Sprintf("%s.%s@%s.%s", firstName, lastName, provider, tld),
		fmt.Sprintf("%s%s@%s.%s", firstName, lastName, provider, tld),
		fmt.Sprintf("%s_%s@%s.%s", firstName, lastName, provider, tld),
		fmt.Sprintf("%s.%s%d@%s.%s", firstName, lastName, g.rand.Intn(100), provider, tld),
	}
	
	return patterns[g.rand.Intn(len(patterns))]
}

// randomChoice picks a random element from a slice
func (g *Generator) randomChoice(slice []string) string {
	return slice[g.rand.Intn(len(slice))]
}

// GenerateRecords generates multiple records for a specific country
func (g *Generator) GenerateRecords(count int, country Country) []*Record {
	records := make([]*Record, count)
	for i := 0; i < count; i++ {
		records[i] = g.GenerateRecord(country)
	}
	return records
}

