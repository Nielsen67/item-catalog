package controller

import (
	"item-catalog/config"
	"item-catalog/model"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	cfg := config.Load()
	store := model.NewItemStore()
	itemController := NewItemController(store)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		log.Println("Serving index.html (Vercel)")
		c.File("./public/index.html")
	})

	api := router.Group("/api")
	{
		api.GET("/items", itemController.GetItems)
		api.POST("/items", itemController.AddItem)
	}

	router.NoRoute(func(c *gin.Context) {
		file := c.Request.URL.Path
		if file == "/" {
			log.Println("Serving index.html (NoRoute fallback, Vercel)")
			c.File("./public/index.html")
			return
		}
		filePath := "./public" + file
		_, err := os.Stat(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				log.Printf("File not found: %s, falling back to index.html (Vercel)", filePath)
				c.File("./public/index.html")
				return
			}
			log.Printf("Error accessing file %s: %v (Vercel)", filePath, err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}
		log.Printf("Serving file: %s (Vercel)", filePath)
		c.File(filePath)
	})

	log.Printf("Vercel controller running in %s mode", cfg.Env)
	router.ServeHTTP(w, r)
}
