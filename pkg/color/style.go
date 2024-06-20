package color

import (
	"strings"
)

// Style is a struct of terminal text style attributes.
type Style struct {
	foreground Color
	background Color
	bold       bool
}

// Apply applies text styling to a string.
//
// For implementation purposes, this is using ANSI SGR sequences.
// https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_(Select_Graphic_Rendition)_parameters
func (s Style) Apply(str string) string {
	var sb strings.Builder
	sb.Grow(32)

	// Attribute: Bold
	if s.bold {
		appendForNextParameter(&sb)
		sb.WriteRune('1')
	}

	// Foreground.
	if s.foreground.kind == colorBasic {
		appendSgrColorBasic(&sb, s.foreground.value, false)
	} else if s.foreground.kind == colorBright {
		// TODO: Replace with real function.
		appendSgrColorNone(&sb, 0, false)
	} else if s.foreground.kind == color256 {
		// TODO: Replace with real function.
		appendSgrColorNone(&sb, 1, false)
	} else if s.foreground.kind == colorTrue {
		// TODO: Replace with real function.
		appendSgrColorNone(&sb, 2, false)
	}

	// Background.
	if s.background.kind == colorBasic {
		appendSgrColorBasic(&sb, s.background.value, true)
	} else if s.background.kind == colorBright {
		// TODO: Replace with real function.
		appendSgrColorNone(&sb, 0, true)
	} else if s.background.kind == color256 {
		// TODO: Replace with real function.
		appendSgrColorNone(&sb, 1, true)
	} else if s.background.kind == colorTrue {
		// TODO: Replace with real function.
		appendSgrColorNone(&sb, 2, true)
	}

	// Return early if there's nothing to be applied.
	if sb.Len() == 0 {
		return str
	}

	// Append the 'm' to the escape sequence and get it as a string.
	sb.WriteRune('m')
	ansi := sb.String()

	// Enable colors inside other colors by replacing the reset color with the parent color.
	// This only works for a depth of 1.
	str = strings.ReplaceAll(str, ansiReset, ansiReset+ansi)

	// Return a string with color codes surrounding it.
	return ansi + str + ansiReset
}

// appendForNextParameter prepares a CSI escape sequence to accept another parameter.
func appendForNextParameter(builder *strings.Builder) {
	if builder.Len() == 0 {
		builder.WriteString("\x1B[")
	} else {
		builder.WriteString(";")
	}
}

// ansiReset is the ANSI SGR escape sequence for resetting all colors back to default.
const ansiReset = "\x1B[m"

var intLookupTable = [...]string{
	0: "0",
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
	9: "9",
}

// appendSgrColorNone appends a color of `colorNone` kind to the end of a string builder.
func appendSgrColorNone(builder *strings.Builder, value uint32, background bool) {
	// No-op
}

// appendSgrColorBasic appends a color of `colorBasic` kind to the end of a string builder.
func appendSgrColorBasic(builder *strings.Builder, value uint32, background bool) {
	appendForNextParameter(builder)

	if background {
		builder.WriteRune('4')
	} else {
		builder.WriteRune('3')
	}

	builder.WriteString(intLookupTable[value])
}
