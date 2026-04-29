package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0

	for {
		fmt.Scanf("%c", &ch)

		if ch == '.' || *n >= NMAX {
			break
		}

		if ch != '\n' && ch != ' ' {
			t[*n] = ch
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Println("Masukkan karakter (akhiri dengan titik):")

	isiArray(&tab, &m)

	fmt.Print("Isi array: ")
	cetakArray(tab, m)

	if palindrom(tab, m) {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Bukan palindrom")
	}

	balikanArray(&tab, m)

	fmt.Print("Setelah dibalik: ")
	cetakArray(tab, m)
}
