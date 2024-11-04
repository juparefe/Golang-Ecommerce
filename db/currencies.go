package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/juparefe/Golang-Ecommerce/models"
)

func SelectCurrencies(BaseCurrency string) (models.Currencies, error) {
	fmt.Println("Executing SelectCurrencies in database")
	var Currencies models.Currencies
	var lastUpdatedString sql.NullString

	err := DbConnect()
	if err != nil {
		return Currencies, err
	}
	defer Db.Close()

	script := "SELECT * FROM exchange_rates WHERE base_currency = '" + BaseCurrency + "';"

	rows, err := Db.Query(script)
	if err != nil {
		fmt.Println("Script SelectCurrencies: ", script)
		fmt.Println("Error getting exchange rates:", err.Error())
		return Currencies, err
	}
	defer rows.Close()

	for rows.Next() {
		var baseCurrency, targetCurrency sql.NullString
		var currencyRate sql.NullFloat64

		err = rows.Scan(&baseCurrency, &targetCurrency, &currencyRate, &lastUpdatedString)
		if err != nil {
			fmt.Println("Error scanning row:", err.Error())
			return Currencies, err
		}

		// Asignar el valor de la tasa de cambio según la moneda objetivo
		if targetCurrency.Valid && currencyRate.Valid {
			switch strings.ToLower(targetCurrency.String) {
			case "cop":
				Currencies.COP = currencyRate.Float64
			case "eur":
				Currencies.EUR = currencyRate.Float64
			case "usd":
				Currencies.USD = currencyRate.Float64
			}
		}
	}

	// Convertir la última fecha de actualización a formato string para incluirla en el struct
	if lastUpdatedString.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", lastUpdatedString.String)
		if err != nil {
			fmt.Println("Error parsing last_updated time: ", err.Error())
			return Currencies, err
		}
		// Verificar si lastUpdated es hoy
		currentDate := time.Now().Truncate(24 * time.Hour)
		if parsedTime.Truncate(24 * time.Hour).Equal(currentDate) {
			Currencies.TimeLastUpdate = lastUpdatedString.String

			fmt.Println("SelectCurrencies > Successful execution")
			return Currencies, nil
		} else {
			errorMessage := fmt.Sprintf("Data is not from today: last updated on %s", lastUpdatedString.String)
			fmt.Println(errorMessage)
			return Currencies, errors.New(errorMessage)
		}
	} else {
		fmt.Println("No valid last_updated date found")
		return Currencies, errors.New("no valid last_updated date found")
	}
}

func SelectCurrencyByTarget(BaseCurrency, TargetCurrency string) (models.Currency, error) {
	fmt.Println("Executing SelectCurrency in database")
	var Currency models.Currency
	err := DbConnect()
	if err != nil {
		return Currency, err
	}
	defer Db.Close()

	script := "SELECT * FROM exchange_rates WHERE base_currency = '" + BaseCurrency + "' AND target_currency = '" + TargetCurrency + "';"

	var row *sql.Row
	row = Db.QueryRow(script)
	var baseCurrency, targetCurrency sql.NullString
	var currencyRate sql.NullFloat64
	var lastUpdatedString sql.NullString

	err = row.Scan(&baseCurrency, &targetCurrency, &currencyRate, &lastUpdatedString)
	if err != nil {
		fmt.Println("Script SelectCurrencyByTarget: ", script)
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
		Currency.BaseCurrency = baseCurrency.String
		Currency.TargetCurrency = targetCurrency.String
		Currency.CurrencyRate = currencyRate.Float64
		Currency.LastUpdated = lastUpdatedString.String

		fmt.Println("SelectCurrency > Successful execution")
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
		err := Db.QueryRow(checkQuery, baseCurrency, targetCurrency).Scan(&exists)
		if err != nil {
			fmt.Println("Script check currencies: ", checkQuery)
			return fmt.Errorf("error checking if row exists: %v", err)
		}

		if exists {
			// Si la fila existe, actualizar la tasa de cambio y la fecha de actualización
			updateQuery := `UPDATE exchange_rates SET rate = ? WHERE base_currency = ? AND target_currency = ?`
			_, err = Db.Exec(updateQuery, rate, baseCurrency, targetCurrency)
			if err != nil {
				fmt.Println("Script UpdateCurrencies: ", updateQuery)
				return fmt.Errorf("error updating row: %v", err)
			}
		} else {
			// Si la fila no existe, insertar una nueva fila
			insertQuery := `
				INSERT INTO exchange_rates (base_currency, target_currency, rate)
				VALUES (?, ?, ?)
			`
			_, err = Db.Exec(insertQuery, baseCurrency, targetCurrency, rate)
			if err != nil {
				fmt.Println("Script InsertCurrencies: ", insertQuery)
				return fmt.Errorf("error inserting new row: %v", err)
			}
		}
	}
	fmt.Println("UpdateCurrencies > Successful execution")
	return nil
}
