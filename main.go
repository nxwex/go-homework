package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userKg float64
	fmt.Println("=== Калькулятор индекса массы тела ===")
	fmt.Print("Введите свой рост в метрах (формат x.xx): ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в кг (формат x.xx): ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Printf("Рост: %2.f, вес: %2.f, ваш индекс массы тела: %2.f\n", userHeight, userKg, IMT)

	fmt.Println("Нажмите Enter, чтобы выйти...")
	fmt.Scanln() // Ждет нажатия Enter
}
