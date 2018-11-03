/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"bufio"
	"encoding/base64"
	"net/http"
	"os"
)

//ConvertPicture : convert the picture in string base64
func ConvertPicture(w http.ResponseWriter, r *http.Request, path string) string {
	var size int64

	FileImg, err := os.Open(path)
	checkError(err)
	defer FileImg.Close()
	fInfo, _ := FileImg.Stat()
	size = fInfo.Size()
	buf := make([]byte, size)
	fReader := bufio.NewReader(FileImg)
	fReader.Read(buf)
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	return (imgBase64Str)
}
