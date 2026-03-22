package main

import "fmt"

func hitungBiaya(jenis string, jamMasuk int, jamKeluar int) int {
	totalBiaya := 0

	for i := jamMasuk; i < jamKeluar; i++ {
		if jenis == "motor" {
			if i < 17 {
				totalBiaya += 4000
			} else {
				totalBiaya += 5000
			}
		} else if jenis == "mobil" {
			if i < 17 {
				totalBiaya += 6000
			} else {
				totalBiaya += 7000
			}
		}
	}

	return totalBiaya
}

func main() {
	var jenis string
	var jamMasuk, jamKeluar int

	fmt.Println("---- Rekap Tarif Parkir Cafe Per Hari ----")

	totalPendapatan := 0
	kendaraanKe := 1

	for {
		fmt.Printf("\n*Kendaraan %d\n", kendaraanKe)

		fmt.Print("Kendaraan apa? (motor/mobil/- untuk selesai): ")
		fmt.Scan(&jenis)

		if jenis == "-" {
			break
		}

		fmt.Print("Masukkan Jam Masuk Kendaraan (0-24): ")
		fmt.Scan(&jamMasuk)

		fmt.Print("Masukkan Jam Keluar Kendaraan (0-24): ")
		fmt.Scan(&jamKeluar)

		biaya := hitungBiaya(jenis, jamMasuk, jamKeluar)

		fmt.Printf("Biaya parkir %s %d: %d\n", jenis, kendaraanKe, biaya)
		fmt.Println("========================================")

		totalPendapatan += biaya
		kendaraanKe++
	}

	fmt.Printf("\n*** Total Pendapatan Parkir Hari Ini Adalah %d ***\n", totalPendapatan)
}
