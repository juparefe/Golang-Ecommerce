package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juparefe/Golang-Ecommerce/models"
	"github.com/juparefe/Golang-Ecommerce/tools"
)

func UpdateUser(u models.User) error {
	fmt.Println("Executing UpdateUser in database")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	script := "UPDATE users SET "
	comma := ""
	script = tools.CreateScript(script, "User_FirstName", "S", u.UserFirstName, 0, 0)
	script = tools.CreateScript(script, "Prod_LastName", "S", u.UserLastName, 0, 0)
	script = tools.CreateScript(script, "Prod_Price", "F", "", 0, u.ProdPrice)
	script = tools.CreateScript(script, "Prod_CategoryId", "N", "", u.ProdCategId, 0)
	script = tools.CreateScript(script, "Prod_Stock", "N", "", u.ProdStock, 0)
	script = tools.CreateScript(script, "Prod_Path", "S", u.ProdPath, 0, 0)
	script += " WHERE Prod_Id = " + strconv.Itoa(p.ProdId) + ";"
	fmt.Println("Script Update: ", script)
	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return err
	}

	fmt.Println("UpdateProduct > Succesfull execution")
	return nil
}

func DeleteProduct(id int) error {
	fmt.Println("Executing DeleteProduct in database")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	script := "DELETE FROM products WHERE Prod_Id = " + strconv.Itoa(id)
	fmt.Println("Script Delete: ", script)
	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return err
	}

	fmt.Println("DeleteProduct > Succesfull execution")
	return nil
}

func SelectProducts(p models.Product, choice, orderType, orderField string, page, pageSize int) (models.ProductRes, error) {
	fmt.Println("Executing SelectProducts in database")
	var ProductRes models.ProductRes
	var Prod []models.Product
	err := DbConnect()
	if err != nil {
		return ProductRes, err
	}
	defer Db.Close()

	var limit, script, scriptCount, where string
	script = "SELECT Prod_Id, Prod_Title, Prod_Description, Prod_CreatedAt, Prod_Updated, Prod_Price, Prod_Path, Prod_CategoryId, Prod_Stock FROM products"
	scriptCount = "SELECT COUNT(*) as records FROM products "
	switch choice {
	case "C":
		where = " WHERE Prod_CategoryId = " + strconv.Itoa(p.ProdCategId)
	case "K":
		join := " JOIN category ON Prod_CategoryId = Categ_Id AND Categ_Path LIKE '%" + strings.ToUpper(p.ProdCategPath) + "%'"
		script += join
		scriptCount += join
	case "P":
		where = " WHERE Prod_Id = " + strconv.Itoa(p.ProdId)
	case "S":
		where = " WHERE UCASE(CONCAT(Prod_Title, Prod_Description)) LIKE '%" + strings.ToUpper(p.ProdSearch) + "%'"
	case "U":
		where = " WHERE UCASE(Prod_Path) LIKE '%" + strings.ToUpper(p.ProdPath) + "%'"
	}
	scriptCount += where
	fmt.Println("Script Select Count: ", scriptCount)

	var rows *sql.Rows
	rows, err = Db.Query(scriptCount)
	defer rows.Close()
	if err != nil {
		fmt.Println("Error getting products:", err.Error())
		return ProductRes, err
	}
	rows.Next()
	var record sql.NullInt32
	err = rows.Scan(&record)
	if err != nil {
		fmt.Println("Error getting products:", err.Error())
		return ProductRes, err
	}
	records := int(record.Int32)
	if page > 0 {
		if records > pageSize {
			limit = " LIMIT " + strconv.Itoa(pageSize)
			if page > 1 {
				offset := (page - 1) * pageSize
				limit += " OFFSET " + strconv.Itoa(offset)
			}
		} else {
			limit = ""
		}
	}

	var orderBy string
	if len(orderField) > 0 {
		switch orderField {
		case "C":
			orderBy = " ORDER BY Prod_CategoryId "
		case "D":
			orderBy = " ORDER BY Prod_Description "
		case "F":
			orderBy = " ORDER BY Prod_CreatedAt "
		case "I":
			orderBy = " ORDER BY Prod_Id "
		case "P":
			orderBy = " ORDER BY Prod_Price "
		case "S":
			orderBy = " ORDER BY Prod_Stock "
		case "T":
			orderBy = " ORDER BY Prod_Title "
		}

		if orderType == "D" {
			orderBy += " DESC"
		}
	}
	script += where + orderBy + limit
	fmt.Println("Script Select: ", script)

	rows, err = Db.Query(script)
	if err != nil {
		fmt.Println("Error getting products:", err.Error())
		return ProductRes, err
	}
	for rows.Next() {
		var p models.Product
		var ProdId sql.NullInt32
		var ProdTitle, ProdDescription sql.NullString
		var ProdCreatedAt, ProdUpdated sql.NullString
		var ProdPrice sql.NullFloat64
		var ProdPath sql.NullString
		var ProdCategoryId, ProdStock sql.NullInt32

		err := rows.Scan(&ProdId, &ProdTitle, &ProdDescription, &ProdCreatedAt, &ProdUpdated, &ProdPrice, &ProdPath, &ProdCategoryId, &ProdStock)
		if err != nil {
			return ProductRes, err
		}

		p.ProdId = int(ProdId.Int32)
		p.ProdTitle = ProdTitle.String
		p.ProdDescription = ProdDescription.String
		p.ProdCreatedAt = ProdCreatedAt.String
		p.ProdUpdated = ProdUpdated.String
		p.ProdPrice = ProdPrice.Float64
		p.ProdPath = ProdPath.String
		p.ProdCategId = int(ProdCategoryId.Int32)
		p.ProdStock = int(ProdStock.Int32)
		fmt.Println("p to append", p)
		Prod = append(Prod, p)
	}
	ProductRes.TotalItems = records
	ProductRes.Data = Prod
	fmt.Println("SelectProducts > Succesfull execution")
	return ProductRes, nil
}

func UpdateStock(p models.Product) error {
	fmt.Println("Executing UpdateStock in database")
	if p.ProdStock == 0 {
		return errors.New("the stock to modify must be greater than 0")
	}
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	script := "UPDATE products SET Prod_Stock = Prod_Stock + " + strconv.Itoa(p.ProdStock) + " WHERE Prod_Id = " + strconv.Itoa(p.ProdId) + ";"
	fmt.Println("Script Update stock: ", script)
	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return err
	}

	fmt.Println("UpdateStock > Succesfull execution")
	return nil
}
