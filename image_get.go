package main

import (
	"database/sql"
	"fmt"
	"log"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "isucon:isucon@/isubata")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT 'name', 'data' FROM image")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer rows.Close()

	type Image struct {
		name string
		data []byte
	}

	for rows.Next() {
		image := Image{}

		err2 := rows.Scan(image.name, image.data)
		if err != nil {
			log.Fatalf(err2.Error())
		}

		fmt.Println(image.name)

		err3 := ioutil.WriteFile(image.name, image.data, 0666)
		if err != nil {
			log.Fatalf(err3.Error())
		}
	}
}