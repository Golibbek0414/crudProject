package controller

import (
	"encoding/json"
	"github.com/exampleEP/crudProject/models"
	"github.com/exampleEP/crudProject/storage/dataFile"
	"github.com/google/uuid"
	"net/http"
)

func (api *api) UserSignUp(w http.ResponseWriter, r *http.Request) {
	var (
		body models.User
		res  models.Successful
	)

	id := uuid.New().String()

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.Write([]byte("error at UserSignUp"))
		return
	}

	res = dataFile.SaveData(models.User{
		Id:    id,
		Name:  body.Name,
		Books: body.Books,
	})

	writeJson(w, res)
}

func Sign(w http.ResponseWriter, r *http.Request) {
	var (
		body models.User
		res  models.Successful
	)

	id := uuid.New().String()

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.Write([]byte("error at UserSignUp"))
		return
	}

	res = dataFile.SaveData(models.User{
		Id:    id,
		Name:  body.Name,
		Books: body.Books,
	})

	writeJson(w, res)
}
