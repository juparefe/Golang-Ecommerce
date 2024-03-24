package routers

import (
	"encoding/json"
	"strconv"

	"github.com/juparefe/Golang-Ecommerce/db"
	"github.com/juparefe/Golang-Ecommerce/models"
)

func InsertAddress(body, User string) (int, string) {
	var t models.Address
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "The request data is incorrect: " + err.Error()
	}

	if t.AddAddress == "" {
		return 400, "You must specify the Address"
	}
	if t.AddName == "" {
		return 400, "You must specify the Name"
	}
	if t.AddTitle == "" {
		return 400, "You must specify the Title"
	}
	if t.AddCity == "" {
		return 400, "You must specify the City"
	}
	if t.AddPhone == "" {
		return 400, "You must specify the Phone"
	}
	if t.AddPostalCode == "" {
		return 400, "You must specify the Postal Code"
	}

	err2 := db.InsertAddress(t, User)
	if err2 != nil {
		return 400, "Error when inserting into the database: " + t.AddAddress + " for the User UUID: " + User + " > " + err2.Error()
	}

	return 200, "Insert address Ok"
}

func UpdateAddress(body, User string, id int) (int, string) {
	var t models.Address
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "The request data is incorrect: " + err.Error()
	}

	t.AddId = id
	var found bool
	err, found = db.AddressExists(User, t.AddId)
	if !found {
		if err != nil {
			return 400, "Error searching address for user: '" + User + "': " + err.Error()
		} else {
			return 204, "There is no user with that UUID asociated to the address ID: '" + strconv.Itoa(id) + "'"
		}
	}
	err = db.UpdateAddress(t)
	if err != nil {
		return 400, "Error when updating into the database: " + strconv.Itoa(id) + " > " + err.Error()
	}

	return 200, "Update Ok"
}

func DeleteAdress(User string, id int) (int, string) {
	if id == 0 {
		return 400, "The request data (ID) is incorrect"
	}

	var found bool
	err, found := db.AddressExists(User, id)
	if !found {
		if err != nil {
			return 400, "Error searching address for user: '" + User + "': " + err.Error()
		} else {
			return 204, "There is no user with that UUID asociated to the address ID: '" + strconv.Itoa(id) + "'"
		}
	}

	err = db.DeleteAddress(id)
	if err != nil {
		return 400, "Error when deleting into the database: " + strconv.Itoa(id) + " > " + err.Error()
	}

	return 200, "Delete Ok"
}

func SelectAdress(User string) (int, string) {
	result, err := db.SelectAddress(User)
	if err != nil {
		return 400, "Error trying to get address for User: '" + User + "', Error: " + err.Error()
	}

	respJson, err := json.Marshal(result)
	if err != nil {
		return 500, "Error trying to convert to JSON adress list" + err.Error()
	}
	return 200, string(respJson)
}
