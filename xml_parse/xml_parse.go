package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type EmailConfig struct {
	//unexpected EOF fault of Unmarshal
	XMLName        xml.Name       `xml:"config"`
	SmtpServer     string         `xml:"smtpServer"`
	SmtpPort       int            `xml:"smtpPort"`
	Sender         string         `xml:"sender"`
	SenderPassword string         `xml:"senderPassword"`
	Receivers      EmailReceivers `xml:"receivers"`
}

type EmailReceivers struct {
	Flag string   `xml:"flag,attr"`
	User []string `xml:"user"`
}

func main() {
	file, err := os.Open("default.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := EmailConfig{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println()
	fmt.Println("SmtpServer is: ", v.SmtpServer)
	fmt.Println("SmtpPort is:", v.SmtpPort)
	fmt.Println("Sender is:", v.Sender)
	fmt.Println("SenderPassword is :", v.SenderPassword)
	for i, element := range v.Receivers.User {
		fmt.Println(i, element)
	}

}
