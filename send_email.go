package main

import (
	"bytes"
	"fmt"
	"github.com/Traineau/gomail/email"
	"github.com/gobuffalo/packr/v2"
	"html/template"
	"net/smtp"
)

var box = packr.New("templateBox", "./templates")

func sendEmail(recipients []*email.Recipient) {
	// Sender data.
	from := "swatxmathis@gmail.com"
	password := "thpatiyapidu72"
	
	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	for _, recipient := range recipients {
		tpl := template.New("welcome.html") // Create a template.

		// Get the string representation of a file, or an error if it doesn't exist:
		html, err := box.FindString("welcome.html")
		if err != nil {
			fmt.Println(err)
			return
		}

		// Authentication.
		auth := smtp.PlainAuth("", from, password, smtpHost)

		t, err := tpl.Parse(html)
		if err != nil {
			fmt.Println(err)
			return
		}

		var body bytes.Buffer

		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body.Write([]byte(fmt.Sprintf("Objet: Mail de bienvenue \n%s\n\n", mimeHeaders)))

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
		err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Email Sent Successfully!")
}
