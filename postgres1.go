package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("here we go")
	db, err := sql.Open("postgres", "host=docker1 user=postgres password=mysecretpassword dbname=postgres sslmode=disable")
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
