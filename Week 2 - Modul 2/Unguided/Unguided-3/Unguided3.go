package main

import "fmt"

func main() {
	var totalBeratGram int

	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&totalBeratGram)

	kg := totalBeratGram / 1000
	sisaGram := totalBeratGram % 1000

	biayaKg := kg * 10000
	biayaSisa := 0

	if totalBeratGram > 10000 {
		biayaSisa = 0
	} else if sisaGram >= 500 {
		biayaSisa = sisaGram * 5
	} else if sisaGram > 0 {
		biayaSisa = sisaGram * 15
	}

	totalBiaya := biayaKg + biayaSisa

	fmt.Println("\n===== Detail Perhitungan =====")
	fmt.Printf("Detail berat : %d kg + %d gram\n", kg, sisaGram)
	fmt.Printf("Detail biaya : Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp %d\n", totalBiaya)
}
