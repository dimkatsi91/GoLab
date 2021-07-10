package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	hostname = "localhost"
	port = 5432
	user = "postgres"
	// passwd = "password"
	// dbName = "postgres"
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
		"host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable",
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

	fmt.Println("Connection to Postgresql DB successfully opened ...")

	

	// No, I want to insert two new records in this intratel DB inside this employees Table
	//
	insertStatement := `insert into employees("name", "age", "salary") values ('Andres', 44, 3500.00), ('Lionel', 36, 4500.00)`
	_, err = postgres_db.Exec(insertStatement)

	if err != nil {
		panic(err)
	}


	// Now, I have a DB called intratel and a table called employees
	// And I want to fetch the entries that I have inside this DB
	// This tables has id, name, age, salary
	//

	rows, err := postgres_db.Query(`SELECT * FROM "employees"`)

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


