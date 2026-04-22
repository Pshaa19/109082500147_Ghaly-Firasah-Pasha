package main

import "fmt"

type angka int

type kata string

type Buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunTerbit   angka
	jumlahHalaman angka
}

func main() {
	var dataBuku Buku

	fmt.Println("=== INPUT BIODATA BUKU ===")
	fmt.Print("Masukkan judul buku : ")
	fmt.Scan(&dataBuku.judul)

	fmt.Print("Masukkan nama penulis : ")
	fmt.Scan(&dataBuku.penulis)

	fmt.Print("Masukkan penerbit : ")
	fmt.Scan(&dataBuku.penerbit)

	fmt.Print("Masukkan tahun terbit : ")
	fmt.Scan(&dataBuku.tahunTerbit)

	fmt.Print("Masukkan jumlah halaman: ")
	fmt.Scan(&dataBuku.jumlahHalaman)

	fmt.Println()

	fmt.Println("=== BIODATA BUKU ===")
	fmt.Printf("Judul Buku : %s\n", dataBuku.judul)
	fmt.Printf("Penulis : %s\n", dataBuku.penulis)
	fmt.Printf("Penerbit : %s\n", dataBuku.penerbit)
	fmt.Printf("Tahun Terbit : %d\n", dataBuku.tahunTerbit)
	fmt.Printf("Jumlah Halaman : %d\n", dataBuku.jumlahHalaman)
}
