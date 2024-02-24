package tabletdetails

import (
	tabletransaksi "challenge-godb/controller/table_transaksi"
	"challenge-godb/lib/db"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Transaksi_details struct {
	Id          int
	Customer_id int
	Entry_date  time.Time
	Finish_date time.Time
	Received_by string
	Total_price int
	Quantity    int
}

type Transaksi_join struct {
	Customer_name  string
	Received_by    string
	Entry_date     time.Time
	Finish_date    time.Time
	Service_name   string
	Quantity       int
	Unit_name      string
	Price_per_unit int
	Total_price    int
}

func InsertDetails(data Transaksi_details, tx *sql.Tx) int {

	query := "INSERT INTO transaksi_details (customer_id,entry_date,finish_date,received_by,total_price,quantity) VALUES ($1,$2,$3,$4,$5,$6) RETURNING ID"
	var id int = 0
	err := tx.QueryRow(query, data.Customer_id, data.Entry_date, data.Finish_date, data.Received_by, data.Total_price, data.Quantity).Scan(&id)

	validate("INSERT transaksi detail", err, tx)

	return id

}

func UpdateInsert(data Transaksi_details, tx *sql.Tx) {
	query := "UPDATE transaksi_details SET total_price = $1,quantity=$2 WHERE id = $3"
	_, err := tx.Exec(query, data.Total_price, data.Quantity, data.Id)
	validate("UPDATE PRICE transaksi_details", err, tx)
}

func ViewTransaksiJoin(id int) []Transaksi_join {
	db := db.Dbcon()
	defer db.Close()
	datas := []Transaksi_join{}
	query := "SELECT c.customer_name,td.received_by,td.entry_date,td.finish_date,s.service_name,t.quantity,s.unit_name,s.price_per_unit,SUM(s.price_per_unit * t.quantity) AS total_price FROM transaksi AS t JOIN mst_service AS s ON t.service_id = s.id JOIN transaksi_details AS td ON td.id = t.transaction_id JOIN mst_customer AS c ON c.id = td.customer_id WHERE t.transaction_id = $1 GROUP BY c.customer_name,td.received_by, td.entry_date,td.finish_date,s.service_name, t.quantity, s.unit_name, s.price_per_unit;"

	res, err := db.Query(query, id)
	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		data := Transaksi_join{}
		var finishDate sql.NullTime
		err := res.Scan(&data.Customer_name, &data.Received_by, &data.Entry_date, &finishDate, &data.Service_name, &data.Quantity, &data.Unit_name, &data.Price_per_unit, &data.Total_price)
		if err != nil {
			log.Fatal(err)
		}
		if finishDate.Valid {
			data.Finish_date = finishDate.Time
		} else {
			data.Finish_date = time.Time{}
		}
		datas = append(datas, data)
	}
	return datas
}

func DeleteTransaksi(id int, tx *sql.Tx) {
	query := "DELETE FROM transaksi_details WHERE id = $1"
	tabletransaksi.DeleteTransaksi(id, tx)
	_, err := tx.Exec(query, id)
	validate("DELETE transaksi details", err, tx)

	tx.Commit()
	fmt.Println("Transaksi_details berhasil dihapus")
}

// validasi
func validate(msg string, err error, tx *sql.Tx) error {
	if err != nil {
		fmt.Printf("[%s] %s akan ke Rollback!\n", msg, err)
		return tx.Rollback()
	}
	return nil
}

func ValidasiId(id int) {
	if id == 0 {
		log.Fatal("Harap perhatikan id yang anda masukan, tidak boleh 0 atau huruf.")
	}
}
