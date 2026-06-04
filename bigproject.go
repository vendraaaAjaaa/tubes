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

// ============================================================
//  CONSTANTS & DATA STRUCTURES
// ============================================================

const MaxStudents = 100

// Student is the main data type (struct) representing
// a single prospective student. (Spec c — custom type)
type Student struct {
	ID     string
	Name   string
	Major  string
	Score  float64
	Status string // "Pending" | "Accepted" | "Rejected"
}

// ============================================================
//  GLOBAL VARIABLES
//  Spec h: hanya array data utama yang boleh global.
// ============================================================

var students [MaxStudents]Student // static array — main data store (Spec c)
var studentCount int              // number of active students

// ============================================================
//  SCREEN UTILITIES
// ============================================================

// clearScreen membersihkan tampilan terminal.
// I.S. : terminal may contain previous output
// F.S. : terminal is blank
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

// printDivider mencetak garis pemisah horizontal.
// I.S. : -
// F.S. : one line of "─" characters printed
func printDivider() {
	fmt.Println(strings.Repeat("─", 60))
}

// printHeader membersihkan layar lalu mencetak judul menu terpusat.
// I.S. : title is a non-empty string
// F.S. : screen cleared, title displayed between dividers
func printHeader(title string) {
	clearScreen()
	printDivider()
	pad := (60 - len(title)) / 2
	fmt.Printf("%s%s\n", strings.Repeat(" ", pad), title)
	printDivider()
}

// printStudentTable menampilkan data mahasiswa dalam format tabel.
// I.S. : arr[0..n-1] holds valid student data, n >= 0
// F.S. : table printed to stdout; arr and n unchanged
func printStudentTable(arr [MaxStudents]Student, n int) {
	if n == 0 {
		fmt.Println("  (No records to display)")
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

// ============================================================
//  INPUT UTILITIES
//  bufio.NewReader is created locally inside readLine — not global
//  to comply with Spec h. In interactive terminal mode (canonical
//  mode), the OS buffers input per line so ReadString('\n') always
//  reads exactly one complete line with no leftover bytes.
// ============================================================

// readLine membaca satu baris input lengkap termasuk spasi.
// I.S. : prompt displayed, user has not yet typed
// F.S. : returns trimmed string of what the user typed
func readLine(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

// readInt membaca satu baris input dan mengkonversinya ke integer.
// I.S. : prompt displayed
// F.S. : returns parsed int, or -99 if input is not a valid integer
func readInt(prompt string) int {
	raw := readLine(prompt)
	n, err := strconv.Atoi(raw)
	if err != nil {
		return -99 // sentinel → triggers "invalid option" branch
	}
	return n
}

// readFloat membaca satu baris input dan mengkonversinya ke float64.
// I.S. : prompt displayed
// F.S. : returns (float64, true) if valid, or (0, false) if not
func readFloat(prompt string) (float64, bool) {
	raw := readLine(prompt)
	f, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, false
	}
	return f, true
}

// pressEnterToContinue menghentikan eksekusi hingga user menekan Enter.
// I.S. : some output is displayed on screen
// F.S. : execution resumes after Enter is pressed
func pressEnterToContinue() {
	fmt.Print("\n  Press Enter to continue...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}

// ============================================================
//  SEARCH ALGORITHMS  (Spec d)
// ============================================================

// sequentialSearchByID melakukan pencarian linear pada students[0..studentCount-1].
// I.S. : students array may be in any order; id is the search target
// F.S. : returns index i where students[i].ID == id, or -1 if not found
// Complexity: O(n)
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

// binarySearchByName melakukan binary search pada array yang sudah terurut by Name.
// Pre-condition: arr[0..n-1] MUST be sorted ascending by Name.
// I.S. : arr is sorted ascending by Name, name is the search target
// F.S. : returns index mid where arr[mid].Name == name (case-insensitive),
//
//	or -1 if not found
//
// Complexity: O(log n)
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

// ============================================================
//  SORT ALGORITHMS  (Spec e)
// ============================================================

// selectionSortByScore mengurutkan arr[0..n-1] berdasarkan Score
// menggunakan algoritma Selection Sort.
// I.S. : arr is a copy of student data, n = number of active elements
// F.S. : arr sorted by Score ascending (if ascending==true) or descending
// Complexity: O(n²)
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

// shouldShiftName adalah fungsi pembantu komparator untuk insertionSortByName.
// I.S. : a and b are two Student elements to compare
// F.S. : returns true if element a should shift right past element b
func shouldShiftName(a, b Student, ascending bool) bool {
	if ascending {
		return strings.ToLower(a.Name) > strings.ToLower(b.Name)
	}
	return strings.ToLower(a.Name) < strings.ToLower(b.Name)
}

// insertionSortByName mengurutkan arr[0..n-1] berdasarkan Name
// menggunakan algoritma Insertion Sort.
// I.S. : arr is a copy of student data
// F.S. : arr sorted by Name ascending (if ascending==true) or descending
// Complexity: O(n²)
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

// shouldShiftMajor adalah fungsi pembantu komparator untuk insertionSortByMajor.
// I.S. : a and b are two Student elements to compare
// F.S. : returns true if element a should shift right past element b
func shouldShiftMajor(a, b Student, ascending bool) bool {
	if ascending {
		return strings.ToLower(a.Major) > strings.ToLower(b.Major)
	}
	return strings.ToLower(a.Major) < strings.ToLower(b.Major)
}

// insertionSortByMajor mengurutkan arr[0..n-1] berdasarkan Major
// menggunakan algoritma Insertion Sort.
// I.S. : arr is a copy of student data
// F.S. : arr sorted by Major ascending (if ascending==true) or descending
// Complexity: O(n²)
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

// ============================================================
//  MENU RENDERING
// ============================================================

func showMainMenu() {
	printHeader("STUDENT REGISTRATION APP")
	fmt.Println("  1. Student Management")
	fmt.Println("  2. Score & Evaluation")
	fmt.Println("  3. Search Student")
	fmt.Println("  4. View Reports")
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

// askSortOrder menampilkan pilihan urutan sort kepada user.
// I.S. : kategori sort sudah ditentukan
// F.S. : returns true for ascending, false for descending
func askSortOrder() bool {
	fmt.Println("\n  Sort Order:")
	fmt.Println("    1. Ascending  (A → Z / Low → High)")
	fmt.Println("    2. Descending (Z → A / High → Low)")
	choice := readInt("  Select: ")
	return choice == 1
}

// ============================================================
//  STUDENT MANAGEMENT  (Spec a & b)
// ============================================================

// addStudent menambahkan data mahasiswa baru ke array students.
// I.S. : studentCount < MaxStudents; all fields entered by user
// F.S. : students[studentCount] filled with new data, studentCount incremented by 1
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

// editStudent memperbarui Nama dan/atau Jurusan mahasiswa yang sudah ada.
// I.S. : studentCount > 0; student with given ID exists in array
// F.S. : students[idx].Name and/or students[idx].Major updated
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
	fmt.Println("  (Press Enter to keep current value)\n")

	newName := readLine("  New Full Name : ")
	if newName != "" {
		students[idx].Name = newName
	}

	newMajor := readLine("  New Major     : ")
	if newMajor != "" {
		students[idx].Major = newMajor
	}

	fmt.Printf("\n  [✓] Student ID '%s' updated successfully.\n", id)
	pressEnterToContinue()
}

// deleteStudent menghapus mahasiswa berdasarkan ID menggunakan teknik swap-and-shrink.
// I.S. : student with given ID exists in array; studentCount > 0
// F.S. : student removed, slot filled by last element, studentCount decremented by 1
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
	students[idx] = students[studentCount-1] // overwrite with last element
	students[studentCount-1] = Student{}     // clear last slot
	studentCount--

	fmt.Printf("\n  [✓] Student '%s' deleted successfully.\n", name)
	pressEnterToContinue()
}

// studentManagementMenu menjalankan loop sub-menu Manajemen Mahasiswa.
// I.S. : -
// F.S. : returns to main menu when user selects 0
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

// ============================================================
//  SCORE & EVALUATION  (Spec c)
// ============================================================

// evaluateStudent menetapkan nilai ujian dan status penerimaan mahasiswa.
// I.S. : student with given ID exists in array
// F.S. : students[idx].Score = score;
//
//	Status = "Accepted" if score >= 75, "Rejected" otherwise
func evaluateStudent() {
	printHeader("SCORE & EVALUATION")

	if studentCount == 0 {
		fmt.Println("\n  [!] No student records yet.")
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

	fmt.Printf("\n  [✓] Score %.2f recorded. Status → %s\n",
		score, students[idx].Status)
	pressEnterToContinue()
}

// ============================================================
//  SEARCH MENU  (Spec d)
// ============================================================

// doSequentialSearch mendemonstrasikan Sequential Search berdasarkan ID.
// I.S. : studentCount > 0
// F.S. : search result displayed; students array unchanged
func doSequentialSearch() {
	printHeader("SEQUENTIAL SEARCH — by ID")

	if studentCount == 0 {
		fmt.Println("\n  [!] No student records yet.")
		pressEnterToContinue()
		return
	}

	fmt.Println("  Student List (original order, unsorted):")
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

// doBinarySearch mendemonstrasikan Binary Search berdasarkan Nama.
// Mengurutkan salinan array terlebih dahulu sebelum binary search dilakukan.
// I.S. : studentCount > 0
// F.S. : sorted copy displayed, binary search result shown;
//
//	original students array unchanged
func doBinarySearch() {
	printHeader("BINARY SEARCH — by Name")

	if studentCount == 0 {
		fmt.Println("\n  [!] No student records yet.")
		pressEnterToContinue()
		return
	}

	name := readLine("  Enter Student Name to search: ")

	// Step 1: copy global array into local copy
	var arrSorted [MaxStudents]Student
	for i := 0; i < studentCount; i++ {
		arrSorted[i] = students[i]
	}

	// Step 2: sort copy by Name ascending (prerequisite for binary search)
	insertionSortByName(&arrSorted, studentCount, true)

	fmt.Println("\n  Array sorted by Name (Ascending) — prerequisite for Binary Search:")
	fmt.Println()
	printStudentTable(arrSorted, studentCount)

	// Step 3: apply binary search on the sorted copy
	fmt.Println("\n  Menjalankan Binary Search...")
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

// searchMenu menjalankan loop sub-menu Pencarian.
// I.S. : -
// F.S. : returns to main menu when user selects 0
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

// ============================================================
//  REPORTS  (Spec b & e)
// ============================================================

// viewByMajorAndStatus menampilkan mahasiswa berdasarkan jurusan dan status.
// I.S. : major and statusFilter entered by user
// F.S. : filtered subset of students displayed;
//
//	no continue used (Spec g) — uses nested if instead
func viewByMajorAndStatus() {
	printHeader("VIEW BY MAJOR & ADMISSION STATUS")

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

// viewSortedByScore menampilkan semua mahasiswa terurut by Score
// menggunakan Selection Sort dengan pilihan ascending/descending.
// I.S. : studentCount > 0
// F.S. : sorted copy displayed; original students array unchanged
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

// viewSortedByName menampilkan semua mahasiswa terurut by Name
// menggunakan Insertion Sort dengan pilihan ascending/descending.
// I.S. : studentCount > 0
// F.S. : sorted copy displayed; original students array unchanged
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

// viewSortedByMajor menampilkan semua mahasiswa terurut by Major
// menggunakan Insertion Sort dengan pilihan ascending/descending.
// I.S. : studentCount > 0
// F.S. : sorted copy displayed; original students array unchanged
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

// reportsMenu menjalankan loop sub-menu Laporan.
// I.S. : -
// F.S. : returns to main menu when user selects 0
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

// ============================================================
//  ENTRY POINT
// ============================================================

// main adalah titik masuk program.
// I.S. : program starts, studentCount = 0
// F.S. : program exits when user selects 0 from main menu
func main() {
	for {
		showMainMenu()
		choice := readInt("  Select option: ")
		if choice == 1 {
			studentManagementMenu()
		} else if choice == 2 {
			evaluateStudent()
		} else if choice == 3 {
			searchMenu()
		} else if choice == 4 {
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
