package main

import (
	"net/http"
	"ngw"
)

func SayName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`manjia`))
}

func main() {
	ngw.R.Addr = "127.0.0.1:6002"
	ngw.R.HandleFunc("/", SayName)
	ngw.Start()
}
