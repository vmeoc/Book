/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"fmt"
	"net/http"
	"os"
)

func displayHTMLPage(get *data, w http.ResponseWriter, r *http.Request) {
	var checkEmpty int
	var test string

	test = ConvertPicture(w, r, "../image/background.jpg")
	background := "<html> <style> body {background-position: center;color: black; background-color: white; background-image: url(data:image/jpeg;base64," + test + "); background-repeat: no-repeat; background-attachment: fixed; -webkit-background-size: cover; -moz-background-size: cover; -o-background-size: cover; background-size: cover;} </style> </html>"
	begin := "<html> <style>  td {background-color: rgba(191,191,191,0.5);} table {border-collapse: collapse;} th {background-color: rgba(71,71,71,1); color: white;} table, th, td {border: 1px solid black} </style> <body> <table align=\"center\"> <th> Image </th> <th> Reservation </th> </body> </html>"
	w.Write([]byte(fmt.Sprintf(background)))
	w.Write([]byte(fmt.Sprintf(begin)))
	for i := 0; i < get.nbLine; i++ {
		if checkBooking(i) == 0 {
			imghtml := "<html> <style> .button-tab { color: black; background-color: transparent; transition-duration: 0.4s; cursor: pointer; padding: 15px 32px; margin: 10px; border-radius: 8px; border: 1px solid black; }  .button-tab:hover { background-color: #01A30A; color: white; } </style> <body> <tr> <td> <img src=\"data:image/png;base64," + get.img[i] + " \" height=\"200\" width=\"240\"> </td> <form method=\"POST\" action=\"" + get.model[i] + "\"> <td> <input type=\"submit\" id=\"" + get.model[i] + "\"value=\"Book\" class=\"button-tab\"> <input type=\"hidden\" id=\"" + get.model[i] + "\"> </form> </td> </tr> </body> </html>"
			w.Write([]byte(fmt.Sprintf(imghtml)))
			checkEmpty++
		}
	}
	end := "<html> </table> </html>"
	w.Write([]byte(fmt.Sprintf(end)))
	if checkEmpty == 0 {
		imgBack := "<html>  <body> <h1 style=\"text-align: center;\">No car Available </h1> </body> </html>"
		refresh := "<html> <script> var timer = setTimeout(function() { window.location='http://"+ os.Getenv("tito_ip") +"' }, 2000); </script> </html>"
		//refresh := "<html> <script> var timer = setTimeout(function() { window.location='http://172.18.12.219:1234' }, 2000); </script> </html>"
		w.Write([]byte(fmt.Sprintf(imgBack)))
		w.Write([]byte(fmt.Sprintf(refresh)))
	} else {
		returnBut := "<html> <style>  .button-ret {border-radius: 8px; color: black; border: 2px solid black; background-color: rgba(191,191,191,0.5); padding: 15px 32px; cursor: pointer; transition-duration: 0.4s;}  .button-ret:hover { background-color: black; color: white; } </style> <body> <form action=\"http://"+ os.Getenv("tito_ip") +"\">  <input type=\"submit\" Value=\"Return\" class=\"button-ret\"> </form> </body> </html>"
		//returnBut := "<html> <style>  .button-ret {border-radius: 8px; color: black; border: 2px solid black; background-color: rgba(191,191,191,0.5); padding: 15px 32px; cursor: pointer; transition-duration: 0.4s;}  .button-ret:hover { background-color: black; color: white; } </style> <body> <form action=\"http://172.18.12.219:1234\">  <input type=\"submit\" Value=\"Return\" class=\"button-ret\"> </form> </body> </html>"
		w.Write([]byte(fmt.Sprintf(returnBut)))
	}
}
