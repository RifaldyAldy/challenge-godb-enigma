Submission Project: Go Database
===

#### Features
- Add Customer
- Get All Customer
- Get Customer by column
- Update Customer
- Delete Customer
- Add Service
- Get All Service
- Get Service by Column
- Update Service
- Delete Service
- Create Transaction
- Get transaction Details by ID
- Get All transaction Information
- Get Transaction Information by ID
- Delete Transaction

#### Information

Untuk setiap tabel, saya menggunakan tampilan tabel agar lebih rapi, jadi mohon saat menjalankan program Terminal dalam keadaan fullscreen, dan baca informasi print pada baris agar tidak salah input

untuk package koneksi database saya berada di lib/db/db.go

#### Penggunaan - Basic

1. File DDL.sql sudah disiapkan maka buat database terlebih dahulu dan table nya
2. Jalankan program dengan `go run .` atau `go run enigma-laundry.go`

## Penggunaan - Menu Customer

Setelah `go run .` akan muncul menu lalu ketik `1` untuk masuk ke menu customer lalu ketik opsi dari 1-5 sesuai keperluan

1. Pada opsi `1` akan menampilkan seluruh data customer
2. Pada opsi `2` akan diminta sebuah ID Customer untuk mencari Customer yang diinginkan
3. Pada opsi `3` mendaftarkan Customer/pelanggan baru untuk keperluan transaksi
    - Nama Customer dapat menggunakan space
    - Pada nomor handphone terdapat 2 validasi yaitu nomor harus berawal 08 dan tidak boleh lebih dari 15
4. Pada opsi `4` anda dapat mengedit data customer berdasarkan kolom dari ID
5. pada opsi `5` menghapus data Customer berdasarkan ID yang diminta

## Penggunaan - Menu Service

Pada menu utama ketik `2` untuk masuk menu service lalu ketik opsi 1-5 sesuai keperluan

1. Pada opsi `1` akan menampilkan seluruh service yang ada pada enigma-laundry
2. Pada opsi `2` fitur untuk mencari sebuah data service berdasarkan kolom yang ada inginkan
    - Pilih kolom
    - Ketik value berdasarkan kolom yang di pilih
3. Pada opsi `3` menambahkan service baru
    - Nama service dapat menggunakan space
    - Unit satuan contoh KG/buah, terdapat validasi jika selain contoh tadi akan terjadi peringatan
    - Price adalah harga sehingga hanya menuliskan angka saja, tidak perlu koma(,) ataupun titik(.)
4. Pada opsi `4` mengedit sebuah service berdasarkan ID yang diminta
    - Terdapat validasi jika service ID yang diinput tidak ada, maka muncul peringatan
5. Pada opsi `5` menghapus sebuah service berdasarkan ID yang diminta
    - Terdapat validasi jika service ID yang diinput tidak ada, maka muncul peringatan

## Penggunaan - Menu Transaksi

Menu ini saling bergantung dengan 2 tabel yaitu One to Many : 
    - Pertama tabel transaksi_details yang memiliki informasi nama customer,tanggal masuk/selesai,total_harga,dan total quantity
    - Kedua tabel transaksi yang memiliki informasi apapun jasa service yang dipilih oleh customer

1. Pada opsi `1` anda membuat transaksi baru dari permintaan customer
    1. Input ID customer yang memesan
    2. Input nama Anda sebagai penerima pesanan customer
    3. Input service ID yang customer pesan ingin jasa apa
    4. Input quantity, contoh Laundry 10KG, maka cukup input 10
    5. Akan ada opsi ketik Y/n untuk menambahkan jasa lagi atau sudah cukup
    6. Jika terdapat kesalahan satu atau lebih input, maka akan terjadi rollback
2. Pada opsi `2` mengambil seluruh data transaksi singkat
3. Pada opsi `3` mengambil sebuah data transaksi singkat berdasarkan transaksi_details ID
4. Pada opsi `4` mengambil data transaksi yang lebih detail berdasarkan transaksi_details ID, jika anda tidak tau ID transaksi anda dapat memeriksanya di opsi `2` dan 
5. Pada opsi `5` menghapus data transaksi berdasarkan transaksi_details ID.
