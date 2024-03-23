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

	fmt.Println("Path to validate: ", path[0:5])
	switch path[0:4] {
	case "addr":
		return ProcessAddresses(body, path, method, user, idn, request)
	case "cate":
		return ProcessCategories(body, path, method, user, idn, request)
	case "orde":
		return ProcessOrders(body, path, method, user, idn, request)
	case "prod":
		return ProcessProducts(body, path, method, user, idn, request)
	case "stoc":
		return ProcessStock(body, path, method, user, idn, request)
	case "user":
		return ProcessUsers(body, path, method, user, id, request)
	}
	return 400, "Method invalid"
}

func ValidateAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	if (path == "product" || path == "category") && method == "GET" {
		return true, 200, ""
	}

	token := headers["authorization"]
	fmt.Println("authorization: ", headers["authorization"])
	if len(token) == 0 {
		return false, 401, "Token required"
	}

	everythingOk, err, msg := auth.ValidateToken(token)
	if !everythingOk {
		if err != nil {
			fmt.Println("Error in the token: ", err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error in the token: ", msg)
			return false, 401, msg
		}
	}
	fmt.Println("Everything ok with authorization")
	return true, 200, msg
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
	}
	if path == "users" && method == "GET" {
		return routers.SelectUsers(body, user, request)
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

func ProcessStock(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return routers.UpdateStock(body, user, id)
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
		return routers.SelectAdress(request)
	}
	return 400, "Method invalid"
}

func ProcessOrders(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method invalid"
}
