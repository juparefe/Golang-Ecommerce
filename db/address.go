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

func InsertAddress(a models.Address, User string) error {
	fmt.Println("Executing InsertAddress in database")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	script := "INSERT INTO addresses (Add_UserId, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name)"
	script += " VALUES ('" + User + "','" + a.AddAddress + "','" + a.AddCity + "','" + a.AddState + "','" + a.AddPostalCode + "','"
	script += a.AddPhone + "','" + a.AddTitle + "','" + a.AddName + "');"
	fmt.Println("Script Insert: ", script)

	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Error inserting:", err.Error())
		return err
	}

	fmt.Println("InsertAddress > Succesfull execution")
	return nil
}

func AddressExists(User string, id int) (error, bool) {
	fmt.Println("Executing AddressExists in database")
	err := DbConnect()
	if err != nil {
		return err, false
	}
	defer Db.Close()

	script := "SELECT 1 FROM addresses WHERE Add_Id='" + strconv.Itoa(id) + "AND Ad_UserId = '" + User + "';"
	fmt.Println("Script Search Address: ", script)

	rows, err := Db.Query(script)
	if err != nil {
		return err, false
	}

	var value string
	rows.Next()
	rows.Scan(&value)
	fmt.Println("AddressExists > Succesfull execution: ", value)

	if value == "1" {
		return nil, true
	}
	return nil, false
}

func UpdateAddress(a models.Address) error {
	fmt.Println("Executing UpdateAddress in database")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	script := "UPDATE addresses SET "
	script = tools.CreateScript(script, "Add_Address", "S", a.AddAddress, 0, 0)
	script = tools.CreateScript(script, "Add_City", "S", a.AddCity, 0, 0)
	script = tools.CreateScript(script, "Add_Phone", "S", a.AddPhone, 0, 0)
	script = tools.CreateScript(script, "Add_PostalCode", "S", a.AddPostalCode, 0, 0)
	script = tools.CreateScript(script, "Add_State", "S", a.AddState, 0, 0)
	script = tools.CreateScript(script, "Add_Title", "S", a.AddTitle, 0, 0)
	script += " WHERE Add_Id = " + strconv.Itoa(a.AddId) + ";"
	fmt.Println("Script Update: ", script)
	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return err
	}

	fmt.Println("UpdateAddress > Succesfull execution")
	return nil
}

func DeleteAddress(id int) error {
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

func SelectAddress(p models.Product, choice, orderType, orderField string, page, pageSize int) (models.ProductRes, error) {
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
