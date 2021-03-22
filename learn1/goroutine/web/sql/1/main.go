package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDb()error  {
	var err error
	dsn := "root:root@tcp(localhost:3306)/test"
	DB, err =sql.Open("mysql",dsn)
	if err != nil {
		fmt.Printf("init failed err: %v\n", err)
		return err
	}
	return nil
}

type User struct {
	Id int
	Name sql.NullString
	Age int
}

func testQuery()  {
	sqlstr := "select * from user"
	var user User
	row := DB.QueryRow(sqlstr,)
	err:=row.Scan(&user.Name,&user.Id,&user.Age,)
	if err != nil {
		fmt.Printf("init failed err: %v\n", err)
		return
	}
	fmt.Println(user.Name,user.Id,user.Age,)
}



func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init failed err: %v\n", err)
		return
	}
	testQuery()
}
