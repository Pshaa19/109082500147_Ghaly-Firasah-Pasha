package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func cariNilaiPertama(arr arrayMahasiswa, n int, targetNIM string) int {
	for i := 0; i < n; i++ {
		if arr[i].NIM == targetNIM {
			return arr[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(arr arrayMahasiswa, n int, targetNIM string) int {
	maxNilai := -1
	for i := 0; i < n; i++ {
		if arr[i].NIM == targetNIM {
			if arr[i].nilai > maxNilai {
				maxNilai = arr[i].nilai
			}
		}
	}
	return maxNilai
}

func main() {
	var dataMhs arrayMahasiswa
	var n int
	var nimDicari string

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	if n > nMax {
		n = nMax
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&dataMhs[i].NIM, &dataMhs[i].nama, &dataMhs[i].nilai)
	}

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimDicari)

	nilaiPertama := cariNilaiPertama(dataMhs, n, nimDicari)
	nilaiTerbesar := cariNilaiTerbesar(dataMhs, n, nimDicari)

	if nilaiPertama != -1 {
		fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", nimDicari, nilaiPertama)
		fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", nimDicari, nilaiTerbesar)
	} else {
		fmt.Printf("Data dengan NIM %s tidak ditemukan.\n", nimDicari)
	}
}
