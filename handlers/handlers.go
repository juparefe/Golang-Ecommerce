package handlers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/juparefe/Golang-Ecommerce/auth"
	"github.com/juparefe/Golang-Ecommerce/routers"
)

func Handlers(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Processing " + path + " > " + method)
	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)
	isOk, statusCode, user := ValidateAuthorization(path, method, headers)
	if !isOk {
		return statusCode, user
	}

	switch path[0:4] {
	case "addr":
		return ProcessAddresses(body, path, method, user, idn, request)
	case "cate":
		return ProcessCategories(body, path, method, user, idn, request)
	case "curr":
		return ProcessCurrencies(body, path, method, user, idn, request)
	case "disc":
		return ProcessDiscount(body, path, method, user, idn, request)
	case "orde":
		return ProcessOrders(body, path, method, user, idn, request)
	case "prod":
		return ProcessProducts(body, path, method, user, idn, request)
	case "stoc":
		return ProcessStock(body, path, method, user, idn, request)
	case "topc":
		return ProcessTopCategories(body, path, method, user, idn, request)
	case "user":
		return ProcessUsers(body, path, method, user, id, request)
	}
	return 400, "Method invalid"
}

func ValidateAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	if (path == "category" || path == "product" || path == "topcategories") && method == "GET" {
		return true, 200, ""
	}

	token := headers["authorization"]

	if len(token) == 0 {
		return false, 401, "Token required"
	}

	everythingOk, err, msg := auth.ValidateToken(token)
	if !everythingOk {
		if err != nil {
			fmt.Println("Authorization: ", headers["authorization"])
			fmt.Println("Error in the token: ", err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Authorization: ", headers["authorization"])
			fmt.Println("Error in the token: ", msg)
			return false, 401, msg
		}
	}
	fmt.Println("Everything ok with authorization")
	return true, 200, msg
}

func ProcessAddresses(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Start ProcessAddresses with method: ", method)
	switch method {
	case "POST":
		return routers.InsertAddress(body, user)
	case "PUT":
		return routers.UpdateAddress(body, user, id)
	case "DELETE":
		return routers.DeleteAdress(user, id)
	case "GET":
		return routers.SelectAdress(user)
	}
	return 400, "Method invalid"
}

func ProcessCategories(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Start ProcessCategories with method: ", method)
	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	case "PUT":
		return routers.UpdateCategory(body, user, id)
	case "DELETE":
		return routers.DeleteCategory(user, id)
	case "GET":
		return routers.SelectCategories(request)
	}
	return 400, "Method invalid"
}

func ProcessCurrencies(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Start ProcessCurrencies with method: ", method)
	switch method {
	case "PUT":
		return routers.UpdateCurrencies(body, user)
	case "GET":
		return routers.SelectCurrencies(request)
	}
	return 400, "Method invalid"
}

func ProcessDiscount(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return routers.UpdateDiscount(body, user, id)
}

func ProcessOrders(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Start ProcessOrders with method: ", method)
	switch method {
	case "POST":
		return routers.InsertOrder(body, user)
	case "GET":
		return routers.SelectOrder(user, request)
	}
	return 400, "Method invalid"
}

func ProcessProducts(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Start ProcessProducts with method: ", method)
	switch method {
	case "POST":
		return routers.InsertProduct(body, user)
	case "PUT":
		return routers.UpdateProduct(body, user, id)
	case "DELETE":
		return routers.DeleteProduct(user, id)
	case "GET":
		return routers.SelectProducts(request)
	}
	return 400, "Method invalid"
}

func ProcessStock(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return routers.UpdateStock(body, user, id)
}

func ProcessTopCategories(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Start ProcessTopCategories with method: ", method)
	switch method {
	case "GET":
		return routers.SelectTopCategories(request)
	}
	return 400, "Method invalid"
}

func ProcessUsers(body, path, method, user, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Start ProcessUsers with method: ", method)
	if path == "user/me" {
		switch method {
		case "PUT":
			return routers.UpdateUser(body, user)
		case "GET":
			return routers.SelectUser(body, user)
		}
	} else if path == "users" && method == "GET" {
		return routers.SelectUsers(body, user, request)
	} else {
		return routers.UpdateUserRole(body, user)
	}
	return 400, "Method invalid"
}
