package color

// colorKind is an enum for differentating the type of color being held by a Color.
type colorKind int

const (
	colorNone   = iota
	colorBasic  = iota
	colorBright = iota
	color256    = iota
	colorTrue   = iota
)

// Color is an abstract representation of a terminal color code.
type Color struct {
	kind  colorKind
	value uint32
}

var (
	None = Color{
		kind: colorNone,
	}
	Red = Color{
		kind:  colorBasic,
		value: 1,
	}
	Green = Color{
		kind:  colorBasic,
		value: 2,
	}
	Yellow = Color{
		kind:  colorBasic,
		value: 3,
	}
	Blue = Color{
		kind:  colorBasic,
		value: 4,
	}
	Magenta = Color{
		kind:  colorBasic,
		value: 5,
	}
	Cyan = Color{
		kind:  colorBasic,
		value: 6,
	}
	White = Color{
		kind:  colorBasic,
		value: 9, // FIXME: Replace with actual white.
	}
)
