package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	runApp()
}

func runApp() {
	clearTerminal()
	fmt.Println("=== Калькулятор ===")
	for {
		fmt.Println("Какую операцию желаете выполнить? (AVG, SUM, MED)")
		fmt.Println("Для выхода напишите 'Exit'")
		userOperationInput := getValidOperation()
		if userOperationInput == "EXIT" {
			clearTerminal()
			return
		}

		userNumbers := getValidNumbers()

		clearTerminal()
		result := fmt.Sprintf("Ваш результат: %v", calculateUserInput(userNumbers, userOperationInput))
		fmt.Println(result)
		fmt.Println("===")
		if !repeatRunApp() {
			clearTerminal()
			break
		}
	}

}

func calculateAvg(userNumbers []float64) float64 {
	var sumNumbers float64

	if len(userNumbers) == 0 {
		return 0
	}

	for _, value := range userNumbers {
		sumNumbers += value
	}
	result := sumNumbers / float64(len(userNumbers))
	return result
}

func calculateSum(userNumbers []float64) float64 {
	var sumNumbers float64
	for _, value := range userNumbers {
		sumNumbers += value
	}
	return sumNumbers
}

// считает медиану в слайсе
func calculateMed(userNumbers []float64) float64 {
	sort.Float64s(userNumbers)
	length := len(userNumbers)

	if length == 0 {
		return 0
	}

	if length%2 != 0 {
		return userNumbers[length/2]
	} else {
		return ((userNumbers[length/2-1] + userNumbers[length/2]) / 2)
	}
}

// преобразует строку с числами через запятую в слайс float64
func transformUserInput(userInput string) []float64 {
	parts := strings.Split(userInput, ",")
	userNumbers := make([]float64, 0, 8)

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)

		if trimmed == "" {
			continue
		}

		val, err := strconv.ParseFloat(trimmed, 64)
		if err != nil {
			fmt.Printf("Ошибка при чтении '%s': %v\n", trimmed, err)
			continue
		}
		userNumbers = append(userNumbers, val)
	}
	return userNumbers
}

// вызывает функцию в зависимости от выбора
func calculateUserInput(userNumbers []float64, operation string) float64 {
	var result float64
	switch operation {
	case "AVG":
		result = calculateAvg(userNumbers)
	case "SUM":
		result = calculateSum(userNumbers)
	case "MED":
		result = calculateMed(userNumbers)
	default:
		fmt.Println("Ошибка! Операция не найдена.")
	}
	return result
}

func scanUserNumInput() string {
	var userInput string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите числа через запятую: ")
	if scanner.Scan() {
		userInput = scanner.Text()
	}
	return userInput
}

func scanUserOperationInput() string {
	var userInput string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nВвод: ")
	if scanner.Scan() {
		userInput = strings.ToUpper(scanner.Text())
	}
	return userInput
}

func getValidOperation() string {
	for {
		operation := scanUserOperationInput()
		if operation == "SUM" || operation == "AVG" || operation == "MED" || operation == "EXIT" {
			return operation
		}
		fmt.Println("Ошибка! Операция не найдена.")
	}
}

func getValidNumbers() []float64 {
	for {
		input := scanUserNumInput()
		nums := transformUserInput(input)
		if len(nums) > 0 {
			return nums
		}
		fmt.Println("Ошибка! Вы не ввели ни одного валидного числа.")
	}
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func repeatRunApp() bool {
	var choice string
	for {
		fmt.Println("Желаете поворить рассчет? (y/n)")
		fmt.Print("\nВвод:")
		fmt.Scan(&choice)

		if strings.EqualFold(choice, "y") {
			clearTerminal()
			return true
		} else if strings.EqualFold(choice, "n") {
			clearTerminal()
			return false
		}

		fmt.Println("Ошибка! Введите 'y' для продолжения или 'n' для выхода.")
	}
}
