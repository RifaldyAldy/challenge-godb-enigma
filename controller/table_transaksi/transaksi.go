package tabletransaksi

import (
	"challenge-godb/lib/db"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Tabel struct {
	Headers []string
	Rows    []Transaksi_join
}

type Transaksi struct {
	Id             int
	Transaction_id int
	Service_id     int
	Quantity       int
	Total_price    int
}

type Transaksi_join struct {
	No_id         int
	Customer_name string
	Entry_date    time.Time
	Finish_date   time.Time
	Received_by   string
}

func InsertTr(data Transaksi, tx *sql.Tx) {
	query := "INSERT into transaksi (transaction_id,service_id,quantity,total_price) VALUES ($1,$2,$3,$4)"

	_, err := tx.Exec(query, data.Transaction_id, data.Service_id, data.Quantity, data.Total_price)

	validate("Insert Transaksi", err, tx)

}

func SearchTransaksi(id int) Transaksi_join {
	db := db.Dbcon()
	defer db.Close()
	query := "SELECT td.id,c.customer_name,td.entry_date,td.finish_date,td.received_by FROM transaksi_details AS td JOIN mst_customer AS c ON td.customer_id = c.id WHERE td.id =$1"

	finishDate := sql.NullTime{}
	data := Transaksi_join{}
	err := db.QueryRow(query, id).Scan(&data.No_id, &data.Customer_name, &data.Entry_date, &finishDate, &data.Received_by)
	if finishDate.Valid {
		data.Finish_date = finishDate.Time
	} else {
		data.Finish_date = time.Time{}
	}
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ViewALl() []Transaksi_join {
	db := db.Dbcon()
	defer db.Close()
	datas := []Transaksi_join{}
	query := "SELECT td.id,c.customer_name,td.entry_date,td.finish_date,td.received_by FROM transaksi_details AS td JOIN mst_customer AS c ON td.customer_id = c.id"

	res, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {
		data := Transaksi_join{}
		res.Scan(&data.No_id, &data.Customer_name, &data.Entry_date, &data.Finish_date, &data.Received_by)
		datas = append(datas, data)
	}

	return datas
}

func DeleteTransaksi(id int, tx *sql.Tx) {
	query := "DELETE FROM transaksi WHERE transaction_id = $1"

	_, err := tx.Exec(query, id)
	validate("DELETE transaksi", err, tx)
}

func validate(msg string, err error, tx *sql.Tx) error {
	if err != nil {
		fmt.Printf("[%s] %s, akan ke Rollback!\n", msg, err)
		return tx.Rollback()
	}
	return nil
}

func (t *Tabel) GenerateTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(t.Headers)
	for _, value := range t.Rows {
		table.Append([]string{strconv.Itoa(value.No_id), value.Customer_name, value.Entry_date.Format("January 02,2006"), value.Finish_date.Format("January 02,2006"), value.Received_by})
	}

	table.Render()
}
