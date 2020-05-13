package data

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStreamableVideoFormat(t *testing.T) {
	var (
		testCases = []struct {
			ext      string
			expected bool
		}{
			{".avi", true}, {".mkv", true}, {".mp4", true}, {".webm", true}, {".wmv", true},
			{".exe", false}, {".mp3", false}, {"", false},
		}
	)
	for _, testCase := range testCases {
		actual := IsStreamableVideoFormat(testCase.ext)
		assert.Equal(t, actual, testCase.expected,
			fmt.Sprintf("IsStreamableVideoFormat(%q) return %v, expect %v",
				testCase.ext, actual, testCase.expected))
	}
}
