package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList []Bin

func main() {
	userMenu()
}

func userMenu() {
	reader := bufio.NewReader(os.Stdin)
	var bins BinList
	clearTerminal()

	for {
		fmt.Println("=== BINS MANAGER ===")
		fmt.Println("1. Создать новый Bin")
		fmt.Println("2. Показать существующие")
		fmt.Println("0. Выход")

		choice, err := prompt(reader, "Ввод >> ")
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}

		switch choice {
		case "1":
			clearTerminal()
			bins = createBin(reader, bins)
		case "2":
			clearTerminal()
			showAllBins(bins)
		case "0":
			clearTerminal()
			return
		default:
			clearTerminal()
			fmt.Println(">> Unknown option. Try 1, 2 or 0.")
		}
	}
}

func createBin(reader *bufio.Reader, bins BinList) BinList {
	fmt.Println("Создание нового Bin")

	id, err := prompt(reader, "Введите ID: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return bins
	}

	name, err := prompt(reader, "Введите name: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return bins
	}

	privateStr, err := prompt(reader, "Private (true/false): ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return bins
	}

	private, err := strconv.ParseBool(privateStr)
	if err != nil {
		fmt.Println("Ошибка! Введите true или false.")
		return bins
	}

	if err := validateUserInput(id, name); err != nil {
		fmt.Println(err)
		return bins
	}

	bin := newBin(id, name, private)
	bins = addBin(bins, bin)
	clearTerminal()
	fmt.Printf(">> Bin создан: %+v\n", bin)
	fmt.Printf("Всего bins: %d\n", len(bins))

	return bins
}

func showAllBins(bins BinList) {
	if len(bins) == 0 {
		fmt.Println(">> Список пуст")
		return
	}

	fmt.Println("=== ALL BINS ===")
	for i, b := range bins {
		fmt.Printf("%d) ID=%s | Name=%s | Private=%t | CreatedAt=%s\n",
			i+1, b.ID, b.Name, b.Private, b.CreatedAt.Format(time.RFC3339))
	}
}

func prompt(r *bufio.Reader, label string) (string, error) {
	fmt.Print(label)
	text, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(text), nil
}

func validateUserInput(id, name string) error {
	if id == "" {
		return fmt.Errorf("ошибка! ID не может быть пустым")
	}
	if name == "" {
		return fmt.Errorf("ошибка! название не может быть пустым")
	}
	return nil
}

func newBin(id, name string, private bool) Bin {
	return Bin{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
		Private:   private,
	}
}

func addBin(list BinList, b Bin) BinList {
	return append(list, b)
}

func clearTerminal() {
	fmt.Print("\x1b[2J\x1b[H")
}
