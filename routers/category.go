package routers

import (
	"encoding/json"
	"strconv"

	"github.com/juparefe/Golang-Ecommerce/db"
	"github.com/juparefe/Golang-Ecommerce/models"
)

func InsertCategory(body, User string) (int, string) {
	var t models.Category
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "The request data is incorrect: " + err.Error()
	}

	if len(t.CategName) == 0 {
		return 400, "You must specify the name of the category"
	}
	if len(t.CategPath) == 0 {
		return 400, "You must specify the path of the category"
	}

	isAdmin, msg := db.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := db.InsertCategory(t)
	if err2 != nil {
		return 400, "Error when inserting into the database: " + t.CategName + " > " + err2.Error()
	}

	return 200, "{ CategID: " + strconv.Itoa(int(result)) + "}"
}
