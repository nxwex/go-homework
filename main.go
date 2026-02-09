package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Калькулятор индекса массы тела ===")
	userKg, userHeight := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)

	fmt.Println("Нажмите Enter, чтобы выйти...")
	fmt.Scanln() // Ждет нажатия Enter
	fmt.Scanln() // костыль
}
func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.2f\n", imt)
	fmt.Println(result)
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
