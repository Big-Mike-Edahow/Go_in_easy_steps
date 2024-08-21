/* main.go */

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/lookup", lookupHandler)
	
	fmt.Println("Starting web server on port 8000")
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./form.html")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Response from the helloHandler!\n")
	io.WriteString(w, "URL.Path = "+r.URL.Path)
}

func lookupHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", 400)
		return
	}

	lang := r.FormValue("Language")
	lang = strings.ToLower(strings.Trim(lang, " "))

	file, err := os.OpenFile("./static/data/data.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print(lang)

	title := "<a href='http://www.ineasysteps.com'>See all titles online</a>"
	book := "logo"

	switch lang {
	case "c":
		title = "C Programming"
		book = "c"
	case "go":
		title = "Go Programming"
		book = "go"
	case "java":
		title = "Java"
		book = "java"
	case "python":
		title = "Python"
		book = "python"
	case "sql":
		title = "SQL"
		book = "sql"
	}

	response := "<!DOCTYPE HTML><title>Web Server Response</title>"
	response += "<link rel='shortcut icon' href='./static/images/favicon.ico' type='image/png'>"
	response += "<link rel='stylesheet' href='./static/css/style.css' type='text/css'>"
	response += "<p>" + title + "<br>in easy steps"
	response += "</p><img src='./static/images/" + book + ".png' alt='Cover'>"
	response += "<img src='./static/images/power.png' alt='Powered by Go'>"
	io.WriteString(w, response)
}
