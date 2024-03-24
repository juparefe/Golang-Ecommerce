package db

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juparefe/Golang-Ecommerce/models"
)

func InsertOrder(o models.Orders) (int64, error) {
	fmt.Println("Executing InsertOrder in database")
	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	script := "INSERT INTO orders (Order_UserUUID, Order_Total, Order_AddId)"
	script += " VALUES ('" + o.Order_UserUUID + "','" + strconv.FormatFloat(o.Order_Total, 'f', -1, 64) + "','" + strconv.Itoa(o.Order_AddID) + "');"
	fmt.Println("Script Insert Orders: ", script)

	var result sql.Result
	result, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Error inserting:", err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}

	for _, od := range o.OrderDetails {
		script = "INSERT INTO orders_detail (OD_OrderId, OD_ProdId, OD_Quantity, OD_Price) VALUES (" + strconv.Itoa(int(LastInsertId)) + ","
		script += strconv.Itoa(od.OD_ProdId) + "," + strconv.Itoa(od.OD_Quantity) + "," + strconv.FormatFloat(od.OD_Price, 'f', -1, 64) + ");"
	}
	fmt.Println("Script Insert Orders Detail: ", script)
	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Error inserting:", err.Error())
		return 0, err
	}

	fmt.Println("InsertOrder > Succesfull execution")
	return LastInsertId, nil
}

func SelectOrder(User string) ([]models.Address, error) {
	fmt.Println("Executing SelectAddress in database")
	Addresses := []models.Address{}
	err := DbConnect()
	if err != nil {
		return Addresses, err
	}
	defer Db.Close()

	script := "SELECT Add_Id, Add_Address, Add_City, Add_Name, Add_Phone, Add_PostalCode, Add_State, Add_Title FROM addresses WHERE Add_UserId = '" + User + "';"
	fmt.Println("Script Select: ", script)

	var rows *sql.Rows
	rows, err = Db.Query(script)
	if err != nil {
		fmt.Println("Error getting addresses:", err.Error())
		return Addresses, err
	}
	for rows.Next() {
		var a models.Address
		var AddId sql.NullInt32
		var AddAddress, AddCity, AddName, AddPhone, AddPostalCode, AddState, AddTitle sql.NullString

		err := rows.Scan(&AddId, &AddAddress, &AddCity, &AddName, &AddPhone, &AddPostalCode, &AddState, &AddTitle)
		if err != nil {
			return Addresses, err
		}

		a.AddId = int(AddId.Int32)
		a.AddAddress = AddAddress.String
		a.AddCity = AddCity.String
		a.AddName = AddName.String
		a.AddPhone = AddPhone.String
		a.AddPostalCode = AddPostalCode.String
		a.AddState = AddState.String
		a.AddTitle = AddTitle.String
		Addresses = append(Addresses, a)
	}

	fmt.Println("SelectAddresses > Succesfull execution")
	return Addresses, nil
}
