package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(res)
	})

	http.ListenAndServe(":8080", nil)
}
