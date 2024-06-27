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
	fmt.Println("Script Update: ", script)
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
	fmt.Println("Script Delete: ", script)
	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return err
	}

	fmt.Println("DeleteCategory > Succesfull execution")
	return nil
}

func SelectCategories(CategId int, Slug string) ([]models.Category, error) {
	fmt.Println("Executing SelectCategory in database")
	var Categ []models.Category
	err := DbConnect()
	if err != nil {
		return Categ, err
	}
	defer Db.Close()

	script := "SELECT Categ_Id, Categ_Name, Categ_Path FROM category "
	if CategId > 0 {
		script += "WHERE Categ_Id = " + strconv.Itoa(CategId)
	}
	if len(Slug) > 0 {
		script += "WHERE Categ_Path LIKE '%" + Slug + "%'"
	}
	fmt.Println("Script Select: ", script)

	var rows *sql.Rows
	rows, err = Db.Query(script)
	if err != nil {
		fmt.Println("Error getting categories:", err.Error())
		return Categ, err
	}
	for rows.Next() {
		var c models.Category
		var categId sql.NullInt32
		var categName sql.NullString
		var categPath sql.NullString

		err = rows.Scan(&categId, &categName, &categPath)
		if err != nil {
			fmt.Println("Error adding row:", err.Error())
			return Categ, err
		}
		c.CategId = int(categId.Int32)
		c.CategName = categName.String
		c.CategPath = categPath.String
		Categ = append(Categ, c)
	}

	fmt.Println("SelectCategory > Succesfull execution")
	return Categ, nil
}

func SelectTopCategories() ([]models.Category, error) {
	fmt.Println("Executing SelectTopCategories in database")
	var Categ []models.Category
	err := DbConnect()
	if err != nil {
		return Categ, err
	}
	defer Db.Close()

	script := `SELECT c.Categ_Id, c.Categ_Name, c.Categ_Path, COALESCE(SUM(od.OD_Quantity), 0) AS TotalSold
				FROM category c
				LEFT JOIN products p ON c.Categ_Id = p.Prod_CategoryId
				LEFT JOIN orders_detail od ON p.Prod_Id = od.OD_ProdId
				LEFT JOIN orders o ON od.OD_OrderId = o.Order_Id
				GROUP BY c.Categ_Id, c.Categ_Name, c.Categ_Path
				ORDER BY TotalSold DESC
				LIMIT 5;`
	fmt.Println("Script Select: ", script)

	var rows *sql.Rows
	rows, err = Db.Query(script)
	if err != nil {
		fmt.Println("Error getting top categories:", err.Error())
		return Categ, err
	}
	for rows.Next() {
		var c models.Category
		var categId sql.NullInt32
		var categName sql.NullString
		var categPath sql.NullString
		var categTotalSold int

		err = rows.Scan(&categId, &categName, &categPath, categTotalSold)
		if err != nil {
			fmt.Println("Error adding row:", err.Error())
			return Categ, err
		}
		c.CategId = int(categId.Int32)
		c.CategName = categName.String
		c.CategPath = categPath.String
		c.CategTotalSold = int(categTotalSold)
		Categ = append(Categ, c)
	}

	fmt.Println("SelectCategory > Succesfull execution")
	return Categ, nil
}
