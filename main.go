package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		io.WriteString(w, "hola go")
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
