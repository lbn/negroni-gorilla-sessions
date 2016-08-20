package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<form method="POST" action="/login"><input name="name" type="text" value="Bob" /><button type="submit">Login</button></form>`)
	}).Methods("GET")

	router.HandleFunc("/login", func(w http.ResponseWriter, req *http.Request) {
		session, _ := store.Get(req, sessionCookieName)
		if !session.IsNew {
			fmt.Fprintf(w, "Already logged in")
			return
		}
		session.Values["name"] = req.FormValue("name")
		session.Save(req, w)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<a href="/secure">/secure</a>`)
	}).Methods("POST")

	authRouter := mux.NewRouter().PathPrefix("/").Subrouter()
	authRouter.HandleFunc("/secure", func(w http.ResponseWriter, req *http.Request) {
		session, _ := store.Get(req, sessionCookieName)
		fmt.Fprintf(w, "Hello %s!", session.Values["name"])
	})

	router.PathPrefix("/").Handler(negroni.New(
		NewAuthMiddleware(),
		negroni.Wrap(authRouter),
	))

	log.Fatal(http.ListenAndServe(":8000", router))
}
