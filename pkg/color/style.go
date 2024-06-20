package color

import (
	"strconv"
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
	if s.foreground.kind != colorNone {
		colorKindJumpTable[s.foreground.kind](&sb, s.foreground.value, false)
	}

	// Background.
	if s.background.kind != colorNone {
		colorKindJumpTable[s.background.kind](&sb, s.background.value, true)
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
const ansiReset = "\x1B[0m"

// colorToAnsi is a jump table for executing the function responsible for appending
// the color to a SGR escape sequence.
var colorKindJumpTable = [...]func(*strings.Builder, uint32, bool){
	colorNone:   appendSgrColorNone,
	colorBasic:  appendSgrColorBasic,
	colorBright: appendSgrColorNone, // TODO:
	color256:    appendSgrColorNone, // TODO:
	colorTrue:   appendSgrColorNone, // TODO:
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

	builder.WriteString(strconv.Itoa(int(value)))
}
