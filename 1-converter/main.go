package main

import "fmt"

func main() {
	const ConvUSDtoEUR float64 = 0.84
	const ConvUSDtoRUB float64 = 77.50
	const ConvEURtoRUB float64 = ConvUSDtoRUB / ConvUSDtoEUR
	var volume float64
	var pairFirst string
	var pairSecond string
	calculatePair(inputUser(volume, pairFirst, pairSecond))
}

func inputUser(volume float64, pairFirst, pairSecond string) (float64, string, string) {
	fmt.Print("Введите количество средств для конвертации: ")
	fmt.Scan(&volume)
	fmt.Print("Из какой валюты (RUB, USD, EUR): ")
	fmt.Scan(&pairFirst)
	fmt.Print("В какую валюту? (RUB, USD, EUR): ")
	fmt.Scan(&pairSecond)
	return volume, pairFirst, pairSecond
}

func calculatePair(volume float64, pairFirst, pairSecond string) {
	fmt.Printf("Сумма: %.2f, из валюты %s в валюту %s = в разработке..\n", volume, pairFirst, pairSecond)
	// Как и просили в курсе, пока без действий просто вывод, поскольку разберем это дальше
}
