package main

import "fmt"

func main() {
	var satu, dua, tiga string
	var temp string

	fmt.Print("Masukkan input: ")
	fmt.Scanln(&satu)
	fmt.Println("Masukkan input: ")
	fmt.Scanln(&dua)
	fmt.Println("Masukkan input: ")
	fmt.Scanln(&tiga)

	fmt.Println("Output Awal: " + satu + " " + dua + " " + tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output Akhir: " + satu + " " + dua + " " + tiga)
}
