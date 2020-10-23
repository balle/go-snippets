package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db_user := "basti"
	db_pass := "secret123"
	db_host := "127.0.0.1"
	db_port := 3306
	db_name := "test"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db_user, db_pass, db_host, db_port, db_name))
	defer db.Close()

	if err != nil {
		log.Fatal(fmt.Sprintf("User %s cannot connect to db %s: %v\n", db_user, db_name, err))
	}

	rows, err := db.Query("SELECT * FROM products")

	if err != nil {
		log.Fatal("Select statement failed: %v\n", err)
	}

	for rows.Next() {
		var (
			id          int
			name        string
			price       float32
			description string
		)

		rows.Scan(&id, &name, &price, &description)
		fmt.Printf("%d %s %f %s", id, name, price, description)
	}

	stmt, _ := db.Prepare("INSERT INTO products(name, price, description) VALUES($1, $2, $3)")

	transaction, _ := db.Begin()

	stmt.Exec("Binary clock", 42.23, "Everything is not the real time")
	stmt.Exec("Some more product", 17.85, "All you need")

	transaction.Commit()
}
