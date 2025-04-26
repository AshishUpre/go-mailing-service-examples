package main

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv" // dotenv doesnt load automatically, need package
	"gopkg.in/gomail.v2"
	"html/template"
	"net/smtp"
	"os"
)

/*
USING NET/SMTP
*/
func sendSimpleMail(subject string, body string, to []string) {
	auth := smtp.PlainAuth("",
		os.Getenv("USERNAME_EMAIL"),
		os.Getenv("GOOGLE_APP_PASSWORD"),
		"smtp.gmail.com",
	)

	// this is how the format should be - "Subject: <subject>\n<body>"
	message := "Subject: " + subject + "\n" + body

	err := smtp.SendMail("smtp.gmail.com:587",
		auth,
		os.Getenv("FROM_EMAIL"),
		to,
		[]byte(message),
	)

	if err != nil {
		fmt.Println("got an error while sending mail ----- ", err)
		return
	}
	fmt.Println("email sent successfully")
}

/*
USING NET/SMTP
*/
func sendHTMLMail(subject string, htmlFilePath string, to []string) {
	// get Html template and send it -- note that we also pass the name variable here to the template
	var body bytes.Buffer
	t, err := template.ParseFiles(htmlFilePath)
	err = t.Execute(&body, struct{ Name string }{Name: "<| receiver name |>"})
	if err != nil {
		fmt.Println("got an error while parsing html ----- ", err)
		return
	}

	auth := smtp.PlainAuth("",
		os.Getenv("USERNAME_EMAIL"),
		os.Getenv("GOOGLE_APP_PASSWORD"),
		"smtp.gmail.com",
	)

	// similarly this is the format for html mails
	Headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	message := "Subject: " + subject + "\n" + Headers + "\n\n" + body.String()

	err = smtp.SendMail("smtp.gmail.com:587",
		auth,
		os.Getenv("FROM_EMAIL"),
		to,
		[]byte(message),
	)

	if err != nil {
		fmt.Println("got an error while sending mail ----- ", err)
		os.Exit(0)
	}
	fmt.Println("email sent successfully")
}

/*
USING GOMAIL
*/
func sendGoMail(templatePath string) {
	// get Html template and send it -- note that we also pass the name variable here to the template
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	err = t.Execute(&body, struct{ Name string }{Name: "<| receiver name |>"})
	if err != nil {
		fmt.Println("got an error while parsing html ----- ", err)
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("FROM_EMAIL"))
	m.SetHeader("To", os.Getenv("TO_EMAIL"), os.Getenv("FROM_EMAIL"))

	// this is for cc
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello, this mail is sent from a Go program")
	m.SetBody("text/html", body.String())
	m.Attach("./image.jpg")

	d := gomail.NewDialer("smtp.gmail.com",
		587,
		os.Getenv("USERNAME_EMAIL"),
		os.Getenv("GOOGLE_APP_PASSWORD"),
	)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("email sent successfully")
}

func main() {
	godotenv.Load()

	//sendSimpleMail("go mailing test",
	//	"hello, this mail was sent from Go program",
	//	[]string{os.Getenv("FROM_EMAIL")},
	//)

	//sendHTMLMail("go mailing test",
	//	"./MailTemplate.html",
	//	[]string{os.Getenv("FROM_EMAIL")},
	//)

	sendGoMail("./MailTemplate.html")
}
