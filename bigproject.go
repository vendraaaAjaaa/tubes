package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const MaxStudents = 100
const MaxDepts = 50

type Student struct {
	ID     string
	Name   string
	Major  string
	Score  float64
	Status string
}

type Department struct {
	ID   string
	Name string
}

var students [MaxStudents]Student
var studentCount int
var departments [MaxDepts]Department
var deptCount int

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func printDivider() {
	fmt.Println(strings.Repeat("─", 60))
}

func printHeader(title string) {
	clearScreen()
	printDivider()
	pad := (60 - len(title)) / 2
	fmt.Printf("%s%s\n", strings.Repeat(" ", pad), title)
	printDivider()
}

func printStudentTable(arr [MaxStudents]Student, n int) {
	if n == 0 {
		fmt.Println("  (No student records to display)")
		return
	}
	fmt.Printf("  %-8s  %-22s  %-15s  %7s  %-10s\n",
		"ID", "Name", "Major", "Score", "Status")
	fmt.Println("  " + strings.Repeat("-", 69))
	for i := 0; i < n; i++ {
		s := arr[i]
		scoreStr := "      —"
		if s.Score > 0 {
			scoreStr = fmt.Sprintf("%7.2f", s.Score)
		}
		fmt.Printf("  %-8s  %-22s  %-15s  %7s  %-10s\n",
			s.ID, s.Name, s.Major, scoreStr, s.Status)
	}
}

func printDeptTable(arr [MaxDepts]Department, n int) {
	if n == 0 {
		fmt.Println("  (No departments registered yet)")
		return
	}
	fmt.Printf("  %-8s  %-30s\n", "ID", "Department Name")
	fmt.Println("  " + strings.Repeat("-", 40))
	for i := 0; i < n; i++ {
		fmt.Printf("  %-8s  %-30s\n", arr[i].ID, arr[i].Name)
	}
}

func readLine(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func readInt(prompt string) int {
	raw := readLine(prompt)
	n, err := strconv.Atoi(raw)
	if err != nil {
		return -99
	}
	return n
}

func readFloat(prompt string) (float64, bool) {
	raw := readLine(prompt)
	f, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, false
	}
	return f, true
}

func pressEnterToContinue() {
	fmt.Print("\n  Press Enter to continue...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}

func sequentialSearchByID(id string) int {
	found := -1
	i := 0
	for i < studentCount && found == -1 {
		if students[i].ID == id {
			found = i
		}
		i++
	}
	return found
}

func sequentialSearchDeptByID(id string) int {
	found := -1
	i := 0
	for i < deptCount && found == -1 {
		if departments[i].ID == id {
			found = i
		}
		i++
	}
	return found
}

func binarySearchByName(arr [MaxStudents]Student, n int, name string) int {
	lo := 0
	hi := n - 1
	result := -1
	target := strings.ToLower(name)
	for lo <= hi && result == -1 {
		mid := (lo + hi) / 2
		midName := strings.ToLower(arr[mid].Name)
		if midName == target {
			result = mid
		} else if midName < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return result
}

func selectionSortByScore(arr *[MaxStudents]Student, n int, ascending bool) {
	for i := 0; i < n-1; i++ {
		extremeIdx := i
		for j := i + 1; j < n; j++ {
			if ascending {
				if arr[j].Score < arr[extremeIdx].Score {
					extremeIdx = j
				}
			} else {
				if arr[j].Score > arr[extremeIdx].Score {
					extremeIdx = j
				}
			}
		}
		if extremeIdx != i {
			arr[i], arr[extremeIdx] = arr[extremeIdx], arr[i]
		}
	}
}

func shouldShiftName(a, b Student, ascending bool) bool {
	if ascending {
		return strings.ToLower(a.Name) > strings.ToLower(b.Name)
	}
	return strings.ToLower(a.Name) < strings.ToLower(b.Name)
}

func insertionSortByName(arr *[MaxStudents]Student, n int, ascending bool) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && shouldShiftName(arr[j], key, ascending) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func shouldShiftMajor(a, b Student, ascending bool) bool {
	if ascending {
		return strings.ToLower(a.Major) > strings.ToLower(b.Major)
	}
	return strings.ToLower(a.Major) < strings.ToLower(b.Major)
}

func insertionSortByMajor(arr *[MaxStudents]Student, n int, ascending bool) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && shouldShiftMajor(arr[j], key, ascending) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func showMainMenu() {
	printHeader("STUDENT REGISTRATION APP")
	fmt.Println("  1. Student Management")
	fmt.Println("  2. Department Management")
	fmt.Println("  3. Score & Evaluation")
	fmt.Println("  4. Search Student")
	fmt.Println("  5. View Reports")
	fmt.Println("  0. Exit")
	printDivider()
}

func showStudentManagementMenu() {
	printHeader("STUDENT MANAGEMENT")
	fmt.Println("  1. Add Student")
	fmt.Println("  2. Edit Student")
	fmt.Println("  3. Delete Student")
	fmt.Println("  0. Back")
	printDivider()
}

func showDeptManagementMenu() {
	printHeader("DEPARTMENT MANAGEMENT")
	fmt.Println("  1. Add Department")
	fmt.Println("  2. Edit Department")
	fmt.Println("  3. Delete Department")
	fmt.Println("  0. Back")
	printDivider()
}

func showScoreMenu() {
	printHeader("SCORE & EVALUATION")
	fmt.Println("  1. Add / Edit Score")
	fmt.Println("  2. Delete Score  (reset to Pending)")
	fmt.Println("  0. Back")
	printDivider()
}

func showSearchMenu() {
	printHeader("SEARCH STUDENT")
	fmt.Println("  1. Sequential Search  (by ID — works on unsorted data)")
	fmt.Println("  2. Binary Search      (by Name — sorts data first)")
	fmt.Println("  0. Back")
	printDivider()
}

func showReportsMenu() {
	printHeader("VIEW REPORTS")
	fmt.Println("  1. View by Major & Admission Status")
	fmt.Println("  2. Sort by Score   (Selection Sort)")
	fmt.Println("  3. Sort by Name    (Insertion Sort)")
	fmt.Println("  4. Sort by Major   (Insertion Sort)")
	fmt.Println("  0. Back")
	printDivider()
}

func askSortOrder() bool {
	fmt.Println("\n  Sort Order:")
	fmt.Println("    1. Ascending  (A → Z / Low → High)")
	fmt.Println("    2. Descending (Z → A / High → Low)")
	choice := readInt("  Select: ")
	return choice == 1
}

func addStudent() {
	printHeader("ADD STUDENT")

	if studentCount >= MaxStudents {
		fmt.Printf("\n  [!] Maximum capacity (%d students) reached.\n", MaxStudents)
		pressEnterToContinue()
		return
	}

	id := readLine("  Student ID   : ")
	if id == "" {
		fmt.Println("\n  [!] ID cannot be empty.")
		pressEnterToContinue()
		return
	}
	if sequentialSearchByID(id) != -1 {
		fmt.Printf("\n  [!] ID '%s' already exists — use Edit Student instead.\n", id)
		pressEnterToContinue()
		return
	}

	name := readLine("  Full Name    : ")

	if deptCount > 0 {
		fmt.Println("\n  Available Departments:")
		fmt.Println()
		printDeptTable(departments, deptCount)
		fmt.Println()
	}
	major := readLine("  Major        : ")

	students[studentCount] = Student{
		ID:     id,
		Name:   name,
		Major:  major,
		Status: "Pending",
	}
	studentCount++

	fmt.Printf("\n  [✓] Student '%s' added successfully.\n", name)
	pressEnterToContinue()
}

func editStudent() {
	printHeader("EDIT STUDENT")

	if studentCount == 0 {
		fmt.Println("\n  [!] No student records yet.")
		pressEnterToContinue()
		return
	}

	fmt.Println("  Current Student List:")
	fmt.Println()
	printStudentTable(students, studentCount)
	fmt.Println()

	id := readLine("  Enter Student ID to edit: ")
	idx := sequentialSearchByID(id)
	if idx == -1 {
		fmt.Printf("\n  [!] No student found with ID '%s'.\n", id)
		pressEnterToContinue()
		return
	}

	fmt.Printf("\n  Selected → %s | Major: %s\n", students[idx].Name, students[idx].Major)
	fmt.Printf("  (Press Enter to keep the current value)\n")

	newName := readLine("  New Full Name : ")
	if newName != "" {
		students[idx].Name = newName
	}

	if deptCount > 0 {
		fmt.Println("\n  Available Departments:")
		fmt.Println()
		printDeptTable(departments, deptCount)
		fmt.Println()
	}
	newMajor := readLine("  New Major     : ")
	if newMajor != "" {
		students[idx].Major = newMajor
	}

	fmt.Printf("\n  [✓] Student ID '%s' updated successfully.\n", id)
	pressEnterToContinue()
}

func deleteStudent() {
	printHeader("DELETE STUDENT")

	if studentCount == 0 {
		fmt.Println("\n  [!] No student records yet.")
		pressEnterToContinue()
		return
	}

	fmt.Println("  Current Student List:")
	fmt.Println()
	printStudentTable(students, studentCount)
	fmt.Println()

	id := readLine("  Enter Student ID to delete: ")
	idx := sequentialSearchByID(id)
	if idx == -1 {
		fmt.Printf("\n  [!] No student found with ID '%s'.\n", id)
		pressEnterToContinue()
		return
	}

	name := students[idx].Name
	students[idx] = students[studentCount-1]
	students[studentCount-1] = Student{}
	studentCount--

	fmt.Printf("\n  [✓] Student '%s' deleted successfully.\n", name)
	pressEnterToContinue()
}

func studentManagementMenu() {
	for {
		showStudentManagementMenu()
		choice := readInt("  Select option: ")
		if choice == 1 {
			addStudent()
		} else if choice == 2 {
			editStudent()
		} else if choice == 3 {
			deleteStudent()
		} else if choice == 0 {
			return
		} else {
			fmt.Println("\n  [!] Invalid option.")
			pressEnterToContinue()
		}
	}
}

func addDepartment() {
	printHeader("ADD DEPARTMENT")

	if deptCount >= MaxDepts {
		fmt.Printf("\n  [!] Maximum capacity (%d departments) reached.\n", MaxDepts)
		pressEnterToContinue()
		return
	}

	id := readLine("  Department ID   : ")
	if id == "" {
		fmt.Println("\n  [!] ID cannot be empty.")
		pressEnterToContinue()
		return
	}
	if sequentialSearchDeptByID(id) != -1 {
		fmt.Printf("\n  [!] Department ID '%s' already exists.\n", id)
		pressEnterToContinue()
		return
	}

	name := readLine("  Department Name : ")
	if name == "" {
		fmt.Println("\n  [!] Department name cannot be empty.")
		pressEnterToContinue()
		return
	}

	departments[deptCount] = Department{ID: id, Name: name}
	deptCount++

	fmt.Printf("\n  [✓] Department '%s' added successfully.\n", name)
	pressEnterToContinue()
}

func editDepartment() {
	printHeader("EDIT DEPARTMENT")

	if deptCount == 0 {
		fmt.Println("\n  [!] No departments registered yet.")
		pressEnterToContinue()
		return
	}

	fmt.Println("  Current Department List:")
	fmt.Println()
	printDeptTable(departments, deptCount)
	fmt.Println()

	id := readLine("  Enter Department ID to edit: ")
	idx := sequentialSearchDeptByID(id)
	if idx == -1 {
		fmt.Printf("\n  [!] No department found with ID '%s'.\n", id)
		pressEnterToContinue()
		return
	}

	fmt.Printf("\n  Selected → %s\n", departments[idx].Name)
	fmt.Printf("  (Press Enter to keep the current value)\n")

	newName := readLine("  New Department Name : ")
	if newName != "" {
		departments[idx].Name = newName
	}

	fmt.Printf("\n  [✓] Department ID '%s' updated successfully.\n", id)
	pressEnterToContinue()
}

func deleteDepartment() {
	printHeader("DELETE DEPARTMENT")

	if deptCount == 0 {
		fmt.Println("\n  [!] No departments registered yet.")
		pressEnterToContinue()
		return
	}

	fmt.Println("  Current Department List:")
	fmt.Println()
	printDeptTable(departments, deptCount)
	fmt.Println()

	id := readLine("  Enter Department ID to delete: ")
	idx := sequentialSearchDeptByID(id)
	if idx == -1 {
		fmt.Printf("\n  [!] No department found with ID '%s'.\n", id)
		pressEnterToContinue()
		return
	}

	name := departments[idx].Name
	departments[idx] = departments[deptCount-1]
	departments[deptCount-1] = Department{}
	deptCount--

	fmt.Printf("\n  [✓] Department '%s' deleted successfully.\n", name)
	pressEnterToContinue()
}

func departmentManagementMenu() {
	for {
		showDeptManagementMenu()
		choice := readInt("  Select option: ")
		if choice == 1 {
			addDepartment()
		} else if choice == 2 {
			editDepartment()
		} else if choice == 3 {
			deleteDepartment()
		} else if choice == 0 {
			return
		} else {
			fmt.Println("\n  [!] Invalid option.")
			pressEnterToContinue()
		}
	}
}

func addEditScore() {
	printHeader("ADD / EDIT SCORE")

	if studentCount == 0 {
		fmt.Println("\n  [!] No student records found.")
		pressEnterToContinue()
		return
	}

	fmt.Println("  Current Student List:")
	fmt.Println()
	printStudentTable(students, studentCount)
	fmt.Println()

	id := readLine("  Enter Student ID: ")
	idx := sequentialSearchByID(id)
	if idx == -1 {
		fmt.Printf("\n  [!] No student found with ID '%s'.\n", id)
		pressEnterToContinue()
		return
	}

	fmt.Printf("\n  Student : %s | Major: %s\n\n", students[idx].Name, students[idx].Major)

	score, ok := readFloat("  Test Score (0 – 100): ")
	if !ok || score < 0 || score > 100 {
		fmt.Println("\n  [!] Invalid score. Must be a number between 0 and 100.")
		pressEnterToContinue()
		return
	}

	students[idx].Score = score
	if score >= 75 {
		students[idx].Status = "Accepted"
	} else {
		students[idx].Status = "Rejected"
	}

	fmt.Printf("\n  [✓] Score %.2f recorded. Status → %s\n", score, students[idx].Status)
	pressEnterToContinue()
}

func deleteScore() {
	printHeader("DELETE SCORE")

	if studentCount == 0 {
		fmt.Println("\n  [!] No student records found.")
		pressEnterToContinue()
		return
	}

	fmt.Println("  Current Student List:")
	fmt.Println()
	printStudentTable(students, studentCount)
	fmt.Println()

	id := readLine("  Enter Student ID to delete score: ")
	idx := sequentialSearchByID(id)
	if idx == -1 {
		fmt.Printf("\n  [!] No student found with ID '%s'.\n", id)
		pressEnterToContinue()
		return
	}

	if students[idx].Score == 0 {
		fmt.Printf("\n  [!] Student '%s' has no score recorded yet.\n", students[idx].Name)
		pressEnterToContinue()
		return
	}

	students[idx].Score = 0
	students[idx].Status = "Pending"

	fmt.Printf("\n  [✓] Score deleted. '%s' status reset to Pending.\n", students[idx].Name)
	pressEnterToContinue()
}

func scoreMenu() {
	for {
		showScoreMenu()
		choice := readInt("  Select option: ")
		if choice == 1 {
			addEditScore()
		} else if choice == 2 {
			deleteScore()
		} else if choice == 0 {
			return
		} else {
			fmt.Println("\n  [!] Invalid option.")
			pressEnterToContinue()
		}
	}
}

func doSequentialSearch() {
	printHeader("SEQUENTIAL SEARCH — by ID")

	if studentCount == 0 {
		fmt.Println("\n  [!] No student records found.")
		pressEnterToContinue()
		return
	}

	fmt.Println("  Student List (original order — unsorted):")
	fmt.Println()
	printStudentTable(students, studentCount)
	fmt.Println()

	id := readLine("  Enter Student ID to search: ")
	fmt.Println("\n  Scanning array sequentially from index 0...")
	fmt.Println()

	idx := sequentialSearchByID(id)

	if idx == -1 {
		fmt.Printf("  [✗] Student with ID '%s' not found.\n", id)
	} else {
		fmt.Printf("  [✓] Found at index [%d]:\n\n", idx)
		var result [MaxStudents]Student
		result[0] = students[idx]
		printStudentTable(result, 1)
	}
	pressEnterToContinue()
}

func doBinarySearch() {
	printHeader("BINARY SEARCH — by Name")

	if studentCount == 0 {
		fmt.Println("\n  [!] No student records found.")
		pressEnterToContinue()
		return
	}

	name := readLine("  Enter Student Name to search: ")

	var arrSorted [MaxStudents]Student
	for i := 0; i < studentCount; i++ {
		arrSorted[i] = students[i]
	}

	insertionSortByName(&arrSorted, studentCount, true)

	fmt.Println("\n  Array sorted by Name (Ascending) — prerequisite for Binary Search:")
	fmt.Println()
	printStudentTable(arrSorted, studentCount)

	fmt.Println("\n  Applying Binary Search...")
	fmt.Println()
	idx := binarySearchByName(arrSorted, studentCount, name)

	if idx == -1 {
		fmt.Printf("  [✗] Student with Name '%s' not found.\n", name)
	} else {
		fmt.Printf("  [✓] Found at sorted index [%d]:\n\n", idx)
		var result [MaxStudents]Student
		result[0] = arrSorted[idx]
		printStudentTable(result, 1)
	}
	pressEnterToContinue()
}

func searchMenu() {
	for {
		showSearchMenu()
		choice := readInt("  Select option: ")
		if choice == 1 {
			doSequentialSearch()
		} else if choice == 2 {
			doBinarySearch()
		} else if choice == 0 {
			return
		} else {
			fmt.Println("\n  [!] Invalid option.")
			pressEnterToContinue()
		}
	}
}

func viewByMajorAndStatus() {
	printHeader("VIEW BY MAJOR & ADMISSION STATUS")

	if deptCount > 0 {
		fmt.Println("  Registered Departments:")
		fmt.Println()
		printDeptTable(departments, deptCount)
		fmt.Println()
	}

	major := readLine("  Enter Major Name: ")
	statusFilter := strings.ToLower(readLine("  Status Filter (Accepted / Rejected / All): "))

	var filtered [MaxStudents]Student
	filteredCount := 0

	for i := 0; i < studentCount; i++ {
		if strings.EqualFold(students[i].Major, major) {
			if statusFilter == "all" || strings.EqualFold(students[i].Status, statusFilter) {
				filtered[filteredCount] = students[i]
				filteredCount++
			}
		}
	}

	fmt.Printf("\n  Major: %s | Filter: %s\n\n", major, statusFilter)
	printStudentTable(filtered, filteredCount)
	pressEnterToContinue()
}

func viewSortedByScore() {
	printHeader("SORT BY SCORE — Selection Sort")

	ascending := askSortOrder()
	order := "Descending"
	if ascending {
		order = "Ascending"
	}

	var arrCopy [MaxStudents]Student
	for i := 0; i < studentCount; i++ {
		arrCopy[i] = students[i]
	}
	selectionSortByScore(&arrCopy, studentCount, ascending)

	fmt.Printf("\n  Students sorted by Score (%s):\n\n", order)
	printStudentTable(arrCopy, studentCount)
	pressEnterToContinue()
}

func viewSortedByName() {
	printHeader("SORT BY NAME — Insertion Sort")

	ascending := askSortOrder()
	order := "Descending"
	if ascending {
		order = "Ascending"
	}

	var arrCopy [MaxStudents]Student
	for i := 0; i < studentCount; i++ {
		arrCopy[i] = students[i]
	}
	insertionSortByName(&arrCopy, studentCount, ascending)

	fmt.Printf("\n  Students sorted by Name (%s):\n\n", order)
	printStudentTable(arrCopy, studentCount)
	pressEnterToContinue()
}

func viewSortedByMajor() {
	printHeader("SORT BY MAJOR — Insertion Sort")

	ascending := askSortOrder()
	order := "Descending"
	if ascending {
		order = "Ascending"
	}

	var arrCopy [MaxStudents]Student
	for i := 0; i < studentCount; i++ {
		arrCopy[i] = students[i]
	}
	insertionSortByMajor(&arrCopy, studentCount, ascending)

	fmt.Printf("\n  Students sorted by Major (%s):\n\n", order)
	printStudentTable(arrCopy, studentCount)
	pressEnterToContinue()
}

func reportsMenu() {
	for {
		showReportsMenu()
		choice := readInt("  Select option: ")
		if choice == 1 {
			viewByMajorAndStatus()
		} else if choice == 2 {
			viewSortedByScore()
		} else if choice == 3 {
			viewSortedByName()
		} else if choice == 4 {
			viewSortedByMajor()
		} else if choice == 0 {
			return
		} else {
			fmt.Println("\n  [!] Invalid option.")
			pressEnterToContinue()
		}
	}
}

func main() {
	for {
		showMainMenu()
		choice := readInt("  Select option: ")
		if choice == 1 {
			studentManagementMenu()
		} else if choice == 2 {
			departmentManagementMenu()
		} else if choice == 3 {
			scoreMenu()
		} else if choice == 4 {
			searchMenu()
		} else if choice == 5 {
			reportsMenu()
		} else if choice == 0 {
			clearScreen()
			printDivider()
			fmt.Println("  Goodbye! Thank you for using the Student Registration App.")
			printDivider()
			return
		} else {
			fmt.Println("\n  [!] Invalid option.")
			pressEnterToContinue()
		}
	}
}
