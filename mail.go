package main

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func SendEachNews(news News) {
	content := strings.Trim(news.Url, "\"") + "\n" + strings.Trim(news.ImageUrl, "\"")
	fmt.Println("Send!")
	SendMail(strings.Trim(news.Title, "\""), content)
}
func SendMail(title string, content string) {
	from := goDotEnvVariable("SENDER")       //"hoangminhisme1@gmail.com" //
	password := goDotEnvVariable("PASSWORD") //"accclone55"           // //input
	toList := getRecipient()

	host := "smtp.gmail.com"
	port := "587"

	title = strings.ToUpper(title)

	msg := []byte(
		"From: uetgetnews@fedora\r\n" +
			"To: mewmew@gmail.com\r\n" +
			"Subject: UET News: " + title + "\r\n\r\n" +
			"Chi tiết thông báo: " + content + "\r\n")

	body := msg
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully sent mail to all user in toList")
}

func getRecipient() []string {
	return []string{
		"20021329@vnu.edu.vn",
		// "20021224@vnu.edu.vn",
	}
}
