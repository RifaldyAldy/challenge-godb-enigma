package tableservice

import (
	"challenge-godb/lib/db"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Tabel struct {
	Rows []Service
}

type Service struct {
	Id             int
	Service_name   string
	Unit_name      string
	Price_per_unit int
}

func ViewAll() []Service {
	db := db.Dbcon()
	defer db.Close()
	datas := []Service{}
	query := "SELECT id, service_name, unit_name, price_per_unit FROM mst_service"
	res, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {
		data := Service{}
		res.Scan(&data.Id, &data.Service_name, &data.Unit_name, &data.Price_per_unit)
		datas = append(datas, data)
	}

	// inisiasi method tampilan tabel
	table := Tabel{Rows: datas}
	table.GenerateTable()
	return datas
}

func SearchServ(column string, value any) []Service {
	db := db.Dbcon()
	defer db.Close()

	query := "SELECT id,service_name,unit_name,price_per_unit FROM mst_service WHERE " + column + " = $1"
	datas := []Service{}
	res, err := db.Query(query, value)
	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {
		data := Service{}
		res.Scan(&data.Id, &data.Service_name, &data.Unit_name, &data.Price_per_unit)
		datas = append(datas, data)
	}
	return datas
}

func InsertServ(data Service) {
	db := db.Dbcon()
	defer db.Close()

	query := "INSERT INTO mst_service (service_name,unit_name,price_per_unit) VALUES ($1,$2,$3)"
	if check := CheckServiceIsExist(data.Service_name); check {
		fmt.Println("Service sudah ada dalam database!")
		return
	}
	if check := validateUnit(data.Unit_name); !check {
		fmt.Println("Unit data harus KG/Buah")
		return
	}
	_, err := db.Exec(query, data.Service_name, data.Unit_name, data.Price_per_unit)
	if err != nil {
		panic(err)
	} else {
		tabel := Tabel{Rows: SearchServ("service_name", data.Service_name)}
		tabel.GenerateTable()
		fmt.Println("Insert service berhasil!")
	}
}

func UpdateServ(id int, column string, value any) {
	db := db.Dbcon()
	defer db.Close()

	query := "UPDATE mst_service SET " + column + " = $1 WHERE id = $2"
	if check := CheckDataisExists(id); !check {
		fmt.Println("Tidak ada data service dengan id = ", id)
		return
	}
	if column == "unit_name" {
		if check := validateUnit(value); !check {
			fmt.Println("Unit data harus KG/Buah")
			return
		}
	}
	_, err := db.Exec(query, value, id)
	if err != nil {
		log.Fatal(err)
	} else {
		tabel := Tabel{Rows: SearchServ("id", id)}
		tabel.GenerateTable()
		fmt.Println("Update service berhasil!")
	}
}

func DeleteServ(id int) {
	db := db.Dbcon()
	defer db.Close()

	if check := CheckDataisExists(id); !check {
		fmt.Println("Tidak ada data service dengan id = ", id)
		return
	}
	query := "DELETE FROM mst_service WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Delete service berhasil!")
	}
}

// validasi

func CheckDataisExists(id int) bool {
	db := db.Dbcon()
	defer db.Close()

	count := 0
	query := "SELECT COUNT(*) FROM mst_service WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}

func CheckServiceIsExist(name string) bool {
	db := db.Dbcon()
	defer db.Close()
	var count int
	query := "SELECT COUNT(*) FROM mst_service WHERE service_name = $1"
	err := db.QueryRow(query, name).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}

func validateUnit(unit interface{}) bool {
	enum := [2]string{"kg", "buah"}
	unit = strings.ToLower(unit.(string))
	for _, value := range enum {
		if value == unit {
			return true
		}
	}
	return false
}

// function table view data

func (t *Tabel) GenerateTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NO", "ID", "Service name", "Unit name", "Price per unit"})
	for i, value := range t.Rows {
		table.Append([]string{strconv.Itoa(i + 1), strconv.Itoa(value.Id), value.Service_name, value.Unit_name, strconv.Itoa(value.Price_per_unit)})
	}
	table.Render()
}
