# logger.go

A lightweight and easy to use Go logging library that includes logging functionalities with different levels and custom formatting.

## Customizing

### Prefix
The prefix, what comes before the message, can be changed with the `SetPrefix` method on a `Logger` structure.

> The default prefix is `${datetime} ${level:color}${level:name}${reset}: `

### Placeholders
There are formatting placeholders that will be replaced in both the message and the prefix that can be seen [here](PLACEHOLDERS.md).

[For example](examples), logging the following message
```
${fg:red}${effect:blink}${effect:bold}${sys:username} says hello!
```
Will print a red blinking message in bold that says `<username> says hello!`, where `<username>` is the username on your system.

### Styling
You can choose whether you want to style your messages or not with the `SetStyling` method on a `Logger` structure. Styling includes foreground colors, background colors and special effects such as bold, and others - see the [terminal package](terminal).

> **Note**: The styling will **not** apply to the message if it is not supported by the terminal.

### Log File

Logs can also be written inside a log file with styling removed. See the [example here](examples/file/main.go).