# Placeholders

This is the list of placeholders and aliases that will get replaced in your logging messages and prefix.

## Placeholders

### Caller

| Name                      | Value                                                      |
|---------------------------|------------------------------------------------------------|
| `${caller:function}`      | The name of the function and file where it has been logged |
| `${caller:shortfunction}` | The name of the function where it has been logged          |
| `${caller:file}`          | The name of the file where it has been logged              |

### Logging Level

| Name                      | Value                                                      |
|---------------------------|------------------------------------------------------------|
| `${level:color}`          | The color representing the logging level                   |
| `${level:name}`           | The name representing the logging level                    |
| `${level:shortname}`      | The short name representing the logging level              |

### Date & Time Now

| Name                      | Value                                                           |
|---------------------------|-----------------------------------------------------------------|
| `${now:date}`             | The current date, default format is `Jan 02, 2006`              |
| `${now:time}`             | The current time, default format is `15:04:05`                  |
| `${now:datetime}`         | The current datetime, default format is `Jan 02, 2006 15:04:05` |

### System

| Name                      | Value                                                      |
|---------------------------|------------------------------------------------------------|
| `${sys:architecture}`     | The runtime architecture of the system                     |
| `${sys:hostname}`         | The hostname of the system                                 |
| `${sys:operating_system}` | The runtime operating system                               |
| `${sys:username}`         | The username of the system                                 |
| `${sys:groupid}`          | The group ID of the current user                           |
| `${sys:userid}`           | The user ID of the current user                            |

### Foreground Colors

| Name                 | Value                              |
|----------------------|------------------------------------|
| `${fg:black}`        | The black foreground color         |
| `${fg:red}`          | The red foreground color           |
| `${fg:green}`        | The green foreground color         |
| `${fg:yellow}`       | The yellow foreground color        |
| `${fg:blue}`         | The blue foreground color          |
| `${fg:purple}`       | The purple foreground color        |
| `${fg:cyan}`         | The cyan foreground color          |
| `${fg:white}`        | The white foreground color         |
| `${fg:gray}`         | The gray foreground color          |
| `${fg:brightred}`    | The bright red foreground color    |
| `${fg:brightgreen}`  | The bright green foreground color  |
| `${fg:brightyellow}` | The bright yellow foreground color |
| `${fg:brightblue}`   | The bright blue foreground color   |
| `${fg:brightpurple}` | The bright purple foreground color |
| `${fg:brightcyan}`   | The bright cyan foreground color   |
| `${fg:brightwhite}`  | The bright white foreground color  |

### Background Colors

| Name                 | Value                              |
|----------------------|------------------------------------|
| `${bg:black}`        | The black background color         |
| `${bg:red}`          | The red background color           |
| `${bg:green}`        | The green background color         |
| `${bg:yellow}`       | The yellow background color        |
| `${bg:blue}`         | The blue background color          |
| `${bg:purple}`       | The purple background color        |
| `${bg:cyan}`         | The cyan background color          |
| `${bg:white}`        | The white background color         |
| `${bg:gray}`         | The gray background color          |
| `${bg:brightred}`    | The bright red background color    |
| `${bg:brightgreen}`  | The bright green background color  |
| `${bg:brightyellow}` | The bright yellow background color |
| `${bg:brightblue}`   | The bright blue background color   |
| `${bg:brightpurple}` | The bright purple background color |
| `${bg:brightcyan}`   | The bright cyan background color   |
| `${bg:brightwhite}`  | The bright white background color  |

### Special Effects

| Name                      | Value                                                      |
|---------------------------|------------------------------------------------------------|
| `${effect:bold}`          | The bold effect                                            |
| `${effect:dim}`           | The dim effect                                             |
| `${effect:underline}`     | The underline effect                                       |
| `${effect:blink}`         | The blink effect                                           |
| `${effect:inverse}`       | The inverse effect                                         |
| `${effect:strikethrough}` | The strikethrough effect                                   |
| `${effect:reset}`         | Resets the styles                                          |

## Aliases

### Caller
| Name               | Alias of                  |
|--------------------|---------------------------|
| `${function}`      | `${caller:function}`      |
| `${shortfunction}` | `${caller:shortfunction}` |
| `${file}`          | `${caller:file}`          |

### Date & Time Now

| Name                  | Alias of                  |
|-----------------------|---------------------------|
| `${date}`             | `${now:date}`             |
| `${time}`             | `${now:time}`             |
| `${datetime}`         | `${now:datetime}`         |

### System

| Name                  | Alias of                  |
|-----------------------|---------------------------|
| `${sys:arch}`         | `${sys:architecture}`     |
| `${architecture}`     | `${sys:architecture}`     |
| `${arch}`             | `${sys:architecture}`     |
| `${sys:os}`           | `${sys:operating_system}` |
| `${operating_system}` | `${sys:operating_system}` |
| `${os}`               | `${sys:operating_system}` |
| `${hostname}`         | `${sys:hostname}`         |
| `${username}`         | `${sys:username}`         |
| `${groupid}`          | `${sys:groupid}`          |
| `${userid}`           | `${sys:userid}`           |

### Foreground Colors

| Name                  | Alias of                  |
|-----------------------|---------------------------|
| `${black}`            | `${fg:black}`             |
| `${red}`              | `${fg:red}`               |
| `${green}`            | `${fg:green}`             |
| `${yellow}`           | `${fg:yellow}`            |
| `${blue}`             | `${fg:blue}`              |
| `${purple}`           | `${fg:purple}`            |
| `${cyan}`             | `${fg:cyan}`              |
| `${white}`            | `${fg:white}`             |
| `${gray}`             | `${fg:gray}`              |
| `${brightred}`        | `${fg:brightred}`         |
| `${brightgreen}`      | `${fg:brightgreen}`       |
| `${brightyellow}`     | `${fg:brightyellow}`      |
| `${brightblue}`       | `${fg:brightblue}`        |
| `${brightpurple}`     | `${fg:brightpurple}`      |
| `${brightcyan}`       | `${fg:brightcyan}`        |
| `${brightwhite}`      | `${fg:brightwhite}`       |

### Background Colors

| Name                  | Alias of                  |
|-----------------------|---------------------------|
| `${bblack}`           | `${bg:black}`             |
| `${bred}`             | `${bg:red}`               |
| `${bgreen}`           | `${bg:green}`             |
| `${byellow}`          | `${bg:yellow}`            |
| `${bblue}`            | `${bg:blue}`              |
| `${bpurple}`          | `${bg:purple}`            |
| `${bcyan}`            | `${bg:cyan}`              |
| `${bwhite}`           | `${bg:white}`             |
| `${bgray}`            | `${bg:gray}`              |
| `${bbrightred}`       | `${bg:brightred}`         |
| `${bbrightgreen}`     | `${bg:brightgreen}`       |
| `${bbrightyellow}`    | `${bg:brightyellow}`      |
| `${bbrightblue}`      | `${bg:brightblue}`        |
| `${bbrightpurple}`    | `${bg:brightpurple}`      |
| `${bbrightcyan}`      | `${bg:brightcyan}`        |
| `${bbrightwhite}`     | `${bg:brightwhite}`       |

### Special Effects

| Name                  | Alias of                  |
|-----------------------|---------------------------|
| `${bold}`             | `${effect:bold}`          |
| `${dim}`              | `${effect:dim}`           |
| `${underline}`        | `${effect:underline}`     |
| `${blink}`            | `${effect:blink}`         |
| `${inverse}`          | `${effect:inverse}`       |
| `${strikethrough}`    | `${effect:strikethrough}` |
| `${reset}`            | `${effect:reset}`         |