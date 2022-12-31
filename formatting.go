package logging

import (
	"github.com/kkrypt0nn/logger.go/terminal"
	"os"
	"os/user"
	"runtime"
	"strings"
	"time"
)

var (
	// Aliases are the aliases for the following placeholders
	Aliases = map[string]string{
		// Variables
		"${date}":          "${now:date}",
		"${time}":          "${now:time}",
		"${datetime}":      "${now:datetime}",
		"${hostname}":      "${sys:hostname}",
		"${username}":      "${sys:username}",
		"${userid}":        "${sys:userid}",
		"${groupid}":       "${sys:groupid}",
		"${function}":      "${caller:function}",
		"${shortfunction}": "${caller:shortfunction}",

		// Styles

		// Colors
		"${black}":  "${fg:black}",
		"${red}":    "${fg:red}",
		"${green}":  "${fg:green}",
		"${yellow}": "${fg:yellow}",
		"${blue}":   "${fg:blue}",
		"${purple}": "${fg:purple}",
		"${cyan}":   "${fg:cyan}",
		"${white}":  "${fg:white}",

		// Backgrounds
		"${bblack}":     "${bg:black}",
		"${bred}":       "${bg:red}",
		"${bgreen}":     "${bg:green}",
		"${byellow}":    "${bg:yellow}",
		"${bblue}":      "${bg:blue}",
		"${bpurple}":    "${bg:purple}",
		"${bcyan}":      "${bg:cyan}",
		"${bwhite}":     "${bg:white}",
		"${bdarkgray}":  "${bg:darkgray}",
		"${blightblue}": "${bg:lightblue}",

		// Effects
		"${bold}":          "${effect:bold}",
		"${dim}":           "${effect:dim}",
		"${underline}":     "${effect:underline}",
		"${blink}":         "${effect:blink}",
		"${inverse}":       "${effect:inverse}",
		"${strikethrough}": "${effect:strikethrough}",

		// Special style
		"${reset}": "${effect:reset}",
	}

	// Variables are the different variables to replace.
	Variables = map[string]func() string{
		// Date & time
		"${now:date}": func() string {
			// Format customization coming soon
			return time.Now().Format("Jan 02, 2006")
		},
		"${now:time}": func() string {
			// Format customization coming soon
			return time.Now().Format("15:04:05")
		},
		"${now:datetime}": func() string {
			// Format customization coming soon
			return time.Now().Format("Jan 02, 2006 15:04:05")
		},
		// System information
		"${sys:hostname}": func() string {
			hostname, err := os.Hostname()
			if err != nil {
				return ""
			}
			return hostname
		},
		"${sys:username}": func() string {
			user, err := user.Current()
			if err != nil {
				return ""
			}
			return user.Username
		},
		"${sys:userid}": func() string {
			user, err := user.Current()
			if err != nil {
				return ""
			}
			return user.Uid
		},
		"${sys:groupid}": func() string {
			user, err := user.Current()
			if err != nil {
				return ""
			}
			return user.Gid
		},
		// Caller
		"${caller:function}": func() string {
			pc, _, _, ok := runtime.Caller(4)
			details := runtime.FuncForPC(pc)
			if ok && details != nil {
				return details.Name()
			}
			return ""
		},
		"${caller:shortfunction}": func() string {
			pc, _, _, ok := runtime.Caller(4)
			details := runtime.FuncForPC(pc)
			split := strings.Split(details.Name(), ".")
			if ok && details != nil && len(split) >= 2 {
				return strings.Split(details.Name(), ".")[1]
			}
			return ""
		},
	}
	// Styles are the different types of styling.
	Styles = map[string]string{
		// Colors
		"${fg:black}":  terminal.BLACK,
		"${fg:red}":    terminal.RED,
		"${fg:green}":  terminal.GREEN,
		"${fg:yellow}": terminal.YELLOW,
		"${fg:blue}":   terminal.BLUE,
		"${fg:purple}": terminal.PURPLE,
		"${fg:cyan}":   terminal.CYAN,
		"${fg:white}":  terminal.WHITE,

		// Backgrounds
		"${bg:black}":     terminal.BACK_BLACK,
		"${bg:red}":       terminal.BACK_RED,
		"${bg:green}":     terminal.BACK_GREEN,
		"${bg:yellow}":    terminal.BACK_YELLOW,
		"${bg:blue}":      terminal.BACK_BLUE,
		"${bg:purple}":    terminal.BACK_PURPLE,
		"${bg:cyan}":      terminal.BACK_CYAN,
		"${bg:white}":     terminal.BACK_WHITE,
		"${bg:darkgray}":  terminal.BACK_DARKGRAY,
		"${bg:lightblue}": terminal.BACK_LIGHTBLUE,

		// Effects
		"${effect:bold}":          terminal.BOLD,
		"${effect:dim}":           terminal.DIM,
		"${effect:underline}":     terminal.UNDERLINE,
		"${effect:blink}":         terminal.BLINK,
		"${effect:inverse}":       terminal.INVERSE,
		"${effect:strikethrough}": terminal.STRIKETHROUGH,

		// Special style
		"${effect:reset}": terminal.RESET,
	}
)

// ReplaceAliases will replace the aliases with the original placeholder.
func ReplaceAliases(message string) string {
	for a, v := range Aliases {
		message = strings.Replace(message, a, v, -1)
	}
	return message
}

// AddStyling will add styling into the message.
func AddStyling(message string) string {
	for e, v := range Styles {
		message = strings.Replace(message, e, v, -1)
	}
	return message
}

// AddVariables will add the various variables into the message.
func AddVariables(message string) string {
	message = ReplaceAliases(message)
	for variable, v := range Variables {
		message = strings.Replace(message, variable, v(), -1)
	}
	return message
}

// AddLoggingLevel will add the logging level variables into the message.
func AddLoggingLevel(message string, level Level) string {
	message = strings.Replace(message, "${level:color}", GetLevelColor(level), -1)
	message = strings.Replace(message, "${level:name}", GetLevelName(level), -1)
	message = strings.Replace(message, "${level:shortname}", GetLevelShortName(level), -1)
	return message
}
