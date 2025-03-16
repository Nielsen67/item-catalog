package controller

import (
	"item-catalog/model"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	store *model.ItemStore
}

func NewItemController(store *model.ItemStore) *ItemController {
	return &ItemController{store: store}
}

func (c *ItemController) GetItems(ctx *gin.Context) {
	items := c.store.GetAll()
	responseItems := make(map[string]model.ItemResponse, len(items))
	for id, item := range items {
		responseItems[id] = model.ItemResponse{
			ID:    item.ID,
			Name:  item.Name,
			Price: item.Price,
		}
	}
	ctx.JSON(200, model.ItemsResponse{Items: responseItems})
}

func (c *ItemController) AddItem(ctx *gin.Context) {
	var req model.ItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	item := model.Item{
		ID:    req.ID,
		Name:  req.Name,
		Price: req.Price,
	}

	c.store.Add(item)

	response := model.ItemResponse{
		ID:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}
	ctx.JSON(201, response)
}
