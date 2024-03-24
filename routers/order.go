package routers

import (
	"encoding/json"
	"strconv"

	"github.com/juparefe/Golang-Ecommerce/db"
	"github.com/juparefe/Golang-Ecommerce/models"
)

func InsertOrder(body, User string) (int, string) {
	var o models.Orders
	err := json.Unmarshal([]byte(body), &o)
	if err != nil {
		return 400, "The request data is incorrect: " + err.Error()
	}

	o.Order_UserUUID = User
	OK, message := ValidOrder(o)
	if !OK {
		return 400, message
	}

	result, err2 := db.InsertOrder(o)
	if err2 != nil {
		return 400, "Error when inserting into the database: " + err2.Error()
	}

	return 200, "{ OrderID: " + strconv.Itoa(int(result)) + "}"
}

func ValidOrder(o models.Orders) (bool, string) {
	if o.Order_Total == 0 {
		return false, "You must specify the total of the order"
	}

	count := 0
	for _, od := range o.OrderDetails {
		if od.OD_ProdId == 0 {
			return false, "You must specify the product Id in the detail of the order"
		}
		if od.OD_Quantity == 0 {
			return false, "You must specify the quantity in the detail of the order"
		}
		count++
	}

	if count == 0 {
		return false, "You must specify items in the order"
	}
	return true, ""
}

func SelectOrder(User string) (int, string) {
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
