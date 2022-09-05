package htcolog

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

//DoGreen prints a string with green color
func DoGreen(str string) {
	fmt.Println(Green + str + Reset)
}

//DoYellow prints a string with yellow color
func DoYellow(str string) {
	fmt.Println(Yellow + str + Reset)
}

//DoBlue prints a string with blue color
func DoBlue(str string) {
	fmt.Println(Blue + str + Reset)
}

//DoRed prints a string with red color
func DoRed(str string) {
	fmt.Println(Red + str + Reset)
}

//DoPurple prints a string with purple color
func DoPurple(str string) {
	fmt.Println(Purple + str + Reset)
}

//DoCyan prints a string with cyan color
func DoCyan(str string) {
	fmt.Println(Cyan + str + Reset)
}

//DoWhite prints a string with white color
func DoWhite(str string) {
	fmt.Println(White + str + Reset)
}

//DoGray prints a string with red color
func DoGray(str string) {
	fmt.Println(Gray + str + Reset)
}

//DoBgYellow prints a string with yellow background
func DoBgYellow(str string) {
	fmt.Println(BackGroundYellow + str + Reset)
}

//DoBgRed prints a string with red background
func DoBgRed(str string) {
	fmt.Println(BackGroundRed + str + Reset)
}

//DoBgGreen prints a string with green background
func DoBgGreen(str string) {
	fmt.Println(BackGroundGreen + str + Reset)
}

//DoBgBlue prints a string with blue background
func DoBgBlue(str string) {
	fmt.Println(BackGroundBlue + str + Reset)
}

//DoBgMagenta prints a string with magenta background
func DoBgMagenta(str string) {
	fmt.Println(BackGroundMagenta + str + Reset)
}

//DoBgCyan prints a string with cyan background
func DoBgCyan(str string) {
	fmt.Println(BackGroundCyan + str + Reset)
}

//DoBgWhite prints a string with white background
func DoBgWhite(str string) {
	fmt.Println(BackGroundWhite + str + Reset)
}

//DoBgBrightBlack prints a string with bright black background
func DoBgBrightBlack(str string) {
	fmt.Println(BackGroundBrightBlack + str + Reset)
}

//DoBgBrightGreen prints a string with bright green background
func DoBgBrightGreen(str string) {
	fmt.Println(BackGroundBrightGreen + str + Reset)
}

//DoBgBrightRed prints a string with bright red background
func DoBgBrightRed(str string) {
	fmt.Println(BackGroundBrightRed + str + Reset)
}

//DoBgBrightYellow prints a string with bright yellow background
func DoBgBrightYellow(str string) {
	fmt.Println(BackGroundBrightYellow + str + Reset)
}

//DoBgBrightMagenta prints a string with bright magenta background
func DoBgBrightMagenta(str string) {
	fmt.Println(BackGroundBrightMagenta + str + Reset)
}

//DoBgBrightBlue prints a string with bright blue background
func DoBgBrightBlue(str string) {
	fmt.Println(BackGroundBrightBlue + str + Reset)
}

//DoBgBrightCyan prints a string with bright cyan background
func DoBgBrightCyan(str string) {
	fmt.Println(BackGroundBrightCyan + str + Reset)
}

//DoBgBrightWhite prints a string with bright white background
func DoBgBrightWhite(str string) {
	fmt.Println(BackGroundBrightWhite + str + Reset)
}

//MakeRed makes a red string
func MakeRed(str string) string {
	if str == "" {
		return ""
	}
	return fmt.Sprintf(Red + str + Reset)
}

//MakeGreen makes a green string
func MakeGreen(str string) string {
	if str == "" {
		return ""
	}
	return fmt.Sprintf(Green + str + Reset)
}

//MakeBlue makes a blue string
func MakeBlue(str string) string {
	if str == "" {
		return ""
	}
	return fmt.Sprintf(Blue + str + Reset)
}
