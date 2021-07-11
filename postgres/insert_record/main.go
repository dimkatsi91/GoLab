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
	// dbName = "postgres"
)

const (
	queryEmployees      = "SELECT * from employees"
	queryInsertEmployee = "INSERT INTO EMPLOYEES(name, age, salary) values ($1, $2, $3)"
)

func main() {
	// Ask user for password & DB name to use in order to open postgresql connection
	//
	var passwd, dbName string

	fmt.Println("Please enter postgresql password: ")
	fmt.Scanln(&passwd)

	fmt.Println("Please enter DB to use name: ")
	fmt.Scanln(&dbName)

	// Attempt to open the connection
	//
	psql_string := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		hostname, port, user, passwd, dbName)
	postgres_db, err := sql.Open("postgres", psql_string)

	error_sanity(err)
	defer postgres_db.Close()

	err = postgres_db.Ping()

	error_sanity(err)

	fmt.Println("Connection to Postgresql DB successfully opened ...")

	// No, I want to insert two new records in this intratel DB inside this employees Table
	//
	_, err = postgres_db.Exec(queryInsertEmployee, "Andres", 44, 3500.00)

	error_sanity(err)

	_, err = postgres_db.Exec(queryInsertEmployee, "Lionel", 36, 4500.00)

	error_sanity(err)

	// Now, I have a DB called intratel and a table called employees
	// And I want to fetch the entries that I have inside this DB
	// This tables has id, name, age, salary
	//

	rows, err := postgres_db.Query(`SELECT * FROM "employees"`)

	error_sanity(err)

	// Now, we have the rows, so lets loop until this iterator ends
	// i.e. until the rows end ..
	//
	for rows.Next() {
		var id int
		var name string
		var age int
		var salary uint16

		err = rows.Scan(&id, &name, &age, &salary)
		error_sanity(err)

		// fmt.Println("ID NAME    AGE    SALARY")
		fmt.Println(id, name, age, salary)
	}

}

func error_sanity(err error) {
	if err != nil {
		panic(err)
	}
}
