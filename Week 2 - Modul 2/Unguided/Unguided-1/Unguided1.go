package main

import "fmt"

func main() {
	var tahun int

	fmt.Print("Tahun : ")
	fmt.Scan(&tahun)

	isKabisat := (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

	if isKabisat {
		fmt.Println("Kabisat : True")
	} else {
		fmt.Println("Kabisat : False")
	}
}
