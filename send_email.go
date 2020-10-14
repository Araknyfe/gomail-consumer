package main

import (
	"fmt"
	"github.com/Traineau/gomail/email"
	"net/smtp"
)

func sendEmail(recipients []*email.Recipient) {
	
	// Sender data.
	from := "swatxmathis@gmail.com"
	password := "<Email Password>"
	
	// Receiver email address.
	
	var to []string
	
	for _, recipient := range recipients {
		to = append(to, recipient.Email)
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println("Email Sent Successfully!")
}
