package main

import "fmt"

func polaBintang(n int) {
	if n == 0 {
		return
	} else {
		polaBintang(n - 1)
		for i := 0; i < n; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var n int

	fmt.Scan(&n)
	polaBintang(n)
}
