package main

import "fmt"

func cekKabisat(tahun int) bool {
	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		return true
	}
	return false
}

func main() {
	var tahun int

	fmt.Print("Tahun : ")
	fmt.Scanln(&tahun)

	isKabisat := cekKabisat(tahun)

	if isKabisat {
		fmt.Println("Kabisat : True")
	} else {
		fmt.Println("Kabisat : False")
	}
}
