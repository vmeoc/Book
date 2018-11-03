/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func fillStruct(get *data) *data {
	db, err := sql.Open(""+ os.Getenv("db_type") +"", ""+ os.Getenv("db_username")+":"+ os.Getenv("db_password") +"@tcp("+ os.Getenv("db_ip") +")/"+ os.Getenv("db_name") +"")
	//db, err := sql.Open("mysql", "root:PASSWORD@tcp(172.18.12.219)/Test")
	checkError(err)
	get.nbLine = getLines(get.nbLine, db)
	get.model = getTabStructFill(get, db, "model")
	get.img = getTabStructFill(get, db, "photo")
	defer db.Close()
	return (get)
}

func fillData(w http.ResponseWriter, r *http.Request) {
	var gif string

	get := &data{nil, nil, 0}
	get = fillStruct(get)
	if r.Method == "GET" {
		displayHTMLPage(get, w, r)
	}
	if r.Method == "POST" {
		url := strings.Split(r.URL.String(), "/")
		correctURL := checkString(url[1])
		bookMySQL(correctURL, get)
		gif = ConvertPicture(w, r, "../image/validation.gif")
		validation := "<html> <body> <p style=\"text-align:center;\"><img src=\"data:image/jpeg;base64," + gif + "\"></p> </body> </html>"
		refresh := "<html> <script> var timer = setTimeout(function() { window.location='http://"+ os.Getenv("tito_ip") +"' }, 3600); </script> </html>"
		//refresh := "<html> <script> var timer = setTimeout(function() { window.location='http://172.18.12.219:1234' }, 3600); </script> </html>"
		w.Write([]byte(fmt.Sprintf(validation)))
		w.Write([]byte(fmt.Sprintf(refresh)))
	}
}
