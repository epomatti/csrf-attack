package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

const SESSION_COOKIE = "SESSION_COOKIE"

var tokens = make(map[string]http.Cookie)

func main() {
	http.HandleFunc("/", handleOk)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/withdraw", handleWithdraw)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Panic(err)
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	accessToken := uuid.NewString()[:8]

	cookie := http.Cookie{
		Name:     SESSION_COOKIE,
		Value:    accessToken,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	tokens[accessToken] = cookie
	http.SetCookie(w, &cookie)

	response := fmt.Sprintf("Authenticated with session token: %s\n", accessToken)
	fmt.Fprint(w, response)
}

func handleWithdraw(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(SESSION_COOKIE)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token := tokens[c.Value]
	if &token == nil {
		log.Print(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else {
		fmt.Fprint(w, "Authorized! Withdraw completed!\n")
	}
}

func handleOk(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK\n")
}
