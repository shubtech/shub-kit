package color

import "fmt"

//Define some color in anscii
const (
	InfoColor    = "\033[1;34m"
	NoticeColor  = "\033[1;36m"
	WarningColor = "\033[1;33m"
	ErrorColor   = "\033[1;31m"
	DebugColor   = "\033[0;36m"
	NormalColor  = "\033[0m"
)

//Println print message with color
func Println(s ...interface{}) {
	messages := ""
	for _, mess := range s {
		messages = messages + mess.(string)
	}

	fmt.Println(messages)
}

//Sprintf buffer message
func Sprintf(s ...interface{}) string {
	messages := ""
	for _, mess := range s {
		messages = messages + mess.(string)
	}

	return messages
}
