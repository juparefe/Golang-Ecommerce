package db

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juparefe/Golang-Ecommerce/models"
	"github.com/juparefe/Golang-Ecommerce/tools"
)

func SelectUser(UserId string) (models.User, error) {
	fmt.Println("Executing SelectUser in database")
	User := models.User{}
	err := DbConnect()
	if err != nil {
		return User, err
	}
	defer Db.Close()

	script := "SELECT * FROM users WHERE User_UUID = '" + UserId + "';"

	var rows *sql.Rows
	rows, err = Db.Query(script)
	if err != nil {
		fmt.Println("Script SelectUser: ", script)
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
	fmt.Println("SelectUser > Successful execution")
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

	if offSet > 0 {
		script += " OFFSET " + strconv.Itoa(offSet)
	}

	var rowsCount *sql.Rows
	rowsCount, err = Db.Query(scriptCount)
	if err != nil {
		fmt.Println("Script SelectUsersCount: ", scriptCount)
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
		fmt.Println("Script SelectUsers: ", script)
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
	fmt.Println("SelectUsers > Successful execution")
	lu.Data = User

	return lu, nil
}

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

	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Script UpdateUser: ", script)
		fmt.Println("Error Updating user:", err.Error())
		return err
	}

	fmt.Println("UpdateUser > Successful execution")
	return nil
}

func UpdateUserRole(u models.User, User string) error {
	fmt.Println("Executing UpdateUserRole in database")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	script := "UPDATE users SET "
	script = tools.CreateScript(script, "User_FirstName", "S", u.UserFirstName, 0, 0)
	script = tools.CreateScript(script, "User_LastName", "S", u.UserLastName, 0, 0)
	script = tools.CreateScript(script, "User_Status", "N", "", u.UserStatus, 0)
	script += " WHERE User_UUID = '" + User + "';"

	_, err = Db.Exec(script)
	if err != nil {
		fmt.Println("Script UpdateUserRole: ", script)
		fmt.Println("Error Updating user role:", err.Error())
		return err
	}

	fmt.Println("UpdateUser > Successful execution")
	return nil
}
