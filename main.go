package handler

import (
	"log"
	"net/http"
	"os"

	"item-catalog/config"
	"item-catalog/controller"
	"item-catalog/model"

	"github.com/gin-gonic/gin"
)

var store *model.ItemStore

func init() {

	store = model.NewItemStore()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	cfg := config.Load()

	itemController := controller.NewItemController(store)

	ginMode := gin.ReleaseMode
	if cfg.Env == "development" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		log.Println("Serving index.html")
		c.File("./public/index.html")
	})

	api := router.Group("/api")
	{
		api.GET("/items", itemController.GetItems)
		api.POST("/items", itemController.AddItem)
		api.DELETE("/items/clear", itemController.ClearItems)
	}

	router.NoRoute(func(c *gin.Context) {
		file := c.Request.URL.Path
		if file == "/" {
			log.Println("Serving index.html (NoRoute fallback)")
			c.File("./public/index.html")
			return
		}
		filePath := "./public" + file
		_, err := os.Stat(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				log.Printf("File not found: %s, falling back to index.html", filePath)
				c.File("./public/index.html")
				return
			}
			log.Printf("Error accessing file %s: %v", filePath, err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}
		log.Printf("Serving file: %s", filePath)
		c.File(filePath)
	})

	router.ServeHTTP(w, r)
}

// func main() {
// 	cfg := config.Load()
// 	log.Printf("Running locally in %s mode on port %s", cfg.Env, cfg.Port)
// 	http.ListenAndServe(":"+cfg.Port, http.HandlerFunc(Handler))
// }
