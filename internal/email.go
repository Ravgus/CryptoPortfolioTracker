package internal

import (
	"github.com/go-mail/mail"
	"log"
	"os"
)

func SendEmail(body string) {
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USER_NAME"))
	m.SetHeader("To", os.Getenv("SMTP_USER_NAME"))
	m.SetHeader("Subject", "Portfolio tracker")
	m.SetBody("text/html", body)

	port := StringToInt(os.Getenv("SMTP_PORT"))

	dialer := mail.NewDialer(
		"smtp.gmail.com",
		port,
		os.Getenv("SMTP_USER_NAME"),
		os.Getenv("SMTP_PASSWORD"),
	)

	if err := dialer.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
}

func CreateEmailBody(percentageDifference float64, date string) string {
	response := "<html><body>"
	response += "<p>" + "Portfolio changed on " + FloatToString(percentageDifference) + "% from " + date + "</p>"
	response += "</body></html>"

	return response
}
