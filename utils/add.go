package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"project/models"
)

func AddNewItem(items *[]models.Item, titleMap map[string]models.Item, filename string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Добавление новой записи:")

	fmt.Print("Название: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Вкус: ")
	flavor, _ := reader.ReadString('\n')
	flavor = strings.TrimSpace(flavor)

	fmt.Print("Рейтинг (0.0 - 10.0): ")
	ratingStr, _ := reader.ReadString('\n')
	ratingStr = strings.TrimSpace(ratingStr)
	rating, err := strconv.ParseFloat(ratingStr, 64)
	if err != nil || rating < 0 || rating > 10 {
		fmt.Println("Некорректный рейтинг.")
		return
	}

	newID := getNextID(*items)

	newItem := models.Item{
		ID:     newID,
		Title:  title,
		Flavor: flavor,
		Rating: rating,
	}

	*items = append(*items, newItem)
	titleMap[strings.ToLower(title)] = newItem

	err = SaveItems(*items, filename)
	if err != nil {
		fmt.Println("Ошибка сохранения в файл:", err)
		return
	}

	fmt.Println("Запись успешно добавлена и сохранена.")
}

func getNextID(items []models.Item) int {
	maxID := 0
	for _, item := range items {
		if item.ID > maxID {
			maxID = item.ID
		}
	}
	return maxID + 1
}
