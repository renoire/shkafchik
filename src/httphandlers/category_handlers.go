package httphandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/renoire/shkafchik/pkg/model"
	"github.com/sirupsen/logrus"
)

func (*HTTPHandlers) GetCategories(w http.ResponseWriter, r *http.Request) {
	// type Response struct {
	// 	ListCategories []model.Category
	// }

	//	writeJSONResponse(w, &res)
}

func (*HTTPHandlers) AddCategory(w http.ResponseWriter, r *http.Request) {
	var (
		category model.Category
		res      interface{}
	)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Error("error reading request body") // log error
		writeError(w, err, 500)
		return
	}

	if err := json.Unmarshal(body, &category); err != nil {
		logrus.Error("error unmarshalling request body")
		writeError(w, err, 500)
		return
	}

	//	call to databse

	writeJSONResponse(w, &res)
}
