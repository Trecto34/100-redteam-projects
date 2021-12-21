package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", reqLog(fs))
	http.ListenAndServe(":1337", reqLog(fs))
}

func reqLog(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[+] Request: %s %s", r.Method, r.URL)
		targetMux.ServeHTTP(w, r)
	})

}
