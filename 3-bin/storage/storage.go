package storage

import (
	"demo/bin/bins"
	"encoding/json"
	"os"
)

type BinStorage bins.BinList

func (b *BinStorage) SaveJSON(path string) error {
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func (b *BinStorage) ReadJSON(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, b)
}
