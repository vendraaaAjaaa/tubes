# Student Registration App

Program ini adalah aplikasi registrasi calon mahasiswa berbasis terminal yang dibuat dengan bahasa Go. Aplikasi digunakan untuk mengelola data calon mahasiswa, data jurusan, nilai tes, status kelulusan, pencarian data, dan laporan data mahasiswa.

## Cara Menjalankan Program

Pastikan Go sudah terpasang di komputer.

```bash
go run bigproject.go
```

Program juga bisa dikompilasi terlebih dahulu:

```bash
go build -o bigproject bigproject.go
./bigproject
```

## Deskripsi Tugas Besar

Application Registration Student adalah aplikasi registrasi mahasiswa universitas yang mengelola data calon mahasiswa dan data jurusan universitas. Pengguna aplikasi terdiri dari admin aplikasi dan calon mahasiswa.

Spesifikasi aplikasi:

1. Pengguna dapat melakukan penambahan, perubahan, dan penghapusan data mahasiswa.
2. Pengguna dapat melakukan penambahan, perubahan, dan penghapusan data jurusan.
3. Pengguna dapat melakukan penambahan, perubahan, dan penghapusan nilai tes mahasiswa.
4. Program menentukan status mahasiswa berdasarkan nilai tes.
5. Pengguna dapat menampilkan data mahasiswa berdasarkan jurusan dan status penerimaan.
6. Pengguna dapat mencari mahasiswa menggunakan Sequential Search dan Binary Search.
7. Pengguna dapat menampilkan data mahasiswa yang diurutkan berdasarkan nilai, nama, dan jurusan.

## Spesifikasi Umum

1. Program dibuat secara modular dengan menggunakan subprogram berupa fungsi dan prosedur.
2. Program mengimplementasikan array dan tipe bentukan struktur.
3. Array yang digunakan adalah array statis, bukan array dinamis atau slice.
4. Program mengimplementasikan Sequential Search dan Binary Search pada proses pencarian data.
5. Program mengimplementasikan Selection Sort dan Insertion Sort untuk pengurutan data.
6. Setiap kategori pengurutan dapat ditampilkan secara ascending atau descending.
7. Program tidak menggunakan statement `break` dan `continue`.
8. Variabel global digunakan untuk array utama mahasiswa dan jurusan beserta jumlah datanya.

## Struktur Data

Data utama disimpan dalam array statis:

```go
const MaxStudents = 100
const MaxDepts = 50

var students [MaxStudents]Student
var studentCount int

var departments [MaxDepts]Department
var deptCount int
```

Tipe bentukan untuk mahasiswa:

```go
type Student struct {
    ID     string
    Name   string
    Major  string
    Score  float64
    Status string
}
```

Keterangan field `Student`:

1. `ID`: nomor atau kode unik mahasiswa.
2. `Name`: nama lengkap calon mahasiswa.
3. `Major`: jurusan yang dipilih mahasiswa.
4. `Score`: nilai tes mahasiswa.
5. `Status`: status seleksi, yaitu `Pending`, `Accepted`, atau `Rejected`.

Tipe bentukan untuk jurusan:

```go
type Department struct {
    ID   string
    Name string
}
```

Keterangan field `Department`:

1. `ID`: kode unik jurusan.
2. `Name`: nama jurusan.

## Flow Program

Saat program dijalankan, fungsi `main()` menampilkan menu utama:

```text
1. Student Management
2. Department Management
3. Score & Evaluation
4. Search Student
5. View Reports
0. Exit
```

Alur utama aplikasi:

1. Program dimulai dari `main()`.
2. User memilih menu utama.
3. Program masuk ke submenu sesuai pilihan user.
4. Setiap submenu menjalankan fungsi terkait.
5. Setelah proses selesai, program kembali ke menu sebelumnya.
6. Program berhenti ketika user memilih menu `0. Exit`.

## Flow Student Management

Menu `Student Management` digunakan untuk mengelola data mahasiswa.

```text
1. Add Student
2. Edit Student
3. Delete Student
0. Back
```

Fungsi yang digunakan:

1. `addStudent()`: menambahkan mahasiswa baru ke array `students`.
2. `editStudent()`: mengubah nama dan jurusan mahasiswa berdasarkan ID.
3. `deleteStudent()`: menghapus mahasiswa berdasarkan ID.
4. `studentManagementMenu()`: mengatur alur submenu manajemen mahasiswa.

Pada proses edit dan delete, program mencari mahasiswa menggunakan `sequentialSearchByID()`.

## Flow Department Management

Menu `Department Management` digunakan untuk mengelola data jurusan.

```text
1. Add Department
2. Edit Department
3. Delete Department
0. Back
```

Fungsi yang digunakan:

1. `addDepartment()`: menambahkan jurusan baru ke array `departments`.
2. `editDepartment()`: mengubah nama jurusan berdasarkan ID.
3. `deleteDepartment()`: menghapus jurusan berdasarkan ID.
4. `departmentManagementMenu()`: mengatur alur submenu manajemen jurusan.
5. `sequentialSearchDeptByID(id string) int`: mencari jurusan berdasarkan ID.

Data jurusan yang sudah ditambahkan akan ditampilkan sebagai referensi saat user menambah atau mengubah data mahasiswa.

## Flow Score & Evaluation

Menu `Score & Evaluation` digunakan untuk mengelola nilai tes mahasiswa.

```text
1. Add / Edit Score
2. Delete Score  (reset to Pending)
0. Back
```

Fungsi yang digunakan:

1. `addEditScore()`: mencari mahasiswa berdasarkan ID, menerima input nilai, lalu menentukan status.
2. `deleteScore()`: menghapus nilai mahasiswa dengan mengatur nilai menjadi `0` dan status menjadi `Pending`.
3. `scoreMenu()`: mengatur alur submenu nilai dan evaluasi.

Aturan status:

1. Jika nilai `>= 75`, status menjadi `Accepted`.
2. Jika nilai `< 75`, status menjadi `Rejected`.
3. Jika nilai dihapus, status kembali menjadi `Pending`.
4. Mahasiswa baru memiliki status awal `Pending`.

## Flow Search Student

Menu `Search Student` digunakan untuk mencari data mahasiswa.

```text
1. Sequential Search  (by ID)
2. Binary Search      (by Name)
0. Back
```

Fungsi yang digunakan:

1. `doSequentialSearch()`: menjalankan pencarian linear berdasarkan ID.
2. `doBinarySearch()`: menyalin data mahasiswa, mengurutkan salinan berdasarkan nama, lalu menjalankan Binary Search.
3. `sequentialSearchByID(id string) int`: mencari data mahasiswa dari indeks awal sampai akhir.
4. `binarySearchByName(arr [MaxStudents]Student, n int, name string) int`: mencari nama mahasiswa pada array yang sudah terurut ascending.

## Flow View Reports

Menu `View Reports` digunakan untuk menampilkan laporan data mahasiswa.

```text
1. View by Major & Admission Status
2. Sort by Score
3. Sort by Name
4. Sort by Major
0. Back
```

Fungsi yang digunakan:

1. `viewByMajorAndStatus()`: menampilkan mahasiswa berdasarkan jurusan dan status.
2. `viewSortedByScore()`: menampilkan mahasiswa terurut berdasarkan nilai.
3. `viewSortedByName()`: menampilkan mahasiswa terurut berdasarkan nama.
4. `viewSortedByMajor()`: menampilkan mahasiswa terurut berdasarkan jurusan.
5. `reportsMenu()`: mengatur alur submenu laporan.

## Algoritma yang Digunakan

### Sequential Search

Sequential Search digunakan pada:

1. `sequentialSearchByID(id string) int`
2. `sequentialSearchDeptByID(id string) int`

Cara kerja:

1. Pencarian dimulai dari indeks `0`.
2. Setiap elemen dibandingkan dengan ID yang dicari.
3. Jika ID ditemukan, fungsi mengembalikan indeks data.
4. Jika tidak ditemukan, fungsi mengembalikan `-1`.

Algoritma ini digunakan pada proses pencarian, edit, delete, dan evaluasi nilai berdasarkan ID.

### Binary Search

Binary Search digunakan pada fungsi `binarySearchByName(arr [MaxStudents]Student, n int, name string) int`.

Cara kerja:

1. Data mahasiswa disalin ke array lokal.
2. Array lokal diurutkan berdasarkan nama secara ascending.
3. Pencarian dilakukan dengan membandingkan nilai tengah array.
4. Jika nama ditemukan, fungsi mengembalikan indeks.
5. Jika tidak ditemukan, fungsi mengembalikan `-1`.

### Selection Sort

Selection Sort digunakan pada fungsi `selectionSortByScore(arr *[MaxStudents]Student, n int, ascending bool)`.

Cara kerja:

1. Program mencari elemen dengan nilai terkecil atau terbesar.
2. Elemen tersebut ditukar ke posisi yang sesuai.
3. Proses diulang sampai seluruh data terurut.

Selection Sort dipakai untuk mengurutkan data berdasarkan nilai tes secara ascending atau descending.

### Insertion Sort

Insertion Sort digunakan pada:

1. `insertionSortByName(arr *[MaxStudents]Student, n int, ascending bool)`
2. `insertionSortByMajor(arr *[MaxStudents]Student, n int, ascending bool)`

Cara kerja:

1. Data diproses dari kiri ke kanan.
2. Elemen aktif disisipkan ke posisi yang tepat.
3. Proses diulang sampai seluruh data terurut.

Insertion Sort dipakai untuk mengurutkan data berdasarkan nama dan jurusan secara ascending atau descending.

## Catatan Implementasi

Program menggunakan pendekatan modular. Setiap bagian dipisahkan menjadi fungsi input, tampilan, pencarian, pengurutan, manajemen mahasiswa, manajemen jurusan, evaluasi nilai, pencarian mahasiswa, laporan, dan menu utama.

Data mahasiswa disimpan di array `students`, sedangkan data jurusan disimpan di array `departments`. Nilai tes dapat ditambahkan, diubah, dan dihapus melalui menu `Score & Evaluation`; status mahasiswa diperbarui otomatis berdasarkan nilai tersebut.

Program berjalan sepenuhnya di terminal dan menyimpan data sementara di memori selama program dijalankan.

## File Utama

```text
bigproject.go
```

File tersebut berisi seluruh implementasi aplikasi, mulai dari deklarasi struct, array statis, fungsi input, fungsi pencarian, fungsi pengurutan, submenu, sampai fungsi `main()`.
