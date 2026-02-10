package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Калькулятор индекса массы тела ===")
	for {
		userKg, userHeight := getUserInput()
		IMT := calculateIMT(userKg, userHeight)
		outputResult(IMT)
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

func calculateIMT(userKg, userHeight float64) float64 {
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
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
