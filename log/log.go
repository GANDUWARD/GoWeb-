package main

import (
	"log"
	"os"
)

func main() {
	fileName := "New.log"
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error")
	}
	debuglog := log.New(logFile, "[Info]", log.Llongfile)
	debuglog.Println("Info Level Message")
	debuglog.SetPrefix("[Debug]")
	debuglog.Println("Debug Level Message")
}
