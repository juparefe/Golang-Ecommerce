package routers

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/juparefe/Golang-Ecommerce/db"
)

func SelectCurrencies(request events.APIGatewayV2HTTPRequest) (int, string) {
	var err error
	var BaseCurrency string
	var TargetCurrency string

	requestBaseCurrency := request.QueryStringParameters["base_currency"]
	requestTargetCurrency := request.QueryStringParameters["target_currency"]
	if len(requestBaseCurrency) > 0 {
		BaseCurrency = requestBaseCurrency
	} else {
		if len(requestTargetCurrency) > 0 {
			TargetCurrency = requestTargetCurrency
		}
	}

	// Obtener los tipos de cambio para el par de la base de datos
	currency, err := db.SelectCurrency(BaseCurrency, TargetCurrency)
	if err != nil {
		return 400, "Error trying to get currency rate: " + err.Error()
	}

	// Verificar si se necesita actualizar los tipos de cambio
	if currency.LastUpdated.IsZero() || currency.LastUpdated.Before(time.Now().Add(-24*time.Hour)) {
		// Actualizar tasas de cambio
		currency, err = db.UpdateRatesFromAPI(BaseCurrency, TargetCurrency)
		if err != nil {
			return 400, "Error updating rates from ExchangeRateAPI: " + err.Error()
		}
	}

	Currency, err2 := json.Marshal(currency)
	if err2 != nil {
		return 500, "Error trying to convert to JSON currency object" + err2.Error()
	}
	return 200, string(Currency)
}
