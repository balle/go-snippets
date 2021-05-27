package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type product struct {
	id          int
	name        string
	price       float32
	description string
}

func do_insert(db *sql.DB, stmt *sql.Stmt, prd product) {
	res, err := stmt.Exec(prd.name, prd.price, prd.description)

	if err != nil {
		log.Fatal("Insert failed! %v", err)
	} else {
		lastId, _ := res.LastInsertId()
		fmt.Printf("Inserted product %d\n", lastId)
	}
}

func main() {
	db_user := "basti"
	db_pass := "secret123"
	db_host := "127.0.0.1"
	db_port := 3306
	db_name := "test"

	// Open() does not connect to the db, it just prepares for a connect
	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db_user, db_pass, db_host, db_port, db_name))
	defer db.Close()

	// Check if db connection is working as expected
	err := db.Ping()

	if err != nil {
		log.Fatal(fmt.Sprintf("User %s cannot connect to db %s: %v\n", db_user, db_name, err))
	}

	// Fetch some rows and iterate over them
	rows, err := db.Query("SELECT * FROM products")

	if err != nil {
		log.Fatal("Select statement failed: %v\n", err)
	}

	defer rows.Close()

	for rows.Next() {
		var prd product

		rows.Scan(&prd.id, &prd.name, &prd.price, &prd.description)
		fmt.Printf("Got %d %s %f %s\n", prd.id, prd.name, prd.price, prd.description)
	}

	// Execute two inserts in a transaction
	stmt, _ := db.Prepare("INSERT INTO products(name, price, description) VALUES(?, ?, ?)")

	transaction, _ := db.Begin()

	do_insert(db, stmt, product{name: "Binary clock", price: 42.23, description: "Everything is not the real time"})
	do_insert(db, stmt, product{name: "Some more product", price: 17.85, description: "All you need"})

	transaction.Commit()
}
