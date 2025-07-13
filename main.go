package main

import (
	"fmt"
	"project/models"
	"project/routes"
	"project/utils"

	"github.com/gin-gonic/gin"
)

var (
	dataFile = "data/sample.json"
	items    []models.Item
)

func main() {
	var err error
	items, err = utils.LoadItems(dataFile)
	if err != nil {
		fmt.Println("Ошибка загрузки данных:", err)
		return
	}

	r := gin.Default()

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		fmt.Println(">>> Обработчик / вызван")
		c.File("./static/index.html")
	})

	routes.SetupRoutes(r, items)

	fmt.Println("Сервер запущен на http://localhost:8080")
	r.Run(":8080")
}
