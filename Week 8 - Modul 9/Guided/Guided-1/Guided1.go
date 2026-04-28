package main

import "fmt"

func main() {
	var mahasiswa [3]string

	mahasiswa[0] = "Dhimas"

	mahasiswa[1] = "Hafizh"

	var i int

	for i = 0; i < 3; i++ {
		fmt.Printf("Masukkan data index ke-%d : ", i)
		fmt.Scan(&mahasiswa[i])
	}

	fmt.Println("index ke-0 : ", mahasiswa[0])
	fmt.Println("index ke-1 : ", mahasiswa[1])

	var j int

	for j = 0; j < 3; j++ {
		fmt.Println("Data index ke-", j, " : ", mahasiswa[j])
	}

	fmt.Println("index [:1] : ", mahasiswa[:3])

}
