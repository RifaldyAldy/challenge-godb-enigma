package customer

import (
	"bufio"
	customer "challenge-godb/controller/table_customer"
	"fmt"
	"os"
)

func Menu() {
	var opsi string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("1. Cari semua customer")
	fmt.Println("2. Cari customer melalui kolom tertentu")
	fmt.Println("3. Daftarkan customer baru")
	fmt.Println("4. Edit data customer")
	fmt.Println("5. Delete data customer")
	fmt.Print("Input nomor 1-5 sebagai pilihan: ")
	fmt.Scan(&opsi)
	switch opsi {
	case "1":
		datas := customer.ViewAll()
		table := customer.Tabel{Headers: []string{"NO", "ID", "NAME", "PHONE NUMBER"}, Rows: datas}
		table.GenerateTable() // method untuk set data menjadi tampilan tabel
	case "2":
		var kolom, value string
		fmt.Println("1. kolom id")
		fmt.Println("2. kolom customer_name")
		fmt.Println("3. kolom phone_number")
		fmt.Print("Masukan 1-3 pilih kolom yang ingin dicari: ")
		fmt.Scan(&kolom)
		if kolom == "1" {
			kolom = "id"
		} else if kolom == "2" {
			kolom = "customer_name"
		} else if kolom == "3" {
			kolom = "phone_number"
		} else {
			fmt.Println("opsi yang kamu pilih tidak ada")
			break
		}
		fmt.Printf("Masukan value yang ingin dicari dari kolom %s: ", kolom)
		scanner.Scan()
		value = scanner.Text()
		datas := customer.SearchCust(kolom, value)
		table := customer.Tabel{Headers: []string{"NO", "ID", "NAME", "Phone number"},
			Rows: datas}
		table.GenerateTable() // method untuk set data menjadi tampilan tabel
	case "3":
		data := customer.Customer{}
		fmt.Print("Masukan nama customer: ")
		scanner.Scan()
		data.Customer_name = scanner.Text()
		fmt.Print("Masukan nomor handphone: ")
		scanner.Scan()
		data.Phone_number = scanner.Text()
		customer.InsertCust(data)
		customer.SearchCust("customer_name", data.Customer_name)
	case "4":
		var id int
		var kolom, value string
		fmt.Print("Masukan id customer yang ingin di edit: ")
		fmt.Scan(&id)
		check := customer.CheckDataisExists(id) // cek apakah id data ada
		if !check {
			fmt.Printf("Data dengan id = %d tidak ada\n", id)
			break
		}
		fmt.Println("1. kolom customer_name")
		fmt.Println("2. kolom phone_number")
		fmt.Print("Kolom yang ingin di edit: ")
		fmt.Scan(&kolom)
		if kolom == "1" {
			kolom = "customer_name"
		} else if kolom == "2" {
			kolom = "phone_number"
		} else {
			fmt.Println("opsi yang kamu pilih tidak ada")
			break
		}
		fmt.Printf("Value %s yang ingin di edit: ", kolom)
		scanner.Scan()
		value = scanner.Text()
		customer.UpdateCust(id, kolom, value)
		defer customer.SearchCust("id", id)
	case "5":
		var id int
		fmt.Print("Masukan id yang ingin di delete: ")
		fmt.Scan(&id)
		customer.DeleteCust(id)
	default:
		fmt.Println("Opsi yang anda masukkan salah!")
	}
	fmt.Print("Kembali ke menu customer? Y/n: ")
	fmt.Scan(&opsi)
	if opsi == "y" || opsi == "Y" {
		Menu()
	} else {
		return
	}
}
