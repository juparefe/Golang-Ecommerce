package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juparefe/Golang-Ecommerce/models"
	"github.com/juparefe/Golang-Ecommerce/secretmngr"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretmngr.GetSecret(os.Getenv("SecretName"))
	return err
}

// Conectarse a la base de datos y hacerle un Ping
func DbConnect() error {
	Db, err = sql.Open("mysql", ConnectionString(SecretModel))
	if err != nil {
		fmt.Println("Error:", err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Successful connection to the database")
	return nil
}

// Obtener el string formateado para conectarse a la base de datos
func ConnectionString(keys models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = keys.Username
	authToken = keys.Password
	dbEndpoint = keys.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println("dsn:", dsn)
	return dsn
}
