package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	events "github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	awsgo "github.com/juparefe/Golang-Ecommerce/awsgo"
	"github.com/juparefe/Golang-Ecommerce/db"
	"github.com/juparefe/Golang-Ecommerce/handlers"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(context context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	// Obtener context y config de AWS
	awsgo.StartAWS()
	// Validar que esten todos los parametros en las variables de entorno
	if !ValidateParameters() {
		panic("Some parameter is missing, it must have SecretName and UrlPrefix")
	}
	var res *events.APIGatewayProxyResponse
	prefix := os.Getenv("UrlPrefix")
	path := strings.Replace(request.RawPath, prefix, "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	headers := request.Headers

	// Leer el secreto de SecretName
	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error reading secret: ", err.Error())
		return res, err

	}

	//Llamar handlers
	status, message := handlers.Handlers(path, method, body, headers, request)

	headersResponse := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headersResponse,
	}

	return res, nil
}

func ValidateParameters() bool {
	_, bringParameter := os.LookupEnv("SecretName")
	if !ValidateParameters() {
		return bringParameter
	}
	_, bringParameter = os.LookupEnv("UrlPrefix")
	if !ValidateParameters() {
		return bringParameter
	}
	return bringParameter
}
