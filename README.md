# Medical Check Up
## General
A Command Line Interface For My Final Project Of The Course Programming Algorithm
Made By :
- Ekmal Reyhan Tarihoran || 1305223079
- Yunita || 1305220067

This Program Is Built According To This:

**Description :**  The application is used to manage data of patients who do medical check up. The data processed is data
service packages as well as patient data that performs medical check up. Application users are hospital officers or
laboratory.

**Specification :**
- Users can add, change (edit), and delete service packages, patient data, recap of medical checkup results.
- Users can display income reports from medical check up services on a certain period.
- Users can search list of patient who choose a particular package, list of patient who did medical check up on a certain period, and search for patient by name.
- Users can display patient data sorted by medical check up time or service packages.

There are also some general specification for this project such as:
- Implement the use of array and type structs. Static array only, no dynamic array or slice.
- Implement the use of Sequential search and Binary search on the process of searching, editing, or deleting particular data.
- Implement the use of Selection sort and Insertion sort algorithm to sort certain data with different category when displaying data. Every category needs to be able to be sorted Ascending and Descending.
- It is not allowed to use Break statements (other than for a repeat-until loop) or continue.
- The use of global variables are only allowed on certain array that will be processed.

** NOTE: Some terminal might not be compatible with how i do transition between menu, try to run as admin or use any terminal that support ANSI escape code.

-------------------
## How The Program Works
**Package Management :**
- See All Package :
  - Select Package :
    - Edit
    - Delete
    - Search Record Of The Particular Package
  - Add Package :
  
      This menu is used to add package which needs a Name, Category (Basic/Standard/Advanced) and the Price for the package
  - Search Record By Package :
  
      Enter Package Name and it will return all the records of all the specified package, if no record was found, it will return a prompt saying no records were found.

**Patient Management :**
- See Patients :

  It will return a prompt to choose which sorting would you prefer or you can enter a patient ID directly which will open up the detail of said patients
  - Select Patient :
  
    This menu will show the detail of the patient and all the records that this patient have that is recorded on the program
      - Edit
      - Delete
      - See Patient Records Detail
      - Add Records To Patient
- Add Patient
- Search Patient :
  - By Patient Name
  - By Package Name
  - By Medical Checkup Time

**Record Management :**
- See All Records :

  This menu will show every records from the newest entry to the oldest, sadly there is no option for sorting. You can also enter the records id to see the detail of the specified record.
  - Select Record :
  
    This menu showed the detail of the records including the detail of the patient, which is linked by the patient id, you can also go to the patient detail page using one of the menu navigation.
    - Edit
    - Delete
    - See Patient Detail
- Add Record
- Search Record :
  - By Patient Name
  - By Package Name
  - By Medical Check Up Time

**Reports :**

This is the simplest menu of them all, this menu shows the total income from every records, or you can specify which month to see the reports of by using one of the menu.

**Quick-add Records :**

This menu is also simple, you can use it to quickly add a records by entering patient name, package name, and all the other things asked by the prompt, sadly you can only use existing patient/package since i havent implemented the feature to add a new patient/package. Also you need a valid date to be inputted like (3-3-2020 OR 17-5-2023).

-----
