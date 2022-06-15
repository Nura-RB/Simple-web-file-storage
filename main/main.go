package main

import (
	"net/http"
	"text/template"
	"utils"
)

func display(w http.ResponseWriter, data interface{}) {
	var templates = template.Must(template.ParseFiles("template/index.html"))
	templates.ExecuteTemplate(w, "index.html", data)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		display(w, nil)
	case "POST":
		_, _, err := r.FormFile("myFile")
		if err != nil {
			utils.Download(w, r)
		} else {
			utils.UploadFile(w, r)
		}
	}
}

//--------------- Main Function --------------------------------

func main() {
	http.HandleFunc("/", uploadHandler)
	http.ListenAndServe(":8080", nil)
}
