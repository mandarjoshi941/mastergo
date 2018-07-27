package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<H1><c>Hello</c></H1>")
	io.WriteString(w, fmt.Sprintf("From URL: %s", r.URL.Path))
}

func main() {
	http.HandleFunc("/", hello)

	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
