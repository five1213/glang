package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	code := 200
	aa := os.Getenv("path")
	w.Header().Set("aa", aa)
	for name, values := range r.Header {
		for _, value := range values {
			w.Header().Set(name, value)
		}
	}

	w.WriteHeader(code)
	remote := r.RemoteAddr
	w.Write([]byte(strconv.Itoa(code)))

	fmt.Printf("remoter:%s---responseCode:%d\n", remote, code)
}

func main() {
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
