package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(0.0.0.0:3306)/field")
	if db.Ping(); err != nil {
		log.Fatalf("Error occured during connection %v", err)
	}
	defer db.Close()

	createTable := `
	CREATE TABLE Employee(
    Emplo_id int NOT NULL ,
	Emplo_Name varchar(255),
	PRIMARY KEY(Emplo_id)
	);`

	createTableRegister := `
	CREATE TABLE Register(
	user_id int AUTO_INCREMENT,
	Username varchar(255),
	Password varchar(255),
	PRIMARY KEY(user_id)
	); `

	_,err3:= db.Exec(createTableRegister)
	if err3 != nil{
		log.Fatalf("failed to create table register %v",err3)
	}
	_,err2 := db.Exec(createTable)

	if err2 != nil{
		log.Fatalf("Error Occured %v",err2)
	}
	fmt.Println("succesfuly create table");
}