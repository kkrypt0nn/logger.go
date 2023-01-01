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
		"${date}":             "${now:date}",
		"${time}":             "${now:time}",
		"${datetime}":         "${now:datetime}",
		"${sys:arch}":         "${sys:architecture}",
		"${architecture}":     "${sys:architecture}",
		"${arch}":             "${sys:architecture}",
		"${sys:os}":           "${sys:operating_system}",
		"${operating_system}": "${sys:operating_system}",
		"${os}":               "${sys:operating_system}",
		"${hostname}":         "${sys:hostname}",
		"${username}":         "${sys:username}",
		"${groupid}":          "${sys:groupid}",
		"${userid}":           "${sys:userid}",
		"${function}":         "${caller:function}",
		"${shortfunction}":    "${caller:shortfunction}",

		// Styles
		"${black}":         "${fg:black}",
		"${red}":           "${fg:red}",
		"${green}":         "${fg:green}",
		"${yellow}":        "${fg:yellow}",
		"${blue}":          "${fg:blue}",
		"${purple}":        "${fg:purple}",
		"${cyan}":          "${fg:cyan}",
		"${white}":         "${fg:white}",
		"${gray}":          "${fg:gray}",
		"${brightred}":     "${fg:brightred}",
		"${brightgreen}":   "${fg:brightgreen}",
		"${brightyellow}":  "${fg:brightyellow}",
		"${brightblue}":    "${fg:brightblue}",
		"${brightpurple}":  "${fg:brightpurple}",
		"${brightcyan}":    "${fg:brightcyan}",
		"${brightwhite}":   "${fg:brightwhite}",
		"${bblack}":        "${bg:black}",
		"${bred}":          "${bg:red}",
		"${bgreen}":        "${bg:green}",
		"${byellow}":       "${bg:yellow}",
		"${bblue}":         "${bg:blue}",
		"${bpurple}":       "${bg:purple}",
		"${bcyan}":         "${bg:cyan}",
		"${bwhite}":        "${bg:white}",
		"${bgray}":         "${bg:gray}",
		"${bbrightred}":    "${bg:brightred}",
		"${bbrightgreen}":  "${bg:brightgreen}",
		"${bbrightyellow}": "${bg:brightyellow}",
		"${bbrightblue}":   "${bg:brightblue}",
		"${bbrightpurple}": "${bg:brightpurple}",
		"${bbrightcyan}":   "${bg:brightcyan}",
		"${bbrightwhite}":  "${bg:brightwhite}",
		"${bold}":          "${effect:bold}",
		"${dim}":           "${effect:dim}",
		"${underline}":     "${effect:underline}",
		"${blink}":         "${effect:blink}",
		"${inverse}":       "${effect:inverse}",
		"${strikethrough}": "${effect:strikethrough}",
		"${reset}":         "${effect:reset}",
	}

	// Variables are the different variables to replace.
	Variables = map[string]func() string{
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
				return split[len(split)-1]
			}
			return ""
		},
		// Logging Level
		"${level:color}": func() string {
			return GetLevelColor(CurrentLoggingLevel)
		},
		"${level:name}": func() string {
			return GetLevelName(CurrentLoggingLevel)
		},
		"${level:shortname}": func() string {
			return GetLevelShortName(CurrentLoggingLevel)
		},
		// Date & Time Now
		"${now:date}": func() string {
			// Format customization coming soon
			return time.Now().Format(CurrentDateFormat)
		},
		"${now:time}": func() string {
			// Format customization coming soon
			return time.Now().Format(CurrentTimeFormat)
		},
		"${now:datetime}": func() string {
			// Format customization coming soon
			return time.Now().Format(CurrentDatetimeFormat)
		},
		// System
		"${sys:architecture}": func() string {
			return runtime.GOARCH
		},
		"${sys:hostname}": func() string {
			hostname, err := os.Hostname()
			if err != nil {
				return ""
			}
			return hostname
		},
		"${sys:operating_system}": func() string {
			return runtime.GOOS
		},
		"${sys:username}": func() string {
			user, err := user.Current()
			if err != nil {
				return ""
			}
			return user.Username
		},
		"${sys:groupid}": func() string {
			user, err := user.Current()
			if err != nil {
				return ""
			}
			return user.Gid
		},
		"${sys:userid}": func() string {
			user, err := user.Current()
			if err != nil {
				return ""
			}
			return user.Uid
		},
	}
	// Styles are the different types of styling.
	Styles = map[string]string{
		// Foreground Colors
		"${fg:black}":        terminal.BLACK,
		"${fg:red}":          terminal.RED,
		"${fg:green}":        terminal.GREEN,
		"${fg:yellow}":       terminal.YELLOW,
		"${fg:blue}":         terminal.BLUE,
		"${fg:purple}":       terminal.PURPLE,
		"${fg:cyan}":         terminal.CYAN,
		"${fg:white}":        terminal.WHITE,
		"${fg:gray}":         terminal.GRAY,
		"${fg:brightred}":    terminal.BRIGHT_RED,
		"${fg:brightgreen}":  terminal.BRIGHT_GREEN,
		"${fg:brightyellow}": terminal.BRIGHT_YELLOW,
		"${fg:brightblue}":   terminal.BRIGHT_BLUE,
		"${fg:brightpurple}": terminal.BRIGHT_PURPLE,
		"${fg:brightcyan}":   terminal.BRIGHT_CYAN,
		"${fg:brightwhite}":  terminal.BRIGHT_WHITE,
		// Background Colors
		"${bg:black}":        terminal.BG_BLACK,
		"${bg:red}":          terminal.BG_RED,
		"${bg:green}":        terminal.BG_GREEN,
		"${bg:yellow}":       terminal.BG_YELLOW,
		"${bg:blue}":         terminal.BG_BLUE,
		"${bg:purple}":       terminal.BG_PURPLE,
		"${bg:cyan}":         terminal.BG_CYAN,
		"${bg:white}":        terminal.BG_WHITE,
		"${bg:gray}":         terminal.BG_GRAY,
		"${bg:brightred}":    terminal.BG_BRIGHT_RED,
		"${bg:brightgreen}":  terminal.BG_BRIGHT_GREEN,
		"${bg:brightyellow}": terminal.BG_BRIGHT_YELLOW,
		"${bg:brightblue}":   terminal.BG_BRIGHT_BLUE,
		"${bg:brightpurple}": terminal.BG_BRIGHT_PURPLE,
		"${bg:brightcyan}":   terminal.BG_BRIGHT_CYAN,
		"${bg:brightwhite}":  terminal.BG_BRIGHT_WHITE,
		// Special Effects
		"${effect:bold}":          terminal.BOLD,
		"${effect:dim}":           terminal.DIM,
		"${effect:underline}":     terminal.UNDERLINE,
		"${effect:blink}":         terminal.BLINK,
		"${effect:inverse}":       terminal.INVERSE,
		"${effect:strikethrough}": terminal.STRIKETHROUGH,
		"${effect:reset}":         terminal.RESET,
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

// RemoveStyling will remove styling from the message.
func RemoveStyling(message string) string {
	for e, v := range Styles {
		message = strings.Replace(message, e, "", -1)
		message = strings.Replace(message, v, "", -1)
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
