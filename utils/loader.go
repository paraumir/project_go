package utils

import (
	"encoding/json"
	"os"
	"project/models"
)

func LoadItems(path string) ([]models.Item, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var items []models.Item
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&items)
	if err != nil {
		return nil, err
	}

	return items, nil
}
