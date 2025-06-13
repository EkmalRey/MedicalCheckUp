package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
)

const NMAX int = 100

type Patient struct {
	ID     int    `json:"patient_id"`
	Name   string `json:"patient_name"`
	Gender string `json:"patient_gender"`
	Age    int    `json:"patient_age"`
}

type tabPatient struct {
	Daftar [NMAX]Patient `json:"tab_patient_daftar"`
	N      int           `json:"tab_patient_n"`
}

type Package struct {
	ID       int    `json:"package_id"`
	Name     string `json:"package_name"`
	Category string `json:"package_category"`
	Price    int    `json:"package_price"`
}

type tabPackage struct {
	Daftar [NMAX]Package `json:"tab_package_daftar"`
	N      int           `json:"tab_package_n"`
}

type Record struct {
	ID      int     `json:"record_id"`
	Patient Patient `json:"record_patient"`
	Package Package `json:"record_package"`
	Year    int     `json:"record_year"`
	Month   int     `json:"record_month"`
	Day     int     `json:"record_day"`
	Result  string  `json:"record_result"`
}

type tabRecord struct {
	Daftar [NMAX]Record `json:"tab_record_daftar"`
	N      int          `json:"tab_record_n"`
}

var patients tabPatient
var packages tabPackage
var records tabRecord

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func bold(sentence string) {
	bold := color.New(color.FgWhite, color.Bold)
	bold.Println(sentence)
}

func underline(sentence string) {
	underline := color.New(color.FgWhite, color.Underline)
	underline.Println(sentence)
}

func validDate(year, month, day int) bool {
	var totalDay int
	var kabisat bool
	if year <= 0 {
		fmt.Println("   Invalid Year!")
		return false
	} else if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		kabisat = true
	}
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		totalDay = 31
	case 2:
		if kabisat {
			totalDay = 29
		} else {
			totalDay = 28
		}
	case 4, 6, 9, 11:
		totalDay = 30
	default:
		fmt.Println("   Invalid Month!")
		return false
	}
	if day > totalDay || day <= 0 {
		fmt.Println(" Invalid Day!")
		return false
	}
	return true
}

func sequentialSearch_nameToRecord(flag string, value string) {
	var i int
	var found, validSelect bool = false, false
	var choice int
	var tempRecord tabRecord = records
	selectionSort_record(&tempRecord)
	for !found {
		switch flag {
		case "patient_name_patient":
			for i = 0; i < patients.N; i++ {
				if patients.Daftar[i].Name == value {
					patient_select(i)
					found = true
				}
			}
			if !found {
				fmt.Print("   Patient not found! Enter anything to return.\n   ")
				fmt.Scan(&choice)
				patient_searchMenu()
				found = true
			}
		case "patient_name_record":
			var record Record
			for i = 0; i < tempRecord.N; i++ {
				record = tempRecord.Daftar[i]
				if record.Patient.Name == value {
					fmt.Printf("   - %d = %s %d %s %s %d-%d-%d %s\n", record.ID, record.Patient.Name, record.Patient.Age, record.Patient.Gender,
						record.Package.Name, record.Day, record.Month, record.Year, record.Result)
					found = true
				}
			}
			if found {
				fmt.Println("  ------------------------------")
				fmt.Print("   Select Records ID (Enter 0 To Return): ")
				for !validSelect {
					var index int
					fmt.Scan(&choice)
					if choice == 0 {
						validSelect = true
					} else {
						index = binarySearch_idToIndex("record", choice)
						if index == -1 {
							fmt.Print("   Records ID Not Found! Select Records ID (Enter 0 To Return): ")
						} else {
							record_select(index)
							validSelect = true
						}
					}
				}
			} else {
				fmt.Printf("   Record of patient %s not found! Enter anything to return.\n   ", value)
				fmt.Scan(&choice)
				found = true
			}
		case "package_name_record":
			var record Record
			for i = 0; i < tempRecord.N; i++ {
				record = tempRecord.Daftar[i]
				if record.Package.Name == value {
					fmt.Printf("   - %d = %s %d %s %s %d-%d-%d %s\n", record.ID, record.Patient.Name, record.Patient.Age, record.Patient.Gender,
						record.Package.Name, record.Day, record.Month, record.Year, record.Result)
					found = true
				}
			}
			if found {
				fmt.Println("  ------------------------------")
				fmt.Print("   Select Records ID (Enter 0 To Return): ")
				for !validSelect {
					var index int
					fmt.Scan(&choice)
					if choice == 0 {
						validSelect = true
					} else {
						index = binarySearch_idToIndex("record", choice)
						if index == -1 {
							fmt.Print("   Records ID Not Found! Select Records ID (Enter 0 To Return): ")
						} else {
							record_select(index)
							validSelect = true
						}
					}
				}
			} else {
				fmt.Printf("   Record of package %s not found! Enter anything to return.\n   ", value)
				fmt.Scan(&choice)
				found = true
			}
		case "record_time_record":
			var record Record
			var year, month string = value[2:6], value[0:2]
			var recordMonth string
			for i = 0; i < tempRecord.N; i++ {
				record = tempRecord.Daftar[i]
				if record.Month < 10 {
					recordMonth = "0" + fmt.Sprint(record.Month)
				} else {
					recordMonth = fmt.Sprint(record.Month)
				}
				if fmt.Sprint(record.Year) == year && recordMonth == month {
					fmt.Printf("   - %d = %s %d %s %s %d-%d-%d %s\n", record.ID, record.Patient.Name, record.Patient.Age, record.Patient.Gender,
						record.Package.Name, record.Day, record.Month, record.Year, record.Result)
					found = true
				}
			}
			if found {
				fmt.Println("  ------------------------------")
				fmt.Print("   Select Records ID (Enter 0 To Return): ")
				for !validSelect {
					var index int
					fmt.Scan(&choice)
					if choice == 0 {
						record_searchMenu()
						validSelect = true
					} else {
						index = binarySearch_idToIndex("record", choice)
						if index == -1 {
							fmt.Print("   Records ID Not Found! Select Records ID (Enter 0 To Return): ")
						} else {
							record_select(index)
							validSelect = true
						}
					}
				}
			} else {
				fmt.Printf("   Record from %s-%s not found! Enter anything to return.\n   ", month, year)
				fmt.Scan(&choice)
				found = true
			}
		}
	}
}

func binarySearch_idToIndex(flag string, value int) int {
	var left, right, mid, index int
	left = 0
	index = -1
	switch flag {
	case "patient":
		right = patients.N
		for left <= right && index == -1 {
			mid = (left + right) / 2
			if value < patients.Daftar[mid].ID {
				right = mid - 1
			} else if value > patients.Daftar[mid].ID {
				left = mid + 1
			} else {
				index = mid
			}
		}
	case "package":
		right = packages.N
		for left <= right && index == -1 {
			mid = (left + right) / 2
			if value < packages.Daftar[mid].ID {
				right = mid - 1
			} else if value > packages.Daftar[mid].ID {
				left = mid + 1
			} else {
				index = mid
			}
		}
	case "record":
		right = records.N
		for left <= right && index == -1 {
			mid = (left + right) / 2
			if value < records.Daftar[mid].ID {
				right = mid - 1
			} else if value > records.Daftar[mid].ID {
				left = mid + 1
			} else {
				index = mid
			}
		}
	}
	return index
}

func insertionSort_record(tab *tabRecord) {
	var pass, i int
	var temp Record
	pass = 1
	for pass < tab.N {
		i = pass
		temp = tab.Daftar[pass]
		for i > 0 && (temp.Year > tab.Daftar[i-1].Year || (temp.Year == tab.Daftar[i-1].Year && temp.Month > tab.Daftar[i-1].Month) || (temp.Year == tab.Daftar[i-1].Year && temp.Month == tab.Daftar[i-1].Month && temp.Day > tab.Daftar[i-1].Day)) {
			tab.Daftar[i] = tab.Daftar[i-1]
			i--
		}
		tab.Daftar[i] = temp
		pass++
	}
}

func selectionSort_record(tab *tabRecord) {
	var pass, idx, i int
	var temp Record
	pass = 1
	for pass < tab.N {
		idx = pass - 1
		i = pass
		for i < tab.N {
			if tab.Daftar[i].Year > tab.Daftar[idx].Year || (tab.Daftar[i].Year == tab.Daftar[idx].Year && tab.Daftar[i].Month > tab.Daftar[idx].Month) || (tab.Daftar[i].Year == tab.Daftar[idx].Year && tab.Daftar[i].Month == tab.Daftar[idx].Month && tab.Daftar[i].Day > tab.Daftar[idx].Day) {
				idx = i
			}
			i++
		}
		temp = tab.Daftar[pass-1]
		tab.Daftar[pass-1] = tab.Daftar[idx]
		tab.Daftar[idx] = temp
		pass++
	}
}

func header() {
	bold("  ------------------------------")
	bold("  *                            *")
	bold("  *  Medical Check Up Program  *")
	bold("  *     By Ekmal && Yunita     *")
	bold("  *                            *")
}

func main_menu() {
	clearScreen()
	header()
	bold("  ----------MAIN  MENU----------")
	var choice int
	var validChoice bool = false
	fmt.Print("   ")
	underline("Menu Navigation :")
	fmt.Println("   1. Package Management")
	fmt.Println("   2. Patient Management")
	fmt.Println("   3. Record Management")
	fmt.Println("   4. Reports")
	fmt.Println("   5. Quick-add Records")
	fmt.Println("   0. Exit")
	fmt.Print("   Choice: ")
	for !validChoice {
		fmt.Scan(&choice)
		switch choice {
		case 1:
			package_management()
			validChoice = true
		case 2:
			patient_management()
			validChoice = true
		case 3:
			record_management()
			validChoice = true
		case 4:
			report_management()
			validChoice = true
		case 5:
			record_add("")
			validChoice = true
			main_menu()
		case 0:
			validChoice = true
		default:
			fmt.Print("   Invalid Choice! Choice : ")
		}
	}
}

func package_management() {
	clearScreen()
	header()
	bold("  ------PACKAGE MANAGEMENT------")
	var choice int
	var value string
	var validChoice bool = false
	fmt.Println("   Total Packages Registered : ", packages.N)
	fmt.Print("   ")
	underline("Menu Navigation :")
	fmt.Println("   1. See All Package")
	fmt.Println("   2. Add Package")
	fmt.Println("   3. Search Record by Package")
	fmt.Println("   0. Return")
	fmt.Print("   Choice: ")
	for !validChoice {
		fmt.Scan(&choice)
		switch choice {
		case 1:
			package_see()
			validChoice = true
		case 2:
			package_add()
			validChoice = true
		case 3:
			fmt.Print("   Enter Package Name : ")
			fmt.Scan(&value)
			sequentialSearch_nameToRecord("package_name_record", value)
			validChoice = true
			package_management()
		case 0:
			main_menu()
			validChoice = true
		default:
			fmt.Print("   Invalid Choice! Choice : ")
		}
	}
}

func package_see() {
	clearScreen()
	header()
	bold("  --------PACKAGES  LIST--------")
	var i, value int
	var valid bool = false
	for i = 0; i < packages.N; i++ {
		fmt.Printf("   %d %s %s %d\n", packages.Daftar[i].ID, packages.Daftar[i].Name, packages.Daftar[i].Category, packages.Daftar[i].Price)
	}
	fmt.Println("  ------------------------------")
	fmt.Print("   Select Package ID (Enter 0 To Return) : ")
	for !valid {
		fmt.Scan(&value)
		if value == 0 {
			package_management()
			valid = true
		} else {
			i = binarySearch_idToIndex("package", value)
			if i == -1 {
				fmt.Print("   Package ID Not Found! Select Package ID (Enter 0 To Return) : ")
			} else {
				package_select(i)
				valid = true
				package_see()
			}
		}
	}
}

func package_select(idx int) {
	clearScreen()
	header()
	bold("  --------PACKAGE DETAIL--------")
	var choice int
	var validChoice bool = false
	fmt.Println("   Package ID       : ", packages.Daftar[idx].ID)
	fmt.Println("   Package Name     : ", packages.Daftar[idx].Name)
	fmt.Println("   Package Category : ", packages.Daftar[idx].Category)
	fmt.Printf("   Package Price    :  Rp.%d\n", packages.Daftar[idx].Price)
	fmt.Println("  ------------------------------")
	fmt.Print("   ")
	underline("Menu Navigation :")
	fmt.Println("   1. Edit Package")
	fmt.Println("   2. Delete Package")
	fmt.Println("   3. Search Record Of This Package")
	fmt.Println("   0. Return")
	fmt.Print("   Choice : ")
	for !validChoice {
		fmt.Scan(&choice)
		switch choice {
		case 1:
			package_edit(idx)
			validChoice = true
		case 2:
			package_delete(idx)
			validChoice = true
		case 3:
			sequentialSearch_nameToRecord("package_name_record", packages.Daftar[idx].Name)
			validChoice = true
			package_select(idx)
		case 0:
			validChoice = true
		default:
			fmt.Print("   Invalid Choice! Choice : ")
		}
	}

}

func package_add() {
	clearScreen()
	header()
	var id int
	var name, category string
	var price int
	bold("  -----CREATING NEW PACKAGE-----")
	id = 10001 + packages.N
	fmt.Print("   Package name : ")
	fmt.Scan(&name)
	fmt.Print("   Package category (Basic/Standard/Advanced) : ")
	fmt.Scan(&category)
	for category != "Basic" && category != "Standard" && category != "Advanced" {
		fmt.Print("   Invalid Category! Package Category (Basic/Standard/Advanced) : ")
		fmt.Scan(&category)
	}
	fmt.Print("   Package price : ")
	fmt.Scan(&price)
	for price < 0 {
		fmt.Print("   Invalid Price! Package Price : ")
		fmt.Scan(&price)
	}
	packages.Daftar[packages.N].ID = id
	packages.Daftar[packages.N].Name = name
	packages.Daftar[packages.N].Category = category
	packages.Daftar[packages.N].Price = price
	packages.N++
	fmt.Printf("   %d %s %s %d Successfully Added!\n", packages.Daftar[packages.N].ID, packages.Daftar[packages.N].Name, packages.Daftar[packages.N].Category, packages.Daftar[packages.N].Price)
	package_management()
}

func package_edit(idx int) {
	clearScreen()
	header()
	bold("  ------EDITING PACKAGES------")
	var name, category, accept string
	var price int
	var validAccept bool = false
	var data Package = packages.Daftar[idx]
	fmt.Println("   ID : ", packages.Daftar[idx].ID)
	fmt.Println("   Old Package Name : ", data.Name)
	fmt.Print("   New Package Name : ")
	fmt.Scan(&name)
	fmt.Println("   Old Package Category : ", data.Category)
	fmt.Print("   New Package Category (Basic/Standard/Advanced) : ")
	fmt.Scan(&category)
	for category != "Basic" && category != "Standard" && category != "Advanced" {
		fmt.Println("   Invalid Category!")
		fmt.Print("   New Package Category (Basic/Standard/Advanced) : ")
		fmt.Scan(&category)
	}
	fmt.Println("   Old Package Price : ", data.Price)
	fmt.Print("   New Package Price : ")
	fmt.Scan(&price)
	for price < 0 {
		fmt.Println("   Invalid Price! Price cannot be less than zero!")
		fmt.Print("   New Package Price : ")
		fmt.Scan(&price)
	}
	fmt.Println("   Package ID       = ", data.ID)
	fmt.Println("   Package Name     = ", data.Name, ">>>", name)
	fmt.Println("   Package Category = ", data.Category, ">>>", category)
	fmt.Println("   Package Price    = ", data.Price, ">>>", price)
	fmt.Print("   Accept Changes? (yes or no) : ")
	for !validAccept {
		fmt.Scan(&accept)
		if accept == "yes" {
			packages.Daftar[idx].Name = name
			packages.Daftar[idx].Category = category
			packages.Daftar[idx].Price = price
			fmt.Println("   Package data saved successfully!")
			validAccept = true
		} else if accept == "no" {
			fmt.Println("   Discarding Changes...")
			validAccept = true
		} else {
			fmt.Print("   Invalid input! Accept Changes? (yes or no) : ")
		}
	}
	package_select(idx)
}

func package_delete(idx int) {
	var accept string
	var validAccept bool = false
	fmt.Print("   Delete Package? (yes or no) : ")
	for !validAccept {
		fmt.Scan(&accept)
		if accept == "yes" {
			var i int
			for i = idx; i < packages.N; i++ {
				temp := packages.Daftar[i].ID
				packages.Daftar[i] = packages.Daftar[i+1]
				packages.Daftar[i].ID = temp
			}
			packages.Daftar[i].ID = 0
			packages.Daftar[i].Name = ""
			packages.Daftar[i].Category = ""
			packages.Daftar[i].Price = 0
			packages.N--
			validAccept = true
			package_management()
		} else if accept == "no" {
			package_select(idx)
			validAccept = true
		} else {
			fmt.Print("   Invalid input! Delete Package? (yes or no) : ")
		}
	}
}

func patient_management() {
	clearScreen()
	header()
	bold("  ------PATIENT MANAGEMENT------")
	var choice int
	var validChoice bool = false
	fmt.Println("   Total patient registered :", patients.N)
	fmt.Print("   ")
	underline("Menu Navigation :")
	fmt.Println("   1. See patients")
	fmt.Println("   2. Add patient")
	fmt.Println("   3. Search Patient")
	fmt.Println("   0. Return")
	fmt.Print("   Choice: ")
	for !validChoice {
		fmt.Scan(&choice)
		switch choice {
		case 1:
			patient_see("id")
			validChoice = true
		case 2:
			patient_add()
			validChoice = true
		case 3:
			patient_searchMenu()
			validChoice = true
		case 0:
			main_menu()
			validChoice = true
		default:
			fmt.Print("   Invalid Choice! Choice : ")
		}
	}
}

func patient_see(flag string) {
	clearScreen()
	header()
	bold("  --------PATIENTS  LIST--------")
	var i, value int
	var validChoice bool = false
	patient_show(flag)
	fmt.Println("  ------------------------------")
	fmt.Print("   ")
	underline("Menu Navigation :")
	fmt.Println("   1. Sort By ID ASCENDING")
	fmt.Println("   2. Sort By ID DESCENDING")
	fmt.Println("   3. Sort By Package ASCENDING")
	fmt.Println("   4. Sort By Package DESCENDING")
	fmt.Println("   5. Sort By Latest Medical Checkup")
	fmt.Println("   6. Sort By Oldest Medical Checkup")
	fmt.Println("   0. Return")
	fmt.Println("   20001-20002. Select Patient ID")
	fmt.Print("   Choice : ")
	for !validChoice {
		fmt.Scan(&value)
		if value == 0 {
			patient_management()
			validChoice = true
		} else if value == 1 {
			validChoice = true
			patient_see("id_asc")
		} else if value == 2 {
			validChoice = true
			patient_see("id_desc")
		} else if value == 3 {
			validChoice = true
			patient_see("package_asc")
		} else if value == 4 {
			validChoice = true
			patient_see("package_desc")
		} else if value == 5 {
			validChoice = true
			patient_see("latest")
		} else if value == 6 {
			validChoice = true
			patient_see("oldest")
		} else {
			i = binarySearch_idToIndex("patient", value)
			if i == -1 {
				fmt.Print("   Patient ID Not Found! Select Patient ID (Enter 0 To Return) : ")
			} else {
				patient_select(i)
				validChoice = true
				patient_see("id")
			}
		}
	}
}

func patient_show(flag string) {
	var j, i int
	var tempRecord tabRecord = records
	var tempID [100]int
	var nTemp int = 0
	var exist bool
	insertionSort_record(&tempRecord)
	switch flag {
	case "id_asc":
		for i = 0; i < patients.N; i++ {
			fmt.Printf("   %d %s %d %s\n", patients.Daftar[i].ID, patients.Daftar[i].Name, patients.Daftar[i].Age, patients.Daftar[i].Gender)
		}
	case "id_desc":
		for i = patients.N - 1; i >= 0; i-- {
			fmt.Printf("   %d %s %d %s\n", patients.Daftar[i].ID, patients.Daftar[i].Name, patients.Daftar[i].Age, patients.Daftar[i].Gender)
		}
	case "package_asc":
		var k int
		for i = 0; i < packages.N; i++ {
			for j = 0; j < tempRecord.N; j++ {
				exist = false
				if packages.Daftar[i].ID == tempRecord.Daftar[j].Package.ID {
					for k = 0; k < nTemp && !exist; k++ {
						if tempRecord.Daftar[j].Patient.ID == tempID[k] {
							exist = true
						}
					}
					if !exist {
						tempID[nTemp] = tempRecord.Daftar[j].Patient.ID
						nTemp++
						fmt.Printf("   %d %s %d %s ", tempRecord.Daftar[j].Patient.ID, tempRecord.Daftar[j].Patient.Name, tempRecord.Daftar[j].Patient.Age, tempRecord.Daftar[j].Patient.Gender)
						fmt.Printf("%d-%d-%d %s\n", tempRecord.Daftar[j].Day, tempRecord.Daftar[j].Month, tempRecord.Daftar[j].Year, tempRecord.Daftar[j].Package.Name)
					}
				}
			}
		}
		for i = 0; i < patients.N; i++ {
			exist = false
			for j = 0; j < nTemp && !exist; j++ {
				if patients.Daftar[i].ID == tempID[j] {
					exist = true
				}
			}
			if !exist {
				tempID[nTemp] = patients.Daftar[i].ID
				nTemp++
				fmt.Printf("   %d %s %d %s NULL NULL\n", patients.Daftar[i].ID, patients.Daftar[i].Name, patients.Daftar[i].Age, patients.Daftar[i].Gender)
			}
		}
	case "package_desc":
		var k int
		for i = packages.N - 1; i >= 0; i-- {
			for j = 0; j < tempRecord.N; j++ {
				exist = false
				if packages.Daftar[i].ID == tempRecord.Daftar[j].Package.ID {
					for k = 0; k < nTemp && !exist; k++ {
						if tempRecord.Daftar[j].Patient.ID == tempID[k] {
							exist = true
						}
					}
					if !exist {
						tempID[nTemp] = tempRecord.Daftar[j].Patient.ID
						nTemp++
						fmt.Printf("   %d %s %d %s ", tempRecord.Daftar[j].Patient.ID, tempRecord.Daftar[j].Patient.Name, tempRecord.Daftar[j].Patient.Age, tempRecord.Daftar[j].Patient.Gender)
						fmt.Printf("%d-%d-%d %s\n", tempRecord.Daftar[j].Day, tempRecord.Daftar[j].Month, tempRecord.Daftar[j].Year, tempRecord.Daftar[j].Package.Name)
					}
				}
			}
		}
		for i = 0; i < patients.N; i++ {
			exist = false
			for j = 0; j < nTemp && !exist; j++ {
				if patients.Daftar[i].ID == tempID[j] {
					exist = true
				}
			}
			if !exist {
				tempID[nTemp] = patients.Daftar[i].ID
				nTemp++
				fmt.Printf("   %d %s %d %s NULL NULL\n", patients.Daftar[i].ID, patients.Daftar[i].Name, patients.Daftar[i].Age, patients.Daftar[i].Gender)
			}
		}
	case "latest":
		for i = 0; i < tempRecord.N; i++ {
			exist = false
			for j = 0; j < nTemp && !exist; j++ {
				if tempRecord.Daftar[i].Patient.ID == tempID[j] {
					exist = true
				}
			}
			if !exist {
				tempID[nTemp] = tempRecord.Daftar[i].Patient.ID
				nTemp++
				fmt.Printf("   %d %s %d %s ", tempRecord.Daftar[i].Patient.ID, tempRecord.Daftar[i].Patient.Name, tempRecord.Daftar[i].Patient.Age, tempRecord.Daftar[i].Patient.Gender)
				fmt.Printf("%d-%d-%d %s\n", tempRecord.Daftar[i].Day, tempRecord.Daftar[i].Month, tempRecord.Daftar[i].Year, tempRecord.Daftar[i].Package.Name)
			}
		}
		for i = 0; i < patients.N; i++ {
			exist = false
			for j = 0; j < nTemp && !exist; j++ {
				if patients.Daftar[i].ID == tempID[j] {
					exist = true
				}
			}
			if !exist {
				tempID[nTemp] = patients.Daftar[i].ID
				nTemp++
				fmt.Printf("   %d %s %d %s NULL NULL\n", patients.Daftar[i].ID, patients.Daftar[i].Name, patients.Daftar[i].Age, patients.Daftar[i].Gender)
			}
		}
	case "oldest":
		for i = tempRecord.N - 1; i >= 0; i-- {
			exist = false
			for j = 0; j < nTemp && !exist; j++ {
				if tempRecord.Daftar[i].Patient.ID == tempID[j] {
					exist = true
				}
			}
			if !exist {
				tempID[nTemp] = tempRecord.Daftar[i].Patient.ID
				nTemp++
				fmt.Printf("   %d %s %d %s ", tempRecord.Daftar[i].Patient.ID, tempRecord.Daftar[i].Patient.Name, tempRecord.Daftar[i].Patient.Age, tempRecord.Daftar[i].Patient.Gender)
				fmt.Printf("%d-%d-%d %s\n", tempRecord.Daftar[i].Day, tempRecord.Daftar[i].Month, tempRecord.Daftar[i].Year, tempRecord.Daftar[i].Package.Name)
			}
		}
		for i = 0; i < patients.N; i++ {
			exist = false
			for j = 0; j < nTemp && !exist; j++ {
				if patients.Daftar[i].ID == tempID[j] {
					exist = true
				}
			}
			if !exist {
				tempID[nTemp] = patients.Daftar[i].ID
				nTemp++
				fmt.Printf("   %d %s %d %s NULL NULL\n", patients.Daftar[i].ID, patients.Daftar[i].Name, patients.Daftar[i].Age, patients.Daftar[i].Gender)
			}
		}
	}
}

func patient_add() {
	clearScreen()
	header()
	bold("  -------CREATING PATIENT-------")
	var name, gender string
	var id int
	var age int
	var data Patient
	var validAge, validGender bool = false, false
	fmt.Println("   Creating a new patient data...")
	id = 20001 + patients.N
	fmt.Print("   Patient name : ")
	fmt.Scan(&name)
	fmt.Print("   Patient age : ")
	for !validAge {
		var err error
		_, err = fmt.Scan(&age)
		if age > 0 && err == nil {
			validAge = true
		} else {
			fmt.Print("   Invalid! Patient age : ")
		}
	}
	fmt.Print("   Patient gender (M/F) : ")
	for !validGender {
		var err error
		_, err = fmt.Scan(&gender)
		if (gender == "M" || gender == "F") && err == nil {
			validGender = true
		} else {
			fmt.Print("   Invalid! Patient gender (M/F) : ")
		}
	}
	data = Patient{
		ID:     id,
		Name:   name,
		Age:    age,
		Gender: gender,
	}
	patients.Daftar[patients.N] = data
	patients.N++
	fmt.Println("   Patient data added to database!")
	patient_management()
}

func patient_select(idx int) {
	clearScreen()
	header()
	bold("  --------PATIENT DETAIL--------")
	var patient Patient
	var choice, i, value, count int
	var tempRecord tabRecord = records
	var validChoice, found bool = false, false
	patient = patients.Daftar[idx]
	fmt.Println("   ID      :", patient.ID)
	fmt.Println("   Name    :", patient.Name)
	fmt.Println("   Age     :", patient.Age)
	fmt.Println("   Gender  :", patient.Gender)
	fmt.Println("   Records :")
	selectionSort_record(&tempRecord)
	for i = 0; i < tempRecord.N; i++ {
		if tempRecord.Daftar[i].Patient.ID == patients.Daftar[idx].ID {
			fmt.Printf("   - %d %s %d-%d-%d %s\n", tempRecord.Daftar[i].ID, tempRecord.Daftar[i].Package.Name, tempRecord.Daftar[i].Day, tempRecord.Daftar[i].Month, tempRecord.Daftar[i].Year, tempRecord.Daftar[i].Result)
			count++
		}
	}
	if count == 0 {
		fmt.Println("   No Record Found!")
	}
	fmt.Println("  ------------------------------")
	fmt.Print("   ")
	underline("Menu Navigation :")
	fmt.Println("   1. Edit Patient Information")
	fmt.Println("   2. Delete Patient")
	fmt.Println("   3. See Patient Records Detail")
	fmt.Println("   4. Add Record To Patient")
	fmt.Println("   0. Return")
	fmt.Print("   Choice : ")
	for !validChoice {
		fmt.Scan(&choice)
		switch choice {
		case 1:
			patient_edit(idx)
			validChoice = true
		case 2:
			patient_delete(idx)
			validChoice = true
		case 3:
			fmt.Print("   Enter Record ID (Press 0 To Return): ")
			for !found {
				fmt.Scan(&value)
				if value == 0 {
					patient_select(idx)
					found = true
				} else {
					i = binarySearch_idToIndex("record", value)
					if i == -1 {
						fmt.Print("   Records ID Not Found! Select Records ID (Enter 0 To Return): ")
					} else {
						record_select(i)
						found = true
						patient_select(idx)
					}
				}
			}
			validChoice = true
		case 4:
			record_add(patient.Name)
			validChoice = true
			patient_select(idx)
		case 0:
			validChoice = true
		default:
			fmt.Print("   Invalid Choice! Choice : ")
		}
	}
}

func patient_edit(idx int) {
	clearScreen()
	header()
	bold("  -------EDITING  PATIENT-------")
	var name, gender, accept string
	var age int
	var validAge, validGender, validAccept bool = false, false, false
	var data Patient = patients.Daftar[idx]
	fmt.Println("   Editing existing patient data...")
	fmt.Println("   Old name :", data.Name)
	fmt.Print("   New name : ")
	fmt.Scan(&name)
	fmt.Println("   Old gender :", data.Gender)
	fmt.Print("   New gender (M/F) : ")
	for !validGender {
		var err error
		_, err = fmt.Scan(&gender)
		if (gender == "M" || gender == "F") && err == nil {
			validGender = true
		} else {
			fmt.Print("   Invalid! New patient gender (M/F) : ")
		}
	}
	fmt.Println("   Old age :", data.Age)
	fmt.Print("   New age : ")
	for !validAge {
		var err error
		_, err = fmt.Scan(&age)
		if age > 0 && err == nil {
			validAge = true
		} else {
			fmt.Print("   Invalid! New patient age : ")
		}
	}
	fmt.Println("   Patient ID     =", data.ID)
	fmt.Println("   Patient Name   =", data.Name, ">>>", name)
	fmt.Println("   Patient Age    =", data.Age, ">>>", age)
	fmt.Println("   Patient Gender =", data.Gender, ">>>", gender)
	fmt.Print("   Accept Changes? (yes or no) : ")
	for !validAccept {
		fmt.Scan(&accept)
		if accept == "yes" {
			patients.Daftar[idx].Name = name
			patients.Daftar[idx].Gender = gender
			patients.Daftar[idx].Age = age
			fmt.Println("   Patient data saved successfully!")
			validAccept = true
		} else if accept == "no" {
			fmt.Println("   Discarding Changes...")
			validAccept = true
		} else {
			fmt.Print("   Invalid input! Accept Changes? (yes or no) : ")
		}
	}
	patient_select(idx)
}

func patient_delete(idx int) {
	var validAccept bool = false
	var accept string
	fmt.Print("   Are you sure to delete this patient? (yes or no) : ")
	for !validAccept {
		fmt.Scan(&accept)
		if accept == "yes" {
			var i int
			for i = idx; i < patients.N; i++ {
				temp := patients.Daftar[i].ID
				patients.Daftar[i] = patients.Daftar[i+1]
				patients.Daftar[i].ID = temp
			}
			patients.Daftar[i].ID = 0
			patients.Daftar[i].Name = ""
			patients.Daftar[i].Gender = ""
			patients.Daftar[i].Age = 0
			patients.N--
			validAccept = true
		} else if accept == "no" {
			patient_select(idx)
			validAccept = true
		} else {
			fmt.Print(" Invalid Choice! Are you sure to delete this patient? (yes or no) : ")
		}
	}
}

func patient_searchMenu() {
	clearScreen()
	header()
	bold("  --------PATIENT SEARCH--------")
	var choice int
	var validChoice bool = false
	var value string
	fmt.Println("   Search patient by :")
	fmt.Println("   1. Patient Name")
	fmt.Println("   2. Package Name")
	fmt.Println("   3. MCU Time")
	fmt.Println("   0. Return")
	for !validChoice {
		fmt.Print("   Choice : ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Print("   Enter Patient Name : ")
			fmt.Scan(&value)
			sequentialSearch_nameToRecord("patient_name_patient", value)
			validChoice = true
			patient_searchMenu()
		case 2:
			fmt.Print("   Enter Package Name : ")
			fmt.Scan(&value)
			sequentialSearch_nameToRecord("package_name_record", value)
			validChoice = true
			patient_searchMenu()
		case 3:
			var year, month string
			fmt.Print("   Enter MCU Year (YYYY) : ")
			fmt.Scan(&year)
			if len(year) != 4 {
				fmt.Print("   Invalid Year Format! : Enter MCU Year (YYYY) :")
				fmt.Scan(&year)
			}
			fmt.Print("   Enter MCU Month (MM) : ")
			fmt.Scan(&month)
			if len(month) != 2 {
				fmt.Print("   Invalid Month Format! : Enter MCU Month (MM) :")
				fmt.Scan(&month)
			}
			value = month + year
			sequentialSearch_nameToRecord("record_time_record", value)
			validChoice = true
			patient_searchMenu()
		case 0:
			patient_management()
			validChoice = true
		default:
			fmt.Print("   Invalid! ")
		}
	}
}

func record_management() {
	clearScreen()
	header()
	bold("  ------RECORD  MANAGEMENT------")
	var choice int
	var validChoice bool = false
	fmt.Println("   Total Records Registered :", records.N)
	fmt.Print("   ")
	underline("Menu Navigation :")
	fmt.Println("   1. See All Records")
	fmt.Println("   2. Add Record")
	fmt.Println("   3. Search Record")
	fmt.Println("   0. Return")
	fmt.Print("   Choice: ")
	for !validChoice {
		fmt.Scan(&choice)
		switch choice {
		case 1:
			record_see()
			validChoice = true
		case 2:
			record_add("")
			validChoice = true
			record_management()
		case 3:
			record_searchMenu()
			validChoice = true
		case 0:
			main_menu()
			validChoice = true
		default:
			fmt.Print("   Invalid Choice! Choice : ")
		}
	}
}

func record_see() {
	clearScreen()
	header()
	bold("  ---------RECORDS LIST---------")
	var choice int
	var validChoice bool = false
	var i, index int
	var record Record
	for i = records.N - 1; i >= 0; i-- {
		record = records.Daftar[i]
		fmt.Printf("   %d %s %d %s %s %d-%d-%d %s\n", record.ID, record.Patient.Name, record.Patient.Age, record.Patient.Gender,
			record.Package.Name, record.Day, record.Month, record.Year, record.Result)
	}
	fmt.Println("  ------------------------------")
	fmt.Print("   Select Records ID (Enter 0 To Return): ")
	for !validChoice {
		fmt.Scan(&choice)
		if choice == 0 {
			record_management()
			validChoice = true
		} else {
			index = binarySearch_idToIndex("record", choice)
			if index == -1 {
				fmt.Print("   Records ID Not Found! Select Records ID (Enter 0 To Return): ")
			} else {
				record_select(index)
				validChoice = true
				record_see()
			}
		}
	}

}

func record_searchMenu() {
	clearScreen()
	header()
	bold("  --------RECORD  SEARCH--------")
	var choice int
	var validChoice bool = false
	var value string
	fmt.Println("   Search record by :")
	fmt.Println("   1. Patient Name")
	fmt.Println("   2. Package Name")
	fmt.Println("   3. MCU Time")
	fmt.Println("   0. Return")
	for !validChoice {
		fmt.Print("   Choice : ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Print("   Enter Name : ")
			fmt.Scan(&value)
			sequentialSearch_nameToRecord("patient_name_record", value)
			validChoice = true
			record_searchMenu()
		case 2:
			fmt.Print("   Enter Package Name : ")
			fmt.Scan(&value)
			sequentialSearch_nameToRecord("package_name_record", value)
			validChoice = true
			record_searchMenu()
		case 3:
			var year, month string
			fmt.Print("   Enter MCU Year (YYYY) : ")
			fmt.Scan(&year)
			if len(year) != 4 {
				fmt.Print("   Invalid Year Format! : Enter MCU Year (YYYY) :")
				fmt.Scan(&year)
			}
			fmt.Print("   Enter MCU Month (MM) : ")
			fmt.Scan(&month)
			if len(month) != 2 {
				fmt.Print("   Invalid Month Format! : Enter MCU Month (MM) :")
				fmt.Scan(&month)
			}
			value = month + year
			sequentialSearch_nameToRecord("record_time_record", value)
			validChoice = true
			record_searchMenu()
		case 0:
			record_management()
			validChoice = true
		default:
			fmt.Print("   Invalid! ")
		}
	}
}

func record_select(idx int) {
	clearScreen()
	header()
	bold("  --------RECORD  DETAIL--------")
	var choice, patient_idx int
	var validChoice bool = false
	fmt.Println("   Record ID              :", records.Daftar[idx].ID)
	fmt.Println("   Record Patient Name    :", records.Daftar[idx].Patient.Name)
	fmt.Println("   Record Patient Age     :", records.Daftar[idx].Patient.Age)
	fmt.Println("   Record Patient Gender  :", records.Daftar[idx].Patient.Gender)
	fmt.Println("   Record Package Name    :", records.Daftar[idx].Package.Name)
	fmt.Printf("   Record Date            : %d-%d-%d\n", records.Daftar[idx].Day, records.Daftar[idx].Month, records.Daftar[idx].Year)
	fmt.Println("   Record Result          :", records.Daftar[idx].Result)
	fmt.Println("  ------------------------------")
	fmt.Print("   ")
	underline("Menu Navigation :")
	fmt.Println("   1. Edit Records")
	fmt.Println("   2. Delete Records")
	fmt.Println("   3. See Patient Detail")
	fmt.Println("   0. Return")
	fmt.Print("   Choice : ")
	for !validChoice {
		fmt.Scan(&choice)
		switch choice {
		case 1:
			record_edit(idx)
			validChoice = true
			record_select(idx)
		case 2:
			record_delete(idx)
			validChoice = true
		case 3:
			fmt.Println(records.Daftar[idx].Patient.ID)
			patient_idx = binarySearch_idToIndex("patient", records.Daftar[idx].Patient.ID)
			patient_select(patient_idx)
			validChoice = true
		case 0:
			validChoice = true
		default:
			fmt.Print("   Invalid Choice! Choice : ")
		}
	}
}

func record_add(patient_name string) {
	clearScreen()
	header()
	bold("  ------ADDING MCU RECORDS------")
	var validPatient, validPackage, stop bool = false, false, false
	var package_name, result string
	var year, month, day, i int = 0, 0, 0, 0
	var patient Patient
	var pkg Package
	fmt.Println("   Enter 0 To Return")
	if patient_name == "" {
		for !validPatient {
			fmt.Print("   Patient Name (Enter \"list\" to see all patients) : ")
			fmt.Scan(&patient_name)
			if patient_name == "list" {
				fmt.Println("  ------------------------------")
				for i = 0; i < patients.N; i++ {
					fmt.Printf("   %d %s %s %d\n", patients.Daftar[i].ID, patients.Daftar[i].Name, patients.Daftar[i].Gender, patients.Daftar[i].Age)
				}
				fmt.Println("  ------------------------------")
			} else if patient_name == "0" {
				validPatient = true
				validPackage = true
				stop = true
			} else {
				for i, row := range patients.Daftar {
					if row.Name == patient_name {
						patient = patients.Daftar[i]
						// records.Daftar[records.N].Patient.ID = patient.ID
						// records.Daftar[records.N].Patient.Name = patient.Name
						// records.Daftar[records.N].Patient.Age = patient.Age
						// records.Daftar[records.N].Patient.Gender = patient.Gender
						validPatient = true
					}
				}
				if !validPatient {
					fmt.Print("   Patient not found! Patient Name : ")
				}
			}
		}
	} else {
		fmt.Println("   Patient Name : ", patient_name)
		for i, row := range patients.Daftar {
			if row.Name == patient_name {
				patient = patients.Daftar[i]
				// records.Daftar[records.N].Patient.ID = patient.ID
				// records.Daftar[records.N].Patient.Name = patient.Name
				// records.Daftar[records.N].Patient.Age = patient.Age
				// records.Daftar[records.N].Patient.Gender = patient.Gender
			}
		}
	}
	for !validPackage {
		fmt.Print("   Package Name (Enter \"list\" to see all packages) : ")
		fmt.Scan(&package_name)
		if package_name == "list" {
			fmt.Println("  ------------------------------")
			for i = 0; i < packages.N; i++ {
				fmt.Printf("   %d %s %s %d\n", packages.Daftar[i].ID, packages.Daftar[i].Name, packages.Daftar[i].Category, packages.Daftar[i].Price)
			}
			fmt.Println("  ------------------------------")
		} else if package_name == "0" {
			validPatient = true
			validPackage = true
			stop = true
		} else {
			for i, row := range packages.Daftar {
				if row.Name == package_name {
					pkg = packages.Daftar[i]
					// records.Daftar[records.N].Package.ID = packages.Daftar[i].ID
					// records.Daftar[records.N].Package.Name = packages.Daftar[i].Name
					// records.Daftar[records.N].Package.Category = packages.Daftar[i].Category
					// records.Daftar[records.N].Package.Price = packages.Daftar[i].Price
					validPackage = true
				}
			}
			if !validPackage {
				fmt.Println("   Package ID not found!")
			}
		}
	}
	for !stop {
		fmt.Print("   Day (DD) : ")
		fmt.Scan(&day)
		fmt.Print("   Month (MM) : ")
		fmt.Scan(&month)
		fmt.Print("   Year (YYYY) : ")
		fmt.Scan(&year)
		for !validDate(year, month, day) {
			fmt.Print("   Day (DD) : ")
			fmt.Scan(&day)
			fmt.Print("   Month (MM) : ")
			fmt.Scan(&month)
			fmt.Print("   Year (YYYY) : ")
			fmt.Scan(&year)
		}
		records.Daftar[records.N].Patient = patient
		records.Daftar[records.N].Package = pkg
		records.Daftar[records.N].ID = 30001 + records.N
		records.Daftar[records.N].Day = day
		records.Daftar[records.N].Month = month
		records.Daftar[records.N].Year = year
		fmt.Print("   Result : ")
		fmt.Scan(&result)
		records.Daftar[records.N].Result = result
		records.N++
		fmt.Println("   Medical Check Up Records Successfully Added!")
		stop = true
	}
}

func record_edit(idx int) {
	clearScreen()
	header()
	bold("  ------EDITING MCU RECORD------")
	var name string
	var i int
	var accept string
	var stop, valid bool = false, false
	var patientID, patientAge, packageID, packagePrice, year, month, day int
	var patientName, patientGender, packageName, packageCategory, result string
	var record Record = records.Daftar[idx]
	fmt.Print("   ")
	fmt.Println(record.ID)
	fmt.Println("   Old Patient Name : ", record.Patient.Name)
	fmt.Print("   New Patient Name (Enter \"list\" to see all patients) : ")
	for !stop {
		fmt.Scan(&name)
		if name == "list" {
			for i = 0; i < patients.N; i++ {
				fmt.Printf("   %d %s %s %d\n", patients.Daftar[i].ID, patients.Daftar[i].Name, patients.Daftar[i].Gender, patients.Daftar[i].Age)
			}
			fmt.Print("   New Patient Name : ")
		} else {
			for i = 0; i < patients.N && !valid; i++ {
				if patients.Daftar[i].Name == name {
					patientID = patients.Daftar[i].ID
					patientName = patients.Daftar[i].Name
					patientAge = patients.Daftar[i].Age
					patientGender = patients.Daftar[i].Gender
					stop = true
				}
			}
			if !stop {
				fmt.Print("   Patient not found! New Patient Name (Enter \"list\" to see all patients) : ")
			}
		}
	}
	stop, valid = false, false
	fmt.Println("   Old Package Name : ", record.Package.Name)
	fmt.Print("   New Package Name (Enter \"list\" to see all patients) : ")
	for !stop {
		fmt.Scan(&name)
		if name == "list" {
			for i = 0; i < packages.N; i++ {
				fmt.Printf("   %d %s %s %d\n", packages.Daftar[i].ID, packages.Daftar[i].Name, packages.Daftar[i].Category, packages.Daftar[i].Price)
			}
			fmt.Print("   New Package Name : ")
		} else {
			for i = 0; i < packages.N && !valid; i++ {
				if packages.Daftar[i].Name == name {
					packageID = packages.Daftar[i].ID
					packageName = packages.Daftar[i].Name
					packageCategory = packages.Daftar[i].Category
					packagePrice = packages.Daftar[i].Price
					stop = true
				}
			}
			if !stop {
				fmt.Print("   Package Name not found! New Package Name (Enter \"list\" to see all packages) : ")
			}
		}
	}
	fmt.Println("   Old Day : ", record.Day)
	fmt.Print("   New Day (DD) : ")
	fmt.Scan(&day)
	fmt.Println("   Old Month : ", record.Month)
	fmt.Print("   New Month (MM) : ")
	fmt.Scan(&month)
	fmt.Println("   Old Year : ", record.Year)
	fmt.Print("   New Year (YYYY) : ")
	fmt.Scan(&year)
	fmt.Println("   Old Result : ", record.Result)
	fmt.Print("   New Result : ")
	fmt.Scan(&result)

	fmt.Println("  ------------------------------")
	fmt.Println("   Patient ID       = ", record.Patient.ID, ">>>", patientID)
	fmt.Println("   Patient Name     = ", record.Patient.Name, ">>>", patientName)
	fmt.Println("   Patient Age      = ", record.Patient.Age, ">>>", patientAge)
	fmt.Println("   Patient Gender   = ", record.Patient.Gender, ">>>", patientGender)
	fmt.Println("   Package ID       = ", record.Package.ID, ">>>", packageID)
	fmt.Println("   Package Name     = ", record.Package.Name, ">>>", packageName)
	fmt.Println("   Package Category = ", record.Package.Category, ">>>", packageCategory)
	fmt.Println("   Package Price    = ", record.Package.Price, ">>>", packagePrice)
	fmt.Println("   Year             = ", record.Year, ">>>", year)
	fmt.Printf("   Date             =  %d-%d-%d >>> %d-%d-%d\n", record.Day, record.Month, record.Year, day, month, year)
	fmt.Println("   Result           = ", record.Result, ">>>", result)
	fmt.Println("  ------------------------------")
	fmt.Print("   Accept Changes? (yes or no) : ")
	stop = false
	for !stop {
		fmt.Scan(&accept)
		if accept == "yes" {
			records.Daftar[idx].Patient.ID = patientID
			records.Daftar[idx].Patient.Name = patientName
			records.Daftar[idx].Patient.Age = patientAge
			records.Daftar[idx].Patient.Gender = patientGender
			records.Daftar[idx].Package.ID = packageID
			records.Daftar[idx].Package.Name = packageName
			records.Daftar[idx].Package.Category = packageCategory
			records.Daftar[idx].Package.Price = packagePrice
			records.Daftar[idx].Year = year
			records.Daftar[idx].Month = month
			records.Daftar[idx].Day = day
			records.Daftar[idx].Result = result
			stop = true
		} else if accept == "no" {
			fmt.Println("   Reversing Changes...")
			stop = true
		} else {
			fmt.Print("   Invalid Input! Accept Changes? (yes or no) : ")
		}
	}
}

func record_delete(idx int) {
	var validAccept bool = false
	var accept string
	fmt.Print("   Are you sure to delete this record? (yes or no) : ")
	for !validAccept {
		fmt.Scan(&accept)
		if accept == "yes" {
			var i int
			for i = idx; i < records.N; i++ {
				temp := records.Daftar[i].ID
				records.Daftar[i] = records.Daftar[i+1]
				records.Daftar[i].ID = temp
			}
			records.Daftar[i].ID = 0
			records.Daftar[i].Patient.ID = 0
			records.Daftar[i].Patient.Name = ""
			records.Daftar[i].Patient.Age = 0
			records.Daftar[i].Patient.Gender = ""
			records.Daftar[i].Package.ID = 0
			records.Daftar[i].Package.Name = ""
			records.Daftar[i].Package.Category = ""
			records.Daftar[i].Package.Price = 0
			records.Daftar[i].Year = 0
			records.Daftar[i].Month = 0
			records.Daftar[i].Day = 0
			records.Daftar[i].Result = ""
			records.N--
			validAccept = true
		} else if accept == "no" {
			record_select(idx)
			validAccept = true
		} else {
			fmt.Print(" Invalid Choice! Are you sure to delete this record? (yes or no) : ")
		}
	}
}

func report_management() {
	clearScreen()
	header()
	bold("  ----------- REPORT -----------")
	var income, i, choice int
	var anything string
	var validChoice bool = false
	for i = 0; i < records.N; i++ {
		income += records.Daftar[i].Package.Price
	}
	fmt.Printf("   Total Income  :  Rp.%d\n", income)
	fmt.Println("   Total Patient : ", patients.N)
	fmt.Println("   Total Records : ", records.N)
	fmt.Println("  ------------------------------")
	fmt.Print("   ")
	underline("Menu Navigation :")
	fmt.Println("   1. Search Income By Date")
	fmt.Println("   0. Return")
	fmt.Print("   Choice : ")
	for !validChoice {
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var year, month, nRecord int
			income = 0
			nRecord = 0
			fmt.Print("   Enter Year (YYYY) : ")
			fmt.Scan(&year)
			fmt.Print("   Enter Month (MM) : ")
			fmt.Scan(&month)
			for i = 0; i < records.N; i++ {
				if records.Daftar[i].Month == month && records.Daftar[i].Year == year {
					income += records.Daftar[i].Package.Price
					nRecord++
				}
			}
			fmt.Println("  ------------------------------")
			fmt.Printf("   Reports On %d-%d : \n", month, year)
			fmt.Printf("   Total Income :  Rp.%d\n", income)
			fmt.Println("   Total Record : ", nRecord)
			fmt.Println("  ------------------------------")
			fmt.Print("   Enter Anything To Return : ")
			fmt.Scan(&anything)
			validChoice = true
			report_management()
		case 0:
			main_menu()
			validChoice = true
		default:
			fmt.Print("   Invalid Choice! Choice : ")
		}
	}
}

func main() {
	fmt.Println("   Starting Program...")
	err := loadArray()
	if err != nil {
		fmt.Println("Error loading data:", err)
	}

	main_menu()

	err = saveArray()
	if err != nil {
		fmt.Println("Error saving data:", err)
	}
	fmt.Println("   Data saved successfully.")
}

func saveArray() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	filePath := dir + "/data.json"
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(patients)
	if err != nil {
		return err
	}
	err = encoder.Encode(packages)
	if err != nil {
		return err
	}
	err = encoder.Encode(records)
	if err != nil {
		return err
	}

	return nil
}

func loadArray() error {
	file, err := os.Open("data.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Data file doesn't exist. Starting with empty data.")
			return nil
		}
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&patients)
	if err != nil {
		return err
	}
	err = decoder.Decode(&packages)
	if err != nil {
		return err
	}
	err = decoder.Decode(&records)
	if err != nil {
		return err
	}

	fmt.Println("   Data loaded successfully.")
	return nil
}
