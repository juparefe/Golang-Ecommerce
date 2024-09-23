package routers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
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

func SelectProducts(request events.APIGatewayV2HTTPRequest) (int, string) {
	var t models.Product
	var page, pageSize int
	var orderType, orderField string

	param := request.QueryStringParameters
	page, _ = strconv.Atoi(param[`page`])
	pageSize, _ = strconv.Atoi(param[`pageSize`])
	orderType = param["orderType"]   // 'D' = Desc, 'A' or nil = Asc
	orderField = param["orderField"] // 'C' = CategId, 'D' = Description, 'F' CreatedAt, 'I' = id, 'P' = Price, 'S' = Stock, 'T' = Title
	if !strings.Contains("CDFIPST", orderField) {
		orderField = ""
	}

	var choice string
	if len(param["categId"]) > 0 {
		choice = "C"
		t.ProdCategId, _ = strconv.Atoi(param["categId"])
	}
	if len(param["slugCateg"]) > 0 {
		choice = "K"
		t.ProdCategPath = param["slugCateg"]
	}
	if len(param["prodId"]) > 0 {
		choice = "P"
		t.ProdId, _ = strconv.Atoi(param["prodId"])
	}
	if len(param["search"]) > 0 {
		choice = "S"
		t.ProdSearch = param["search"]
	}
	if len(param["slug"]) > 0 {
		choice = "U"
		t.ProdPath = param["slug"]
	}

	result, err := db.SelectProducts(t, choice, orderType, orderField, page, pageSize)
	if err != nil {
		fmt.Println("Search parameters for products: ", param)
		return 400, "Error trying to get product/s of Type: '" + choice + "', Error: " + err.Error()
	}

	Product, err2 := json.Marshal(result)
	if err2 != nil {
		return 500, "Error trying to convert to JSON products list" + err2.Error()
	}
	return 200, string(Product)
}

func UpdateStock(body, User string, id int) (int, string) {
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
	err2 := db.UpdateStock(t)
	if err2 != nil {
		return 400, "Error when updating stock into the database: " + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, "Update Ok"
}
