package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	specials := map[string]int{
		"recommendation": '❗',
		"search": '🔍',
		"weather": '☀',
		}
	for _, char := range log {
		for app, cp := range specials {
			if rune(cp) == char {
				return app
			}
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newlogline string
	for _, char := range log {
		if char == oldRune {
			newlogline += string(newRune)
		} else {
			newlogline += string(char)
		}
	}
	return newlogline
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
