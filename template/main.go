package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"os"
	"github.com/lazureykis/dotenv"
	"html/template"
)

type Manifest struct {
	JS  string `json:"/public/js/app.js"`
	CSS string `json:"/public/css/app.css"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	file, e := ioutil.ReadFile("./mix-manifest.json")
	if e != nil {
		log.Println("Manifest file not found!")
		return
	}

	var manifest Manifest
	json.Unmarshal(file, &manifest)

	if _, err := os.Stat("./hot"); !os.IsNotExist(err) {
		manifest.JS = "//localhost:8080" + manifest.JS
		manifest.CSS = "//localhost:8080" + manifest.CSS
	}

	index := template.Must(template.ParseFiles(
		"src/views/index.html",
	))
	index.Execute(w, manifest)
}

func main() {
	// Load .env file
	// How to use: os.Getenv("VAR")
	dotenv.Go()

	// Create a simple file server
	http.HandleFunc("/", indexHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	// Start the server on localhost port 8000 and log any errors
	port := os.Getenv("APP_PORT")
	log.Println("Listening on port " + port)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
