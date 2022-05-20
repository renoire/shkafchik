package httphandlers

import (
	"github.com/renoire/shkafchik/src/categories"
	"github.com/renoire/shkafchik/src/items"
)

type HTTPHandlers struct {
	categoriesService categories.CategoriesService
	itemsService      items.ItemsService
}

func New(categoriesService categories.CategoriesService, itemsService items.ItemsService) *HTTPHandlers {
	return &HTTPHandlers{
		categoriesService: categoriesService,
		itemsService:      itemsService,
	}
}
