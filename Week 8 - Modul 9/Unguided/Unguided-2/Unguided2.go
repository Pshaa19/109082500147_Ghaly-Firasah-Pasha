package main

import (
	"fmt"
	"math"
)

const NMAX = 100

func main() {
	var arr [NMAX]int
	var n int

	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("Semua elemen: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)

	fmt.Print("Indeks kelipatan ", x, ": ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var idx int
	fmt.Print("Masukkan indeks yang akan dihapus: ")
	fmt.Scan(&idx)

	for i := idx; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--

	fmt.Print("Array setelah dihapus: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var sum int
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	mean := float64(sum) / float64(n)
	fmt.Println("Rata-rata:", mean)

	var total float64
	for i := 0; i < n; i++ {
		selisih := float64(arr[i]) - mean
		total += selisih * selisih
	}
	stddev := math.Sqrt(total / float64(n))
	fmt.Println("Standar deviasi:", stddev)

	var cari, freq int
	fmt.Print("Masukkan bilangan yang dicari: ")
	fmt.Scan(&cari)

	for i := 0; i < n; i++ {
		if arr[i] == cari {
			freq++
		}
	}
	fmt.Println("Frekuensi:", freq)
}
