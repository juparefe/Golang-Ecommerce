package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/juparefe/Golang-Ecommerce/db"
	"github.com/juparefe/Golang-Ecommerce/models"
)

func UpdateUser(body, User string) (int, string) {
	var t models.User
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "The request data is incorrect: " + err.Error()
	}
	if (len(t.UserFirstName) == 0) && (len(t.UserLastName) == 0) {
		return 400, "You must specify the name (FirstName) or the lastName (LastName) of the user to update"
	}

	_, found := db.UserExists(User)
	if found {
		return 400, "There is no user with that UUID '" + User + "'"
	}

	err = db.UpdateUser(t)
	if err != nil {
		return 400, "Error when updating into the database: " + User + " > " + err.Error()
	}

	return 200, "Update Ok"
}

func DeleteUser(User, id string) (int, string) {
	if id == "0" {
		return 400, "The request data (ID) is incorrect"
	}

	isAdmin, msg := db.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	return 200, "Delete Ok"
}

func SelectUsers(request events.APIGatewayV2HTTPRequest) (int, string) {
	var err error
	var CategId int
	var Slug string

	requestCategId := request.QueryStringParameters["categId"]
	requestSlug := request.QueryStringParameters["slug"]
	if len(requestCategId) > 0 {
		CategId, err = strconv.Atoi(requestCategId)
		if err != nil {
			return 500, "Error when converting the value to an integer: " + requestCategId
		}
	} else {
		if len(requestSlug) > 0 {
			Slug = requestSlug
		}
	}

	list, err2 := db.SelectCategories(CategId, Slug)
	if err2 != nil {
		return 400, "Error trying to get category/ies: " + err2.Error()
	}

	Categ, err3 := json.Marshal(list)
	if err3 != nil {
		return 500, "Error trying to convert to JSON categories list" + err3.Error()
	}
	return 200, string(Categ)
}
