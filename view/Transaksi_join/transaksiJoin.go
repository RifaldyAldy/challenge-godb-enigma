package transaksijoin

import (
	tabletdetails "challenge-godb/controller/table_t_details"
	"fmt"
	"os"

	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"
)

func TransaksiJoin(id int) {
	tabletdetails.ValidasiId(id)
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	transactions := tabletdetails.ViewTransaksiJoin(id)

	// Membuat tabel baru
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Customer Name", "Received By", "Service Name", "Quantity", "Unit Name", "Price Per Unit", "Total Price"})
	total := 0
	// Iterasi melalui data transaksi dan menambahkan baris ke tabel
	for _, t := range transactions {
		row := []string{
			t.Customer_name,
			t.Received_by,
			t.Service_name,
			fmt.Sprintf("%d", t.Quantity),
			t.Unit_name,
			ac.FormatMoney(t.Price_per_unit),
			ac.FormatMoney(t.Total_price),
		}
		total += t.Total_price
		table.Append(row)
	}
	table.SetFooter([]string{"", "", "", "", "", "TOTAL HARGA : ", ac.FormatMoney(total)})
	// Mengatur beberapa properti tabel
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	table.SetAutoMergeCells(true)
	table.SetAutoFormatHeaders(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	// Menampilkan tabel ke konsol
	table.Render()
}
