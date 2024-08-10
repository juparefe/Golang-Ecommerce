package db

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/juparefe/Golang-Ecommerce/models"
)

func SelectCurrency(BaseCurrency, TargetCurrency string) (models.Currency, error) {
	fmt.Println("Executing SelectCurrency in database")
	var Currency models.Currency
	err := DbConnect()
	if err != nil {
		return Currency, err
	}
	defer Db.Close()

	script := "SELECT * FROM currencies WHERE base_currency = '" + BaseCurrency + "' AND target_currency = '" + TargetCurrency + "';"
	fmt.Println("Script Select: ", script)

	var row *sql.Row
	row = Db.QueryRow(script)
	var baseCurrency, targetCurrency sql.NullString
	var currencyRate sql.NullFloat64
	var lastUpdated time.Time

	err = row.Scan(&baseCurrency, &currencyRate, &lastUpdated, &targetCurrency)
	if err != nil {
		fmt.Println("Error scanning row with base_currency = '" + BaseCurrency + "' and target_currency = '" + TargetCurrency + ", " + err.Error())
		UpdateRatesFromAPI(BaseCurrency, TargetCurrency)
		return Currency, err
	}
	Currency.BaseCurrency = baseCurrency.String
	Currency.CurrencyRate = currencyRate.Float64
	Currency.LastUpdated = lastUpdated
	Currency.TargetCurrency = targetCurrency.String

	fmt.Println("SelectCurrency > Successfull execution")
	return Currency, nil
}

func UpdateRatesFromAPI(BaseCurrency, TargetCurrency string) (models.Currency, error) {
	apiKey := os.Getenv("UrlPrefix")
	apiUrl := "https://v6.exchangerate-api.com/v6/" + apiKey + "/latest/" + BaseCurrency
	fmt.Println("Executing UpdateRatesFromAPI with url: ", apiUrl)

	resp, err := http.Get(apiUrl)
	fmt.Println("Respuesta obtenida:", resp)
	if err != nil {
		fmt.Println("Error fetching data from API: ", err.Error())
		return models.Currency{}, err
	}
	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return fmt.Errorf("error reading response: %v", err)
	// }

	// var rates ExchangeRates
	// if err := json.Unmarshal(body, &rates); err != nil {
	// 	return fmt.Errorf("error unmarshalling response: %v", err)
	// }

	// tx, err := Db.Begin()
	// if err != nil {
	// 	return fmt.Errorf("error starting transaction: %v", err)
	// }

	// _, err = tx.Exec("DELETE FROM exchange_rates")
	// if err != nil {
	// 	tx.Rollback()
	// 	return fmt.Errorf("error clearing exchange rates: %v", err)
	// }

	// for currency, rate := range rates.Rates {
	// 	_, err := tx.Exec("REPLACE INTO exchange_rates (currency, rate, last_updated) VALUES (?, ?, ?)", currency, rate, time.Now())
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return fmt.Errorf("error updating database: %v", err)
	// 	}
	// }

	// return tx.Commit()
	return models.Currency{}, nil
}
