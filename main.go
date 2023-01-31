package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"

	"github.com/k0kubun/pp"
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

func handleHeaders(w http.ResponseWriter, r *http.Request) {
    headers := r.Header

    fmt.Println("------------")
    pp.Println(headers)
    fmt.Println("------------")

    w.Header().Set("Content-Type", "application/json")

    res, err := json.Marshal(headers)
    if err != nil {
        panic(err)
    }

    _, err = io.WriteString(w, string(res))
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
    http.HandleFunc("/headers", handleHeaders)

	fmt.Printf(":: Echoback listening at %s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}
