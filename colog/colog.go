package colog

import (
	"fmt"
	"runtime"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var BackGroundRed = "\u001b[41m"
var BackGroundGreen = "\u001b[42m"
var BackGroundYellow = "\u001b[43m"
var BackGroundBlue = "\u001b[44m"
var BackGroundMagenta = "\u001b[45m"
var BackGroundCyan = "\u001b[46m"
var BackGroundWhite = "\u001b[47m"

var BackGroundBrightBlack = "\u001b[40;1m"
var BackGroundBrightRed = "\u001b[41;1m"
var BackGroundBrightGreen = "\u001b[42;1m"
var BackGroundBrightYellow = "\u001b[43;1m"
var BackGroundBrightBlue = "\u001b[44;1m"
var BackGroundBrightMagenta = "\u001b[45;1m"
var BackGroundBrightCyan = "\u001b[46;1m"
var BackGroundBrightWhite = "\u001b[47;1m"

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}

}

func DoGreen(str string) {
	fmt.Println(Green + str + Reset)
}

func DoYellow(str string) {
	fmt.Println(Yellow + str + Reset)
}

func DoBlue(str string) {
	fmt.Println(Blue + str + Reset)
}

func DoRed(str string) {
	fmt.Println(Red + str + Reset)
}

func DoPurple(str string) {
	fmt.Println(Purple + str + Reset)
}

func DoCyan(str string) {
	fmt.Println(Cyan + str + Reset)
}

func DoWhite(str string) {
	fmt.Println(White + str + Reset)
}

func DoGray(str string) {
	fmt.Println(Gray + str + Reset)
}

func DoBgYellow(str string) {
	fmt.Println(BackGroundYellow + str + Reset)
}

func DoBgRed(str string) {
	fmt.Println(BackGroundRed + str + Reset)
}

func DoBgGreen(str string) {
	fmt.Println(BackGroundGreen + str + Reset)
}

func DoBgBlue(str string) {
	fmt.Println(BackGroundBlue + str + Reset)
}

func DoBgMagenta(str string) {
	fmt.Println(BackGroundMagenta + str + Reset)
}

func DoBgCyan(str string) {
	fmt.Println(BackGroundCyan + str + Reset)
}

func DoBgWhite(str string) {
	fmt.Println(BackGroundWhite + str + Reset)
}

func DoBgBrightBlack(str string) {
	fmt.Println(BackGroundBrightBlack + str + Reset)
}

func DoBgBrightGreen(str string) {
	fmt.Println(BackGroundBrightGreen + str + Reset)
}

func DoBgBrightRed(str string) {
	fmt.Println(BackGroundBrightRed + str + Reset)
}

func DoBgBrightYellow(str string) {
	fmt.Println(BackGroundBrightYellow + str + Reset)
}

func DoBgBrightMagenta(str string) {
	fmt.Println(BackGroundBrightMagenta + str + Reset)
}

func DoBgBrightBlue(str string) {
	fmt.Println(BackGroundBrightBlue + str + Reset)
}

func DoBgBrightCyan(str string) {
	fmt.Println(BackGroundBrightCyan + str + Reset)
}

func DoBgBrightWhite(str string) {
	fmt.Println(BackGroundBrightWhite + str + Reset)
}

func MakeRed(str string) string {
	if str == "" {
		return ""
	}
	return fmt.Sprintf(Red + str + Reset)
}

func MakeGreen(str string) string {
	if str == "" {
		return ""
	}
	return fmt.Sprintf(Green + str + Reset)
}

func MakeBlue(str string) string {
	if str == "" {
		return ""
	}
	return fmt.Sprintf(Blue + str + Reset)
}
