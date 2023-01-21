package main

import "github.com/kkrypt0nn/logger.go"

func main() {
	logger := logger.NewLogger()
	logger.Println("${bold}${underline}Styling:")
	logger.Println("${bold}Foreground colors:")
	logger.Println("${fg:black}Black${reset}\t${fg:red}Red${reset}\t\t${fg:green}Green${reset}\t\t${fg:yellow}Yellow${reset}\t\t${fg:blue}Blue${reset}\t\t${fg:purple}Purple${reset}\t\t${fg:cyan}Cyan${reset}\t\t${fg:white}White")
	logger.Println("${fg:gray}Gray${reset}\t${fg:brightred}Bright red${reset}\t${fg:brightgreen}Bright green${reset}\t${fg:brightyellow}Bright yellow\t${fg:brightblue}Bright blue${reset}\t${fg:brightpurple}Bright purple${reset}\t${fg:brightcyan}Bright cyan${reset}\t${fg:brightwhite}Bright white")
	logger.Println("${bold}Background colors:")
	logger.Println("${bg:black}Black${reset}\t${bg:red}Red${reset}\t\t${bg:green}Green${reset}\t\t${bg:yellow}Yellow${reset}\t\t${bg:blue}Blue${reset}\t\t${bg:purple}Purple${reset}\t\t${bg:cyan}Cyan${reset}\t\t${bg:white}White")
	logger.Println("${bg:gray}Gray${reset}\t${bg:brightred}Bright red${reset}\t${bg:brightgreen}Bright green${reset}\t${bg:brightyellow}Bright yellow\t${bg:brightblue}Bright blue${reset}\t${bg:brightpurple}Bright purple${reset}\t${bg:brightcyan}Bright cyan${reset}\t${bg:brightwhite}Bright white")
	logger.Println("${bold}Special Effects:")
	logger.Println("${effect:bold}Bold${reset}\t${effect:dim}Dim${reset}\t\t${effect:underline}Underline${reset}\t${effect:blink}Blink${reset}\t\t${effect:inverse}Inverse${reset}\t\t${effect:strikethrough}Strikethrough${reset}")

	logger.Println("\n${bold}${underline}Variables:")
	logger.Println("${bold}Caller:")
	logger.Println("Function: ${caller:function}\tShort function: ${caller:shortfunction}\t\tFile: ${caller:file}")
	logger.Println("${bold}Logging Level:")
	logger.SetLoggingLevel(logging.FATAL)
	logger.Println("Level Color: ${level:color}Color${reset}\tLevel Name: ${level:name}\t\tLevel Short Name: ${level:shortname}")
	logger.SetLoggingLevel(logging.NONE)
	logger.Println("${bold}Date & Time Now:")
	logger.Println("Date: ${now:date}\tTime: ${now:time}\t\t\tDate & Time: ${now:datetime}")
	logger.Println("${bold}System:")
	logger.Println("Architecture: ${sys:architecture}\tHostname: ${sys:hostname}\tOperating System: ${sys:operating_system}\t\tUsername: ${username}")
	logger.Println("Group ID: ${sys:groupid}\t\tUser ID: ${sys:userid}")
}
