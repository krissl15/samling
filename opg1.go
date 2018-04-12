package main

import (
"net/http"
)

func main() {
	http.HandleFunc("/1", foo1)
	http.HandleFunc("/2", foo2)
	http.HandleFunc("/3", foo3)
	http.HandleFunc("/4", foo4)
	http.ListenAndServe(":8080", nil)
}

func foo1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nummer 1"))
}

func foo2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nummer2"))
}

func foo3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nummer3"))
}

func foo4(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nummer34"))
}

func foo5(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nummer5"))
}
