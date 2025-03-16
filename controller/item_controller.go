package controller

import (
	"item-catalog/model"
	"log"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	store *model.ItemStore
}

func NewItemController(store *model.ItemStore) *ItemController {
	return &ItemController{store: store}
}

func (ic *ItemController) GetItems(c *gin.Context) {
	items := ic.store.GetAll()
	responseItems := make(map[string]model.ItemResponse)
	for id, item := range items {
		responseItems[id] = model.ItemResponse{
			ID:    item.ID,
			Name:  item.Name,
			Price: item.Price,
		}
	}
	c.JSON(200, model.ItemsResponse{Items: responseItems})
}

func (ic *ItemController) AddItem(c *gin.Context) {
	var itemReq model.ItemRequest
	if err := c.ShouldBindJSON(&itemReq); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	item := model.Item{
		ID:    itemReq.ID,
		Name:  itemReq.Name,
		Price: itemReq.Price,
	}

	ic.store.Add(item)
	c.JSON(201, gin.H{"message": "Item added successfully"})
}
