package model

type Item struct {
	ID    string `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}

type ItemRequest struct {
	ID    string `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required,gt=0"`
}

type ItemResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ItemsResponse struct {
	Items map[string]ItemResponse `json:"items"`
}

type ItemStore struct {
	items map[string]Item
}

func NewItemStore() *ItemStore {
	return &ItemStore{
		items: make(map[string]Item),
	}
}

func (s *ItemStore) GetAll() map[string]Item {
	return s.items
}

func (s *ItemStore) Add(item Item) {
	s.items[item.ID] = item
}
