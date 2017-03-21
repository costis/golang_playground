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

type cookieAuthHandler struct {
	next http.Handler
}

func (h *cookieAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	_, err := r.Cookie(AuthCookieName)

	if err == http.ErrNoCookie {
		w.Header().Add("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	h.next.ServeHTTP(w, r)
}

func MustAuthWithCookie(nextHandler http.Handler) http.Handler {
	return &cookieAuthHandler{next: nextHandler}
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
