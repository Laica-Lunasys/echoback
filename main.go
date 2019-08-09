package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type Page struct {
	Message string
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	msg := os.Getenv("ECHOBACK_MESSAGE")
	if msg == "" {
		msg = "<not defined>"
	}

	page := Page{
		Message: msg,
	}
	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handleIndex)

	fmt.Printf(":: Echoback listening at %s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}
