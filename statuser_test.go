package statuser

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

// Test function for the separateBySpaces function
func TestSeparateBySpaces(t *testing.T) {
	tt := []struct {
		items    []interface{}
		expected string
	}{
		{
			items:    []interface{}{},
			expected: "",
		},
		{
			items:    []interface{}{"Hello"},
			expected: "Hello",
		},
		{
			items:    []interface{}{0, "Hello", true},
			expected: "0 Hello true",
		},
	}

	for _, test := range tt {
		assert.Equal(t, test.expected, separateBySpaces(test.items))
	}
}

func TestWarning(t *testing.T) {
	Emojis = true
	Warning("Test")
	Emojis = false
	Warning("Test")
}

func TestSuccess(t *testing.T) {
	Emojis = true
	Success("Test")
	Emojis = false
	Success("Test")
}
