package dataFile

import (
	"encoding/json"
	"github.com/exampleEP/crudProject/models"
	"io/ioutil"
	"log"
)

func SaveData(user models.User) models.Successful {
	var users []models.User

	users = append(users, user)

	file, err := json.MarshalIndent(users, "", " ")

	if err != nil {
		log.Println("Error at SaveData")
		return models.Successful{
			Success: "",
			Err:     err,
		}
	}

	err = ioutil.WriteFile("Base/base.json", file, 0644)
	if err != nil {
		log.Println("error at WriteFile")
		return models.Successful{
			Success: "",
			Err:     err,
		}
	}

	return models.Successful{
		Success: "successful",
		Err:     nil,
	}

}
