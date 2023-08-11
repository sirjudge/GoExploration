package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
		PrintString("this is a test")
	}
	printPortInfo(port)
}

func printPortInfo(port string) {
	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	if r.URL.Path == "/" {
		_, err := fmt.Fprint(w, "home page")
		if err != nil {
			logError(err, r)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if r.URL.Path == "/test" {
		_, err := fmt.Fprint(w, "test page")
		if err != nil {
			logError(err, r)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if r.URL.Path == "/RealController" {
		RealController(w, r)
	} else {
		http.NotFound(w, r)
		return
	}
}

func logError(err error, r *http.Request) {
	log.Printf("error occurred when accessing path: %s Message: %s}", r.URL.Path, err)
}

func logRequest(r *http.Request) {
	log.Printf("URL Path: %s", r.URL.Path)
}
