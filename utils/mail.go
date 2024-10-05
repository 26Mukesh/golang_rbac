package utils

import (
	"os"

	"gopkg.in/gomail.v2"
)

func SendVerificationEmail(to string, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_FROM"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Verify your email")
	m.SetBody("text/html", "Please verify your email by clicking the link: <a href=\""+os.Getenv("BASE_URL")+""+token+"\">Verify Email</a>")

	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		587,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	return d.DialAndSend(m)
}
