package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func save(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_name")
	session.Values["foo"] = "bar"
	session.Values["42"] = 43
	session.Save(r, w)
}

func get(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_name")
	fmt.Println(session)
}
func main() {
	http.HandleFunc("/save", save)
	http.HandleFunc("/get", get)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("err", err)
	} else {
		fmt.Print("OK")
	}
}
