package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetLastUpdateDate() (time.Time, error) {
	var lastUpdated time.Time
	err := Db.QueryRow("SELECT last_updated FROM exchange_rates ORDER BY last_updated DESC LIMIT 1").Scan(&lastUpdated)
	if err != nil && err != sql.ErrNoRows {
		return time.Time{}, err
	}
	return lastUpdated, nil
}

func UpdateRatesFromAPI() error {
	resp, err := http.Get(apiURL)
	if err != nil {
		return fmt.Errorf("error fetching rates: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	var rates ExchangeRates
	if err := json.Unmarshal(body, &rates); err != nil {
		return fmt.Errorf("error unmarshalling response: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}

	_, err = tx.Exec("DELETE FROM exchange_rates")
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error clearing exchange rates: %v", err)
	}

	for currency, rate := range rates.Rates {
		_, err := tx.Exec("REPLACE INTO exchange_rates (currency, rate, last_updated) VALUES (?, ?, ?)", currency, rate, time.Now())
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error updating database: %v", err)
		}
	}

	return tx.Commit()
}
