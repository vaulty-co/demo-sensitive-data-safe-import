package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

type user struct {
	Name        string
	Email       string
	PhoneNumber string
	IP          string
	CardNumber  string
	SSN         string
}

const EXT_API_URL = "http://api:3001/users"

func main() {
	db := &DB{}
	mux := http.NewServeMux()

	mux.HandleFunc("/import", func(w http.ResponseWriter, req *http.Request) {
		var users []user

		res, err := http.Get(EXT_API_URL)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Error: %s", err.Error())
		}

		dec := json.NewDecoder(res.Body)
		err = dec.Decode(&users)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Error: %s", err.Error())
		}
		db.AddList(users)

		fmt.Fprintf(w, "Pulled %d users from external API", len(users))
	})

	mux.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		type pageData struct {
			Users []user
		}
		tmpl, err := template.ParseFiles("./index.tpl")
		if err != nil {
			log.Fatal(err)
		}
		data := &pageData{
			Users: db.List(),
		}
		tmpl.Execute(w, data)
	})

	// serve static content
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	mux.Handle("/", http.FileServer(http.Dir(filesDir)))

	var port = flag.String("port", "3000", "web server port")
	flag.Parse()
	http.ListenAndServe(fmt.Sprintf(":%s", *port), mux)
}
