package db

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juparefe/Golang-Ecommerce/models"
	"github.com/juparefe/Golang-Ecommerce/tools"
)

func InsertProduct(p models.Product) (int64, error) {
	fmt.Println("Executing InsertProduct in database")
	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	script := "INSERT INTO productos (Prod_Title "
	if len(p.ProdDescription) > 0 {
		script += ", Prod_Description"
	}
	if p.ProdPrice > 0 {
		script += ", Prod_Price"
	}
	if p.ProdCategId > 0 {
		script += ", Prod_CategoryId"
	}
	if p.ProdStock > 0 {
		script += ", Prod_Stock"
	}
	if len(p.ProdPath) > 0 {
		script += ", Prod_Path"
	}
	script += ") VALUES ('" + tools.EscapeString(p.ProdTitle) + "'"
	if len(p.ProdDescription) > 0 {
		script += ",'" + tools.EscapeString(p.ProdDescription) + "'"
	}
	if p.ProdPrice > 0 {
		script += ",'" + strconv.FormatFloat(p.ProdPrice, 'e', -1, 64) + "'"
	}
	if p.ProdCategId > 0 {
		script += ",'" + strconv.Itoa(p.ProdCategId) + "'"
	}
	if p.ProdStock > 0 {
		script += ",'" + strconv.Itoa(p.ProdStock) + "'"
	}
	if len(p.ProdPath) > 0 {
		script += ",'" + tools.EscapeString(p.ProdPath) + "'"
	}
	script += ")"
	fmt.Println("Script Insert: ", script)

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

	fmt.Println("InsertProduct > Succesfull execution: ", LastInsertId)
	return LastInsertId, nil

}
