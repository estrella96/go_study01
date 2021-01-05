package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("localhost:8888", nil)
}
func index(rw http.ResponseWriter, req *http.Request) {
	message := ""
	if req.Method == "POST" {
		req.ParseMultipartForm(32 << 20)
		file, handler, err := req.FormFile("upfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		message = "success"

	}
	t, _ := template.ParseFiles("index.tpl")
	t.Execute(rw, message)
}
