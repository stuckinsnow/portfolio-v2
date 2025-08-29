package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/", handleAPI)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fs := http.FileServer(http.Dir("../dist/"))

		if _, err := http.Dir("../dist/").Open(r.URL.Path); err != nil {
			http.ServeFile(w, r, "../dist/index.html")
			return
		}

		fs.ServeHTTP(w, r)
	})

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
