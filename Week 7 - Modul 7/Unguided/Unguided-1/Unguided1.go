package main

import "fmt"

type suhu float64

func CelciusToReamur(c suhu) suhu {
	return (4.0 / 5.0) * c
}

func CelciusToFahrenheit(c suhu) suhu {
	return (9.0/5.0)*c + 32.0
}

func CelciusToKelvin(c suhu) suhu {
	return c + 273.15
}

func main() {
	var c suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&c)

	fmt.Println()

	fmt.Printf("%v celcius = %v reamur\n", c, CelciusToReamur(c))
	fmt.Printf("%v celcius = %v fahrenheit\n", c, CelciusToFahrenheit(c))
	fmt.Printf("%v celcius = %v kelvin\n", c, CelciusToKelvin(c))
}
