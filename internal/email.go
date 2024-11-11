package internal

import (
	"log"
	"os"

	"github.com/go-mail/mail"
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

func CreatePercentEmailBody(percentageDifference float64, date string) string {
	response := "<html><body>"
	response += "<p>" + "Portfolio changed on " + FloatToString(percentageDifference) + "% from " + date + "</p>"
	response += "</body></html>"

	return response
}

func CreatePriceEmailBody(currentPrice float64, trackedPrice float64) string {
	response := "<html><body>"
	response += "<p>" + "Portfolio reached " + FloatToString(trackedPrice) + "$. Current price is " + FloatToString(currentPrice) + "$</p>"
	response += "</body></html>"

	return response
}
