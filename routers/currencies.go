package routers

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/juparefe/Golang-Ecommerce/db"
)

func SelectCurrencies(request events.APIGatewayV2HTTPRequest) (int, string) {
	var err error
	var Currency string

	requestCurrency := request.QueryStringParameters["currency"]
	if len(requestCurrency) > 0 {
		Currency = requestCurrency
	}

	// Obtener la fecha de la última actualización de la base de datos
	lastUpdated, err := db.GetLastCurrenciesUpdateDate()
	if err != nil {
		return 400, "Error trying to get the last update date: " + err.Error()
	}

	// Verificar si se necesita actualizar los tipos de cambio
	if lastUpdated.IsZero() || lastUpdated.Before(time.Now().Add(-24*time.Hour)) {
		// Actualizar tasas de cambio
		err = db.UpdateRatesFromAPI()
		if err != nil {
			return 400, "Error updating rates from ExchangeRateAPI: " + err.Error()
		}
	}

	list, err := db.SelectCurrencies(Currency)
	if err != nil {
		return 400, "Error trying to get currencies: " + err.Error()
	}

	Currencies, err2 := json.Marshal(list)
	if err2 != nil {
		return 500, "Error trying to convert to JSON currencies list" + err2.Error()
	}
	return 200, string(Currencies)
}
