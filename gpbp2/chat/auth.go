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

	// TODO: associate cookie with a session entry
	cookie, err := r.Cookie(AuthCookieName)

	if err == http.ErrNoCookie {
		w.Header().Add("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if cookie.Value == "12345" {
		h.next.ServeHTTP(w, r)
	} else {
		w.Header().Add("Location", "/doLogin")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
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
