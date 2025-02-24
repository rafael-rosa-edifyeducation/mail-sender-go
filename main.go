package main

import (
	"log"
	"net/http"

	"github.com/RafaelCruzRosa/mail-sender-go/health"
	"github.com/RafaelCruzRosa/mail-sender-go/mails"
)

func main() {
	fs := http.FileServer(http.Dir("./assets"))

	http.Handle("/", fs)
	http.HandleFunc("/mails", mails.HandleSendMail)
	http.HandleFunc("/health", health.HandleHealth)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
