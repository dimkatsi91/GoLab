package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	hostname = "localhost"
	port     = 5432
	user     = "postgres"
	// passwd = "password"
	dbName = "intratel"
)

const (
	queryDeleteEmployeeById   = "DELETE FROM employees WHERE id=$1"
	queryDeleteEmployeeByAge  = "DELETE FROM employees WHERE age=$1"
	queryDeleteEmployeeByName = "DELETE FROM employees WHERE name=$1"
	queryEmployees            = "SELECT * from employees"
)

func main() {
	// Ask user for password & DB name to use in order to open postgresql connection
	//
	//var passwd, dbName string
	var passwd string

	fmt.Println("Please enter postgresql password: ")
	fmt.Scanln(&passwd)

	//fmt.Println("Please enter DB to use name: ")
	//fmt.Scanln(&dbName)

	// Attempt to open the connection
	//
	psql_string := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		hostname, port, user, passwd, dbName)
	postgres_db, err := sql.Open("postgres", psql_string)

	if err != nil {
		panic(err)
	}
	defer postgres_db.Close()

	err = postgres_db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection to Postgres successfully opened ...")

	// No, I want to insert two new records in this intratel DB inside this employees Table
	//
	var value string
	fmt.Println("Please enter name to delete: ")
	fmt.Scanln(&value)

	_, err = postgres_db.Exec(queryDeleteEmployeeByName, value)

	if err != nil {
		panic(err)
	}

	// Now, I have a DB called intratel and a table called employees
	// And I want to fetch the entries that I have inside this DB
	// This tables has id, name, age, salary
	//

	rows, err := postgres_db.Query(queryEmployees)

	if err != nil {
		panic(err)
	}

	// Now, we have the rows, so lets loop until this iterator ends
	// i.e. until the rows end ..
	//
	for rows.Next() {
		var id int
		var name string
		var age int
		var salary uint16

		err = rows.Scan(&id, &name, &age, &salary)
		if err != nil {
			panic(err)
		}

		// fmt.Println("ID NAME    AGE    SALARY")
		fmt.Println(id, name, age, salary)
	}

}
