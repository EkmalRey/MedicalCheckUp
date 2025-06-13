package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Constants
const (
	NMAX             = 100
	PATIENT_ID_START = 20001
	PACKAGE_ID_START = 10001
	RECORD_ID_START  = 30001
	DATA_FILE        = "data.json"

	// ANSI color codes
	RESET  = "\033[0m"
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	BLUE   = "\033[34m"
	PURPLE = "\033[35m"
	CYAN   = "\033[36m"
	WHITE  = "\033[37m"
	BOLD   = "\033[1m"
)

// Data structures
type Patient struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

type Package struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

type Record struct {
	ID      int     `json:"id"`
	Patient Patient `json:"patient"`
	Package Package `json:"package"`
	Date    string  `json:"date"`
}

type PatientArray struct {
	Daftar [NMAX]Patient `json:"daftar"`
	N      int           `json:"n"`
}

type PackageArray struct {
	Daftar [NMAX]Package `json:"daftar"`
	N      int           `json:"n"`
}

type RecordArray struct {
	Daftar [NMAX]Record `json:"daftar"`
	N      int          `json:"n"`
}

type DataStore struct {
	Patients PatientArray `json:"patients"`
	Packages PackageArray `json:"packages"`
	Records  RecordArray  `json:"records"`
}

// Global variables
var (
	patients PatientArray
	packages PackageArray
	records  RecordArray
	scanner  = bufio.NewScanner(os.Stdin)
)

// Utility functions
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func pause() {
	fmt.Print("\nPress Enter to continue...")
	scanner.Scan()
}

func printHeader(title string) {
	clearScreen()
	fmt.Printf("%s%s=== %s ===%s\n\n", BOLD, CYAN, title, RESET)
}

func printSuccess(message string) {
	fmt.Printf("%s%s✓ %s%s\n", BOLD, GREEN, message, RESET)
}

func printError(message string) {
	fmt.Printf("%s%s✗ %s%s\n", BOLD, RED, message, RESET)
}

func printWarning(message string) {
	fmt.Printf("%s%s⚠ %s%s\n", BOLD, YELLOW, message, RESET)
}

// Input validation functions
func getValidInput(prompt string) string {
	for {
		fmt.Print(prompt)
		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())
			if input != "" {
				return input
			}
		}
		printError("Input cannot be empty. Please try again.")
	}
}

func getValidInt(prompt string, min, max int) int {
	for {
		input := getValidInput(prompt)
		if num, err := strconv.Atoi(input); err == nil {
			if num >= min && num <= max {
				return num
			}
			printError(fmt.Sprintf("Please enter a number between %d and %d.", min, max))
		} else {
			printError("Please enter a valid number.")
		}
	}
}

func getValidFloat(prompt string, min float64) float64 {
	for {
		input := getValidInput(prompt)
		if num, err := strconv.ParseFloat(input, 64); err == nil {
			if num >= min {
				return num
			}
			printError(fmt.Sprintf("Please enter a number greater than or equal to %.2f.", min))
		} else {
			printError("Please enter a valid number.")
		}
	}
}

func getValidGender() string {
	for {
		gender := strings.ToUpper(getValidInput("Enter gender (M/F): "))
		if gender == "M" || gender == "F" {
			return gender
		}
		printError("Please enter 'M' for Male or 'F' for Female.")
	}
}

func getValidCategory() string {
	categories := []string{"Basic", "Standard", "Premium", "Executive"}

	fmt.Println("Available categories:")
	for i, cat := range categories {
		fmt.Printf("%d. %s\n", i+1, cat)
	}

	choice := getValidInt("Select category (1-4): ", 1, 4)
	return categories[choice-1]
}

// Date validation functions
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func isValidDate(day, month, year int) bool {
	if year < 1900 || year > 2100 {
		return false
	}
	if month < 1 || month > 12 {
		return false
	}

	daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if isLeapYear(year) {
		daysInMonth[1] = 29
	}

	return day >= 1 && day <= daysInMonth[month-1]
}

func getValidDate(prompt string) string {
	for {
		fmt.Print(prompt + " (DD/MM/YYYY): ")
		if scanner.Scan() {
			dateStr := strings.TrimSpace(scanner.Text())
			parts := strings.Split(dateStr, "/")

			if len(parts) == 3 {
				day, err1 := strconv.Atoi(parts[0])
				month, err2 := strconv.Atoi(parts[1])
				year, err3 := strconv.Atoi(parts[2])

				if err1 == nil && err2 == nil && err3 == nil {
					if isValidDate(day, month, year) {
						return fmt.Sprintf("%02d/%02d/%04d", day, month, year)
					}
				}
			}
		}
		printError("Please enter a valid date in DD/MM/YYYY format.")
	}
}

// Data persistence functions
func saveData() error {
	data := DataStore{
		Patients: patients,
		Packages: packages,
		Records:  records,
	}

	file, err := os.Create(DATA_FILE)
	if err != nil {
		return fmt.Errorf("failed to create data file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to encode data: %v", err)
	}

	return nil
}

func loadData() error {
	file, err := os.Open(DATA_FILE)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, start with empty data
			return nil
		}
		return fmt.Errorf("failed to open data file: %v", err)
	}
	defer file.Close()

	var data DataStore
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return fmt.Errorf("failed to decode data: %v", err)
	}

	patients = data.Patients
	packages = data.Packages
	records = data.Records

	return nil
}

// Search functions
func sequentialSearchPatientByName(name string) int {
	searchName := strings.ToLower(name)
	for i := 0; i < patients.N; i++ {
		if strings.Contains(strings.ToLower(patients.Daftar[i].Name), searchName) {
			return i
		}
	}
	return -1
}

func sequentialSearchPackageByName(name string) int {
	searchName := strings.ToLower(name)
	for i := 0; i < packages.N; i++ {
		if strings.Contains(strings.ToLower(packages.Daftar[i].Name), searchName) {
			return i
		}
	}
	return -1
}

func binarySearchPatientByID(id int) int {
	left, right := 0, patients.N-1
	for left <= right {
		mid := (left + right) / 2
		if patients.Daftar[mid].ID == id {
			return mid
		} else if patients.Daftar[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func binarySearchPackageByID(id int) int {
	left, right := 0, packages.N-1
	for left <= right {
		mid := (left + right) / 2
		if packages.Daftar[mid].ID == id {
			return mid
		} else if packages.Daftar[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// Sorting functions
func selectionSortPatientsByName() {
	for i := 0; i < patients.N-1; i++ {
		minIdx := i
		for j := i + 1; j < patients.N; j++ {
			if strings.ToLower(patients.Daftar[j].Name) < strings.ToLower(patients.Daftar[minIdx].Name) {
				minIdx = j
			}
		}
		if minIdx != i {
			patients.Daftar[i], patients.Daftar[minIdx] = patients.Daftar[minIdx], patients.Daftar[i]
		}
	}
}

func selectionSortPackagesByPrice() {
	for i := 0; i < packages.N-1; i++ {
		minIdx := i
		for j := i + 1; j < packages.N; j++ {
			if packages.Daftar[j].Price < packages.Daftar[minIdx].Price {
				minIdx = j
			}
		}
		if minIdx != i {
			packages.Daftar[i], packages.Daftar[minIdx] = packages.Daftar[minIdx], packages.Daftar[i]
		}
	}
}

func insertionSortRecordsByDate() {
	for i := 1; i < records.N; i++ {
		key := records.Daftar[i]
		j := i - 1

		// Sort by date (newest first)
		for j >= 0 && records.Daftar[j].Date < key.Date {
			records.Daftar[j+1] = records.Daftar[j]
			j--
		}
		records.Daftar[j+1] = key
	}
}

// ID generation functions
func getNextPatientID() int {
	maxID := PATIENT_ID_START - 1
	for i := 0; i < patients.N; i++ {
		if patients.Daftar[i].ID > maxID {
			maxID = patients.Daftar[i].ID
		}
	}
	return maxID + 1
}

func getNextPackageID() int {
	maxID := PACKAGE_ID_START - 1
	for i := 0; i < packages.N; i++ {
		if packages.Daftar[i].ID > maxID {
			maxID = packages.Daftar[i].ID
		}
	}
	return maxID + 1
}

func getNextRecordID() int {
	maxID := RECORD_ID_START - 1
	for i := 0; i < records.N; i++ {
		if records.Daftar[i].ID > maxID {
			maxID = records.Daftar[i].ID
		}
	}
	return maxID + 1
}

// Patient management functions
func addPatient() {
	printHeader("Add New Patient")

	if patients.N >= NMAX {
		printError("Cannot add more patients. Maximum capacity reached.")
		pause()
		return
	}

	name := getValidInput("Enter patient name: ")
	gender := getValidGender()
	age := getValidInt("Enter age: ", 0, 150)

	patient := Patient{
		ID:     getNextPatientID(),
		Name:   name,
		Gender: gender,
		Age:    age,
	}

	patients.Daftar[patients.N] = patient
	patients.N++

	printSuccess(fmt.Sprintf("Patient added successfully with ID: %d", patient.ID))
	pause()
}

func displayPatients() {
	printHeader("Patient List")

	if patients.N == 0 {
		printWarning("No patients found.")
		pause()
		return
	}

	fmt.Printf("Sort by: 1. Name  2. Age  3. ID\n")
	choice := getValidInt("Choose sorting option: ", 1, 3)

	// Create a copy for sorting
	tempPatients := patients

	switch choice {
	case 1:
		selectionSortPatientsByName()
	case 2:
		// Sort by age using insertion sort
		for i := 1; i < tempPatients.N; i++ {
			key := tempPatients.Daftar[i]
			j := i - 1
			for j >= 0 && tempPatients.Daftar[j].Age > key.Age {
				tempPatients.Daftar[j+1] = tempPatients.Daftar[j]
				j--
			}
			tempPatients.Daftar[j+1] = key
		}
		patients = tempPatients
	case 3:
		// Sort by ID (already in order typically)
	}

	fmt.Printf("\n%s%-10s %-30s %-8s %-5s%s\n", BOLD, "ID", "Name", "Gender", "Age", RESET)
	fmt.Println(strings.Repeat("-", 60))

	for i := 0; i < patients.N; i++ {
		p := patients.Daftar[i]
		fmt.Printf("%-10d %-30s %-8s %-5d\n", p.ID, p.Name, p.Gender, p.Age)
	}

	pause()
}

func searchPatient() {
	printHeader("Search Patient")

	fmt.Println("Search by: 1. Name  2. ID")
	choice := getValidInt("Choose search method: ", 1, 2)

	var foundIdx int = -1

	switch choice {
	case 1:
		name := getValidInput("Enter patient name to search: ")
		foundIdx = sequentialSearchPatientByName(name)
	case 2:
		id := getValidInt("Enter patient ID to search: ", 1, 999999)
		foundIdx = binarySearchPatientByID(id)
	}

	if foundIdx != -1 {
		p := patients.Daftar[foundIdx]
		fmt.Printf("\n%sPatient Found:%s\n", GREEN, RESET)
		fmt.Printf("ID: %d\n", p.ID)
		fmt.Printf("Name: %s\n", p.Name)
		fmt.Printf("Gender: %s\n", p.Gender)
		fmt.Printf("Age: %d\n", p.Age)
	} else {
		printError("Patient not found.")
	}

	pause()
}

func updatePatient() {
	printHeader("Update Patient")

	id := getValidInt("Enter patient ID to update: ", 1, 999999)
	idx := binarySearchPatientByID(id)

	if idx == -1 {
		printError("Patient not found.")
		pause()
		return
	}

	p := &patients.Daftar[idx]
	fmt.Printf("\nCurrent patient data:\n")
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Gender: %s\n", p.Gender)
	fmt.Printf("Age: %d\n", p.Age)

	fmt.Println("\nWhat would you like to update?")
	fmt.Println("1. Name")
	fmt.Println("2. Gender")
	fmt.Println("3. Age")
	fmt.Println("4. All fields")

	choice := getValidInt("Choose option: ", 1, 4)

	switch choice {
	case 1:
		p.Name = getValidInput("Enter new name: ")
	case 2:
		p.Gender = getValidGender()
	case 3:
		p.Age = getValidInt("Enter new age: ", 0, 150)
	case 4:
		p.Name = getValidInput("Enter new name: ")
		p.Gender = getValidGender()
		p.Age = getValidInt("Enter new age: ", 0, 150)
	}

	printSuccess("Patient updated successfully.")
	pause()
}

func deletePatient() {
	printHeader("Delete Patient")

	id := getValidInt("Enter patient ID to delete: ", 1, 999999)
	idx := binarySearchPatientByID(id)

	if idx == -1 {
		printError("Patient not found.")
		pause()
		return
	}

	p := patients.Daftar[idx]
	fmt.Printf("\nPatient to be deleted:\n")
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Gender: %s\n", p.Gender)
	fmt.Printf("Age: %d\n", p.Age)

	confirm := getValidInput("\nAre you sure you want to delete this patient? (y/N): ")
	if strings.ToLower(confirm) != "y" && strings.ToLower(confirm) != "yes" {
		printWarning("Deletion cancelled.")
		pause()
		return
	}

	// Shift elements to remove the patient
	for i := idx; i < patients.N-1; i++ {
		patients.Daftar[i] = patients.Daftar[i+1]
	}
	patients.N--

	printSuccess("Patient deleted successfully.")
	pause()
}

// Package management functions
func addPackage() {
	printHeader("Add New Package")

	if packages.N >= NMAX {
		printError("Cannot add more packages. Maximum capacity reached.")
		pause()
		return
	}

	name := getValidInput("Enter package name: ")
	category := getValidCategory()
	price := getValidFloat("Enter price: ", 0.0)

	pkg := Package{
		ID:       getNextPackageID(),
		Name:     name,
		Category: category,
		Price:    price,
	}

	packages.Daftar[packages.N] = pkg
	packages.N++

	printSuccess(fmt.Sprintf("Package added successfully with ID: %d", pkg.ID))
	pause()
}

func displayPackages() {
	printHeader("Package List")

	if packages.N == 0 {
		printWarning("No packages found.")
		pause()
		return
	}

	fmt.Printf("Sort by: 1. Name  2. Price  3. Category\n")
	choice := getValidInt("Choose sorting option: ", 1, 3)

	switch choice {
	case 1:
		// Sort by name using selection sort
		for i := 0; i < packages.N-1; i++ {
			minIdx := i
			for j := i + 1; j < packages.N; j++ {
				if strings.ToLower(packages.Daftar[j].Name) < strings.ToLower(packages.Daftar[minIdx].Name) {
					minIdx = j
				}
			}
			if minIdx != i {
				packages.Daftar[i], packages.Daftar[minIdx] = packages.Daftar[minIdx], packages.Daftar[i]
			}
		}
	case 2:
		selectionSortPackagesByPrice()
	case 3:
		// Sort by category using insertion sort
		for i := 1; i < packages.N; i++ {
			key := packages.Daftar[i]
			j := i - 1
			for j >= 0 && strings.ToLower(packages.Daftar[j].Category) > strings.ToLower(key.Category) {
				packages.Daftar[j+1] = packages.Daftar[j]
				j--
			}
			packages.Daftar[j+1] = key
		}
	}

	fmt.Printf("\n%s%-10s %-30s %-15s %-12s%s\n", BOLD, "ID", "Name", "Category", "Price", RESET)
	fmt.Println(strings.Repeat("-", 70))

	for i := 0; i < packages.N; i++ {
		p := packages.Daftar[i]
		fmt.Printf("%-10d %-30s %-15s $%-11.2f\n", p.ID, p.Name, p.Category, p.Price)
	}

	pause()
}

func searchPackage() {
	printHeader("Search Package")

	fmt.Println("Search by: 1. Name  2. ID  3. Category")
	choice := getValidInt("Choose search method: ", 1, 3)

	var foundIdx int = -1

	switch choice {
	case 1:
		name := getValidInput("Enter package name to search: ")
		foundIdx = sequentialSearchPackageByName(name)
	case 2:
		id := getValidInt("Enter package ID to search: ", 1, 999999)
		foundIdx = binarySearchPackageByID(id)
	case 3:
		category := getValidCategory()
		for i := 0; i < packages.N; i++ {
			if packages.Daftar[i].Category == category {
				foundIdx = i
				break
			}
		}
	}

	if foundIdx != -1 {
		p := packages.Daftar[foundIdx]
		fmt.Printf("\n%sPackage Found:%s\n", GREEN, RESET)
		fmt.Printf("ID: %d\n", p.ID)
		fmt.Printf("Name: %s\n", p.Name)
		fmt.Printf("Category: %s\n", p.Category)
		fmt.Printf("Price: $%.2f\n", p.Price)
	} else {
		printError("Package not found.")
	}

	pause()
}

func updatePackage() {
	printHeader("Update Package")

	id := getValidInt("Enter package ID to update: ", 1, 999999)
	idx := binarySearchPackageByID(id)

	if idx == -1 {
		printError("Package not found.")
		pause()
		return
	}

	p := &packages.Daftar[idx]
	fmt.Printf("\nCurrent package data:\n")
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Category: %s\n", p.Category)
	fmt.Printf("Price: $%.2f\n", p.Price)

	fmt.Println("\nWhat would you like to update?")
	fmt.Println("1. Name")
	fmt.Println("2. Category")
	fmt.Println("3. Price")
	fmt.Println("4. All fields")

	choice := getValidInt("Choose option: ", 1, 4)

	switch choice {
	case 1:
		p.Name = getValidInput("Enter new name: ")
	case 2:
		p.Category = getValidCategory()
	case 3:
		p.Price = getValidFloat("Enter new price: ", 0.0)
	case 4:
		p.Name = getValidInput("Enter new name: ")
		p.Category = getValidCategory()
		p.Price = getValidFloat("Enter new price: ", 0.0)
	}

	printSuccess("Package updated successfully.")
	pause()
}

func deletePackage() {
	printHeader("Delete Package")

	id := getValidInt("Enter package ID to delete: ", 1, 999999)
	idx := binarySearchPackageByID(id)

	if idx == -1 {
		printError("Package not found.")
		pause()
		return
	}

	p := packages.Daftar[idx]
	fmt.Printf("\nPackage to be deleted:\n")
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Category: %s\n", p.Category)
	fmt.Printf("Price: $%.2f\n", p.Price)

	confirm := getValidInput("\nAre you sure you want to delete this package? (y/N): ")
	if strings.ToLower(confirm) != "y" && strings.ToLower(confirm) != "yes" {
		printWarning("Deletion cancelled.")
		pause()
		return
	}

	// Shift elements to remove the package
	for i := idx; i < packages.N-1; i++ {
		packages.Daftar[i] = packages.Daftar[i+1]
	}
	packages.N--

	printSuccess("Package deleted successfully.")
	pause()
}

// Record management functions
func addRecord() {
	printHeader("Add New Medical Record")

	if records.N >= NMAX {
		printError("Cannot add more records. Maximum capacity reached.")
		pause()
		return
	}

	if patients.N == 0 {
		printError("No patients available. Please add patients first.")
		pause()
		return
	}

	if packages.N == 0 {
		printError("No packages available. Please add packages first.")
		pause()
		return
	}

	// Select patient
	fmt.Println("Available patients:")
	for i := 0; i < patients.N; i++ {
		p := patients.Daftar[i]
		fmt.Printf("%d. %s (ID: %d)\n", i+1, p.Name, p.ID)
	}

	patientChoice := getValidInt("Select patient: ", 1, patients.N)
	selectedPatient := patients.Daftar[patientChoice-1]

	// Select package
	fmt.Println("\nAvailable packages:")
	for i := 0; i < packages.N; i++ {
		p := packages.Daftar[i]
		fmt.Printf("%d. %s - %s ($%.2f)\n", i+1, p.Name, p.Category, p.Price)
	}

	packageChoice := getValidInt("Select package: ", 1, packages.N)
	selectedPackage := packages.Daftar[packageChoice-1]

	// Get date
	date := getValidDate("Enter checkup date")

	record := Record{
		ID:      getNextRecordID(),
		Patient: selectedPatient,
		Package: selectedPackage,
		Date:    date,
	}

	records.Daftar[records.N] = record
	records.N++

	printSuccess(fmt.Sprintf("Medical record added successfully with ID: %d", record.ID))
	pause()
}

func displayRecords() {
	printHeader("Medical Records")

	if records.N == 0 {
		printWarning("No medical records found.")
		pause()
		return
	}

	fmt.Println("Sort by: 1. Date (Newest first)  2. Patient Name  3. Package Name")
	choice := getValidInt("Choose sorting option: ", 1, 3)

	switch choice {
	case 1:
		insertionSortRecordsByDate()
	case 2:
		// Sort by patient name using selection sort
		for i := 0; i < records.N-1; i++ {
			minIdx := i
			for j := i + 1; j < records.N; j++ {
				if strings.ToLower(records.Daftar[j].Patient.Name) < strings.ToLower(records.Daftar[minIdx].Patient.Name) {
					minIdx = j
				}
			}
			if minIdx != i {
				records.Daftar[i], records.Daftar[minIdx] = records.Daftar[minIdx], records.Daftar[i]
			}
		}
	case 3:
		// Sort by package name using insertion sort
		for i := 1; i < records.N; i++ {
			key := records.Daftar[i]
			j := i - 1
			for j >= 0 && strings.ToLower(records.Daftar[j].Package.Name) > strings.ToLower(key.Package.Name) {
				records.Daftar[j+1] = records.Daftar[j]
				j--
			}
			records.Daftar[j+1] = key
		}
	}

	fmt.Printf("\n%s%-8s %-20s %-20s %-15s %-12s%s\n", BOLD, "ID", "Patient", "Package", "Category", "Date", RESET)
	fmt.Println(strings.Repeat("-", 80))

	for i := 0; i < records.N; i++ {
		r := records.Daftar[i]
		fmt.Printf("%-8d %-20s %-20s %-15s %-12s\n",
			r.ID, r.Patient.Name, r.Package.Name, r.Package.Category, r.Date)
	}

	pause()
}

func searchRecords() {
	printHeader("Search Medical Records")

	fmt.Println("Search by: 1. Patient Name  2. Package Name  3. Date  4. Record ID")
	choice := getValidInt("Choose search method: ", 1, 4)

	var found bool = false

	switch choice {
	case 1:
		name := getValidInput("Enter patient name to search: ")
		searchName := strings.ToLower(name)
		fmt.Printf("\n%sRecords found for patient containing '%s':%s\n", GREEN, name, RESET)
		fmt.Printf("%-8s %-20s %-20s %-15s %-12s\n", "ID", "Patient", "Package", "Category", "Date")
		fmt.Println(strings.Repeat("-", 80))

		for i := 0; i < records.N; i++ {
			if strings.Contains(strings.ToLower(records.Daftar[i].Patient.Name), searchName) {
				r := records.Daftar[i]
				fmt.Printf("%-8d %-20s %-20s %-15s %-12s\n",
					r.ID, r.Patient.Name, r.Package.Name, r.Package.Category, r.Date)
				found = true
			}
		}

	case 2:
		name := getValidInput("Enter package name to search: ")
		searchName := strings.ToLower(name)
		fmt.Printf("\n%sRecords found for package containing '%s':%s\n", GREEN, name, RESET)
		fmt.Printf("%-8s %-20s %-20s %-15s %-12s\n", "ID", "Patient", "Package", "Category", "Date")
		fmt.Println(strings.Repeat("-", 80))

		for i := 0; i < records.N; i++ {
			if strings.Contains(strings.ToLower(records.Daftar[i].Package.Name), searchName) {
				r := records.Daftar[i]
				fmt.Printf("%-8d %-20s %-20s %-15s %-12s\n",
					r.ID, r.Patient.Name, r.Package.Name, r.Package.Category, r.Date)
				found = true
			}
		}

	case 3:
		date := getValidDate("Enter date to search")
		fmt.Printf("\n%sRecords found for date %s:%s\n", GREEN, date, RESET)
		fmt.Printf("%-8s %-20s %-20s %-15s %-12s\n", "ID", "Patient", "Package", "Category", "Date")
		fmt.Println(strings.Repeat("-", 80))

		for i := 0; i < records.N; i++ {
			if records.Daftar[i].Date == date {
				r := records.Daftar[i]
				fmt.Printf("%-8d %-20s %-20s %-15s %-12s\n",
					r.ID, r.Patient.Name, r.Package.Name, r.Package.Category, r.Date)
				found = true
			}
		}

	case 4:
		id := getValidInt("Enter record ID to search: ", 1, 999999)
		for i := 0; i < records.N; i++ {
			if records.Daftar[i].ID == id {
				r := records.Daftar[i]
				fmt.Printf("\n%sRecord Found:%s\n", GREEN, RESET)
				fmt.Printf("Record ID: %d\n", r.ID)
				fmt.Printf("Patient: %s (ID: %d)\n", r.Patient.Name, r.Patient.ID)
				fmt.Printf("Package: %s (%s)\n", r.Package.Name, r.Package.Category)
				fmt.Printf("Price: $%.2f\n", r.Package.Price)
				fmt.Printf("Date: %s\n", r.Date)
				found = true
				break
			}
		}
	}

	if !found {
		printError("No records found.")
	}

	pause()
}

func deleteRecord() {
	printHeader("Delete Medical Record")

	id := getValidInt("Enter record ID to delete: ", 1, 999999)
	var idx int = -1

	for i := 0; i < records.N; i++ {
		if records.Daftar[i].ID == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		printError("Record not found.")
		pause()
		return
	}

	r := records.Daftar[idx]
	fmt.Printf("\nRecord to be deleted:\n")
	fmt.Printf("Record ID: %d\n", r.ID)
	fmt.Printf("Patient: %s\n", r.Patient.Name)
	fmt.Printf("Package: %s\n", r.Package.Name)
	fmt.Printf("Date: %s\n", r.Date)

	confirm := getValidInput("\nAre you sure you want to delete this record? (y/N): ")
	if strings.ToLower(confirm) != "y" && strings.ToLower(confirm) != "yes" {
		printWarning("Deletion cancelled.")
		pause()
		return
	}

	// Shift elements to remove the record
	for i := idx; i < records.N-1; i++ {
		records.Daftar[i] = records.Daftar[i+1]
	}
	records.N--

	printSuccess("Record deleted successfully.")
	pause()
}

// Report functions
func generatePatientReport() {
	printHeader("Patient Statistics Report")

	if patients.N == 0 {
		printWarning("No patients available for report.")
		pause()
		return
	}

	var maleCount, femaleCount int
	var totalAge, minAge, maxAge int
	minAge = 999
	maxAge = 0

	for i := 0; i < patients.N; i++ {
		p := patients.Daftar[i]
		if p.Gender == "M" {
			maleCount++
		} else {
			femaleCount++
		}

		totalAge += p.Age
		if p.Age < minAge {
			minAge = p.Age
		}
		if p.Age > maxAge {
			maxAge = p.Age
		}
	}

	averageAge := float64(totalAge) / float64(patients.N)

	fmt.Printf("%sPatient Statistics:%s\n\n", BOLD, RESET)
	fmt.Printf("Total Patients: %d\n", patients.N)
	fmt.Printf("Male Patients: %d (%.1f%%)\n", maleCount, float64(maleCount)*100/float64(patients.N))
	fmt.Printf("Female Patients: %d (%.1f%%)\n", femaleCount, float64(femaleCount)*100/float64(patients.N))
	fmt.Printf("Average Age: %.1f years\n", averageAge)
	fmt.Printf("Youngest Patient: %d years\n", minAge)
	fmt.Printf("Oldest Patient: %d years\n", maxAge)

	pause()
}

func generatePackageReport() {
	printHeader("Package Statistics Report")

	if packages.N == 0 {
		printWarning("No packages available for report.")
		pause()
		return
	}

	categoryCount := make(map[string]int)
	var totalPrice, minPrice, maxPrice float64
	minPrice = 999999
	maxPrice = 0

	for i := 0; i < packages.N; i++ {
		p := packages.Daftar[i]
		categoryCount[p.Category]++

		totalPrice += p.Price
		if p.Price < minPrice {
			minPrice = p.Price
		}
		if p.Price > maxPrice {
			maxPrice = p.Price
		}
	}

	averagePrice := totalPrice / float64(packages.N)

	fmt.Printf("%sPackage Statistics:%s\n\n", BOLD, RESET)
	fmt.Printf("Total Packages: %d\n", packages.N)
	fmt.Printf("Average Price: $%.2f\n", averagePrice)
	fmt.Printf("Cheapest Package: $%.2f\n", minPrice)
	fmt.Printf("Most Expensive Package: $%.2f\n", maxPrice)

	fmt.Printf("\n%sPackages by Category:%s\n", BOLD, RESET)
	for category, count := range categoryCount {
		percentage := float64(count) * 100 / float64(packages.N)
		fmt.Printf("%s: %d (%.1f%%)\n", category, count, percentage)
	}

	pause()
}

func generateRevenueReport() {
	printHeader("Revenue Report")

	if records.N == 0 {
		printWarning("No records available for revenue calculation.")
		pause()
		return
	}

	var totalRevenue float64
	monthlyRevenue := make(map[string]float64)
	categoryRevenue := make(map[string]float64)

	for i := 0; i < records.N; i++ {
		r := records.Daftar[i]
		price := r.Package.Price
		totalRevenue += price

		// Extract month/year from date (DD/MM/YYYY)
		dateParts := strings.Split(r.Date, "/")
		if len(dateParts) == 3 {
			monthYear := dateParts[1] + "/" + dateParts[2]
			monthlyRevenue[monthYear] += price
		}

		categoryRevenue[r.Package.Category] += price
	}

	averagePerRecord := totalRevenue / float64(records.N)

	fmt.Printf("%sRevenue Statistics:%s\n\n", BOLD, RESET)
	fmt.Printf("Total Revenue: $%.2f\n", totalRevenue)
	fmt.Printf("Total Records: %d\n", records.N)
	fmt.Printf("Average Revenue per Record: $%.2f\n", averagePerRecord)

	fmt.Printf("\n%sRevenue by Month:%s\n", BOLD, RESET)
	for month, revenue := range monthlyRevenue {
		fmt.Printf("%s: $%.2f\n", month, revenue)
	}

	fmt.Printf("\n%sRevenue by Category:%s\n", BOLD, RESET)
	for category, revenue := range categoryRevenue {
		percentage := revenue * 100 / totalRevenue
		fmt.Printf("%s: $%.2f (%.1f%%)\n", category, revenue, percentage)
	}

	pause()
}

// Menu functions
func patientManagement() {
	for {
		printHeader("Patient Management")
		fmt.Printf("%s1.%s Add Patient\n", YELLOW, RESET)
		fmt.Printf("%s2.%s Display Patients\n", YELLOW, RESET)
		fmt.Printf("%s3.%s Search Patient\n", YELLOW, RESET)
		fmt.Printf("%s4.%s Update Patient\n", YELLOW, RESET)
		fmt.Printf("%s5.%s Delete Patient\n", YELLOW, RESET)
		fmt.Printf("%s0.%s Back to Main Menu\n", RED, RESET)

		choice := getValidInt("\nSelect option: ", 0, 5)

		switch choice {
		case 1:
			addPatient()
		case 2:
			displayPatients()
		case 3:
			searchPatient()
		case 4:
			updatePatient()
		case 5:
			deletePatient()
		case 0:
			return
		}
	}
}

func packageManagement() {
	for {
		printHeader("Package Management")
		fmt.Printf("%s1.%s Add Package\n", YELLOW, RESET)
		fmt.Printf("%s2.%s Display Packages\n", YELLOW, RESET)
		fmt.Printf("%s3.%s Search Package\n", YELLOW, RESET)
		fmt.Printf("%s4.%s Update Package\n", YELLOW, RESET)
		fmt.Printf("%s5.%s Delete Package\n", YELLOW, RESET)
		fmt.Printf("%s0.%s Back to Main Menu\n", RED, RESET)

		choice := getValidInt("\nSelect option: ", 0, 5)

		switch choice {
		case 1:
			addPackage()
		case 2:
			displayPackages()
		case 3:
			searchPackage()
		case 4:
			updatePackage()
		case 5:
			deletePackage()
		case 0:
			return
		}
	}
}

func recordManagement() {
	for {
		printHeader("Medical Record Management")
		fmt.Printf("%s1.%s Add Medical Record\n", YELLOW, RESET)
		fmt.Printf("%s2.%s Display Medical Records\n", YELLOW, RESET)
		fmt.Printf("%s3.%s Search Medical Records\n", YELLOW, RESET)
		fmt.Printf("%s4.%s Delete Medical Record\n", YELLOW, RESET)
		fmt.Printf("%s0.%s Back to Main Menu\n", RED, RESET)

		choice := getValidInt("\nSelect option: ", 0, 4)

		switch choice {
		case 1:
			addRecord()
		case 2:
			displayRecords()
		case 3:
			searchRecords()
		case 4:
			deleteRecord()
		case 0:
			return
		}
	}
}

func reportManagement() {
	for {
		printHeader("Reports & Analytics")
		fmt.Printf("%s1.%s Patient Statistics Report\n", YELLOW, RESET)
		fmt.Printf("%s2.%s Package Statistics Report\n", YELLOW, RESET)
		fmt.Printf("%s3.%s Revenue Report\n", YELLOW, RESET)
		fmt.Printf("%s0.%s Back to Main Menu\n", RED, RESET)

		choice := getValidInt("\nSelect option: ", 0, 3)

		switch choice {
		case 1:
			generatePatientReport()
		case 2:
			generatePackageReport()
		case 3:
			generateRevenueReport()
		case 0:
			return
		}
	}
}

func mainMenu() {
	for {
		printHeader("Medical Check-Up Management System")
		fmt.Printf("%s1.%s Patient Management\n", CYAN, RESET)
		fmt.Printf("%s2.%s Package Management\n", CYAN, RESET)
		fmt.Printf("%s3.%s Medical Record Management\n", CYAN, RESET)
		fmt.Printf("%s4.%s Reports & Analytics\n", CYAN, RESET)
		fmt.Printf("%s5.%s Save Data\n", GREEN, RESET)
		fmt.Printf("%s0.%s Exit\n", RED, RESET)

		choice := getValidInt("\nSelect option: ", 0, 5)

		switch choice {
		case 1:
			patientManagement()
		case 2:
			packageManagement()
		case 3:
			recordManagement()
		case 4:
			reportManagement()
		case 5:
			if err := saveData(); err != nil {
				printError(fmt.Sprintf("Failed to save data: %v", err))
			} else {
				printSuccess("Data saved successfully!")
			}
			pause()
		case 0:
			fmt.Printf("\n%sSaving data before exit...%s\n", YELLOW, RESET)
			if err := saveData(); err != nil {
				printError(fmt.Sprintf("Failed to save data: %v", err))
			} else {
				printSuccess("Data saved successfully!")
			}
			fmt.Printf("\n%sThank you for using Medical Check-Up Management System!%s\n", GREEN, RESET)
			return
		}
	}
}

func main() {
	// Initialize scanner
	scanner = bufio.NewScanner(os.Stdin)

	// Load existing data
	fmt.Printf("%sLoading data...%s\n", YELLOW, RESET)
	if err := loadData(); err != nil {
		printError(fmt.Sprintf("Failed to load data: %v", err))
		fmt.Println("Starting with empty data...")
	} else {
		printSuccess(fmt.Sprintf("Data loaded successfully! Patients: %d, Packages: %d, Records: %d",
			patients.N, packages.N, records.N))
	}

	time.Sleep(2 * time.Second)

	// Start main menu
	mainMenu()
}
