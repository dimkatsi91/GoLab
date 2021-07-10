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

	if err!= nil {
		panic(err)
	}

	fmt.Println("Connection to Postgresql DB successfully opened ...")
}
