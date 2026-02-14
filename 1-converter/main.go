package main

import (
	"fmt"
	"strings"
)

const ConvUSDEUR float64 = 0.84
const ConvUSDRUB float64 = 77.50
const ConvEURRUB float64 = 98.0

func main() {
	userMenu()
}

func userMenu() {
	for {
		var choice int
		clearTerminal()
		fmt.Println("=== Конвертер валют - МЕНЮ ===")
		fmt.Println("1. Конвертировать валюты")
		fmt.Println("2. Актуальный курс валют")
		fmt.Println("3. Github")
		fmt.Println("0. Выход")
		fmt.Println("============")
		fmt.Print("\nВыберите пункт меню: ")

		if _, err := fmt.Scan(&choice); err != nil {
			continue
		}

		switch choice {
		case 1:
			volume, pairFirst, pairSecond := inputUser()
			calculatePair(volume, pairFirst, pairSecond)
		case 2:
			currentExchange()
		case 3:
			showGithub()
		case 0:
			clearTerminal()
			return
		}
	}
}

func inputUser() (float64, string, string) {
	clearTerminal()
	pairFirst := inputCurrency("Из какой валюты ", "")
	volume := inputAmount()
	pairSecond := inputCurrency("В какую валюту: ", pairFirst)

	fmt.Println("=== Итог ===")
	fmt.Printf("%.2f %s в валюту %s.\n", volume, pairFirst, pairSecond)
	fmt.Println("=========")
	return volume, pairFirst, pairSecond
}

func inputCurrency(label string, pairFirst string) string {
	availableCurrency := "В какую валюту? "
	var current string
	for {
		if pairFirst == "" {
			fmt.Print(label, "(RUB, USD, EUR): ")
		} else {
			switch strings.ToLower(pairFirst) {
			case "rub":
				fmt.Print(availableCurrency, "(USD, EUR): ")
			case "usd":
				fmt.Print(availableCurrency, "(RUB, EUR): ")
			case "eur":
				fmt.Print(availableCurrency, "(RUB, USD): ")
			}
		}

		fmt.Scan(&current)
		current = strings.ToLower(current)

		switch current {
		case "rub", "usd", "eur":
			if current == strings.ToLower(pairFirst) {
				fmt.Printf("Ошибка! Пожалуйста, выберите, из доступных валют.\n")
				continue
			}
			fmt.Printf("Выбрана валюта \"%s\"\n", strings.ToUpper(current))
			return strings.ToUpper(current)
		default:
			if pairFirst == "" {
				fmt.Printf("Ошибка! Пожалуйста, выберите, из доступных валют.\n")
			} else {
				fmt.Println("Ошибка! Выбрана недоступная валюта. Повторите попытку заново.")
			}
			continue
		}
	}
}

func inputAmount() float64 {
	var volume float64
	for {
		fmt.Print("Введите количество средств для конвертации: ")
		if _, err := fmt.Scan(&volume); err != nil {
			fmt.Println("Ошибка! Нужно ввести целое число.")
			// добавил защиту от спама, если пользователь введет символ, отличающийся от числа, то fmt.Scanln скушает и выведет ошибку
			var trash string
			fmt.Scanln(&trash)
			continue
		}

		if volume <= 0 {
			fmt.Println("Ошибка! Число должно быть больше нуля.")
			continue
		}
		return volume
	}
}

func calculatePair(volume float64, pairFirst, pairSecond string) {
	for {
		var choice int
		convVolume := volume

		rates := map[string]float64{
			"USDEUR": ConvUSDEUR,
			"USDRUB": ConvUSDRUB,
			"EURRUB": ConvEURRUB,
			"EURUSD": 1 / ConvUSDEUR, // обратные пары, сразу записываем результат деления
			"RUBUSD": 1 / ConvUSDRUB,
			"RUBEUR": 1 / ConvEURRUB,
		}

		pair := pairFirst + pairSecond
		if rate, ok := rates[pair]; ok {
			convVolume *= rate
		} else {
			fmt.Println("Валютная пара не найдена")
		}

		clearTerminal()
		fmt.Printf("Сумма: %.2f, из валюты %s в валюту %s = %.2f\n", volume, pairFirst, pairSecond, convVolume)
		fmt.Print("\n1. Конвертировать еще раз")
		fmt.Print("\n0. Вернуться в меню\n")
		fmt.Print("\nВыберите пункт меню: ")
		fmt.Scan(&choice)

		if choice == 1 {
			clearTerminal()
			fmt.Println("=== Конвертер валют ===")
			volume, pairFirst, pairSecond = inputUser()
			continue
		} else if choice == 0 {
			return
		}
	}
}

func currentExchange() {
	var choice int
	clearTerminal()
	fmt.Println("=== Актуальный курс валют к рублю ===")
	fmt.Printf("USD = %.2f\n", ConvUSDRUB)
	fmt.Printf("EUR = %.2f\n", ConvEURRUB)
	fmt.Printf("\n0. Вернуться в меню\n")
	fmt.Print("\nВыберите пункт меню: ")
	fmt.Scan(&choice)
	if choice == 0 {
		return
	}
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func showGithub() {
	var choice int
	clearTerminal()
	fmt.Print("link - https://github.com/nxwex\n")
	fmt.Printf("\n0. Вернуться в меню\n")
	fmt.Print("\nВыберите пункт меню: ")
	fmt.Scan(&choice)
	if choice == 0 {
		return
	}
}
