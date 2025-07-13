package routes

import (
	"net/http"
	"project/models"
	"project/sort"
	"project/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, items []models.Item) {

	r.GET("/items", func(c *gin.Context) {
		sortBy := c.Query("sort")
		search := c.Query("search")

		filtered := []models.Item{}

		for _, item := range items {
			if search == "" || strings.Contains(strings.ToLower(item.Title), strings.ToLower(search)) {
				filtered = append(filtered, item)
			}
		}

		if sortBy != "" {
			filtered = sort.QSort(filtered, sortBy)
		}

		c.JSON(http.StatusOK, filtered)
	})

	r.GET("/sort", func(c *gin.Context) {
		sortBy := c.Query("by")
		sorted := sort.QSort(items, strings.ToLower(sortBy))
		c.JSON(http.StatusOK, sorted)
	})

	r.POST("/add", func(c *gin.Context) {
		var newItem models.Item
		if err := c.ShouldBindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		items = append(items, newItem)
		err := utils.SaveItems(items, "data/sample.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить"})
			return
		}
		c.JSON(http.StatusOK, newItem)
	})
}
