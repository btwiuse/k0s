package config

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	inputs := [][]string{
		[]string{"-hub", "example.com"},
		[]string{"-hub", "example.com:8080"},
		[]string{"-hub", "example.com:443"},
		[]string{"-hub", "http://example.com:443"},
		[]string{"-hub", "https://example.com:80"},
		[]string{"-hub", "https://example.com"},
		[]string{"-hub", ":80"},
		[]string{"-hub", ":443"},
	}
	expects := []string{
		"http://example.com:80",
		"http://example.com:8080",
		"http://example.com:443",
		"http://example.com:443",
		"https://example.com:80",
		"https://example.com:443",
		"http://127.0.0.1:80",
		"https://127.0.0.1:443",
	}
	for i, input := range inputs {
		var (
			c        = Parse(input)
			expected = expects[i]
			output   = fmt.Sprintf("%s://%s", c.Scheme(), c.Addr())
		)
		if output != expected {
			t.Errorf("f(%s) = %s, expected: %s\n", input, output, expected)
		}
	}
}
