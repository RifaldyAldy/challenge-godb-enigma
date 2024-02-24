package menu

import (
	customerMenu "challenge-godb/menu/customer_menu"
	servicemenu "challenge-godb/menu/service_menu"
	transaksimenu "challenge-godb/menu/transaksi_menu"
	"fmt"
)

func Menu() {
	var opsi string
	fmt.Println("1. Menu Customer")
	fmt.Println("2. Menu Service")
	fmt.Println("3. Menu transaksi")
	fmt.Print("Pilih 1-2 untuk pergi ke menu: ")
	fmt.Scan(&opsi)
	switch opsi {
	case "1":
		customerMenu.Menu()
	case "2":
		servicemenu.Menu()
	case "3":
		transaksimenu.Menu()
	}
}
