//https://golang.org/doc/articles/wiki/
package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func Port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func main() {
	http.HandleFunc("/", handler)

	port := Port()

	fmt.Printf("listening on %s ...", port)
	http.ListenAndServe(":" + port, nil)
}
