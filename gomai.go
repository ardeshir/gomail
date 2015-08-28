package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "ardeshir.org@gmail.com", "XXXXXXXX", "smpt.gmail.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"asepahsalar@hennepintech.edu"}
	msg := []byte("To: asepahsalar@hennepintech.edu\r\n" +
		"Subject: Love from Gomail!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("smtp.gmail.com:465", auth, "ardeshir.org@gmail.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
