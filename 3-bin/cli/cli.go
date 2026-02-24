package cli

import (
	"bufio"
	"demo/bin/bins"
	"demo/bin/storage"
	"fmt"
	"os"
)

func UserMenu() {
	reader := bufio.NewReader(os.Stdin)
	s := storage.NewBinStorage("list.json")
	var actBins bins.BinList
	if err := s.ReadJSON(&actBins); err != nil {
		fmt.Println("info: Локальная база не найдена, создаю новую")
	}
	ClearTerminal()

	fmt.Println("=== bins менеджер ===")
	for {
		fmt.Println("1. Создать новый bin")
		fmt.Println("2. Показать существующие")
		fmt.Println("0. Выход")

		choice, err := bins.Prompt(reader, "\nВвод >> ")
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}

		switch choice {
		case "1":
			ClearTerminal()
			if err := bins.CreateBin(reader, &actBins, s); err != nil {
				fmt.Println("Ошибка:", err)
			}
		case "2":
			ClearTerminal()
			bins.ShowAllBins(actBins)
		case "0":
			ClearTerminal()
			return
		default:
			ClearTerminal()
			fmt.Println(">> Ошибка! Введите существующий пункт меню")
		}
	}
}

func ClearTerminal() {
	fmt.Print("\x1b[2J\x1b[H")
}
