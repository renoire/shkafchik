package httphandlers

import (
	"net/http"

	"github.com/renoire/shkafchik/src/model"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		ListCategories []model.Category
	}

	//	writeJSONResponse(w, &res)
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	type Response struct {
	}
	type Request struct {
		cat model.Category
	}

	//	writeJSONResponse(w, &res)
}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	type Response struct {
	}
	type Request struct {
	}

	//	writeJSONResponse(w, &res)
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	type Response struct {
	}
	type Request struct {
	}

	//	writeJSONResponse(w, &res)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	type Response struct {
	}
	type Request struct {
	}

	//	writeJSONResponse(w, &res)
}
