package tabletdetails

import (
	tableservice "challenge-godb/controller/table_service"
	tabletransaksi "challenge-godb/controller/table_transaksi"
	"challenge-godb/lib/db"
	"log"
)

func CommitInsert(T_Details Transaksi_details, T []tabletransaksi.Transaksi) {
	db := db.Dbcon()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	T_Details.Id = InsertDetails(T_Details, tx)
	for _, value := range T {
		value.Transaction_id = T_Details.Id
		pricePerUnit := tableservice.SearchServ("id", value.Service_id)[0].Price_per_unit
		value.Total_price = pricePerUnit * value.Quantity
		tabletransaksi.InsertTr(value, tx)
		T_Details.Total_price += value.Total_price
		T_Details.Quantity += value.Quantity
	}
	UpdateInsert(T_Details, tx)

	tx.Commit()
}
