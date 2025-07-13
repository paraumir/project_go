package sort

import (
	"project/models"
	"strings"
)

func QSort(items []models.Item, sortBy string) []models.Item {
	if len(items) < 2 {
		return items
	}
	pivot := items[0]
	var less, greater []models.Item

	for _, item := range items[1:] {
		switch sortBy {
		case "rating":
			if item.Rating < pivot.Rating {
				less = append(less, item)
			} else {
				greater = append(greater, item)
			}
		case "title":
			if strings.ToLower(item.Title) < strings.ToLower(pivot.Title) {
				less = append(less, item)
			} else {
				greater = append(greater, item)
			}
		case "flavor":
			if strings.ToLower(item.Flavor) < strings.ToLower(pivot.Flavor) {
				less = append(less, item)
			} else {
				greater = append(greater, item)
			}
		default:
			if item.ID < pivot.ID {
				less = append(less, item)
			} else {
				greater = append(greater, item)
			}
		}
	}

	return append(
		append(QSort(less, sortBy), pivot),
		QSort(greater, sortBy)...,
	)
}
