package db

import (
	"database/sql"
	"fmt"
	"strings"
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

	// Base currency será siempre "COP"
	baseCurrency := "COP"

	for targetCurrency, rate := range currencies {
		// Convertir targetCurrency a mayúsculas
		targetCurrency = strings.ToUpper(targetCurrency)

		// Verificar si la fila ya existe en la base de datos
		var exists bool
		checkQuery := `
			SELECT EXISTS (
				SELECT 1 FROM exchange_rates 
				WHERE base_currency = ? AND target_currency = ?
			)
		`
		fmt.Println("Script check currencies: ", checkQuery)
		err := Db.QueryRow(checkQuery, baseCurrency, targetCurrency).Scan(&exists)
		if err != nil {
			return fmt.Errorf("error checking if row exists: %v", err)
		}

		if exists {
			// Si la fila existe, actualizar la tasa de cambio y la fecha de actualización
			updateQuery := `
				UPDATE exchange_rates
				SET rate = ?, last_updated = ?
				WHERE base_currency = ? AND target_currency = ?
			`
			fmt.Println("Script update currencies: ", updateQuery)
			_, err = Db.Exec(updateQuery, rate, timeLastUpdate, baseCurrency, targetCurrency)
			if err != nil {
				return fmt.Errorf("error updating row: %v", err)
			}
		} else {
			// Si la fila no existe, insertar una nueva fila
			insertQuery := `
				INSERT INTO exchange_rates (base_currency, target_currency, rate, last_updated)
				VALUES (?, ?, ?, ?)
			`
			fmt.Println("Script insert currencies: ", insertQuery)
			_, err = Db.Exec(insertQuery, baseCurrency, targetCurrency, rate, timeLastUpdate)
			if err != nil {
				return fmt.Errorf("error inserting new row: %v", err)
			}
		}
	}
	fmt.Println("UpdateCurrencies > Successful execution")
	return nil
}
