package main

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

// Import the main package functions by copying the main file content
// Since we can't import from parent directory easily, we'll test the copied functions

// Constants for testing
const (
	NMAX             = 100
	PATIENT_ID_START = 20001
	PACKAGE_ID_START = 10001
	RECORD_ID_START  = 30001
	DATA_FILE        = "test_data.json"
)

// Test data structures
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

// Global test variables
var (
	patients PatientArray
	packages PackageArray
	records  RecordArray
)

// Copy of utility functions from main.go for testing
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

		for j >= 0 && records.Daftar[j].Date < key.Date {
			records.Daftar[j+1] = records.Daftar[j]
			j--
		}
		records.Daftar[j+1] = key
	}
}

func saveData() error {
	data := DataStore{
		Patients: patients,
		Packages: packages,
		Records:  records,
	}

	file, err := os.Create(DATA_FILE)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func loadData() error {
	file, err := os.Open(DATA_FILE)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	var data DataStore
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	patients = data.Patients
	packages = data.Packages
	records = data.Records
	return nil
}

// COMPREHENSIVE UNIT TESTS

// Test date validation functions
func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		year     int
		expected bool
	}{
		{2000, true},  // Divisible by 400
		{2004, true},  // Divisible by 4, not by 100
		{1900, false}, // Divisible by 100, not by 400
		{2001, false}, // Not divisible by 4
		{2020, true},  // Divisible by 4, not by 100
		{1600, true},  // Divisible by 400
		{1700, false}, // Divisible by 100, not by 400
	}

	for _, test := range tests {
		result := isLeapYear(test.year)
		if result != test.expected {
			t.Errorf("isLeapYear(%d) = %v; want %v", test.year, result, test.expected)
		}
	}
}

func TestIsValidDate(t *testing.T) {
	tests := []struct {
		day      int
		month    int
		year     int
		expected bool
	}{
		{29, 2, 2020, true},   // Leap year February 29
		{29, 2, 2021, false},  // Non-leap year February 29
		{31, 1, 2020, true},   // Valid January 31
		{31, 4, 2020, false},  // Invalid April 31
		{1, 1, 2020, true},    // Valid date
		{0, 1, 2020, false},   // Invalid day 0
		{32, 1, 2020, false},  // Invalid day 32
		{15, 0, 2020, false},  // Invalid month 0
		{15, 13, 2020, false}, // Invalid month 13
		{15, 6, 1899, false},  // Year too early
		{15, 6, 2101, false},  // Year too late
		{28, 2, 2021, true},   // Valid February 28 non-leap year
		{30, 11, 2020, true},  // Valid November 30
		{31, 11, 2020, false}, // Invalid November 31
	}

	for _, test := range tests {
		result := isValidDate(test.day, test.month, test.year)
		if result != test.expected {
			t.Errorf("isValidDate(%d, %d, %d) = %v; want %v",
				test.day, test.month, test.year, result, test.expected)
		}
	}
}

// Test ID generation functions
func TestGetNextPatientID(t *testing.T) {
	originalPatients := patients

	patients = PatientArray{N: 0}
	expected := PATIENT_ID_START
	result := getNextPatientID()
	if result != expected {
		t.Errorf("getNextPatientID() with empty array = %d; want %d", result, expected)
	}

	patients.Daftar[0] = Patient{ID: 20005}
	patients.Daftar[1] = Patient{ID: 20003}
	patients.Daftar[2] = Patient{ID: 20010}
	patients.N = 3
	expected = 20011
	result = getNextPatientID()
	if result != expected {
		t.Errorf("getNextPatientID() with existing patients = %d; want %d", result, expected)
	}

	patients = originalPatients
}

func TestGetNextPackageID(t *testing.T) {
	originalPackages := packages

	packages = PackageArray{N: 0}
	expected := PACKAGE_ID_START
	result := getNextPackageID()
	if result != expected {
		t.Errorf("getNextPackageID() with empty array = %d; want %d", result, expected)
	}

	packages.Daftar[0] = Package{ID: 10005}
	packages.Daftar[1] = Package{ID: 10003}
	packages.Daftar[2] = Package{ID: 10010}
	packages.N = 3
	expected = 10011
	result = getNextPackageID()
	if result != expected {
		t.Errorf("getNextPackageID() with existing packages = %d; want %d", result, expected)
	}

	packages = originalPackages
}

func TestGetNextRecordID(t *testing.T) {
	originalRecords := records

	records = RecordArray{N: 0}
	expected := RECORD_ID_START
	result := getNextRecordID()
	if result != expected {
		t.Errorf("getNextRecordID() with empty array = %d; want %d", result, expected)
	}

	records.Daftar[0] = Record{ID: 30005}
	records.Daftar[1] = Record{ID: 30003}
	records.Daftar[2] = Record{ID: 30010}
	records.N = 3
	expected = 30011
	result = getNextRecordID()
	if result != expected {
		t.Errorf("getNextRecordID() with existing records = %d; want %d", result, expected)
	}

	records = originalRecords
}

// Test search functions
func TestSequentialSearchPatientByName(t *testing.T) {
	originalPatients := patients

	patients = PatientArray{
		Daftar: [NMAX]Patient{
			{ID: 1, Name: "John Doe", Gender: "M", Age: 30},
			{ID: 2, Name: "Jane Smith", Gender: "F", Age: 25},
			{ID: 3, Name: "Bob Johnson", Gender: "M", Age: 35},
		},
		N: 3,
	}

	tests := []struct {
		searchName string
		expected   int
	}{
		{"John", 0},         // Partial match
		{"jane", 1},         // Case insensitive
		{"johnson", 2},      // Partial match
		{"Doe", 0},          // Partial match
		{"NonExistent", -1}, // No match
		{"", 0},             // Empty string matches first
	}

	for _, test := range tests {
		result := sequentialSearchPatientByName(test.searchName)
		if result != test.expected {
			t.Errorf("sequentialSearchPatientByName(%s) = %d; want %d",
				test.searchName, result, test.expected)
		}
	}

	patients = originalPatients
}

func TestSequentialSearchPackageByName(t *testing.T) {
	originalPackages := packages

	packages = PackageArray{
		Daftar: [NMAX]Package{
			{ID: 1, Name: "Basic Health Check", Category: "Basic", Price: 100.0},
			{ID: 2, Name: "Complete Blood Test", Category: "Standard", Price: 150.0},
			{ID: 3, Name: "Premium Package", Category: "Premium", Price: 300.0},
		},
		N: 3,
	}

	tests := []struct {
		searchName string
		expected   int
	}{
		{"Basic", 0},        // Partial match
		{"blood", 1},        // Case insensitive
		{"premium", 2},      // Case insensitive
		{"health", 0},       // Partial match
		{"NonExistent", -1}, // No match
	}

	for _, test := range tests {
		result := sequentialSearchPackageByName(test.searchName)
		if result != test.expected {
			t.Errorf("sequentialSearchPackageByName(%s) = %d; want %d",
				test.searchName, result, test.expected)
		}
	}

	packages = originalPackages
}

func TestBinarySearchPatientByID(t *testing.T) {
	originalPatients := patients

	patients = PatientArray{
		Daftar: [NMAX]Patient{
			{ID: 20001, Name: "John Doe", Gender: "M", Age: 30},
			{ID: 20002, Name: "Jane Smith", Gender: "F", Age: 25},
			{ID: 20005, Name: "Bob Johnson", Gender: "M", Age: 35},
		},
		N: 3,
	}

	tests := []struct {
		searchID int
		expected int
	}{
		{20001, 0},  // First element
		{20002, 1},  // Middle element
		{20005, 2},  // Last element
		{20003, -1}, // Not found
		{19999, -1}, // Too small
		{30000, -1}, // Too large
	}

	for _, test := range tests {
		result := binarySearchPatientByID(test.searchID)
		if result != test.expected {
			t.Errorf("binarySearchPatientByID(%d) = %d; want %d",
				test.searchID, result, test.expected)
		}
	}

	patients = originalPatients
}

func TestBinarySearchPackageByID(t *testing.T) {
	originalPackages := packages

	packages = PackageArray{
		Daftar: [NMAX]Package{
			{ID: 10001, Name: "Basic Health Check", Category: "Basic", Price: 100.0},
			{ID: 10002, Name: "Complete Blood Test", Category: "Standard", Price: 150.0},
			{ID: 10005, Name: "Premium Package", Category: "Premium", Price: 300.0},
		},
		N: 3,
	}

	tests := []struct {
		searchID int
		expected int
	}{
		{10001, 0},  // First element
		{10002, 1},  // Middle element
		{10005, 2},  // Last element
		{10003, -1}, // Not found
		{9999, -1},  // Too small
		{20000, -1}, // Too large
	}

	for _, test := range tests {
		result := binarySearchPackageByID(test.searchID)
		if result != test.expected {
			t.Errorf("binarySearchPackageByID(%d) = %d; want %d",
				test.searchID, result, test.expected)
		}
	}

	packages = originalPackages
}

// Test sorting functions
func TestSelectionSortPatientsByName(t *testing.T) {
	originalPatients := patients

	patients = PatientArray{
		Daftar: [NMAX]Patient{
			{ID: 1, Name: "Charlie", Gender: "M", Age: 30},
			{ID: 2, Name: "Alice", Gender: "F", Age: 25},
			{ID: 3, Name: "Bob", Gender: "M", Age: 35},
		},
		N: 3,
	}

	expected := []string{"Alice", "Bob", "Charlie"}
	selectionSortPatientsByName()

	for i := 0; i < patients.N; i++ {
		if patients.Daftar[i].Name != expected[i] {
			t.Errorf("After sorting, patient at index %d = %s; want %s",
				i, patients.Daftar[i].Name, expected[i])
		}
	}

	patients = originalPatients
}

func TestSelectionSortPackagesByPrice(t *testing.T) {
	originalPackages := packages

	packages = PackageArray{
		Daftar: [NMAX]Package{
			{ID: 1, Name: "Premium", Category: "Premium", Price: 300.0},
			{ID: 2, Name: "Basic", Category: "Basic", Price: 100.0},
			{ID: 3, Name: "Standard", Category: "Standard", Price: 200.0},
		},
		N: 3,
	}

	expected := []float64{100.0, 200.0, 300.0}
	selectionSortPackagesByPrice()

	for i := 0; i < packages.N; i++ {
		if packages.Daftar[i].Price != expected[i] {
			t.Errorf("After sorting, package at index %d price = %.2f; want %.2f",
				i, packages.Daftar[i].Price, expected[i])
		}
	}

	packages = originalPackages
}

func TestInsertionSortRecordsByDate(t *testing.T) {
	originalRecords := records

	records = RecordArray{
		Daftar: [NMAX]Record{
			{ID: 1, Date: "15/06/2023"},
			{ID: 2, Date: "20/06/2023"},
			{ID: 3, Date: "10/06/2023"},
		},
		N: 3,
	}

	expected := []string{"20/06/2023", "15/06/2023", "10/06/2023"}
	insertionSortRecordsByDate()

	for i := 0; i < records.N; i++ {
		if records.Daftar[i].Date != expected[i] {
			t.Errorf("After sorting, record at index %d date = %s; want %s",
				i, records.Daftar[i].Date, expected[i])
		}
	}

	records = originalRecords
}

// Test data persistence functions
func TestSaveAndLoadData(t *testing.T) {
	originalPatients := patients
	originalPackages := packages
	originalRecords := records

	testPatients := PatientArray{
		Daftar: [NMAX]Patient{
			{ID: 20001, Name: "Test Patient", Gender: "M", Age: 30},
		},
		N: 1,
	}

	testPackages := PackageArray{
		Daftar: [NMAX]Package{
			{ID: 10001, Name: "Test Package", Category: "Basic", Price: 100.0},
		},
		N: 1,
	}

	testRecords := RecordArray{
		Daftar: [NMAX]Record{
			{ID: 30001, Patient: testPatients.Daftar[0], Package: testPackages.Daftar[0], Date: "15/06/2023"},
		},
		N: 1,
	}

	patients = testPatients
	packages = testPackages
	records = testRecords

	err := saveData()
	if err != nil {
		t.Fatalf("saveData() failed: %v", err)
	}

	if _, err := os.Stat(DATA_FILE); os.IsNotExist(err) {
		t.Fatal("Data file was not created")
	}

	patients = PatientArray{N: 0}
	packages = PackageArray{N: 0}
	records = RecordArray{N: 0}

	err = loadData()
	if err != nil {
		t.Fatalf("loadData() failed: %v", err)
	}

	if patients.N != 1 || patients.Daftar[0].Name != "Test Patient" {
		t.Error("Patient data not loaded correctly")
	}

	if packages.N != 1 || packages.Daftar[0].Name != "Test Package" {
		t.Error("Package data not loaded correctly")
	}

	if records.N != 1 || records.Daftar[0].Date != "15/06/2023" {
		t.Error("Record data not loaded correctly")
	}

	os.Remove(DATA_FILE)
	patients = originalPatients
	packages = originalPackages
	records = originalRecords
}

// Test edge cases and boundary conditions
func TestEmptyArrayOperations(t *testing.T) {
	originalPatients := patients
	originalPackages := packages
	originalRecords := records

	patients = PatientArray{N: 0}
	packages = PackageArray{N: 0}
	records = RecordArray{N: 0}

	// Test searches on empty arrays
	if result := sequentialSearchPatientByName("any"); result != -1 {
		t.Errorf("sequentialSearchPatientByName on empty array = %d; want -1", result)
	}

	if result := binarySearchPatientByID(20001); result != -1 {
		t.Errorf("binarySearchPatientByID on empty array = %d; want -1", result)
	}

	if result := sequentialSearchPackageByName("any"); result != -1 {
		t.Errorf("sequentialSearchPackageByName on empty array = %d; want -1", result)
	}

	if result := binarySearchPackageByID(10001); result != -1 {
		t.Errorf("binarySearchPackageByID on empty array = %d; want -1", result)
	}

	// Test sorting on empty arrays (should not crash)
	selectionSortPatientsByName()
	selectionSortPackagesByPrice()
	insertionSortRecordsByDate()

	patients = originalPatients
	packages = originalPackages
	records = originalRecords
}

// Test integration workflow
func TestCompleteWorkflow(t *testing.T) {
	originalPatients := patients
	originalPackages := packages
	originalRecords := records

	// Start fresh
	patients = PatientArray{N: 0}
	packages = PackageArray{N: 0}
	records = RecordArray{N: 0}

	// Add patients
	patientData := []Patient{
		{ID: getNextPatientID(), Name: "Alice Johnson", Gender: "F", Age: 25},
		{ID: getNextPatientID(), Name: "Bob Smith", Gender: "M", Age: 30},
		{ID: getNextPatientID(), Name: "Charlie Brown", Gender: "M", Age: 35},
	}

	for _, patient := range patientData {
		patients.Daftar[patients.N] = patient
		patients.N++
	}

	// Add packages
	packageData := []Package{
		{ID: getNextPackageID(), Name: "Basic Health Check", Category: "Basic", Price: 100.0},
		{ID: getNextPackageID(), Name: "Standard Health Check", Category: "Standard", Price: 200.0},
		{ID: getNextPackageID(), Name: "Premium Health Check", Category: "Premium", Price: 300.0},
	}

	for _, pkg := range packageData {
		packages.Daftar[packages.N] = pkg
		packages.N++
	}

	// Create records
	recordData := []Record{
		{ID: getNextRecordID(), Patient: patients.Daftar[0], Package: packages.Daftar[0], Date: "10/06/2023"},
		{ID: getNextRecordID(), Patient: patients.Daftar[1], Package: packages.Daftar[1], Date: "15/06/2023"},
		{ID: getNextRecordID(), Patient: patients.Daftar[2], Package: packages.Daftar[2], Date: "20/06/2023"},
	}

	for _, record := range recordData {
		records.Daftar[records.N] = record
		records.N++
	}

	// Test searches work correctly
	aliceIndex := sequentialSearchPatientByName("Alice")
	if aliceIndex == -1 {
		t.Error("Could not find Alice")
	}

	basicIndex := sequentialSearchPackageByName("Basic")
	if basicIndex == -1 {
		t.Error("Could not find Basic package")
	}

	// Test sorting
	selectionSortPatientsByName()
	selectionSortPackagesByPrice()
	insertionSortRecordsByDate()

	// Verify patients are sorted by name
	if patients.Daftar[0].Name != "Alice Johnson" {
		t.Errorf("First patient after name sort = %s; want Alice Johnson", patients.Daftar[0].Name)
	}

	// Verify packages are sorted by price
	if packages.Daftar[0].Price != 100.0 {
		t.Errorf("First package after price sort = %.2f; want 100.00", packages.Daftar[0].Price)
	}

	// Verify records are sorted by date (newest first)
	if records.Daftar[0].Date != "20/06/2023" {
		t.Errorf("First record after date sort = %s; want 20/06/2023", records.Daftar[0].Date)
	}

	// Test data persistence
	err := saveData()
	if err != nil {
		t.Fatalf("Failed to save complete system data: %v", err)
	}

	// Clear all data
	patients = PatientArray{N: 0}
	packages = PackageArray{N: 0}
	records = RecordArray{N: 0}

	// Load data back
	err = loadData()
	if err != nil {
		t.Fatalf("Failed to load complete system data: %v", err)
	}

	// Verify everything is back
	if patients.N != 3 || packages.N != 3 || records.N != 3 {
		t.Errorf("Data counts after load: patients=%d, packages=%d, records=%d; want 3,3,3",
			patients.N, packages.N, records.N)
	}

	// Cleanup
	os.Remove(DATA_FILE)

	// Restore original state
	patients = originalPatients
	packages = originalPackages
	records = originalRecords
}

// Benchmark tests
func BenchmarkSequentialSearchPatient(b *testing.B) {
	originalPatients := patients

	// Setup large dataset
	patients = PatientArray{N: 1000}
	for i := 0; i < 1000; i++ {
		patients.Daftar[i] = Patient{
			ID:     PATIENT_ID_START + i,
			Name:   "Patient " + string(rune(65+i%26)),
			Gender: "M",
			Age:    20 + i%50,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sequentialSearchPatientByName("Patient Z")
	}

	patients = originalPatients
}

func BenchmarkBinarySearchPatient(b *testing.B) {
	originalPatients := patients

	// Setup large sorted dataset
	patients = PatientArray{N: 1000}
	for i := 0; i < 1000; i++ {
		patients.Daftar[i] = Patient{
			ID:     PATIENT_ID_START + i,
			Name:   "Patient",
			Gender: "M",
			Age:    20,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		binarySearchPatientByID(PATIENT_ID_START + 999)
	}

	patients = originalPatients
}

func BenchmarkSelectionSort(b *testing.B) {
	originalPatients := patients

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		patients = PatientArray{N: 100}
		for j := 0; j < 100; j++ {
			patients.Daftar[j] = Patient{
				ID:     PATIENT_ID_START + j,
				Name:   "Patient " + string(rune(90-j%26)),
				Gender: "M",
				Age:    20,
			}
		}
		b.StartTimer()
		selectionSortPatientsByName()
	}

	patients = originalPatients
}
