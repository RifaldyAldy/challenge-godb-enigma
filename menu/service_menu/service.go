package servicemenu

import (
	"bufio"
	service "challenge-godb/controller/table_service"
	"fmt"
	"os"
	"strconv"
)

func Menu() {
	scanner := bufio.NewScanner(os.Stdin)
	var opsi string
	fmt.Println("1. Cari semua service")
	fmt.Println("2. Cari service melalui kolom tertentu")
	fmt.Println("3. Daftarkan service baru")
	fmt.Println("4. Edit data service")
	fmt.Println("5. Delete data service")
	fmt.Print("Input nomor 1-5 sebagai pilihan: ")
	fmt.Scan(&opsi)
	switch opsi {
	case "1":
		service.ViewAll()
	case "2":
		var kolom, value string
		fmt.Println("1. kolom id")
		fmt.Println("2. kolom service_name")
		fmt.Println("3. kolom unit_name")
		fmt.Println("4. kolom price_per_unit")
		fmt.Print("Masukan 1-4 pilih kolom yang ingin dicari: ")
		fmt.Scan(&kolom)
		if kolom == "1" {
			kolom = "id"
		} else if kolom == "2" {
			kolom = "service_name"
		} else if kolom == "3" {
			kolom = "unit_name"
		} else if kolom == "4" {
			kolom = "price_per_unit"
		} else {
			fmt.Println("opsi yang kamu pilih tidak ada")
			break
		}
		fmt.Printf("Masukan value yang ingin dicari dari kolom %s: ", kolom)
		scanner.Scan()
		value = scanner.Text()
		datas := service.SearchServ(kolom, value)
		table := service.Tabel{
			Rows: datas}
		table.GenerateTable() // method untuk set data menjadi tampilan tabel
	case "3":
		data := service.Service{}
		fmt.Print("Masukan nama service: ")
		scanner.Scan()
		data.Service_name = scanner.Text()
		fmt.Print("Masukan unit satuan: ")
		scanner.Scan()
		data.Unit_name = scanner.Text()
		fmt.Print("Masukan Price: ")
		scanner.Scan()
		data.Price_per_unit, _ = strconv.Atoi(scanner.Text())
		service.InsertServ(data)
		service.SearchServ("service_name", data.Service_name)
	case "4":
		var id int
		var kolom, value string
		fmt.Print("Masukan id service yang ingin di edit: ")
		fmt.Scan(&id)
		check := service.CheckDataisExists(id) // cek apakah id data ada
		if !check {
			fmt.Printf("Data dengan id = %d tidak ada\n", id)
			break
		}
		fmt.Println("1. kolom service_name")
		fmt.Println("2. kolom unit_name")
		fmt.Println("3. kolom price_per_unit")
		fmt.Print("Kolom yang ingin di edit: ")
		fmt.Scan(&kolom)
		if kolom == "1" {
			kolom = "service_name"
		} else if kolom == "2" {
			kolom = "unit_name"
		} else if kolom == "3" {
			kolom = "price_per_unit"
		} else {
			fmt.Println("opsi yang kamu pilih tidak ada")
			break
		}
		fmt.Printf("Value %s yang ingin di edit: ", kolom)
		scanner.Scan()
		value = scanner.Text()
		service.UpdateServ(id, kolom, value)
		defer service.SearchServ("id", id)
	case "5":
		var id int
		fmt.Print("Masukan id yang ingin di delete: ")
		fmt.Scan(&id)
		service.DeleteServ(id)
	default:
		fmt.Println("Opsi yang anda masukkan salah!")
	}
	fmt.Print("Kembali ke menu service? Y/n: ")
	fmt.Scan(&opsi)
	if opsi == "y" || opsi == "Y" {
		Menu()
	} else {
		return
	}
}
