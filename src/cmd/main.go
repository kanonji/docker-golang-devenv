package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("App started.")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
			fmt.Println("Hello world")
			_, err := w.Write([]byte("Hello world"))
			if err != nil {
				log.Fatal("main HandleFunc /", err)
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("http.ListenAndServe:", err)
	}
}
