package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// ============================================================
//  KONSTANTA & TIPE BENTUKAN  (General Spec c)
// ============================================================

const MaxStudents = 100
const MaxDepts    = 50

// Student : Tipe bentukan untuk data calon mahasiswa
type Student struct {
	ID     string
	Name   string
	Major  string
	Score  float64
	Status string // "Pending" | "Accepted" | "Rejected"
}

// Department : Tipe bentukan untuk data jurusan universitas
type Department struct {
	ID   string
	Name string
}

// ============================================================
//  VARIABEL GLOBAL  (General Spec h: hanya array data utama)
// ============================================================

var students     [MaxStudents]Student // array statis mahasiswa
var studentCount int                  // jumlah mahasiswa aktif
var departments  [MaxDepts]Department // array statis jurusan
var deptCount    int                  // jumlah jurusan terdaftar

// ============================================================
//  UTILITAS LAYAR
// ============================================================

// clearScreen : Membersihkan tampilan terminal
// I.S. : Terminal mungkin menampilkan output sebelumnya
// F.S. : Terminal bersih dan kosong
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// printDivider : Mencetak garis pemisah sepanjang 60 karakter
// I.S. : -
// F.S. : Satu baris karakter "─" dicetak ke stdout
func printDivider() {
	fmt.Println(strings.Repeat("─", 60))
}

// printHeader : Membersihkan layar dan menampilkan judul section
// I.S. : title adalah string non-kosong
// F.S. : Layar bersih, title tampil terpusat di antara dua garis pemisah
func printHeader(title string) {
	clearScreen()
	printDivider()
	pad := (60 - len(title)) / 2
	fmt.Printf("%s%s\n", strings.Repeat(" ", pad), title)
	printDivider()
}

// printStudentTable : Menampilkan data mahasiswa dalam format tabel
// I.S. : arr[0..n-1] berisi data mahasiswa valid; n >= 0
// F.S. : Tabel dicetak ke stdout; arr dan n tidak berubah
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
		scoreStr := "      -"
		if s.Score > 0 {
			scoreStr = fmt.Sprintf("%7.2f", s.Score)
		}
		fmt.Printf("  %-8s  %-22s  %-15s  %7s  %-10s\n",
			s.ID, s.Name, s.Major, scoreStr, s.Status)
	}
}

// printDeptTable : Menampilkan data jurusan dalam format tabel
// I.S. : arr[0..n-1] berisi data jurusan valid; n >= 0
// F.S. : Tabel dicetak ke stdout; arr dan n tidak berubah
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

// ============================================================
//  UTILITAS INPUT
//  - Tidak menggunakan strconv; menggunakan fmt.Sscan
//  - bufio.Scanner dibuat lokal di readLine (General Spec h)
//  - scanner.Scan() mengembalikan bool (bisa diabaikan tanpa _)
// ============================================================

// readLine : Membaca satu baris input lengkap termasuk spasi
// I.S. : prompt ditampilkan; user belum mengetik
// F.S. : Mengembalikan string hasil input yang sudah di-trim
func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// readInt : Membaca satu baris input dan mengkonversi ke integer
// I.S. : prompt ditampilkan
// F.S. : Mengembalikan integer dari input; 0 jika input tidak valid
func readInt(prompt string) int {
	var n int
	fmt.Sscan(readLine(prompt), &n)
	return n
}

// readFloat : Membaca satu baris input dan mengkonversi ke float64
// I.S. : prompt ditampilkan
// F.S. : Mengembalikan float64 dari input; 0 jika input tidak valid
func readFloat(prompt string) float64 {
	var f float64
	fmt.Sscan(readLine(prompt), &f)
	return f
}

// pressEnterToContinue : Menahan eksekusi sampai user menekan Enter
// I.S. : Ada output di layar yang perlu dibaca user
// F.S. : Eksekusi dilanjutkan setelah Enter ditekan
func pressEnterToContinue() {
	fmt.Print("\n  Press Enter to continue...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}

// ============================================================
//  ALGORITMA PENCARIAN  (General Spec d)
// ============================================================

// sequentialSearchByID : Pencarian linear pada students[0..studentCount-1]
// I.S. : Array students boleh dalam urutan apapun; id = target pencarian
// F.S. : Mengembalikan index i jika students[i].ID == id, atau -1
// Kompleksitas: O(n)
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

// sequentialSearchDeptByID : Pencarian linear pada departments[0..deptCount-1]
// I.S. : Array departments boleh dalam urutan apapun; id = target pencarian
// F.S. : Mengembalikan index i jika departments[i].ID == id, atau -1
// Kompleksitas: O(n)
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

// binarySearchByName : Binary search pada array mahasiswa terurut by Name
// Pre-kondisi : arr[0..n-1] HARUS sudah terurut ascending by Name
// I.S. : arr terurut ascending by Name; name = target pencarian
// F.S. : Mengembalikan index jika arr[idx].Name == name, atau -1
// Kompleksitas: O(log n)
func binarySearchByName(arr [MaxStudents]Student, n int, name string) int {
	lo     := 0
	hi     := n - 1
	result := -1
	target := strings.ToLower(name)
	for lo <= hi && result == -1 {
		mid     := (lo + hi) / 2
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
//  ALGORITMA PENGURUTAN  (General Spec e)
// ============================================================

// selectionSortByScore : Mengurutkan arr[0..n-1] by Score (Selection Sort)
// I.S. : arr adalah salinan data mahasiswa; n = jumlah elemen aktif
// F.S. : arr terurut by Score ascending jika ascending=true, descending jika false
// Kompleksitas: O(n²)
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
			temp              := arr[i]
			arr[i]             = arr[extremeIdx]
			arr[extremeIdx]    = temp
		}
	}
}

// shouldShiftName : Fungsi komparator pembantu untuk insertionSortByName
// I.S. : a dan b adalah dua elemen Student yang akan dibandingkan
// F.S. : Mengembalikan true jika elemen a harus digeser melewati elemen b
func shouldShiftName(a, b Student, ascending bool) bool {
	if ascending {
		return strings.ToLower(a.Name) > strings.ToLower(b.Name)
	}
	return strings.ToLower(a.Name) < strings.ToLower(b.Name)
}

// insertionSortByName : Mengurutkan arr[0..n-1] by Name (Insertion Sort)
// I.S. : arr adalah salinan data mahasiswa
// F.S. : arr terurut by Name ascending jika ascending=true, descending jika false
// Kompleksitas: O(n²)
func insertionSortByName(arr *[MaxStudents]Student, n int, ascending bool) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j   := i - 1
		for j >= 0 && shouldShiftName(arr[j], key, ascending) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// shouldShiftMajor : Fungsi komparator pembantu untuk insertionSortByMajor
// I.S. : a dan b adalah dua elemen Student yang akan dibandingkan
// F.S. : Mengembalikan true jika elemen a harus digeser melewati elemen b
func shouldShiftMajor(a, b Student, ascending bool) bool {
	if ascending {
		return strings.ToLower(a.Major) > strings.ToLower(b.Major)
	}
	return strings.ToLower(a.Major) < strings.ToLower(b.Major)
}

// insertionSortByMajor : Mengurutkan arr[0..n-1] by Major (Insertion Sort)
// I.S. : arr adalah salinan data mahasiswa
// F.S. : arr terurut by Major ascending jika ascending=true, descending jika false
// Kompleksitas: O(n²)
func insertionSortByMajor(arr *[MaxStudents]Student, n int, ascending bool) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j   := i - 1
		for j >= 0 && shouldShiftMajor(arr[j], key, ascending) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// ============================================================
//  TAMPILAN MENU
// ============================================================

// showMainMenu : Menampilkan menu utama aplikasi
// I.S. : -
// F.S. : Menu utama dengan 5 pilihan ditampilkan ke layar
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

// showStudentManagementMenu : Menampilkan sub-menu Student Management
// I.S. : -
// F.S. : Sub-menu dengan pilihan Add/Edit/Delete ditampilkan
func showStudentManagementMenu() {
	printHeader("STUDENT MANAGEMENT")
	fmt.Println("  1. Add Student")
	fmt.Println("  2. Edit Student")
	fmt.Println("  3. Delete Student")
	fmt.Println("  0. Back")
	printDivider()
}

// showDeptManagementMenu : Menampilkan sub-menu Department Management
// I.S. : -
// F.S. : Sub-menu dengan pilihan Add/Edit/Delete ditampilkan
func showDeptManagementMenu() {
	printHeader("DEPARTMENT MANAGEMENT")
	fmt.Println("  1. Add Department")
	fmt.Println("  2. Edit Department")
	fmt.Println("  3. Delete Department")
	fmt.Println("  0. Back")
	printDivider()
}

// showScoreMenu : Menampilkan sub-menu Score & Evaluation
// I.S. : -
// F.S. : Sub-menu dengan pilihan Add/Edit Score dan Delete Score ditampilkan
func showScoreMenu() {
	printHeader("SCORE & EVALUATION")
	fmt.Println("  1. Add / Edit Score")
	fmt.Println("  2. Delete Score  (reset to Pending)")
	fmt.Println("  0. Back")
	printDivider()
}

// showSearchMenu : Menampilkan sub-menu Search Student
// I.S. : -
// F.S. : Sub-menu dengan pilihan Sequential dan Binary Search ditampilkan
func showSearchMenu() {
	printHeader("SEARCH STUDENT")
	fmt.Println("  1. Sequential Search  (by ID)")
	fmt.Println("  2. Binary Search      (by Name)")
	fmt.Println("  0. Back")
	printDivider()
}

// showReportsMenu : Menampilkan sub-menu View Reports
// I.S. : -
// F.S. : Sub-menu dengan pilihan tampilan dan pengurutan ditampilkan
func showReportsMenu() {
	printHeader("VIEW REPORTS")
	fmt.Println("  1. View by Major & Admission Status")
	fmt.Println("  2. Sort by Score   (Selection Sort)")
	fmt.Println("  3. Sort by Name    (Insertion Sort)")
	fmt.Println("  4. Sort by Major   (Insertion Sort)")
	fmt.Println("  0. Back")
	printDivider()
}

// askSortOrder : Menampilkan pilihan urutan sort kepada user
// I.S. : Kategori sort sudah ditentukan oleh pemanggil
// F.S. : Mengembalikan true untuk ascending, false untuk descending
func askSortOrder() bool {
	fmt.Println("\n  Sort Order:")
	fmt.Println("    1. Ascending  (A-Z / Low-High)")
	fmt.Println("    2. Descending (Z-A / High-Low)")
	choice := readInt("  Select: ")
	return choice == 1
}

// ============================================================
//  STUDENT MANAGEMENT  (App Spec a)
// ============================================================

// addStudent : Menambah data mahasiswa baru ke array students
// I.S. : studentCount < MaxStudents; data akan diinput oleh user
// F.S. : students[studentCount] terisi data baru; studentCount bertambah 1
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
		fmt.Printf("\n  [!] ID '%s' already exists.\n", id)
		pressEnterToContinue()
		return
	}

	name := readLine("  Full Name    : ")
	if name == "" {
		fmt.Println("\n  [!] Name cannot be empty.")
		pressEnterToContinue()
		return
	}

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

	fmt.Printf("\n  [v] Student '%s' added successfully.\n", name)
	pressEnterToContinue()
}

// editStudent : Memperbarui Nama dan/atau Jurusan mahasiswa yang ada
// I.S. : studentCount > 0; mahasiswa dengan ID yang diberikan ada di array
// F.S. : students[idx].Name dan/atau students[idx].Major diperbarui
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

	fmt.Printf("\n  Selected: %s | Major: %s\n", students[idx].Name, students[idx].Major)
	fmt.Println("  (Press Enter to keep current value)\n")

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

	fmt.Printf("\n  [v] Student ID '%s' updated successfully.\n", id)
	pressEnterToContinue()
}

// deleteStudent : Menghapus data mahasiswa by ID menggunakan swap-and-shrink
// I.S. : Mahasiswa dengan ID yang diberikan ada di array; studentCount > 0
// F.S. : Mahasiswa dihapus; elemen terakhir mengisi posisinya; studentCount-1
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

	name                       := students[idx].Name
	students[idx]               = students[studentCount-1]
	students[studentCount-1]    = Student{}
	studentCount--

	fmt.Printf("\n  [v] Student '%s' deleted successfully.\n", name)
	pressEnterToContinue()
}

// studentManagementMenu : Menjalankan loop sub-menu Student Management
// I.S. : -
// F.S. : Kembali ke menu utama ketika user memilih 0 (selesai = true)
func studentManagementMenu() {
	selesai := false
	for !selesai {
		showStudentManagementMenu()
		choice := readInt("  Select option: ")
		if choice == 1 {
			addStudent()
		} else if choice == 2 {
			editStudent()
		} else if choice == 3 {
			deleteStudent()
		} else if choice == 0 {
			selesai = true
		} else {
			fmt.Println("\n  [!] Invalid option.")
			pressEnterToContinue()
		}
	}
}

// ============================================================
//  DEPARTMENT MANAGEMENT  (App Spec a)
// ============================================================

// addDepartment : Menambah data jurusan baru ke array departments
// I.S. : deptCount < MaxDepts; data akan diinput oleh user
// F.S. : departments[deptCount] terisi data baru; deptCount bertambah 1
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
		fmt.Println("\n  [!] Name cannot be empty.")
		pressEnterToContinue()
		return
	}

	departments[deptCount] = Department{ID: id, Name: name}
	deptCount++

	fmt.Printf("\n  [v] Department '%s' added successfully.\n", name)
	pressEnterToContinue()
}

// editDepartment : Memperbarui nama jurusan yang ada
// I.S. : deptCount > 0; jurusan dengan ID yang diberikan ada di array
// F.S. : departments[idx].Name diperbarui
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

	fmt.Printf("\n  Selected: %s\n", departments[idx].Name)
	fmt.Println("  (Press Enter to keep current value)\n")

	newName := readLine("  New Department Name : ")
	if newName != "" {
		departments[idx].Name = newName
	}

	fmt.Printf("\n  [v] Department ID '%s' updated successfully.\n", id)
	pressEnterToContinue()
}

// deleteDepartment : Menghapus data jurusan by ID menggunakan swap-and-shrink
// I.S. : Jurusan dengan ID yang diberikan ada di array; deptCount > 0
// F.S. : Jurusan dihapus; elemen terakhir mengisi posisinya; deptCount-1
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

	name                          := departments[idx].Name
	departments[idx]               = departments[deptCount-1]
	departments[deptCount-1]       = Department{}
	deptCount--

	fmt.Printf("\n  [v] Department '%s' deleted successfully.\n", name)
	pressEnterToContinue()
}

// departmentManagementMenu : Menjalankan loop sub-menu Department Management
// I.S. : -
// F.S. : Kembali ke menu utama ketika user memilih 0 (selesai = true)
func departmentManagementMenu() {
	selesai := false
	for !selesai {
		showDeptManagementMenu()
		choice := readInt("  Select option: ")
		if choice == 1 {
			addDepartment()
		} else if choice == 2 {
			editDepartment()
		} else if choice == 3 {
			deleteDepartment()
		} else if choice == 0 {
			selesai = true
		} else {
			fmt.Println("\n  [!] Invalid option.")
			pressEnterToContinue()
		}
	}
}

// ============================================================
//  SCORE & EVALUATION  (App Spec c)
// ============================================================

// addEditScore : Memasukkan atau memperbarui nilai ujian mahasiswa
// I.S. : Mahasiswa dengan ID yang diberikan ada di array students
// F.S. : students[idx].Score = score;
//        students[idx].Status = "Accepted" jika score >= 75, "Rejected" jika tidak
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

	fmt.Printf("\n  Student: %s | Major: %s\n\n", students[idx].Name, students[idx].Major)

	score := readFloat("  Test Score (0 - 100): ")
	if score < 0 || score > 100 {
		fmt.Println("\n  [!] Score must be between 0 and 100.")
		pressEnterToContinue()
		return
	}

	students[idx].Score = score
	if score >= 75 {
		students[idx].Status = "Accepted"
	} else {
		students[idx].Status = "Rejected"
	}

	fmt.Printf("\n  [v] Score %.2f recorded. Status: %s\n", score, students[idx].Status)
	pressEnterToContinue()
}

// deleteScore : Mereset nilai ujian mahasiswa ke 0 dan status ke "Pending"
// I.S. : Mahasiswa ada di array; students[idx].Score > 0
// F.S. : students[idx].Score = 0; students[idx].Status = "Pending"
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
		fmt.Printf("\n  [!] Student '%s' has no score yet.\n", students[idx].Name)
		pressEnterToContinue()
		return
	}

	students[idx].Score  = 0
	students[idx].Status = "Pending"

	fmt.Printf("\n  [v] Score deleted. '%s' reset to Pending.\n", students[idx].Name)
	pressEnterToContinue()
}

// scoreMenu : Menjalankan loop sub-menu Score & Evaluation
// I.S. : -
// F.S. : Kembali ke menu utama ketika user memilih 0 (selesai = true)
func scoreMenu() {
	selesai := false
	for !selesai {
		showScoreMenu()
		choice := readInt("  Select option: ")
		if choice == 1 {
			addEditScore()
		} else if choice == 2 {
			deleteScore()
		} else if choice == 0 {
			selesai = true
		} else {
			fmt.Println("\n  [!] Invalid option.")
			pressEnterToContinue()
		}
	}
}

// ============================================================
//  PENCARIAN  (General Spec d)
// ============================================================

// doSequentialSearch : Demo Sequential Search berdasarkan ID mahasiswa
// I.S. : studentCount > 0
// F.S. : Hasil pencarian ditampilkan ke layar; array students tidak berubah
func doSequentialSearch() {
	printHeader("SEQUENTIAL SEARCH - by ID")

	if studentCount == 0 {
		fmt.Println("\n  [!] No student records found.")
		pressEnterToContinue()
		return
	}

	fmt.Println("  Student List (original order):")
	fmt.Println()
	printStudentTable(students, studentCount)
	fmt.Println()

	id := readLine("  Enter Student ID to search: ")
	fmt.Println("\n  Scanning array from index 0...")
	fmt.Println()

	idx := sequentialSearchByID(id)

	if idx == -1 {
		fmt.Printf("  [x] Student with ID '%s' not found.\n", id)
	} else {
		fmt.Printf("  [v] Found at index [%d]:\n\n", idx)
		var result [MaxStudents]Student
		result[0] = students[idx]
		printStudentTable(result, 1)
	}
	pressEnterToContinue()
}

// doBinarySearch : Demo Binary Search berdasarkan Nama (sort salinan dulu)
// I.S. : studentCount > 0
// F.S. : Salinan terurut dan hasil pencarian ditampilkan; array asli tidak berubah
func doBinarySearch() {
	printHeader("BINARY SEARCH - by Name")

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

	fmt.Println("\n  Array sorted by Name (Ascending) - required for Binary Search:")
	fmt.Println()
	printStudentTable(arrSorted, studentCount)

	fmt.Println("\n  Applying Binary Search...")
	fmt.Println()
	idx := binarySearchByName(arrSorted, studentCount, name)

	if idx == -1 {
		fmt.Printf("  [x] Student '%s' not found.\n", name)
	} else {
		fmt.Printf("  [v] Found at sorted index [%d]:\n\n", idx)
		var result [MaxStudents]Student
		result[0] = arrSorted[idx]
		printStudentTable(result, 1)
	}
	pressEnterToContinue()
}

// searchMenu : Menjalankan loop sub-menu Search Student
// I.S. : -
// F.S. : Kembali ke menu utama ketika user memilih 0 (selesai = true)
func searchMenu() {
	selesai := false
	for !selesai {
		showSearchMenu()
		choice := readInt("  Select option: ")
		if choice == 1 {
			doSequentialSearch()
		} else if choice == 2 {
			doBinarySearch()
		} else if choice == 0 {
			selesai = true
		} else {
			fmt.Println("\n  [!] Invalid option.")
			pressEnterToContinue()
		}
	}
}

// ============================================================
//  LAPORAN  (App Spec b & d, General Spec e)
// ============================================================

// viewByMajorAndStatus : Menampilkan mahasiswa berdasarkan jurusan dan status
// I.S. : major dan statusFilter diinput oleh user; studentCount >= 0
// F.S. : Subset mahasiswa yang sesuai filter ditampilkan (App Spec b);
//        Tidak menggunakan continue (General Spec g) — if bersarang digunakan
func viewByMajorAndStatus() {
	printHeader("VIEW BY MAJOR & ADMISSION STATUS")

	if deptCount > 0 {
		fmt.Println("  Registered Departments:")
		fmt.Println()
		printDeptTable(departments, deptCount)
		fmt.Println()
	}

	major        := readLine("  Enter Major Name: ")
	statusFilter := strings.ToLower(readLine("  Status Filter (Accepted/Rejected/All): "))

	var filtered  [MaxStudents]Student
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

// viewSortedByScore : Menampilkan mahasiswa terurut by Score (Selection Sort)
// I.S. : studentCount >= 0
// F.S. : Salinan array terurut ditampilkan (App Spec d); array asli tidak berubah
func viewSortedByScore() {
	printHeader("SORT BY SCORE - Selection Sort")

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

// viewSortedByName : Menampilkan mahasiswa terurut by Name (Insertion Sort)
// I.S. : studentCount >= 0
// F.S. : Salinan array terurut ditampilkan (App Spec d); array asli tidak berubah
func viewSortedByName() {
	printHeader("SORT BY NAME - Insertion Sort")

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

// viewSortedByMajor : Menampilkan mahasiswa terurut by Major (Insertion Sort)
// I.S. : studentCount >= 0
// F.S. : Salinan array terurut ditampilkan (App Spec d); array asli tidak berubah
func viewSortedByMajor() {
	printHeader("SORT BY MAJOR - Insertion Sort")

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

// reportsMenu : Menjalankan loop sub-menu View Reports
// I.S. : -
// F.S. : Kembali ke menu utama ketika user memilih 0 (selesai = true)
func reportsMenu() {
	selesai := false
	for !selesai {
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
			selesai = true
		} else {
			fmt.Println("\n  [!] Invalid option.")
			pressEnterToContinue()
		}
	}
}

// ============================================================
//  ENTRY POINT
// ============================================================

// main : Titik masuk program utama
// I.S. : Program dimulai; studentCount = 0, deptCount = 0
// F.S. : Program berakhir ketika user memilih 0 (selesai = true)
func main() {
	selesai := false
	for !selesai {
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
			selesai = true
		} else {
			fmt.Println("\n  [!] Invalid option.")
			pressEnterToContinue()
		}
	}
}
