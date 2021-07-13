package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	empNo int
	firstName string
	lastName string
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/employees")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select emp_no, first_name, last_name from employees where emp_no = ?", 10001)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&empNo, &firstName, &lastName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(empNo, firstName, lastName)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("select emp_no, first_name, last_name from employees where emp_no = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err = stmt.Query(10001)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		log.Println(empNo, firstName, lastName)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
