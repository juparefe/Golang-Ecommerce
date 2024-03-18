package db

import (
	"database/sql"
	"fmt"

	//"strconv"
	//"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juparefe/Golang-Ecommerce/models"
	//"github.com/juparefe/Golang-Ecommerce/tools"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Executing InsertCategory in database")
	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	script := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('" + c.CategName + "', '" + c.CategPath + "')"
	fmt.Println("Script: ", script)

	var result sql.Result
	result, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}

	fmt.Println("InsertCategory > Succesfull execution: ", LastInsertId)
	return LastInsertId, nil

}
