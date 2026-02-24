package storage

import (
	"demo/bin/bins"
	"encoding/json"
	"os"
)

type BinStorage struct {
	filename string
}

func NewBinStorage(filename string) *BinStorage {
	return &BinStorage{
		filename: filename,
	}
}

func (b *BinStorage) SaveJSON(list bins.BinList) error {
	data, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return os.WriteFile(b.filename, data, 0644)
}

func (b *BinStorage) ReadJSON(list *bins.BinList) error {
	data, err := os.ReadFile(b.filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, list)
	if err != nil {
		return err
	}
	return nil
}
