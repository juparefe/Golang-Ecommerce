package routers

import (
	"encoding/json"
	"strconv"

	"github.com/juparefe/Golang-Ecommerce/db"
	"github.com/juparefe/Golang-Ecommerce/models"
)

func InsertProduct(body, User string) (int, string) {
	var t models.Product
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "The request data is incorrect: " + err.Error()
	}

	if len(t.ProdTitle) == 0 {
		return 400, "You must specify the name (Title) of the product"
	}

	isAdmin, msg := db.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := db.InsertProduct(t)
	if err2 != nil {
		return 400, "Error when inserting into the database: " + t.ProdTitle + " > " + err2.Error()
	}

	return 200, "{ ProductID: " + strconv.Itoa(int(result)) + "}"
}

func UpdateProduct(body, User string, id int) (int, string) {
	var t models.Product
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "The request data is incorrect: " + err.Error()
	}

	isAdmin, msg := db.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.ProdId = id
	err2 := db.UpdateProduct(t)
	if err2 != nil {
		return 400, "Error when updating into the database: " + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, "Update Ok"
}

func DeleteProduct(User string, id int) (int, string) {
	if id == 0 {
		return 400, "The request data (ID) is incorrect"
	}

	isAdmin, msg := db.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err := db.DeleteProduct(id)
	if err != nil {
		return 400, "Error when deleting into the database: " + strconv.Itoa(id) + " > " + err.Error()
	}

	return 200, "Delete Ok"
}
