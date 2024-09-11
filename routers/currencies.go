package routers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/juparefe/Golang-Ecommerce/db"
)

func SelectCurrencies(request events.APIGatewayV2HTTPRequest) (int, string) {
	var err error
	var BaseCurrency string
	var TargetCurrency string

	requestBaseCurrency := request.QueryStringParameters["base_currency"]
	requestTargetCurrency := request.QueryStringParameters["target_currency"]
	if len(requestBaseCurrency) > 0 && len(requestTargetCurrency) > 0 {
		BaseCurrency = requestBaseCurrency
		TargetCurrency = requestTargetCurrency
	} else {
		return 400, "The request data is incorrect: base_currency: " + requestBaseCurrency + ", target_currency: " + requestTargetCurrency
	}

	// Obtener los tipos de cambio para el par de la base de datos
	currency, err := db.SelectCurrency(BaseCurrency, TargetCurrency)
	if err != nil {
		return 400, "Error trying to get currency rate: " + err.Error()
	}

	Currency, err2 := json.Marshal(currency)
	if err2 != nil {
		return 500, "Error trying to convert to JSON currency object" + err2.Error()
	}
	return 200, string(Currency)
}

func UpdateCurrencies(body, User string) (int, string) {
	// Deserializar el body en un map[string]interface{}
	var requestData map[string]interface{}
	err := json.Unmarshal([]byte(body), &requestData)
	if err != nil {
		return 400, "The request data is incorrect: " + err.Error()
	}
	// Extraer las tasas de cambio y la fecha de actualización
	currencies := make(map[string]float64)
	if copRate, ok := requestData["cop"].(float64); ok {
		currencies["cop"] = copRate
	} else {
		return 400, "The COP rate is missing or incorrect"
	}

	if usdRate, ok := requestData["usd"].(float64); ok {
		currencies["usd"] = usdRate
	} else {
		return 400, "The USD rate is missing or incorrect"
	}

	if eurRate, ok := requestData["eur"].(float64); ok {
		currencies["eur"] = eurRate
	} else {
		return 400, "The EUR rate is missing or incorrect"
	}

	// Obtener la fecha de última actualización
	timeLastUpdate, ok := requestData["timeLastUpdate"].(string)
	if !ok || len(timeLastUpdate) == 0 {
		return 400, "The timeLastUpdate is missing or incorrect"
	}

	// Llamar a la función para actualizar las tasas de cambio
	err2 := db.UpdateCurrencies(currencies, timeLastUpdate)
	if err2 != nil {
		return 400, "Error when updating into the database: " + err2.Error()
	}

	return 200, "Update Ok"
}
