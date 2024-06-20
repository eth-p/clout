package color

// Foreground creates a new Style with a foreground Color.
func Foreground(color Color) Style {
	return Plain().Foreground(color)
}

// Background creates a new Style with a background Color.
func Background(color Color) Style {
	return Plain().Background(color)
}

// Plain creates a new empty Style.
func Plain() Style {
	return Style{
		foreground: None,
		background: None,
		bold:       false,
	}
}
