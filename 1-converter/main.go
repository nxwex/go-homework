package main

import (
	"fmt"
	"strings"
)

func main() {
	userMenu()
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
		default:
			fmt.Println("Ошибка! Выбрана недоступная валюта. Повторите попытку заново.")
			continue
		}

		pairFirst, pairSecond = strings.ToUpper(pairFirst), strings.ToUpper(pairSecond)
		clearTerminal()
		fmt.Println("=== Итог ===")
		fmt.Printf("%.2f %s в валюту %s.\n", volume, pairFirst, pairSecond)
		fmt.Println("=========")
		fmt.Print("Подтвердите ваш выбор (y/n): ")
		var choice string
		fmt.Scan(&choice)
		if choice == "y" || choice == "Y" {
			return volume, pairFirst, pairSecond
		} else {
			clearTerminal()
			continue
		}
	}
}

func calculatePair(volume float64, pairFirst, pairSecond string) {
	var choice int
	const ConvUSDtoEUR float64 = 0.84
	const ConvUSDtoRUB float64 = 77.50
	const ConvEURtoRUB float64 = ConvUSDtoRUB / ConvUSDtoEUR
	convVolume := volume

	switch {
	case pairFirst == "USD" && pairSecond == "EUR":
		convVolume *= ConvUSDtoEUR
	case pairFirst == "USD" && pairSecond == "RUB":
		convVolume *= ConvUSDtoRUB
	case pairFirst == "EUR" && pairSecond == "RUB":
		convVolume *= ConvEURtoRUB
	case pairFirst == "EUR" && pairSecond == "USD":
		convVolume /= ConvUSDtoEUR
	case pairFirst == "RUB" && pairSecond == "EUR":
		convVolume /= ConvEURtoRUB
	case pairFirst == "RUB" && pairSecond == "USD":
		convVolume /= ConvUSDtoRUB
	default:
		fmt.Println("Ошибка! Не найден курс для конвертации.")
	}
	clearTerminal()
	fmt.Printf("Сумма: %.2f, из валюты %s в валюту %s = %.2f\n", volume, pairFirst, pairSecond, convVolume)
	fmt.Print("\n1. Конвертировать еще раз")
	fmt.Print("\n2. Вернуться в меню")
	fmt.Print("\n0. Выход\n")
	fmt.Print("\nВыберите пункт меню: ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		clearTerminal()
		fmt.Println("=== Конвертер валют ===")
		calculatePair(inputUser(volume, pairFirst, pairSecond))
	case 2:
		userMenu()
	case 0:
		break
	default:
		fmt.Println("Ошибка! Такого пункта меню не существует.")
	}
}

func userMenu() {
	var volume float64
	var pairFirst string
	var pairSecond string
	var choice int
	clearTerminal()
	fmt.Println("=== Конвертер валют - МЕНЮ ===")
	fmt.Println("1. Конвертировать валюты")
	fmt.Println("2. Актуальный курс валют")
	fmt.Println("0. Выход")
	fmt.Println("============")
	fmt.Print("\nВыберите пункт меню: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		clearTerminal()
		fmt.Println("=== Конвертер валют ===")
		calculatePair(inputUser(volume, pairFirst, pairSecond))
	case 2:
		currentExchange()
	case 0:
		break
	default:
		fmt.Println("Ошибка! Такого пункта меню не существует.")
	}
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J") // ANSI код для очистки консоли
}

func currentExchange() {
	const ConvUSDtoEUR float64 = 0.84
	const ConvUSDtoRUB float64 = 77.50
	const ConvEURtoRUB float64 = ConvUSDtoRUB / ConvUSDtoEUR
	var choice int

	clearTerminal()
	fmt.Println("=== Актуальный курс валют к рублю ===")
	fmt.Printf("USD = %.2f\n", ConvUSDtoRUB)
	fmt.Printf("EUR = %.2f\n", ConvEURtoRUB)
	fmt.Printf("\n1. Вернуться в меню\n")
	fmt.Printf("0. Выход\n")
	fmt.Print("\nВыберите пункт меню: ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		userMenu()
	case 0:
		break
	default:
		fmt.Println("Ошибка! Такого пункта меню не существует.")
	}
}
