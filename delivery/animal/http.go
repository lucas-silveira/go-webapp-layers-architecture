package animal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"webapp/datastore/animal"
	"webapp/entities"
)

// Controller ...
type Controller struct {
	datastore animal.IAnimalRepository
}

// New ...
func New(animal animal.IAnimalRepository) Controller {
	return Controller{datastore: animal}
}

// Handler ...
func (controller Controller) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controller.get(w, r)
	case http.MethodPost:
		controller.create(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (controller Controller) get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idParsed, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid parameter id"))

		return
	}

	res, err := controller.datastore.Get(idParsed)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Could not retrieve animal"))

		return
	}

	body, _ := json.Marshal(res)
	_, _ = w.Write(body)
}

func (controller Controller) create(w http.ResponseWriter, r *http.Request) {
	var animal entities.Animal

	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &animal)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid body"))

		return
	}

	res, err := controller.datastore.Create(animal)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Could not create animal"))

		return
	}

	body, _ = json.Marshal(res)
	_, _ = w.Write(body)
}
