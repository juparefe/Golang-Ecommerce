package handlers

import (
	"fmt"
	//"strconv"
	"github.com/aws/aws-lambda-go/events"
)

func Handlers(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Processing " + path + " > " + method)
	//id := request.PathParameters["id"]
	//idn, _ := strconv.Atoi(id)
	// test
	return 400, "Method invalid"
}
