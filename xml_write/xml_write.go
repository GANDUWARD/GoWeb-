package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Languages struct {
	XMLNam  xml.Name   `xml:"languages"`
	Version string     `xml:"version,attr"`
	Lang    []Language `xml:"language"`
}

type Language struct {
	Name string `xml:"name"`
	Site string `xml:site`
}

func main() {
	v := &Languages{Version: "2"}
	v.Lang = append(v.Lang, Language{"JAVA", "https://www.java.com/"})
	v.Lang = append(v.Lang, Language{"Go", "https://golang.org/"})
	output, err := xml.MarshalIndent(v, " ", " ")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	file, _ := os.Create("languages.xml")
	defer file.Close()
	file.Write([]byte(xml.Header))
	file.Write(output)
}
