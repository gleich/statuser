package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/fatih/color"
)

var emojis = true

var blockChar = "░"

func generateBlock(message, surroundingChar string) string {
	messageLen := utf8.RuneCountInString(message)
	var topAndBottom string
	for i := 0; i < messageLen+4; i++ {
		topAndBottom = topAndBottom + surroundingChar
	}
	return fmt.Sprintf("%v\n%v%v%v\n%v", topAndBottom, surroundingChar, message, surroundingChar, topAndBottom)
}

// Error ... Output an error to the user
func Error(message string, err error) {
	title := "ERROR"
	if emojis {
		title = "🚨 ERROR 🚨"
	}
	color.Red(generateBlock(title, "░"))
	color.Red("\n" + message)
	color.Red("\nGOLANG ERROR (SHOW DEVELOPER):\n" + err.Error())
}

// Warning ... Output a warning to the user
func Warning(message string) {
	title := "WARNING"
	if emojis {
		title = "⚠️ WARNING ⚠️"
	}
	color.Yellow(title + "\n" + message)
}
