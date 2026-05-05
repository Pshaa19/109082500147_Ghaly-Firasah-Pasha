package main

import "fmt"

func main() {
	var x, y int
	var ikan [1000]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	i := 0
	jumlahWadah := 0
	totalSemua := 0.0

	for i < x {
		total := 0.0
		count := 0

		for count < y && i < x {
			total += ikan[i]
			i++
			count++
		}

		fmt.Printf("%.2f ", total)
		totalSemua += total
		jumlahWadah++
	}

	fmt.Println()

	rata := totalSemua / float64(jumlahWadah)
	fmt.Printf("%.2f\n", rata)
}
