package utils

import (
	"encoding/json"
	"os"
	"project/models"
)

func SaveItems(items []models.Item, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(items)
}
