package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"realserver/envs"
	"strconv"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}
	COOKIE_PATH = envs.GetString("COOKIE_PATH")
	COOKIE_HTTP_ONLY, err = envs.GetBool("COOKIE_HTTP_ONLY")
	if err != nil {
		log.Panic(err)
	}
	COOKIE_SECURE, err = envs.GetBool("COOKIE_SECURE")
	if err != nil {
		log.Panic(err)
	}
	CSRF_PROTECTION_ENABLED, err = envs.GetBool("CSRF_PROTECTION_ENABLED")
	if err != nil {
		log.Panic(err)
	}
}

const SESSION_COOKIE = "SESSION_COOKIE"

var COOKIE_PATH string
var COOKIE_HTTP_ONLY bool
var COOKIE_SECURE bool
var CSRF_PROTECTION_ENABLED bool

var authTokens = make(map[string]http.Cookie)
var CSRFTokens = make(map[string]string)

func main() {
	http.HandleFunc("/", handleOk)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/withdraw", handleWithdraw)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Panic(err)
	}
}

type LoginResponse struct {
	CSRFToken   string `json:"CSRFToken"`
	AccessToken string `json:"accessToken"`
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	accessToken := uuid.NewString()[:8]

	cookie := http.Cookie{
		Name:     SESSION_COOKIE,
		Value:    accessToken,
		Path:     COOKIE_PATH,
		MaxAge:   3600,
		HttpOnly: COOKIE_HTTP_ONLY,
		Secure:   COOKIE_SECURE,
		SameSite: http.SameSiteLaxMode,
	}

	authTokens[accessToken] = cookie
	http.SetCookie(w, &cookie)

	var responseJson LoginResponse

	if CSRF_PROTECTION_ENABLED == true {
		responseJson.CSRFToken = strconv.Itoa(rand.Int())
	}

	responseJson.AccessToken = accessToken
	response, _ := json.Marshal(responseJson)

	fmt.Fprint(w, string(response))
}

func handleWithdraw(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(SESSION_COOKIE)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token := authTokens[c.Value]
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
