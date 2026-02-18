package bins

import (
	"bufio"
	"fmt"
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

func CreateBin(reader *bufio.Reader, bins BinList) BinList {
	fmt.Println("Создание нового bin")

	id, err := Prompt(reader, "Введите id: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return bins
	}

	name, err := Prompt(reader, "Введите название: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return bins
	}

	privateStr, err := Prompt(reader, "Приватный бин? (true/false): ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return bins
	}

	private, err := strconv.ParseBool(privateStr)
	if err != nil {
		ClearTerminal()
		fmt.Println("Ошибка! Значение должно быть true или false.")
		return bins
	}

	if err := ValidateUserInput(id, name); err != nil {
		fmt.Println(err)
		return bins
	}

	bin := NewBin(id, name, private)
	bins = AddBin(bins, bin)
	ClearTerminal()
	fmt.Printf(">> Bin создан: %+v\n", bin)
	fmt.Printf("Всего bins: %d\n", len(bins))

	return bins
}

func ShowAllBins(bins BinList) {
	if len(bins) == 0 {
		fmt.Println(">> Список пуст")
		return
	}

	fmt.Println("=== Список бинов ===")
	for i, b := range bins {
		fmt.Printf("%d) ID=%s | Name=%s | Private=%t | CreatedAt=%s\n",
			i+1, b.ID, b.Name, b.Private, b.CreatedAt.Format(time.RFC3339))
	}
}

func Prompt(r *bufio.Reader, label string) (string, error) {
	fmt.Print(label)
	text, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(text), nil
}

func ValidateUserInput(id, name string) error {
	if id == "" {
		return fmt.Errorf("ошибка! ID не может быть пустым")
	}
	if name == "" {
		return fmt.Errorf("ошибка! название не может быть пустым")
	}
	return nil
}

func NewBin(id, name string, private bool) Bin {
	return Bin{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
		Private:   private,
	}
}

func AddBin(list BinList, b Bin) BinList {
	return append(list, b)
}

func ClearTerminal() {
	fmt.Print("\x1b[2J\x1b[H")
}
