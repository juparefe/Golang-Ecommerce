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

func InsertProduct(p models.Product) (int64, error) {
	fmt.Println("Executing InsertProduct in database")
	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	script := "INSERT INTO products (Prod_Title "
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
		script += "," + strconv.FormatFloat(p.ProdPrice, 'e', -1, 64)
	}
	if p.ProdCategId > 0 {
		script += "," + strconv.Itoa(p.ProdCategId)
	}
	if p.ProdStock > 0 {
		script += "," + strconv.Itoa(p.ProdStock)
	}
	if len(p.ProdPath) > 0 {
		script += ",'" + tools.EscapeString(p.ProdPath) + "'"
	}
	script += ");"

	var result sql.Result
	result, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Script InsertProduct: ", script)
		fmt.Println("Error inserting product: ", err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}

	fmt.Println("InsertProduct > Successful execution: ", LastInsertId)
	return LastInsertId, nil
}

func UpdateProduct(p models.Product) error {
	fmt.Println("Executing UpdateProduct in database")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	script := "UPDATE products SET "
	script = tools.CreateScript(script, "Prod_Title", "S", p.ProdTitle, 0, 0)
	script = tools.CreateScript(script, "Prod_Description", "S", p.ProdDescription, 0, 0)
	script = tools.CreateScript(script, "Prod_Price", "F", "", 0, p.ProdPrice)
	script = tools.CreateScript(script, "Prod_CategoryId", "N", "", p.ProdCategId, 0)
	script = tools.CreateScript(script, "Prod_Stock", "N", "", p.ProdStock, 0)
	script = tools.CreateScript(script, "Prod_Path", "S", p.ProdPath, 0, 0)
	script += " WHERE Prod_Id = " + strconv.Itoa(p.ProdId) + ";"

	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Script UpdateProduct: ", script)
		fmt.Println("Error updating product:", err.Error())
		return err
	}

	fmt.Println("UpdateProduct > Successful execution")
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

	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Script DeleteProduct: ", script)
		fmt.Println("Error deleting product:", err.Error())
		return err
	}

	fmt.Println("DeleteProduct > Successful execution")
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
	script = `
		SELECT 
			p.Prod_Id,
			p.Prod_Title,
			p.Prod_Description,
			p.Prod_CreatedAt,
			p.Prod_Updated,
			p.Prod_Discount,
			p.Prod_Price,
			p.Prod_Path,
			p.Prod_CategoryId,
			p.Prod_Stock,
			(SELECT COUNT(od.OD_ProdId) FROM orders_detail AS od WHERE od.OD_ProdId = p.Prod_Id) AS SalesCount,
            (
                SELECT CASE WHEN RANK() OVER (ORDER BY COUNT(od.OD_ProdId) DESC) <= 15 THEN TRUE ELSE FALSE END
                FROM orders_detail AS od
                WHERE od.OD_ProdId = p.Prod_Id
                GROUP BY od.OD_ProdId
            ) AS ProdTop
		FROM products AS p`

	scriptCount = "SELECT COUNT(*) as records FROM products "
	switch choice {
	case "C":
		where = " WHERE p.Prod_CategoryId = " + strconv.Itoa(p.ProdCategId)
	case "K":
		joinCategory := " JOIN category AS c ON p.Prod_CategoryId = c.Categ_Id AND c.Categ_Path LIKE '%" + strings.ToUpper(p.ProdCategPath) + "%'"
		script += joinCategory
		scriptCount += joinCategory
		script += joinCategory
	case "P":
		where = " WHERE p.Prod_Id = " + strconv.Itoa(p.ProdId)
	case "S":
		where = " WHERE UCASE(CONCAT(p.Prod_Title, p.Prod_Description)) LIKE '%" + strings.ToUpper(p.ProdSearch) + "%'"
	case "U":
		where = " WHERE UCASE(p.Prod_Path) LIKE '%" + strings.ToUpper(p.ProdPath) + "%'"
	}
	scriptCount += where

	var rows *sql.Rows
	rows, err = Db.Query(scriptCount)
	defer rows.Close()
	if err != nil {
		fmt.Println("Script SelectProductsCount: ", scriptCount)
		fmt.Println("Error getting products count:", err.Error())
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
			orderBy = " ORDER BY p.Prod_CategoryId "
		case "D":
			orderBy = " ORDER BY p.Prod_Description "
		case "F":
			orderBy = " ORDER BY p.Prod_CreatedAt "
		case "I":
			orderBy = " ORDER BY p.Prod_Id "
		case "P":
			orderBy = " ORDER BY p.Prod_Price "
		case "S":
			orderBy = " ORDER BY p.Prod_Stock "
		case "T":
			orderBy = " ORDER BY p.Prod_Title "
		}

		if orderType == "D" {
			orderBy += " DESC"
		}
	}
	script += where + orderBy + limit

	rows, err = Db.Query(script)
	if err != nil {
		fmt.Println("Script SelectProducts: ", script)
		fmt.Println("Error getting products:", err.Error())
		return ProductRes, err
	}
	for rows.Next() {
		var p models.Product
		var ProdId sql.NullInt32
		var ProdTitle, ProdDescription sql.NullString
		var ProdCreatedAt, ProdUpdated sql.NullString
		var ProdDiscount, ProdPrice sql.NullFloat64
		var ProdPath sql.NullString
		var ProdCategoryId, ProdStock sql.NullInt32
		var ProdTop sql.NullBool

		err := rows.Scan(&ProdId, &ProdTitle, &ProdDescription, &ProdCreatedAt, &ProdUpdated, &ProdDiscount, &ProdPrice, &ProdPath, &ProdCategoryId, &ProdStock, &ProdTop)
		if err != nil {
			return ProductRes, err
		}

		p.ProdId = int(ProdId.Int32)
		p.ProdTitle = ProdTitle.String
		p.ProdDescription = ProdDescription.String
		p.ProdCreatedAt = ProdCreatedAt.String
		p.ProdUpdated = ProdUpdated.String
		p.ProdDiscount = ProdDiscount.Float64
		p.ProdPrice = ProdPrice.Float64
		p.ProdPath = ProdPath.String
		p.ProdCategId = int(ProdCategoryId.Int32)
		p.ProdStock = int(ProdStock.Int32)
		p.ProdTop = ProdTop.Bool
		Prod = append(Prod, p)
	}
	ProductRes.TotalItems = records
	ProductRes.Data = Prod
	fmt.Println("SelectProducts > Successful execution")
	return ProductRes, nil
}

func UpdateDiscount(p models.Product) error {
	fmt.Println("Executing UpdateDiscount in database")
	if p.ProdDiscount < 0 {
		return errors.New("the discount to modify must be greater than 0")
	}
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	script := "UPDATE products SET Prod_Discount = " + strconv.FormatFloat(p.ProdDiscount, 'f', 2, 64) + " WHERE Prod_Id = " + strconv.Itoa(p.ProdId) + ";"

	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Script UpdateDiscount: ", script)
		fmt.Println("Error updating discount:", err.Error())
		return err
	}

	fmt.Println("UpdateDiscount > Successful execution")
	return nil
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

	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Script UpdateStock: ", script)
		fmt.Println("Error updating stock:", err.Error())
		return err
	}

	fmt.Println("UpdateStock > Successful execution")
	return nil
}
