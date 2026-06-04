# Student Registration App

Program ini adalah aplikasi registrasi calon mahasiswa berbasis terminal yang dibuat dengan bahasa Go. Aplikasi digunakan untuk mengelola data calon mahasiswa, jurusan yang dipilih, nilai tes, status kelulusan, pencarian data, dan laporan data mahasiswa.

## Cara Menjalankan Program

Pastikan Go sudah terpasang di komputer.

```bash
go run bigproject.go
```

Program juga bisa dikompilasi terlebih dahulu:

```bash
go build bigproject.go
./bigproject
```

## Deskripsi Tugas Besar

Application Registration Student adalah aplikasi registrasi mahasiswa universitas yang mengelola data calon mahasiswa dan data jurusan universitas. Pengguna aplikasi terdiri dari admin aplikasi dan calon mahasiswa.

Spesifikasi aplikasi:

1. Pengguna dapat melakukan penambahan, perubahan, dan penghapusan data mahasiswa serta data jurusan pada universitas.
2. Pengguna dapat menampilkan data mahasiswa yang mendaftar pada jurusan tertentu, termasuk mahasiswa yang diterima dan ditolak.
3. Pengguna dapat melakukan penambahan, perubahan, dan penghapusan nilai tes mahasiswa yang akan menentukan status diterima atau ditolak.
4. Pengguna dapat menampilkan data mahasiswa yang diurutkan berdasarkan nilai tes, jurusan, dan nama mahasiswa.

## Spesifikasi Umum

1. Program dibuat secara modular dengan menggunakan subprogram berupa fungsi dan prosedur.
2. Setiap subprogram dilengkapi parameter dan spesifikasi berupa kondisi awal atau I.S. dan kondisi akhir atau F.S.
3. Program mengimplementasikan array dan tipe bentukan struktur.
4. Array yang digunakan adalah array statis, bukan array dinamis atau slice.
5. Program mengimplementasikan Sequential Search dan Binary Search pada proses pencarian data.
6. Program mengimplementasikan Selection Sort dan Insertion Sort untuk pengurutan data dengan kategori berbeda.
7. Setiap kategori pengurutan dapat ditampilkan secara ascending maupun descending.
8. Rekursif tidak wajib diimplementasikan, tetapi dapat menjadi nilai tambahan.
9. Program tidak menggunakan statement `break` dan `continue`.
10. Variabel global hanya digunakan untuk data array utama yang diolah.

## Struktur Data

Data utama disimpan dalam array statis:

```go
const MaxStudents = 100

var students [MaxStudents]Student
var studentCount int
```

Tipe bentukan yang digunakan:

```go
type Student struct {
    ID     string
    Name   string
    Major  string
    Score  float64
    Status string
}
```

Keterangan field:

1. `ID`: nomor atau kode unik mahasiswa.
2. `Name`: nama lengkap calon mahasiswa.
3. `Major`: jurusan yang dipilih mahasiswa.
4. `Score`: nilai tes mahasiswa.
5. `Status`: status seleksi, yaitu `Pending`, `Accepted`, atau `Rejected`.

## Flow Program

Saat program dijalankan, fungsi `main()` menampilkan menu utama:

```text
1. Student Management
2. Score & Evaluation
3. Search Student
4. View Reports
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

## Flow Score & Evaluation

Menu `Score & Evaluation` digunakan untuk memberikan atau mengubah nilai tes mahasiswa.

Fungsi yang digunakan:

1. `evaluateStudent()`: mencari mahasiswa berdasarkan ID, menerima input nilai, lalu menentukan status.

Aturan status:

1. Jika nilai `>= 75`, status menjadi `Accepted`.
2. Jika nilai `< 75`, status menjadi `Rejected`.
3. Mahasiswa baru memiliki status awal `Pending`.

## Flow Search Student

Menu `Search Student` digunakan untuk mencari data mahasiswa.

```text
1. Sequential Search by ID
2. Binary Search by Name
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

Sequential Search digunakan pada fungsi `sequentialSearchByID(id string) int`.

Cara kerja:

1. Pencarian dimulai dari indeks `0`.
2. Setiap elemen dibandingkan dengan ID yang dicari.
3. Jika ID ditemukan, fungsi mengembalikan indeks data.
4. Jika tidak ditemukan, fungsi mengembalikan `-1`.

Algoritma ini digunakan pada proses pencarian, edit, delete, dan evaluasi nilai mahasiswa berdasarkan ID.

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

Program menggunakan pendekatan modular. Setiap bagian dipisahkan menjadi fungsi input, tampilan, pencarian, pengurutan, manajemen mahasiswa, evaluasi nilai, pencarian mahasiswa, laporan, dan menu utama.

Data jurusan pada kode ini direpresentasikan melalui field `Major` pada struct `Student`. Nilai tes dapat ditambahkan dan diubah melalui menu `Score & Evaluation`; status mahasiswa diperbarui otomatis berdasarkan nilai tersebut.

## File Utama

```text
bigproject.go
```

File tersebut berisi seluruh implementasi aplikasi, mulai dari deklarasi struct, array statis, fungsi input, fungsi pencarian, fungsi pengurutan, submenu, sampai fungsi `main()`.
