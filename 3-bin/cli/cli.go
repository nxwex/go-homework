package cli

import (
	"bufio"
	"demo/bin/bins"
	"fmt"
	"os"
)

func UserMenu() {
	reader := bufio.NewReader(os.Stdin)
	var actBins bins.BinList
	bins.ClearTerminal()

	for {
		fmt.Println("=== bins менеджер ===")
		fmt.Println("1. Создать новый bin")
		fmt.Println("2. Показать существующие")
		fmt.Println("0. Выход")

		choice, err := bins.Prompt(reader, "Ввод >> ")
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}

		switch choice {
		case "1":
			bins.ClearTerminal()
			actBins = bins.CreateBin(reader, actBins)
		case "2":
			bins.ClearTerminal()
			bins.ShowAllBins(actBins)
		case "0":
			bins.ClearTerminal()
			return
		default:
			bins.ClearTerminal()
			fmt.Println(">> Ошибка! Введите существующий пункт меню")
		}
	}
}
