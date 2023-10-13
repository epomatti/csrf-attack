package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleOk)
	http.HandleFunc("/cookies", handleCookies)

	err := http.ListenAndServe(":3666", nil)
	if err != nil {
		log.Panic(err)
	}
}

func handleCookies(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("SESSION_COOKIE")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("Session cookie captured: %s\n", c.Value)
	fmt.Fprint(w, response)
}

func handleOk(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK\n")
}
