package main

import (
	"time"
	"fmt"
	"log"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("waiting for db...")
	time.Sleep(2000 * time.Millisecond)	

	fmt.Println("here we go")

	var (
		postgresHost string = "docker1"
		postgresPort string = "5432"
	)

	if os.Getenv("POSTGRES_PORT_5432_TCP_ADDR") != "" {
		postgresHost = os.Getenv("POSTGRES_PORT_5432_TCP_ADDR")
	}

	if os.Getenv("POSTGRES_PORT_5432_TCP_PORT") != "" {
		postgresPort = os.Getenv("POSTGRES_PORT_5432_TCP_PORT")
	}

	var connectionString string

	connectionString = fmt.Sprintf("host=%s port=%s user=postgres password=mysecretpassword dbname=postgres sslmode=disable", postgresHost, postgresPort)

	fmt.Println("connectionString: ", connectionString)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	var (
		name string
	)

	fmt.Println("iterating rows")
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Name:", name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("and we're done")
}
