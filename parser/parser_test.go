package parser

import (
	"testing"
)

func Test_parseURL(t *testing.T) {
	urls := []struct {
		raw    string
		parsed string
		err    bool
	}{
		{"google.com", "http://google.com", false},
		{"http://adjust.com", "http://adjust.com", false},
		{"abcde", "", true},
		{"reddit.com/r/funny", "http://reddit.com/r/funny", false},
	}

	for _, u := range urls {
		parsed, err := parseURL(u.raw)
		if (err != nil) != u.err {
			t.Errorf("parsing failed: expected %t, got %t", u.err, (err != nil))
		}

		if parsed != u.parsed {
			t.Errorf("parsing failed: expected %s, got %s", u.raw, parsed)
		}
	}
}
