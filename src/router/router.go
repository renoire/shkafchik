package router

import (
	"github.com/go-chi/chi"

	"github.com/renoire/shkafchik/src/httphandlers"
)

func getRouter() {
	r := chi.NewRouter()

	r.Route("/public/api/v1/categories", func(r chi.Router) {
		r.Get("/", httphandlers.GetCategories)
		r.Post("/", httphandlers.AddCategory)
	})

	r.Route("/public/api/v1/items", func(r chi.Router) {
		r.Get("/", httphandlers.GetAllItems)
		r.Post("/", httphandlers.AddItem)
		r.Post("/", httphandlers.GetItem)
	})

	r.Get("/public/api/v1/items", httphandlers.GetAllItems)
}
