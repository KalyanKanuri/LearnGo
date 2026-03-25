package extensions

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type EmailBox struct {
	Subject     string
	FromAddress string
	Body        string
	IsUnread    bool
	UnreadItems int
	EmailItems  []string
}

/*
backticks enables multi line strings in Go,
in templates we can use conditions like if else if, range as well
the syntax is to use in {{}} brackets as mentioned while using conditions or loops
we have to end them explicitly like {{end}} to access any item from current object
which we passed to template we can use {{.Item}}
*/
func TextTemplatesInGo() {
	fmt.Println("-- Email template Processor --")
	mailBody := `
	Subject: {{.Subject}}

	Hi Kalyan Kanuri,

	We are delighted to inform you that you have successfully cleared
	Google's third round of technical interviews.

	We are pleased to extend you an offer. Please find the attached
	details for your review.

	Best regards,
	{{.FromAddress}}

	------------------Gmail Footer-------------------
	Current Email Status: {{if .IsUnread}}Unread{{else}}Read{{end}}
	Inbox Items: {{range .EmailItems}}
		-{{.}}
	{{end}}
	Total Unread: {{.UnreadItems}}
	`
	demoMailBox := EmailBox{
		Subject:     "Congrats on clearing Google SDE III Go Engineer",
		FromAddress: "googlehiring@gmail.com",
		IsUnread:    true,
		Body:        mailBody,
		UnreadItems: 20,
		EmailItems: []string{
			"Reports",
			"Offer Letters",
			"Important Documents",
		},
	}

	tmpl, err := template.New("email-template").Parse(mailBody)

	var output strings.Builder
	if err != nil {
		fmt.Println("Error while parsing mail body", err)
		os.Exit(1)
	}
	tmpl.Execute(&output, demoMailBox)
	fmt.Printf("%s", output.String())
}
