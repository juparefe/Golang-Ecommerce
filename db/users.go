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
	script = tools.CreateScript(script, "User_DateUpg", "S", tools.DateMySQL(), 0, 0)
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

func UpdateUserRole(u models.User, User string) error {
	fmt.Println("Executing UpdateUser in database")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	script := "UPDATE users SET "

	script = tools.CreateScript(script, "User_FirstName", "S", u.UserFirstName, 0, 0)
	script = tools.CreateScript(script, "User_LastName", "S", u.UserLastName, 0, 0)
	script = tools.CreateScript(script, "User_DateUpg", "S", tools.DateMySQL(), 0, 0)
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

func SelectUser(UserId string) (models.User, error) {
	fmt.Println("Executing SelectUser in database")
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
		fmt.Println("Error getting user:", err.Error())
		return User, err
	}
	defer rows.Close()

	rows.Next()
	var firstName, lastName, dateUpg sql.NullString

	rows.Scan(&User.UserUUID, &User.UserEmail, &firstName, &lastName, &User.UserStatus, &User.UserDateAdd, &dateUpg)

	User.UserFirstName = firstName.String
	User.UserLastName = lastName.String
	User.UserDateUpg = dateUpg.String
	fmt.Println("SelectUser > Succesfull execution")
	return User, nil
}

func SelectUsers(Page int) (models.ListUsers, error) {
	fmt.Println("Executing SelectUsers in database")
	var lu models.ListUsers
	User := []models.User{}
	err := DbConnect()
	if err != nil {
		return lu, err
	}
	defer Db.Close()

	var offSet int = (Page * 10) - 10
	script := "SELECT * FROM users LIMIT 10"
	scriptCount := "SELECT COUNT(*) as records FROM users;"
	fmt.Println("Script Select count: ", scriptCount)

	if offSet > 0 {
		script += " OFFSET " + strconv.Itoa(offSet)
	}
	fmt.Println("Script Select: ", script)

	var rowsCount *sql.Rows
	rowsCount, err = Db.Query(scriptCount)
	if err != nil {
		fmt.Println("Error getting users count:", err.Error())
		return lu, err
	}
	defer rowsCount.Close()

	rowsCount.Next()
	var records int

	rowsCount.Scan(&records)

	lu.TotalItems = records

	var rows *sql.Rows
	rows, err = Db.Query(script)
	if err != nil {
		fmt.Println("Error getting users:", err.Error())
		return lu, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User
		var firstName, lastName, dateUpg sql.NullString

		rows.Scan(&u.UserUUID, &u.UserEmail, &firstName, &lastName, &u.UserStatus, &u.UserDateAdd, &dateUpg)

		u.UserFirstName = firstName.String
		u.UserLastName = lastName.String
		u.UserDateUpg = dateUpg.String
		User = append(User, u)
	}
	fmt.Println("SelectUsers > Succesfull execution")
	lu.Data = User

	return lu, nil
}
