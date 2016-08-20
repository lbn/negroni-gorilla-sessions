package main

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

type AuthMiddleware struct {
}

const sessionCookieName = "session-name"

var store = sessions.NewCookieStore([]byte("something-very-secret"))

// AuthMiddleware is a struct that has a ServeHTTP method
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

// The middleware handler
func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	log.Println("Checking session")
	session, _ := store.Get(req, "session-name")
	if session.IsNew {
		w.WriteHeader(401)
		return
	}
	next(w, req)
}
