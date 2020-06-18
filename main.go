package main

import (
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/fatih/color"
)

// Emojis ... If the output should use emojis
var Emojis = true

var blockChar = "░"

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
	return fmt.Sprintf("%v\n%v%v%v\n%v", topAndBottom, surroundingChar, message, surroundingChar, topAndBottom)
}

// Error ... Output an error to the user
func Error(message string, err error, exitCode int) {
	title := "ERROR"
	if Emojis {
		title = "🚨 ERROR 🚨"
	}
	color.Red(generateBlock(title, "░"))
	color.Red("\n" + message)
	color.Red("\nGOLANG ERROR (SHOW DEVELOPER):\n" + err.Error())
	os.Exit(exitCode)
}

// Warning ... Output a warning to the user
func Warning(message string) {
	title := "WARNING"
	if Emojis {
		title = "⚠️ WARNING ⚠️"
	}
	color.Yellow(title + "\n" + message)
}

// Success ... Output a success to the user
func Success(message string) {
	title := "SUCCESS"
	if Emojis {
		title = "✅ SUCCESS ✅"
	}
	color.Green(title + "\n" + message)
}
