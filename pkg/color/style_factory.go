package color

// Foreground applies a foreground color to the Style.
func (s Style) Foreground(color Color) Style {
	s.foreground = color
	return s
}

// Background applies a background color to the Style.
func (s Style) Background(color Color) Style {
	s.background = color
	return s
}

// Bold applies a bold attribute to the Style.
func (s Style) Bold(bold bool) Style {
	s.bold = bold
	return s
}
