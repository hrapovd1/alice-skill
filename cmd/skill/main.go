package main

import (
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	return http.ListenAndServe(`:8080`, http.HandlerFunc(webhook))
}

func webhook(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rw.Header().Set("Content-Type", "application/json")

	_, _ = rw.Write([]byte(`
	  {
	    "response": {
	      "text": "Извините, я пока ничего не умею"
	    },
	    "version": "1.0"
	  }
	`))
}
