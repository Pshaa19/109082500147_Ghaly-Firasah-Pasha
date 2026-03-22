package main

import "fmt"

func hitungPersegi(sisi int) {
	luas := sisi * sisi
	keliling := 4 * sisi
	fmt.Printf("Luas persegi : %d\n", luas)
	fmt.Printf("Keliling persegi : %d\n", keliling)
}

func hitungPersegiPanjang(panjang, lebar int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	fmt.Printf("Luas persegi panjang : %d\n", luas)
	fmt.Printf("Keliling persegi panjang : %d\n", keliling)
}

func hitungLingkaran(jarijari float64) {
	var pi float64 = 3.14
	luas := pi * jarijari * jarijari
	keliling := 2 * pi * jarijari

	fmt.Printf("Luas lingkaran : %v\n", luas)
	fmt.Printf("Keliling lingkaran : %v\n", keliling)
}

func main() {
	var pilihan, sisi int
	var panjang, lebar int
	var jarijari float64

	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)

	fmt.Println()

	switch pilihan {
	case 1:
		fmt.Print("Masukkan sisi : ")
		fmt.Scan(&sisi)
		hitungPersegi(sisi)
	case 2:
		fmt.Print("Masukkan panjang : ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan lebar : ")
		fmt.Scan(&lebar)
		hitungPersegiPanjang(panjang, lebar)
	case 3:
		fmt.Print("Masukkan jari-jari : ")
		fmt.Scan(&jarijari)
		hitungLingkaran(jarijari)
	default:
		fmt.Println("Pilihan tidak valid. Silakan masukkan angka 1, 2, atau 3.")
	}
}
