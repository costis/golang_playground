package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type authHander struct {
	next http.Handler
}

func (h *authHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("auth")

	// not authenticated
	if err == http.ErrNoCookie {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}

	// other error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// success: user is authenticated
	h.next.ServeHTTP(w, r)
}

func MustAuth(nextHandler http.Handler) http.Handler {
	return &authHander{next: nextHandler}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]

	switch action {
	case "login":
		log.Println("TODO handle login for", provider)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action not supported %s", action)
	}
}
