package bins

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

type Storage interface {
	SaveJSON(BinList) error
	ReadJSON(*BinList) error
}

type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}
type BinRecord struct {
	Bin
	UpdatedAt time.Time `json:"updated_at"`
}

type BinList []BinRecord

func (l *BinList) AddNew(id, name string, private bool) BinRecord {
	b := NewBin(id, name, private)
	*l = append(*l, b)
	return b
}

func (l *BinList) isRepeat(id string) bool {
	for _, b := range *l {
		if b.ID == id {
			return true
		}
	}
	return false
}

func CreateBin(reader *bufio.Reader, bins *BinList, s Storage) error {
	fmt.Println("=== Создание нового bin ===")

	id, err := promptUniqueID(reader, *bins)
	if err != nil {
		return err
	}

	name, err := Prompt(reader, "Введите название: ")
	if err != nil {
		return err
	}

	isPrivate, err := promptPrivate(reader)
	if err != nil {
		return err
	}

	newBin := bins.AddNew(id, name, isPrivate)

	fmt.Println("===")
	fmt.Printf("* Успешно! Bin создан: %+v\n", newBin)
	fmt.Printf("Всего бинов: %d\n\n", len(*bins))
	return s.SaveJSON(*bins)
}

func promptUniqueID(reader *bufio.Reader, bins BinList) (string, error) {
	for {
		id, err := Prompt(reader, "Введите id: ")
		if err != nil {
			return "", err
		}

		if bins.isRepeat(id) {
			fmt.Printf("Ошибка! ID '%s' уже занят.\n", id)
			continue
		}

		return id, nil
	}
}

func promptPrivate(reader *bufio.Reader) (bool, error) {
	for {
		input, err := Prompt(reader, "Сделать приватным? (y/n): ")
		if err != nil {
			return false, err
		}

		switch strings.ToLower(input) {
		case "y", "yes":
			return true, nil
		case "n", "no":
			return false, nil
		default:
			fmt.Println("Пожалуйста, введите 'y' или 'n'")
		}
	}
}

func ShowAllBins(bins BinList) {
	if len(bins) == 0 {
		fmt.Print("* Список пуст\n\n")
		return
	}
	fmt.Println("* Список бинов")
	for i, b := range bins {
		fmt.Printf("%d) ID = %s | Name = %s | Private = %t | CreatedAt = %s\n",
			i+1, b.ID, b.Name, b.Private, b.CreatedAt.Format(time.RFC3339))
	}
	fmt.Println("---")
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

func NewBin(id, name string, private bool) BinRecord {
	now := time.Now()
	return BinRecord{
		Bin: Bin{
			ID:        id,
			Name:      name,
			CreatedAt: now,
			Private:   private,
		},
		UpdatedAt: now,
	}
}
