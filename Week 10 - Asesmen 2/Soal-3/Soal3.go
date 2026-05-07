package main

import "fmt"

const nProv int = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("==== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	maxIdx := 0
	maxTumbuh := tumbuh[0]
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > maxTumbuh {
			maxTumbuh = tumbuh[i]
			maxIdx = i
		}
	}
	return maxIdx
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("\n==== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := float64(pop[i]) * (tumbuh[i] + 1.0)
			fmt.Printf("%s %d\n", prov[i], int(prediksi))
		}
	}
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv
	var cariProvinsi string

	InputData(&prov, &pop, &tumbuh)

	fmt.Scan(&cariProvinsi)

	idxTercepat := ProvinsiTercepat(tumbuh)
	fmt.Printf("\nProvinsi dengan angka pertumbuhan tercepat : %s\n", prov[idxTercepat])

	idxDicari := IndeksProvinsi(prov, cariProvinsi)
	if idxDicari != -1 {
		fmt.Printf("\nData provinsi yang dicari : %s\n", prov[idxDicari])
	} else {
		fmt.Printf("\nData provinsi yang dicari : %s tidak ditemukan\n", cariProvinsi)
	}

	Prediksi(prov, pop, tumbuh)
}
