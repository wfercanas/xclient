package main

import "net/http"

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("service up and running"))
}
