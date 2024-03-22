package routers

import (
	"encoding/json"
	"fmt"
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

	err = db.UpdateUser(t, User)
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

func SelectUser(body, User string) (int, string) {
	_, found := db.UserExists(User)
	if found {
		return 400, "There is no user with that UUID '" + User + "'"
	}

	row, err := db.SelectUser(User)
	fmt.Println("Row with user: ", row)
	if err != nil {
		return 400, "Error trying to get user: " + err.Error()
	}

	respJson, err2 := json.Marshal(row)
	if err2 != nil {
		return 500, "Error trying to convert to JSON the user" + err2.Error()
	}
	return 200, string(respJson)
}

func SelectUsers(body, User string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var Page int
	if len(request.QueryStringParameters["page"]) == 0 {
		Page = 1
	} else {
		Page, _ = strconv.Atoi(request.QueryStringParameters["page"])
	}

	isAdmin, msg := db.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	row, err := db.SelectUsers(Page)
	if err != nil {
		return 400, "Error trying to get users: " + err.Error()
	}

	respJson, err2 := json.Marshal(row)
	if err2 != nil {
		return 500, "Error trying to convert to JSON the users" + err2.Error()
	}
	return 200, string(respJson)
}
