/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"fmt"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func bookMySQL(url string, get *data) {
	db, err := sql.Open(""+ os.Getenv("db_type") +"", ""+ os.Getenv("db_username")+":"+ os.Getenv("db_password") +"@tcp("+ os.Getenv("db_ip") +")/"+ os.Getenv("db_name") +"")
	//db, err := sql.Open("mysql", "root:PASSWORD@tcp(172.18.12.219)/Test")
	checkError(err)
	for i := 0; i < get.nbLine; i++ {
		if url == get.model[i] {
			query, err := db.Query("UPDATE garage SET Book = 1 WHERE id=?", i+1)
			checkError(err)
			fmt.Printf("Booking Done\n")
			defer query.Close()
		}
	}
	defer db.Close()
}
