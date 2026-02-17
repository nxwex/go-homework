package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Калькулятор индекса массы тела ===")
	for {
		userKg, userHeight := getUserInput()
		imt, err := calculateIMT(userKg, userHeight)
		if err != nil {
			fmt.Println(err)
			continue
		}
		outputResult(imt)
		if !checkRepeatCalculation() {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.2f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calculateIMT(userKg, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("101: Invalid input: weight or height should be greater than zero")
	}
	const IMTPower = 2
	imt := userKg / math.Pow(userHeight/100, IMTPower)
	return imt, nil
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Printf("Хотите проверить еще раз? (y/n):")
	fmt.Scan(&userChoice)
	fmt.Println("======")
	if userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false
}
