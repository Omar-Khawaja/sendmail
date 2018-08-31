package main

import (
	"log"
	"net/smtp"
)

var host = "127.0.0.1"
var auth = smtp.PlainAuth("", "", "", host)
var from = "<sender e-mail address"
var msg = "This is a test. Thanks!"

var to = []string{
	"<recipient e-mail address",
}

func main() {
	err := smtp.SendMail(host+":25", auth, from, to, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}
