package main

import (
	"bytes"
	"fmt"
	"github.com/Traineau/gomail/email"
	"net/smtp"
	"text/template"
)

func sendEmail(recipients []*email.Recipient) {

	// Sender data.
	from := "swatxmathis@gmail.com"
	password := ""

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, err := template.ParseFiles("templates/email.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	var body bytes.Buffer

	for _, recipient := range recipients {
		t.Execute(&body, struct {
			FirstName   string
			LastName    string
			URL			string
		}{
			FirstName:	recipient.FirstName,
			LastName:	recipient.LastName,
			URL:	 "www.google.com",
		})

		to := []string{
			recipient.Email,
		}

		// Sending email.
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Email Sent Successfully!")
}
