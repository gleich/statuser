package statuser

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/enescakir/emoji"
	"github.com/fatih/color"
)

var (
	// If the output should use emojis
	Emojis = true

	// Error emoji
	ErrorEmoji = emoji.PoliceCarLight
	// Error text
	ErrorText = " ERROR "
	// Error box character
	ErrorBoxChar = "░"
	// Error character if emojis are turned off
	ErrorChar = "✗"

	// Warning emoji
	WarningEmoji = emoji.Warning
	// Warning text
	WarningText = " WARNING "
	// Warning character if emojis are turned off
	WarningChar = "◈"

	// Success emoji
	SuccessEmoji = emoji.CheckMarkButton
	// Warning text
	SuccessText = " "
	// Warning character if emojis are turned off
	SuccessChar = "✔"
)

func generateBlock(message, surroundingChar string) string {
	messageLen := utf8.RuneCountInString(message)
	var topAndBottom string
	var extension int
	if Emojis {
		extension = 4
	} else {
		extension = 2
	}
	for i := 0; i < messageLen+extension; i++ {
		topAndBottom = topAndBottom + surroundingChar
	}
	return fmt.Sprintf(
		"%v\n%v%v%v\n%v",
		topAndBottom,
		surroundingChar,
		message,
		surroundingChar,
		topAndBottom,
	)
}

func separateBySpaces(items []interface{}) string {
	stringItems := []string{}
	for _, item := range items {
		stringItems = append(stringItems, fmt.Sprint(item))
	}
	return strings.Join(stringItems, " ")
}

// Output an error to the user
func Error(message string, err error, exitCode int) {
	title := emoji.Sprint(ErrorChar, ErrorText, ErrorChar)
	if Emojis {
		title = emoji.Sprint(ErrorEmoji, ErrorText, ErrorEmoji)
	}
	color.Red(generateBlock(title, ErrorBoxChar))
	color.Red("\n" + message)
	color.Red("\nGOLANG ERROR (SHOW DEVELOPER):\n" + err.Error())
	os.Exit(exitCode)
}

// Output an error to the user
func ErrorMsg(message string, exitCode int) {
	title := emoji.Sprint(ErrorChar, ErrorText, ErrorChar)
	if Emojis {
		title = emoji.Sprint(ErrorEmoji, ErrorText, ErrorEmoji)
	}
	color.Red(generateBlock(title, ErrorBoxChar))
	color.Red("\n" + message)
	os.Exit(exitCode)
}

// Output a warning to the user
func Warning(msgItems ...interface{}) {
	message := separateBySpaces(msgItems)
	title := emoji.Sprint(WarningChar, WarningText, WarningChar)
	if Emojis {
		title = emoji.Sprint(WarningEmoji, WarningText, WarningEmoji)
	}
	color.Yellow(title + "\n" + message)
}

// Output a success to the user
func Success(msgItems ...interface{}) {
	message := separateBySpaces(msgItems)
	prefix := emoji.Sprint(SuccessChar, SuccessText)
	if Emojis {
		prefix = emoji.Sprint(SuccessEmoji, SuccessText)
	}
	color.Green(prefix + message)
}
