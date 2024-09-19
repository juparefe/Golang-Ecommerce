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
		script = "INSERT INTO orders_detail (OD_Currency, OD_Currency_Last_Symbol, OD_Currency_Symbol, OD_OrderId, OD_ProdId, OD_Quantity, OD_Price) "
		script += "VALUES ('" + od.OD_Currency + "','" + od.OD_Currency_Last_Symbol + "','" + od.OD_Currency_Symbol + "'," + strconv.Itoa(int(LastInsertId)) + ","
		script += strconv.Itoa(od.OD_ProdId) + "," + strconv.Itoa(od.OD_Quantity) + "," + strconv.FormatFloat(od.OD_Price, 'f', -1, 64) + ");"
		fmt.Println("Script Insert Orders Detail: ", script)
		_, err = Db.Exec(script)
		if err != nil {
			fmt.Println("Error inserting:", err.Error())
			return 0, err
		}
	}
	fmt.Println("InsertOrder > Successfull execution")
	return LastInsertId, nil
}

func SelectOrders(user, startDate, endDate string, orderId, page int) ([]models.Orders, error) {
	fmt.Println("Executing SelectOrders in database")
	var Orders []models.Orders

	script := "SELECT Order_Id, Order_UserUUID, Order_AddId, Order_Date, Order_Total FROM orders "
	if orderId > 0 {
		script += " WHERE Order_Id = " + strconv.Itoa(orderId)
	} else {
		offset := 0
		if page <= 0 {
			page = 1
		} else {
			offset = (10 * (page - 1))
		}
		if len(startDate) == 10 {
			startDate += " 23:59:59"
		}
		var where string
		var whereUser string = " Order_UserUUID = '" + user + "'"
		if len(startDate) > 0 && len(endDate) > 0 {
			where += " WHERE Order_Date BETWEEN '" + startDate + "' AND '" + endDate
		}
		if len(where) > 0 {
			where += " AND " + whereUser
		} else {
			where += " WHERE " + whereUser
		}
		limit := " LIMIT 10 "
		if offset > 0 {
			limit += " OFFSET " + strconv.Itoa(offset)
		}
		script += where + limit
	}

	fmt.Println("Script Select: ", script)
	err := DbConnect()
	if err != nil {
		return Orders, err
	}
	defer Db.Close()

	var rows *sql.Rows
	rows, err = Db.Query(script)
	if err != nil {
		return Orders, err
	}
	defer rows.Close()

	for rows.Next() {
		var o models.Orders
		var OrderAddId sql.NullInt32

		err := rows.Scan(&o.Order_Id, &o.Order_UserUUID, &OrderAddId, &o.Order_Date, &o.Order_Total)
		if err != nil {
			return Orders, err
		}

		o.Order_AddID = int(OrderAddId.Int32)

		var rowsD *sql.Rows
		scriptD := "SELECT OD_Currency, OD_Currency_Last_Symbol, OD_Currency_Symbol, OD_Id, OD_ProdId, OD_Quantity, OD_Price FROM orders_detail WHERE OD_OrderID = " + strconv.Itoa(o.Order_Id)
		fmt.Println("Script Select Order details: ", scriptD)

		rowsD, err = Db.Query(scriptD)
		if err != nil {
			return Orders, err
		}
		for rowsD.Next() {
			var OD_Currency, OD_Currency_Last_Symbol, OD_Currency_Symbol string
			var OD_Id, OD_ProdId, OD_Quantity int64
			var OD_Price float64
			err = rowsD.Scan(&OD_Currency, &OD_Currency_Last_Symbol, &OD_Currency_Symbol, &OD_Id, &OD_ProdId, &OD_Quantity, &OD_Price)
			if err != nil {
				return Orders, err
			}

			var od models.OrdersDetails
			od.OD_Currency = OD_Currency
			od.OD_Currency_Last_Symbol = OD_Currency_Last_Symbol
			od.OD_Currency_Symbol = OD_Currency_Symbol
			od.OD_Id = int(OD_Id)
			od.OD_ProdId = int(OD_ProdId)
			od.OD_Quantity = int(OD_Quantity)
			od.OD_Price = OD_Price
			o.OrderDetails = append(o.OrderDetails, od)
		}
		Orders = append(Orders, o)
		rowsD.Close()
	}

	fmt.Println("SelectOrders > Successfull execution")
	return Orders, nil
}
