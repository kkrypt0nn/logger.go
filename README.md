# logger.go

A lightweight and easy to use Go logging library that includes logging functionalities with different levels and custom formatting.

## Customizing

### Prefix
The prefix, what comes before the message, can be change with the `SetPrefix` method on a `Logger` structure.

> Default prefix is `${datetime} ${level:color}${level:name}${reset}: `

### Placeholders
These are formatting placeholders that will be replaced in both the message and the prefix.

| Name                 | Value                                                      |
|----------------------|------------------------------------------------------------|
| `${date}`            | The current date formatted in `Jan 02, 2006`               |
| `${time}`            | The current time formatted in `15:04:05`                   |
| `${datetime}`        | The current datetime formatted in `Jan 02, 2006 15:04:05`  |
| `${hostname}`        | The hostname of the system                                 |
| `${username}`        | The username of the system                                 |
| `${function}`        | The name of the function and file where it has been logged |
| `${shortfunction}`   | The name of the function where it has been logged          |
| `${black}`           | The black foreground color                                 |
| `${red}`             | The red foreground color                                   |
| `${green}`           | The green foreground color                                 |
| `${yellow}`          | The yellow foreground color                                |
| `${blue}`            | The blue foreground color                                  |
| `${purple}`          | The purple foreground color                                |
| `${cyan}`            | The cyan foreground color                                  |
| `${white}`           | The white foreground color                                 |
| `${bg:black}`        | The black background color                                 |
| `${bg:red}`          | The red background color                                   |
| `${bg:green}`        | The green background color                                 |
| `${bg:yellow}`       | The yellow background color                                |
| `${bg:blue}`         | The blue background color                                  |
| `${bg:purple}`       | The purple background color                                |
| `${bg:cyan}`         | The cyan background color                                  |
| `${bg:white}`        | The white background color                                 |
| `${bg:darkgray}`     | The darkgray background color                              |
| `${bg:lightblue}`    | The lightblue background color                             |
| `${bold}`            | The bold effect                                            |
| `${dim}`             | The dim effect                                             |
| `${underline}`       | The underline effect                                       |
| `${blink}`           | The blink effect                                           |
| `${inverse}`         | The inverse effect                                         |
| `${strikethrough}`   | The strikethrough effect                                   |
| `${reset}`           | Resets the styles                                          |
| `${level:color}`     | The color representing the logging level                   |
| `${level:name}`      | The name representing the logging level                    |
| `${level:shortname}` | The short name representing the logging level              |

> **Note**: There are aliases listed in [`formatting.go`](formatting.go), documentation will be rewritten for these placeholders.

### Styling
You can choose whether you want to style your messages or not with the `SetStyling` method on a `Logger` structure. Styling includes foreground colors, background colors and special effects such as bold, and others - see the [terminal package](terminal).

> **Note**: The styling will **not** apply to the message if it is not supported by the terminal.
