package tools

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
)

func SendMail(email string) string {
	var Receivers []string

	Receivers = append(Receivers, email)

	auth := smtp.PlainAuth("", GSender, GPassword, GHost)

	password := RandomGenerator()

	t, err := template.ParseFiles("./client/mail.html")

	if err != nil {
		log.Fatalf("Error in parsing files: %v", err)
	}

	var bodyBuffer bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	bodyBuffer.Write([]byte(fmt.Sprintf("Test: \n%s\n\n", mimeHeaders)))

	t.Execute(&bodyBuffer, struct {
		Password string
	}{
		Password: password,
	})
	err = smtp.SendMail(GHost+":"+GPort, auth, GSender, Receivers, bodyBuffer.Bytes())

	if err != nil {
		log.Fatalf("Error in sending mail: %v", err)
	}

	return password

}
