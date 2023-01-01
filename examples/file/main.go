package main

import (
	"github.com/kkrypt0nn/logger.go"
	"os"
)

func main() {
	file, err := os.OpenFile("example.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		createdFile, err := os.Create("example.log")
		if err != nil {
			panic(err)
		}
		file = createdFile
	}
	logger := logging.NewLogger()
	logger.SetLogFile(file)
	// The styling will get removed when writing into the file.
	logger.Debug("${fg:red}${effect:blink}${effect:bold}${sys:username} says hello!")
}
