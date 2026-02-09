package main

import "fmt"

func main() {
	const ConvUSDtoEUR float64 = 0.84
	const ConvUSDtoRUB float64 = 77.50
	var eur float64 = 10.0
	rub := eur / ConvUSDtoEUR * ConvUSDtoRUB
	fmt.Printf("Курс доллара к евро = %v, курс доллара к рублю = %v, количество евро = %2.f, получается %2.f руб.\n", ConvUSDtoEUR, ConvUSDtoRUB, eur, rub)
}
