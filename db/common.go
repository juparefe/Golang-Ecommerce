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
	SecretModel, err = secretmngr.GetSecretLambdaProxy(os.Getenv("SecretName"))
	return err
}

// Conectarse a la base de datos y hacerle un Ping
func DbConnect() error {
	Db, err = sql.Open("mysql", ConnectionString(SecretModel))
	if err != nil {
		fmt.Println("Error trying connection to the database:", err.Error())
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
	return dsn
}

func UserIsAdmin(userUUID string) (bool, string) {
	err := DbConnect()
	if err != nil {
		return false, err.Error()
	}
	defer Db.Close()
	script := "SELECT 1 FROM users WHERE User_UUID='" + userUUID + "' AND User_Status IN (1,2);"

	rows, err := Db.Query(script)
	if err != nil {
		fmt.Println("Script Select user role: ", script)
		return false, err.Error()
	}

	var value string
	rows.Next()
	rows.Scan(&value)
	fmt.Println("UserIsAdmin > Successful execution: ", value)

	if value == "1" {
		return true, ""
	}
	return false, "User is not admin"
}

func UserIsSuperAdmin(userUUID string) (bool, string) {
	err := DbConnect()
	if err != nil {
		return false, err.Error()
	}
	defer Db.Close()
	script := "SELECT 1 FROM users WHERE User_UUID='" + userUUID + "' AND User_Status=1;"

	rows, err := Db.Query(script)
	if err != nil {
		fmt.Println("Script Select user role: ", script)
		return false, err.Error()
	}

	var value string
	rows.Next()
	rows.Scan(&value)
	fmt.Println("UserIsSuperAdmin > Successful execution: ", value)

	if value == "1" {
		return true, ""
	}
	return false, "User is not super admin"
}

func UserExists(userUUID string) (error, bool) {
	err := DbConnect()
	if err != nil {
		return err, false
	}
	defer Db.Close()
	script := "SELECT 1 FROM users WHERE User_UUID='" + userUUID + "'"

	rows, err := Db.Query(script)
	if err != nil {
		fmt.Println("Script Search User: ", script)
		return err, false
	}

	var value string
	rows.Next()
	rows.Scan(&value)
	fmt.Println("UserExists > Successful execution: ", value)

	if value == "1" {
		return nil, true
	}
	return nil, false
}
