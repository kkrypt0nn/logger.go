package main

import "github.com/kkrypt0nn/logger.go"

func main() {
	logger := logger.NewLogger()
	logger.Println("${fg:red}${effect:blink}${effect:bold}${sys:username} says hello!")
}
