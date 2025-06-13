# 🏥 Medical Check-Up Management System

A **CLI-based medical management application** built in **Go** for managing patient data, medical packages, and health records. This project demonstrates fundamental programming concepts and data structure implementation.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go&logoColor=white)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-lightgrey)
![Status](https://img.shields.io/badge/Status-Educational%20Project-blue)

## 🎯 **What This Project Shows**

This is my Final Project Of The Course Programming Algorithm project that demonstrates my understanding of:
- **Data structures** (arrays, structs)
- **Search algorithms** (binary search, sequential search)
- **Sorting algorithms** (selection sort, insertion sort)
- **File I/O and JSON handling**
- **User interface design** for CLI applications
- **Input validation and error handling**

### **The Problem**
Create a simple system to manage medical check-up data including patients, medical packages, and records.

### **My Solution**
A menu-driven CLI program that handles basic CRUD operations with data persistence and some nice quality-of-life features.

## ⚡ **Features Implemented**

### 👥 **Patient Management**
- Add, view, edit, and delete patients
- Search by name or ID
- Sort by name, age, or ID
- Basic input validation

### 📦 **Package Management**
- Manage medical packages with categories and pricing
- Search and filter options
- CRUD operations with validation

### 📋 **Medical Records**
- Link patients with packages and dates
- Date validation (including leap years)
- Search by various criteria
- Basic record management

### 📊 **Simple Reports**
- Patient statistics (age, gender distribution)
- Package analytics
- Revenue tracking
- Basic data summaries

## 🛠️ **Technical Stuff**

### **What I Used**
- **Language**: Go (learning project)
- **Data Storage**: JSON files
- **Data Structures**: Arrays and structs
- **Algorithms**: Binary search, selection sort, insertion sort
- **Interface**: CLI with colored output

### **Programming Concepts Applied**
- ✅ **Struct definitions** and data organization
- ✅ **Function design** and modular programming
- ✅ **Input validation** and error handling
- ✅ **File I/O** with JSON serialization
- ✅ **Search algorithms** (O(log n) binary search)
- ✅ **Sorting algorithms** (selection and insertion sort)

### **Code Example**
```go
// Binary search implementation for patient lookup
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
```

## 🚀 **How to Run It**

### **You'll Need**
- Go installed on your computer ([Get it here](https://golang.org/dl/))
- Any terminal/command prompt

### **Running the Program**
```bash
# Clone or download the project
cd MedicalCheckUp

# Run the main program
go run main.go

# Or build an executable
go build -o medical.exe main.go
./medical.exe
```

### **What You Can Do**
1. Start the program and follow the menus
2. Add some patients and packages first
3. Create medical records linking them together
4. Try the search and sorting features
5. Check out the reports section
6. Data saves automatically when you exit

## 📁 **Project Structure**

```
MedicalCheckUp/
├── 📄 main.go                     # Main program (optimized version)
├── 📄 go.mod                      # Go module file
├── 📁 Archive/
│   ├── 📄 main_old.go             # Original version (with color dependency)
│   └── 📄 MedicalCheckUp.exe      # Pre-built executable
├── 📁 tests/
│   ├── 📄 unit_test.go            # Comprehensive test suite
│   ├── 📄 run_tests.bat           # Test runner script
│   ├── 📄 coverage.out            # Test coverage data (generated)
│   └── 📄 test_results.txt        # Test results (generated)
├── 📄 README.md                   # This file
└── 📄 data.json                   # Data storage (created automatically)
```

## 🎓 **What This Demonstrates**

### **Programming Skills**
- **Go Language**: Syntax, data types, functions, structs
- **Algorithms**: Search and sorting algorithm implementation
- **Data Structures**: Working with arrays and structured data
- **File Handling**: JSON reading/writing and data persistence
- **User Interface**: CLI design with menus and colored output

### **Problem-Solving Approach**
- Breaking down a complex problem into manageable parts
- Designing data structures to represent real-world entities
- Implementing user-friendly interfaces for data manipulation
- Adding validation and error handling for robustness

### **Code Organization**
- Modular function design
- Proper variable naming and code structure
- Input validation and error handling
- Documentation and comments

## � **For Recruiters**

This project shows I can:
- **Write clean, functional Go code** with proper structure
- **Implement basic algorithms** (binary search, sorting)
- **Handle user input and validation** effectively
- **Work with JSON data** and file persistence
- **Create user-friendly interfaces** even for CLI applications
- **Organize code** into logical, maintainable functions
- **Add testing** and documentation to projects

It's a solid demonstration of fundamental programming skills and problem-solving approach using Go.

## 🧪 **Testing**

The project includes a comprehensive test suite in the `tests/` folder. You can run it to see the code quality and functionality:

```bash
# Navigate to tests folder and run the test suite
cd tests
./run_tests.bat

# Or run tests manually
go test -v ./tests/

# Generate test coverage report
go test -cover -coverprofile=tests/coverage.out ./tests/
go tool cover -html=tests/coverage.out
```

The test suite covers:
- Date validation functions
- ID generation algorithms  
- Search functions (sequential and binary)
- Sorting algorithms (selection and insertion)
- Data persistence (save/load)
- Edge cases and boundary conditions
- Performance benchmarks

---

**Built with Go** • **Educational Project** • **Demonstrates Programming Fundamentals**
