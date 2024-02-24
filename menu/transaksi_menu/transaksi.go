package transaksimenu

import (
	tabletdetails "challenge-godb/controller/table_t_details"
	tabletransaksi "challenge-godb/controller/table_transaksi"
	"challenge-godb/lib/db"
	transaksijoin "challenge-godb/view/Transaksi_join"
	"fmt"
	"strings"
	"time"
)

func Menu() {
	var opsi, repeat string
	t_details := tabletdetails.Transaksi_details{}
	transaksi := []tabletransaksi.Transaksi{}
	fmt.Println("1. Buat Transaksi")
	fmt.Println("2. View semua transaksi info")
	fmt.Println("3. Search Transaksi info")
	fmt.Println("4. View transaksi details")
	fmt.Println("5. Delete transaksi")
	fmt.Print("Masukan opsi: ")
	fmt.Scan(&opsi)

	switch opsi {
	case "1":
		fmt.Print("Masukan Customer_id: ")
		fmt.Scan(&t_details.Customer_id)
		fmt.Print("Pesanan diterima oleh: ")
		fmt.Scan(&t_details.Received_by)
		t_details.Entry_date = time.Now()
		t_details.Finish_date = t_details.Entry_date.AddDate(0, 0, 2)
		t := tabletransaksi.Transaksi{}
		repeat := true
		for repeat {
			fmt.Print("Masukan jasa service_id: ")
			fmt.Scan(&t.Service_id)
			fmt.Print("Masukan quantity: ")
			fmt.Scan(&t.Quantity)
			transaksi = append(transaksi, t)
			fmt.Print("Tambah service lagi? : Y/n: ")
			fmt.Scan(&opsi)
			if strings.ToLower(opsi) != "y" {
				repeat = false

			}
		}
		tabletdetails.CommitInsert(t_details, transaksi)
		fmt.Println("Pesan berhasil di commit dan masuk database")
	case "2":
		data := tabletransaksi.ViewALl()
		table := tabletransaksi.Tabel{Headers: []string{"ID", "Customer Name", "Tanggal masuk", "Tanggal selesai", "Diterima oleh"}, Rows: data}
		table.GenerateTable()
	case "3":
		var id int
		fmt.Print("Masukkan transaksi ID: ")
		fmt.Scan(&id)
		data := []tabletransaksi.Transaksi_join{tabletransaksi.SearchTransaksi(id)}
		table := tabletransaksi.Tabel{Headers: []string{"ID", "Customer Name", "Tanggal masuk", "Tanggal selesai", "Diterima oleh"}, Rows: data}
		table.GenerateTable()
	case "4":
		var id int
		fmt.Print("Masukkan transaksi ID: ")
		fmt.Scan(&id)
		transaksijoin.TransaksiJoin(id)
	case "5":
		var id int
		db := db.Dbcon()
		tx, _ := db.Begin()
		defer db.Close()
		fmt.Print("Masukan id transaksi_details: ")
		fmt.Scan(&id)
		tabletdetails.DeleteTransaksi(id, tx)
	default:
		fmt.Printf("Ospi %s tidak ada!\n", opsi)
	}
	fmt.Print("Kembali ke menu transaksi? Y/n: ")
	fmt.Scan(&repeat)
	if strings.ToLower(repeat) == "y" {
		Menu()
	}
}
