package router

import (
	"github.com/go-chi/chi"

	"github.com/renoire/shkafchik/src/httphandlers"
)

func GetRouter(h httphandlers.HTTPHandlers) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/public/api/v1/categories", func(r chi.Router) {
		r.Get("/", h.GetCategories)
		r.Post("/", h.AddCategory)
	})

	r.Route("/public/api/v1/items", func(r chi.Router) {
		r.Get("/", h.GetAllItems)
		r.Post("/", h.AddItem)
		r.Post("/", h.GetItem)
	})

	r.Get("/public/api/v1/items", h.GetAllItems)

	return r
}
