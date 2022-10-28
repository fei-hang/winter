package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() (err error) {
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("iii")
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	res, err := db.Query("show databases;")
	if err != nil {
		log.Fatal(err)
		return err
	}
	var databases []string = make([]string, 0, 20)
	// 拿到数据库
	for res.Next() {
		var name string
		res.Scan(&name)
		databases = append(databases, name)
	}
	log.Println(databases)

	data, err := db.Query("select * from user")
	if err != nil {
		return err
	}

	for data.Next() {
		var name map[string]string = make(map[string]string, 10)
		if err := data.Scan(&name); err != nil {
			log.Println("..")
			log.Fatal(err)
		}
		fmt.Printf("%s\n", name)
	}

	log.Println(data)
	return nil
}
