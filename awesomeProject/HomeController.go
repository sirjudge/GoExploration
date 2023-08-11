package main

import (
	"log"
	"net/http"
)

func PrintString(input string) {
	log.Println(input)
}
func RealController(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		logError(err, r)
		return
	}
	key1 := r.Form.Get("key1")
	key2 := r.Form.Get("key2")
	log.Printf("key1:%s  key2:%s", key1, key2)
}
