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

	script := "SELECT * FROM exchange_rates WHERE base_currency = '" + BaseCurrency + "' AND target_currency = '" + TargetCurrency + "';"
	fmt.Println("Script Select: ", script)

	var row *sql.Row
	row = Db.QueryRow(script)
	var baseCurrency, targetCurrency sql.NullString
	var currencyRate sql.NullFloat64
	var lastUpdatedString sql.NullString

	err = row.Scan(&baseCurrency, &targetCurrency, &currencyRate, &lastUpdatedString)
	if err != nil {
		fmt.Println("Error scanning row with base_currency = '" + BaseCurrency + "' and target_currency = '" + TargetCurrency + ", " + err.Error())
		return Currency, err
	}

	// Convertir lastUpdatedString a time.Time
	var lastUpdated time.Time
	if lastUpdatedString.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", lastUpdatedString.String)
		if err != nil {
			fmt.Println("Error parsing last_updated time: ", err.Error())
			return Currency, err
		}
		lastUpdated = parsedTime
	}

	// Verificar si lastUpdated es hoy
	currentDate := time.Now().Truncate(24 * time.Hour)
	if lastUpdated.Truncate(24 * time.Hour).Equal(currentDate) {
		fmt.Println("Data is from today")
		Currency.BaseCurrency = baseCurrency.String
		Currency.TargetCurrency = targetCurrency.String
		Currency.CurrencyRate = currencyRate.Float64
		Currency.LastUpdated = lastUpdatedString.String

		fmt.Println("SelectCurrency > Successfull execution")
		return Currency, nil
	} else {
		fmt.Println("Data is not from today")
		return Currency, nil
	}
}

func UpdateCurrencies(currencies map[string]float64) error {
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
		checkQuery := `SELECT EXISTS (SELECT 1 FROM exchange_rates WHERE base_currency = ? AND target_currency = ?)`
		fmt.Println("Script check currencies: ", checkQuery)
		err := Db.QueryRow(checkQuery, baseCurrency, targetCurrency).Scan(&exists)
		if err != nil {
			return fmt.Errorf("error checking if row exists: %v", err)
		}

		if exists {
			// Si la fila existe, actualizar la tasa de cambio y la fecha de actualización
			updateQuery := `UPDATE exchange_rates SET rate = ? WHERE base_currency = ? AND target_currency = ?`
			fmt.Println("Script update currencies: ", updateQuery)
			_, err = Db.Exec(updateQuery, rate, baseCurrency, targetCurrency)
			if err != nil {
				return fmt.Errorf("error updating row: %v", err)
			}
		} else {
			// Si la fila no existe, insertar una nueva fila
			insertQuery := `
				INSERT INTO exchange_rates (base_currency, target_currency, rate)
				VALUES (?, ?, ?)
			`
			fmt.Println("Script insert currencies: ", insertQuery)
			_, err = Db.Exec(insertQuery, baseCurrency, targetCurrency, rate)
			if err != nil {
				return fmt.Errorf("error inserting new row: %v", err)
			}
		}
	}
	fmt.Println("UpdateCurrencies > Successful execution")
	return nil
}
