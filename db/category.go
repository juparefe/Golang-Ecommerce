package db

import (
	"database/sql"
	"fmt"

	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juparefe/Golang-Ecommerce/models"
	"github.com/juparefe/Golang-Ecommerce/tools"
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

func UpdateCategory(c models.Category) error {
	fmt.Println("Executing UpdateCategory in database")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	script := "UPDATE category SET "
	if len(c.CategName) > 0 {
		script += " Categ_Name = '" + tools.EscapeString(c.CategName) + "'"
	}
	if len(c.CategPath) > 0 {
		if !strings.HasSuffix(script, "SET ") {
			script += ", "
		}
		script += " Categ_Path = '" + tools.EscapeString(c.CategPath) + "'"
	}
	script += " WHERE Categ_ID = " + strconv.Itoa(c.CategId)
	fmt.Println("Script: ", script)
	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return err
	}

	fmt.Println("UpdateCategory > Succesfull execution")
	return nil
}

func DeleteCategory(id int) error {
	fmt.Println("Executing DeleteCategory in database")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	script := "DELETE FROM category WHERE Categ_Id = " + strconv.Itoa(id)
	fmt.Println("Script: ", script)
	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return err
	}

	fmt.Println("DeleteCategory > Succesfull execution")
	return nil
}
