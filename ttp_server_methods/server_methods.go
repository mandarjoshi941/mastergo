package main

import (
	"io"
	"log"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Friend!")
	io.WriteString(w, " URL: "+r.URL.Path)
}

func handleQuote(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Quotes Yet to implement! Please work with below joke, W@rning: PUN intended - \n")
	switch r.Method {
	case "GET":
		io.WriteString(w, "GET set go!")
	case "POST":
		io.WriteString(w, "POST the HOST!")
	case "PUT":
		io.WriteString(w, "This is writer's Block! You 'PUT' it on top of your desk and then you can't write there any more!\n ---------------- \n ---------------- \n ---------------- \n ---------------- \n ---------------- \n ---------------- \n ---------------- \n ---------------- \n ---------------- \n ---------------- \n")
	case "DELETE":
		io.WriteString(w, "DELETE Cookies?!")
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handleQuotes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, "Congrats! you are here early. You should GET somthing here in some days(like 365ish)")
	} else {
		io.WriteString(w, " I was supposed to send 404 or any other code, but couldn't. Hold on!")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	const PREFIX string = "/api/v1/"
	http.HandleFunc("/", handleHello)
	http.HandleFunc(PREFIX+"quote", handleQuote)
	http.HandleFunc(PREFIX+"quotes", handleQuotes)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalln("ListenAndServe Failed:", err)
	}
}
