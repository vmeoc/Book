/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"fmt"
	"log"
	"net/http"
)

type data struct {
	img    []string
	model  []string
	nbLine int
}

func main() {
	http.HandleFunc("/", fillData)

	fmt.Printf("Connection\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
