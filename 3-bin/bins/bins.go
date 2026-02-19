package bins

import (
	"bufio"
	"fmt"
	"io"
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

func (l *BinList) AddNew(id, name string, private bool) Bin {
	b := NewBin(id, name, private)
	*l = append(*l, b)
	return b
}

func CreateBin(reader *bufio.Reader, bins *BinList) error {
	fmt.Println("Создание нового bin")

	id, err := Prompt(reader, "Введите id: ")
	if err != nil {
		return err
	}

	name, err := Prompt(reader, "Введите название: ")
	if err != nil {
		return err
	}

	var private bool
	for {
		privateStr, err := Prompt(reader, "Приватный бин? (true/false): ")
		if err != nil {
			return err
		}

		v, err := strconv.ParseBool(privateStr)
		if err != nil {
			fmt.Println("Ошибка! Значение должно быть true или false.")
			continue
		}
		private = v
		break
	}

	if err := ValidateUserInput(id, name); err != nil {
		return err
	}

	bin := bins.AddNew(id, name, private)
	fmt.Printf("* Bin создан: %+v\n", bin)
	fmt.Printf("Всего бинов: %d\n\n", len(*bins))

	return nil
}

func ShowAllBins(bins BinList) {
	if len(bins) == 0 {
		fmt.Print("* Список пуст\n\n")
		return
	}

	fmt.Println("=== Список бинов ===")
	for i, b := range bins {
		fmt.Printf("%d) ID = %s | Name = %s | Private = %t | CreatedAt = %s\n\n",
			i+1, b.ID, b.Name, b.Private, b.CreatedAt.Format(time.RFC3339))
	}
}

func Prompt(r *bufio.Reader, label string) (string, error) {
	for {
		fmt.Print(label)
		text, err := r.ReadString('\n')
		if err != nil && err != io.EOF {

			return "", err
		}
		text = strings.TrimSpace(text)
		if text == "" {
			if err == io.EOF {
				return "", io.EOF
			}
			fmt.Println("Ошибка! Ввод не должен быть пустым")
			continue
		}
		return text, nil
	}
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
