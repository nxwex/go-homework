package main

import (
	"fmt"
	"strings"
)

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
	availableCurrency := "В какую валюту? "

	for {
		fmt.Print("Из какой валюты (RUB, USD, EUR): ")
		fmt.Scan(&pairFirst)
		pairFirst = strings.ToLower(pairFirst)
		switch pairFirst {
		case "rub":
			fmt.Println("Выбрана валюта \"RUB\"")
		case "usd":
			fmt.Println("Выбрана валюта \"USD\"")
		case "eur":
			fmt.Println("Выбрана валюта \"EUR\"")
		default:
			fmt.Printf("Ошибка! Пожалуйста, выберите, из доступных валют.\n")
			continue
		}

		fmt.Print("Введите количество средств для конвертации: ")
		if _, err := fmt.Scan(&volume); err != nil {
			fmt.Println("Ошибка! Нужно ввести целое число.")
			continue
		}

		if volume <= 0 {
			fmt.Println("Ошибка! Число должно быть больше нуля.")
			continue
		}

		switch pairFirst {
		case "rub":
			fmt.Print(availableCurrency, "(USD, EUR): ")
		case "usd":
			fmt.Print(availableCurrency, "(RUB, EUR): ")
		case "eur":
			fmt.Print(availableCurrency, "(RUB, USD): ")
		}

		fmt.Scan(&pairSecond)
		pairSecond = strings.ToLower(pairSecond)
		switch pairSecond {
		case pairFirst:
			fmt.Printf("Ошибка! Пожалуйста, выберите, из доступных валют.\n")
			continue
		case "rub":
			fmt.Println("Выбрана валюта \"RUB\"")
		case "usd":
			fmt.Println("Выбрана валюта \"USD\"")
		case "eur":
			fmt.Println("Выбрана валюта \"EUR\"")
		}

		pairFirst, pairSecond = strings.ToUpper(pairFirst), strings.ToUpper(pairSecond)
		fmt.Println("=== Итог ===")
		fmt.Printf("%.2f %s в валюту %s.\n", volume, pairFirst, pairSecond)
		fmt.Println("=========")
		fmt.Print("Подтвердите ваш выбор (y/n): ")
		var choice string
		fmt.Scan(&choice)
		if choice == "y" || choice == "Y" {
			return volume, pairFirst, pairSecond
		} else {
			continue
		}
	}
}

func calculatePair(volume float64, pairFirst, pairSecond string) {
	fmt.Printf("Сумма: %.2f, из валюты %s в валюту %s = в разработке..\n", volume, pairFirst, pairSecond)
	// Как и просили в курсе, пока без действий просто вывод, поскольку разберем это дальше
}
