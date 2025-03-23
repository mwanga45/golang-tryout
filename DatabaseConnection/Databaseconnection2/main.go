package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"log"
)


func main(){
	db,err := sql.Open("mysql","root@tcp(0.0.0.0:3306)/field")

	if db.Ping(); err != nil{
		log.Fatalf("Failed to connect %v",err)
	}
    defer db.Close()
	createTb := `
	CREATE TABLE Attendeance(
	Student_id int not null,
	StudentName Varchar(255),
	PRIMARY KEY(Student_id));
	`
	_,err2 := db.Exec(createTb)

	if err2 != nil{
		log.Fatalf("Failed to execute query %v",err2)
	}
    fmt.Println("succefuly create table")

}