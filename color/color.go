package color

import (
    "fmt"
    "strings"
)

const (
    Reset = "\x1b[0m"
    Bright = "\x1b[1m"
    Dim = "\x1b[2m"
    Underscore = "\x1b[4m"
    Blink = "\x1b[5m"
    Reverse = "\x1b[7m"
    Hidden = "\x1b[8m"

    FgBlack = "\x1b[30m"
    FgRed = "\x1b[31m"
    FgGreen = "\x1b[32m"
    FgYellow = "\x1b[33m"
    FgBlue = "\x1b[34m"
    FgMagenta = "\x1b[35m"
    FgCyan = "\x1b[36m"
    FgWhite = "\x1b[37m"

    BgBlack = "\x1b[40m"
    BgRed = "\x1b[41m"
    BgGreen = "\x1b[42m"
    BgYellow = "\x1b[43m"
    BgBlue = "\x1b[44m"
    BgMagenta = "\x1b[45m"
    BgCyan = "\x1b[46m"
    BgWhite = "\x1b[47m"
)

func Colorize(s string, color string) string {
    if len(s) > 2 && s[:2] == "\x1b[" {
        return s
    } else {
        return color + s + Reset
    }
}

func ColorizeAll(color string, args ...interface{}) string {
    var parts []string
    for _, arg := range args {
        parts = append(parts, Colorize(fmt.Sprintf("%v", arg), color))
    }
    return strings.Join(parts, "")
}

func Black(args ...interface{}) string {
    return ColorizeAll(FgBlack, args...)
}

func Red(args ...interface{}) string {
    return ColorizeAll(FgRed, args...)
}

func Green(args ...interface{}) string {
    return ColorizeAll(FgGreen, args...)
}

func Yellow(args ...interface{}) string {
    return ColorizeAll(FgYellow, args...)
}

func Blue(args ...interface{}) string {
    return ColorizeAll(FgBlue, args...)
}

func Magenta(args ...interface{}) string {
    return ColorizeAll(FgMagenta, args...)
}

func Cyan(args ...interface{}) string {
    return ColorizeAll(FgCyan, args...)
}

func White(args ...interface{}) string {
    return ColorizeAll(FgWhite, args...)
}


func BlackBG(args ...interface{}) string {
    return ColorizeAll(BgBlack, args...)
}

func RedBG(args ...interface{}) string {
    return ColorizeAll(BgRed, args...)
}

func GreenBG(args ...interface{}) string {
    return ColorizeAll(BgGreen, args...)
}

func YellowBG(args ...interface{}) string {
    return ColorizeAll(BgYellow, args...)
}

func BlueBG(args ...interface{}) string {
    return ColorizeAll(BgBlue, args...)
}

func MagentaBG(args ...interface{}) string {
    return ColorizeAll(BgMagenta, args...)
}

func CyanBG(args ...interface{}) string {
    return ColorizeAll(BgCyan, args...)
}

func WhiteBG(args ...interface{}) string {
    return ColorizeAll(BgWhite, args...)
}