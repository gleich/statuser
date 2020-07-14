package statuser

import (
	"testing"
)

func TestGenerateBlock(t *testing.T) {
	Emojis = true
	emojiBlock := generateBlock("🚨 ERROR 🚨", "░")
	if emojiBlock != "░░░░░░░░░░░░░\n░🚨 ERROR 🚨░\n░░░░░░░░░░░░░" {
		t.Errorf(emojiBlock, "!= ░░░░░░░░░░░░░\n░🚨 ERROR 🚨░\n░░░░░░░░░░░░░")
	}

	Emojis = false
	block := generateBlock("ERROR", "░")
	if block != "░░░░░░░\n░ERROR░\n░░░░░░░" {
		t.Errorf(block, "!= ░░░░░░░\n░ERROR░\n░░░░░░░")
	}
}

func TestWarning(t *testing.T) {
	Emojis = true
	Warning("Test")
}
