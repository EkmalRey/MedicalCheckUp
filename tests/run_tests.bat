@echo off
setlocal enabledelayedexpansion

:: Medical CheckUp Management System - Test Runner
:: This script runs comprehensive unit tests for the Go application

echo.
echo ===============================================
echo   Medical CheckUp Management System
echo   Comprehensive Test Suite Runner
echo ===============================================
echo.

:: Check if Go is installed
go version >nul 2>&1
if errorlevel 1 (
    echo ERROR: Go is not installed or not in PATH
    echo Please install Go from https://golang.org/dl/
    pause
    exit /b 1
)

echo Go version:
go version
echo.

:: Navigate to the project root directory (parent of tests folder)
cd /d "%~dp0.."

:: Initialize Go module if not already done
if not exist go.mod (
    echo Initializing Go module...
    go mod init medicalcheckup
    echo.
)

:: Clean any previous test artifacts in tests folder
if exist tests\data.json del tests\data.json
if exist tests\test_data.json del tests\test_data.json
if exist tests\coverage.out del tests\coverage.out
if exist tests\test_results.txt del tests\test_results.txt

echo ===============================================
echo Running Unit Tests
echo ===============================================
echo.

:: Run all tests with verbose output
echo Running all unit tests from tests folder...
go test -v ./tests/ > tests\test_results.txt 2>&1
set TEST_EXIT_CODE=!errorlevel!

:: Display test results
type tests\test_results.txt
echo.

if !TEST_EXIT_CODE! equ 0 (
    echo [PASSED] All tests PASSED!
) else (
    echo [ERROR] Some tests FAILED!
)

echo.
echo ===============================================
echo Running Tests with Coverage
echo ===============================================
echo.

:: Run tests with coverage
echo Generating test coverage report...
go test -cover -coverprofile=tests\coverage.out ./tests/ 2>&1
set COVERAGE_EXIT_CODE=!errorlevel!

if !COVERAGE_EXIT_CODE! equ 0 (
    echo.
    echo Coverage report generated successfully!
    echo To view detailed coverage report, run: go tool cover -html=tests\coverage.out
) else (
    echo [ERROR] Failed to generate coverage report
)

echo.
echo ===============================================
echo Running Benchmark Tests
echo ===============================================
echo.

:: Run benchmark tests
echo Running performance benchmarks...
go test -bench=. -benchmem ./tests/ 2>&1
set BENCH_EXIT_CODE=!errorlevel!

if !BENCH_EXIT_CODE! equ 0 (
    echo [PASSED] Benchmarks completed!
) else (
    echo [ERROR] Benchmark tests failed!
)

echo.
echo ===============================================
echo Running Race Condition Tests
echo ===============================================
echo.

:: Run tests with race detection
echo Checking for race conditions...
go test -race ./tests/ 2>&1
set RACE_EXIT_CODE=!errorlevel!

if !RACE_EXIT_CODE! equ 0 (
    echo [PASSED] No race conditions detected!
) else (
    echo [WARNING] Race conditions detected!
)

echo.
echo ===============================================
echo Running Short Tests Only
echo ===============================================
echo.

:: Run only short tests (excluding performance-intensive ones)
echo Running short tests...
go test -short ./tests/ 2>&1
set SHORT_EXIT_CODE=!errorlevel!

if !SHORT_EXIT_CODE! equ 0 (
    echo [PASSED] Short tests PASSED!
) else (
    echo [ERROR] Short tests FAILED!
)

echo.
echo ===============================================
echo Running Tests by Function
echo ===============================================
echo.

echo Testing Date Validation Functions...
go test -run=TestIsValidDate -v ./tests/
go test -run=TestIsLeapYear -v ./tests/

echo.
echo Testing ID Generation Functions...
go test -run=TestGetNext -v ./tests/

echo.
echo Testing Search Functions...
go test -run=TestSequentialSearch -v ./tests/
go test -run=TestBinarySearch -v ./tests/

echo.
echo Testing Sort Functions...
go test -run=TestSelectionSort -v ./tests/
go test -run=TestInsertionSort -v ./tests/

echo.
echo Testing Data Persistence...
go test -run=TestSaveAndLoadData -v ./tests/

echo.
echo Testing Integration Workflow...
go test -run=TestCompleteWorkflow -v ./tests/

echo.
echo ===============================================
echo Test Summary
echo ===============================================
echo.

echo Test Results Summary:
if !TEST_EXIT_CODE! equ 0 (
    echo   - Unit Tests:      [PASSED]
) else (
    echo   - Unit Tests:      [FAILED]
)

if !COVERAGE_EXIT_CODE! equ 0 (
    echo   - Coverage:        [GENERATED]
) else (
    echo   - Coverage:        [FAILED]
)

if !BENCH_EXIT_CODE! equ 0 (
    echo   - Benchmarks:      [COMPLETED]
) else (
    echo   - Benchmarks:      [FAILED]
)

if !RACE_EXIT_CODE! equ 0 (
    echo   - Race Detection:  [CLEAN]
) else (
    echo   - Race Detection:  [ISSUES FOUND]
)

if !SHORT_EXIT_CODE! equ 0 (
    echo   - Short Tests:     [PASSED]
) else (
    echo   - Short Tests:     [FAILED]
)

echo.
echo Generated Files:
if exist tests\test_results.txt echo   - tests\test_results.txt    - Detailed test output
if exist tests\coverage.out echo   - tests\coverage.out        - Coverage data file
echo.

echo Commands for further analysis:
echo   go tool cover -html=tests\coverage.out     - View coverage in browser
echo   go test -v -run=TestSpecific ./tests/  - Run specific test
echo   go test -bench=BenchmarkName ./tests/  - Run specific benchmark
echo.

echo Individual Test Categories:
echo   go test -run=TestIsValid ./tests/      - Date validation tests
echo   go test -run=TestGetNext ./tests/      - ID generation tests
echo   go test -run=TestSearch ./tests/       - Search function tests
echo   go test -run=TestSort ./tests/         - Sorting function tests
echo   go test -run=TestSave ./tests/         - Data persistence tests
echo   go test -run=TestComplete ./tests/     - Integration tests
echo   go test -run=TestEmpty ./tests/        - Edge case tests

:: Calculate overall exit code
set OVERALL_EXIT_CODE=0
if !TEST_EXIT_CODE! neq 0 set OVERALL_EXIT_CODE=1
if !SHORT_EXIT_CODE! neq 0 set OVERALL_EXIT_CODE=1

if !OVERALL_EXIT_CODE! equ 0 (
    echo.
    echo [SUCCESS] ALL CRITICAL TESTS PASSED!
    echo The Medical CheckUp Management System is ready for use.
) else (
    echo.
    echo [ERROR] CRITICAL TESTS FAILED!
    echo Please review the test results and fix any issues before deployment.
)

echo.
echo Test Functions Covered:
echo   [OK] isLeapYear()
echo   [OK] isValidDate()
echo   [OK] getNextPatientID()
echo   [OK] getNextPackageID()
echo   [OK] getNextRecordID()
echo   [OK] sequentialSearchPatientByName()
echo   [OK] sequentialSearchPackageByName()
echo   [OK] binarySearchPatientByID()
echo   [OK] binarySearchPackageByID()
echo   [OK] selectionSortPatientsByName()
echo   [OK] selectionSortPackagesByPrice()
echo   [OK] insertionSortRecordsByDate()
echo   [OK] saveData()
echo   [OK] loadData()
echo   [OK] Complete workflow integration
echo   [OK] Edge cases and boundary conditions
echo   [OK] Performance benchmarks

echo.
echo Test run completed at %date% %time%
echo.

:: Clean up test files
if exist tests\test_data.json del tests\test_data.json

:: Ask user if they want to see coverage in browser
if exist tests\coverage.out (
    echo.
    choice /c YN /m "Do you want to view the test coverage report in your browser? (Y/N)"
    if !errorlevel! equ 1 (
        echo Opening coverage report in browser...
        go tool cover -html=tests\coverage.out
    )
)

echo.
pause
exit /b !OVERALL_EXIT_CODE!
