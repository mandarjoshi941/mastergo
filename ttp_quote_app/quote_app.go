package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Quote ghune ja re marbat
type Quote struct {
	Author string `json:"author"`
	Text   string `json:"text"`
	Source string `json:"source,omitempty"`
}

// QuoteApp quotes ka khajana
type QuoteApp struct {
	Storage map[string]*Quote
}

func (qa *QuoteApp) handleHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Friend!")
	io.WriteString(w, " URL: "+r.URL.Path)
}

func (qa *QuoteApp) handleQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		var q Quote
		val := r.URL.Query()
		author := val["author"][0]
		if len(author) > 0 {
			db := getDatabase()
			if err := db.DB("lessons").C("quote").Find(bson.M{"author": author}).One(&q); err != nil {
				log.Fatalln(err)
			}
			db.Close()
			jsonVal, err := json.Marshal(q)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			io.WriteString(w, string(jsonVal))
		}
	case "POST", "PUT":
		var q Quote
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &q)
		db := getDatabase()
		if err := db.DB("lessons").C("quote").Insert(&q); err != nil {
			log.Fatalln(err)
		}
		db.Close()
		io.WriteString(w, string(body))
	case "DELETE":
		val := r.URL.Query()
		author := val["author"][0]
		if len(author) > 0 {
			jsonVal, err := json.Marshal(qa.Storage[author])
			delete(qa.Storage, author)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			io.WriteString(w, "Deleted:"+string(jsonVal))
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (qa *QuoteApp) handleQuotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var quotes []Quote
		db := getDatabase()
		if err := db.DB("lessons").C("quote").Find(nil).All(&quotes); err != nil {
			log.Fatalln(err)
		}
		db.Close()
		jsonVal, _ := json.Marshal(quotes)
		io.WriteString(w, string(jsonVal))
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {

	qa := QuoteApp{
		Storage: map[string]*Quote{},
	}

	const PREFIX string = "/api/v1/"
	http.HandleFunc("/", qa.handleHello)
	http.HandleFunc(PREFIX+"quote", qa.handleQuote)
	http.HandleFunc(PREFIX+"quotes", qa.handleQuotes)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalln("ListenAndServe Failed:", err)
	}
}

func getDatabase() *mgo.Session {
	db, err := mgo.Dial("localhost")
	if err != nil {
		return nil
	}
	return db
}
