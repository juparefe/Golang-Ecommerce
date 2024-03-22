package db

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juparefe/Golang-Ecommerce/models"
	"github.com/juparefe/Golang-Ecommerce/tools"
)

func UpdateUser(u models.User, User string) error {
	fmt.Println("Executing UpdateUser in database")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	script := "UPDATE users SET "

	script = tools.CreateScript(script, "User_FirstName", "S", u.UserFirstName, 0, 0)
	script = tools.CreateScript(script, "User_LastName", "S", u.UserLastName, 0, 0)
	script = tools.CreateScript(script, "User_DataUpg", "S", tools.DateMySQL(), 0, 0)
	script += " WHERE User_UUID = '" + User + "';"
	fmt.Println("Script Update: ", script)

	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Error Updating:", err.Error())
		return err
	}

	fmt.Println("UpdateUser > Succesfull execution")
	return nil
}

func DeleteUser(id int) error {
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

func SelectUsers(UserId string) (models.User, error) {
	fmt.Println("Executing SelectUsers in database")
	User := models.User{}
	err := DbConnect()
	if err != nil {
		return User, err
	}
	defer Db.Close()

	script := "SELECT * FROM users WHERE User_UUID = '" + UserId + "';"
	fmt.Println("Script Select: ", script)

	var rows *sql.Rows
	rows, err = Db.Query(script)
	if err != nil {
		fmt.Println("Error getting users:", err.Error())
		return User, err
	}
	defer rows.Close()

	rows.Next()
	var firstName, lastName, dateUpg sql.NullString

	rows.Scan(&User.UserUUID, &User.UserEmail, &firstName, &lastName, &User.UserStatus, &User.UserDateAdd, &dateUpg)

	User.UserFirstName = firstName.String
	User.UserLastName = lastName.String
	User.UserDateUpd = dateUpg.String
	fmt.Println("SelectUsers > Succesfull execution")
	return User, nil
}
