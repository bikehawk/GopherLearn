package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var (
	id         int
	testField1 string
	testField2 string
)

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	stmt, err := db.Prepare("SELECT id, test_field_1, test_field_2 FROM test_table")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &testField1, &testField2)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, testField1, testField2)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

}
