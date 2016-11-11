package main

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "cogito3373@gmail.com")
	m.SetHeader("To", "mike.prinsloo@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>World</b>!")

	d := gomail.NewPlainDialer("mail.localhost", 25, "testuser", "testpass")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
