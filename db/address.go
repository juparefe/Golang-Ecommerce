package db

import (
	"database/sql"
	"fmt"
	"strconv"

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

	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Script InsertAddress: ", script)
		fmt.Println("Error inserting address:", err.Error())
		return err
	}

	fmt.Println("InsertAddress > Successful execution")
	return nil
}

func AddressExists(User string, id int) (error, bool) {
	fmt.Println("Executing AddressExists in database")
	err := DbConnect()
	if err != nil {
		return err, false
	}
	defer Db.Close()

	script := "SELECT 1 FROM addresses WHERE Add_Id='" + strconv.Itoa(id) + "' AND Add_UserId = '" + User + "';"

	rows, err := Db.Query(script)
	if err != nil {
		fmt.Println("Script Search Address: ", script)
		return err, false
	}

	var value string
	rows.Next()
	rows.Scan(&value)
	fmt.Println("AddressExists > Successful execution: ", value)

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
	script = tools.CreateScript(script, "Add_Name", "S", a.AddName, 0, 0)
	script = tools.CreateScript(script, "Add_Phone", "S", a.AddPhone, 0, 0)
	script = tools.CreateScript(script, "Add_PostalCode", "S", a.AddPostalCode, 0, 0)
	script = tools.CreateScript(script, "Add_State", "S", a.AddState, 0, 0)
	script = tools.CreateScript(script, "Add_Title", "S", a.AddTitle, 0, 0)
	script += " WHERE Add_Id = " + strconv.Itoa(a.AddId) + ";"
	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Script UpdateAddress: ", script)
		fmt.Println("Error updating address:", err.Error())
		return err
	}

	fmt.Println("UpdateAddress > Successful execution")
	return nil
}

func DeleteAddress(id int) error {
	fmt.Println("Executing DeleteAddress in database")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	script := "DELETE FROM addresses WHERE Add_Id = " + strconv.Itoa(id)
	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Script DeleteAddress: ", script)
		fmt.Println("Error deleting address:", err.Error())
		return err
	}

	fmt.Println("DeleteAddress > Successful execution")
	return nil
}

func SelectAddress(User string) ([]models.Address, error) {
	fmt.Println("Executing SelectAddress in database")
	Addresses := []models.Address{}
	err := DbConnect()
	if err != nil {
		return Addresses, err
	}
	defer Db.Close()

	script := "SELECT Add_Id, Add_Address, Add_City, Add_Name, Add_Phone, Add_PostalCode, Add_State, Add_Title FROM addresses WHERE Add_UserId = '" + User + "';"

	var rows *sql.Rows
	rows, err = Db.Query(script)
	if err != nil {
		fmt.Println("Script SelectAddress: ", script)
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

	fmt.Println("SelectAddresses > Successful execution")
	return Addresses, nil
}
