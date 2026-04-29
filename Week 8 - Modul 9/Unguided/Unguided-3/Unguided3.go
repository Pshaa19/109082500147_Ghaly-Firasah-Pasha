package main

import "fmt"

const NMAX = 100

func main() {
	var klubA, klubB string
	var hasil [NMAX]string
	var n int = 0

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	var skorA, skorB int
	i := 1

	for {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		if skorA > skorB {
			fmt.Println("Hasil:", klubA)
			hasil[n] = klubA
			n++
		} else if skorB > skorA {
			fmt.Println("Hasil:", klubB)
			hasil[n] = klubB
			n++
		} else {
			fmt.Println("Hasil: Draw")
		}

		i++
	}

	fmt.Println("\nDaftar pemenang:")
	for j := 0; j < n; j++ {
		fmt.Println(hasil[j])
	}
}
