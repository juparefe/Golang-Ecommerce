package db

import (
	"database/sql"
	"fmt"
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
		return Currency, err
	}
	Currency.BaseCurrency = baseCurrency.String
	Currency.CurrencyRate = currencyRate.Float64
	Currency.LastUpdated = lastUpdated
	Currency.TargetCurrency = targetCurrency.String

	fmt.Println("SelectCurrency > Successfull execution")
	return Currency, nil
}

func UpdateCurrencies(currencies map[string]float64, timeLastUpdate string) error {
	fmt.Println("Executing UpdateCurrencies in database")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	// Preparar el statement para actualizar la tabla
	stmt, err := Db.Prepare("UPDATE exchange_rates SET rate = ?, last_updated = ? WHERE base_currency = 'COP' AND target_currency = ?")
	if err != nil {
		fmt.Println("Error prepare", err)
		return err
	}
	defer stmt.Close()

	fmt.Println("Stmt", stmt)
	// Actualizar la tasa para COP -> COP
	_, err = stmt.Exec(currencies["cop"], timeLastUpdate, "COP")
	if err != nil {
		return err
	}

	// Actualizar la tasa para COP -> USD
	_, err = stmt.Exec(currencies["usd"], timeLastUpdate, "USD")
	if err != nil {
		return err
	}

	// Actualizar la tasa para COP -> EUR
	_, err = stmt.Exec(currencies["eur"], timeLastUpdate, "EUR")
	if err != nil {
		return err
	}

	fmt.Println("UpdateCurrencies > Successful execution")
	return nil
}
