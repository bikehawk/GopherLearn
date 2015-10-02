package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	id         int
	testField1 string
	testField2 string
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("select id, test_field_1, test_field_2 from test_table where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(1)
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
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
