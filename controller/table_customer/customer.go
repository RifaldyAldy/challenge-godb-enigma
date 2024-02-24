package customer

import (
	"challenge-godb/lib/db"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type Tabel struct {
	Headers []string
	Rows    []Customer
}
type Customer struct {
	Id            int
	Customer_name string
	Phone_number  string
}

func InsertCust(data Customer) {
	db := db.Dbcon()
	var err error
	defer db.Close()
	err = validate(data)
	if err != nil {
		log.Fatal(err)
	}
	query := "INSERT INTO mst_customer (customer_name,phone_number) VALUES ($1,$2)"
	_, err = db.Exec(query, data.Customer_name, data.Phone_number)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Insert data customer berhasil!")
	}
}

func SearchCust(column string, value any) []Customer {
	db := db.Dbcon()
	defer db.Close()
	query := "SELECT id,customer_name,phone_number FROM mst_customer WHERE " + column + " = $1"

	res, err := db.Query(query, value)
	if err != nil {
		log.Fatal(err)
	}
	datas := []Customer{}
	for res.Next() {
		data := Customer{}
		res.Scan(&data.Id, &data.Customer_name, &data.Phone_number)
		datas = append(datas, data)
	}

	return datas
}

func ViewAll() []Customer {
	db := db.Dbcon()
	defer db.Close()
	query := "SELECT id,customer_name,phone_number FROM mst_customer"

	res, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	datas := []Customer{}
	for res.Next() {
		data := Customer{}
		res.Scan(&data.Id, &data.Customer_name, &data.Phone_number)
		datas = append(datas, data)
	}
	return datas
}

func UpdateCust(id int, column, value string) {
	db := db.Dbcon()
	var err error
	defer db.Close()
	data := Customer{}
	if column == "phone_number" {
		data.Phone_number = value
		if err = validate(data); err != nil {
			log.Fatal(err)
		}
	}
	query := "UPDATE mst_customer SET " + column + " = $1 WHERE id = $2"
	_, err = db.Exec(query, value, id)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Update id=%d kolom %s menjadi %s telah berhasil!\n", id, column, value)
	}
}

func DeleteCust(id int) {
	db := db.Dbcon()
	defer db.Close()

	query := "DELETE FROM mst_customer WHERE id = $1"
	check := CheckDataisExists(id)
	if !check {
		fmt.Println("Tidak ada data id = " + strconv.Itoa(id))
		return
	}
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Delete id=%d telah berhasil!\n", id)
	}
}

// validasi

func CheckDataisExists(id int) bool {
	db := db.Dbcon()
	defer db.Close()
	var count int
	query := "SELECT COUNT(*) FROM mst_customer WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}

func validate(data Customer) error {
	if err := lenHp(data); err != nil {
		return err
	}
	if err := formatHp(data); err != nil {
		return err
	}
	return nil
}

func lenHp(data Customer) error {
	lenHp := len(data.Phone_number)
	if lenHp > 15 {
		return errors.New("nomor hp tidak boleh lebih dari 15 karakter")
	}
	return nil
}

func formatHp(data Customer) error {
	format := data.Phone_number[:2]
	if format != "08" {
		return errors.New("nomor hp harus diawali 08+++++++++")
	}
	return nil
}

// method tabel

func (t *Tabel) GenerateTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(t.Headers)
	for i, value := range t.Rows {
		table.Append([]string{strconv.Itoa(i + 1), strconv.Itoa(value.Id), value.Customer_name, value.Phone_number})
	}
	table.Render()
}
