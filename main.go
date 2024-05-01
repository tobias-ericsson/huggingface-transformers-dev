package main

import (
	"fmt"
	"net/http"
	"os"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("pong: %v\n", r.RemoteAddr)
	w.Write([]byte("pong"))
}

func main() {

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "80"
		fmt.Println("env PORT=XXX to change port")
	}

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/ping", ping)
	fmt.Printf("listening on port: %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
