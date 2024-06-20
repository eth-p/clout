package highlight

import (
	"fmt"
)

// New creates a new Highlight that will output a hyperlink.
func Hyperlink(value interface{}, href string) Highlight {
	return hyperlinkHighlight{
		value: value,
		href:  href,
	}
}

// hyperlinkHighlight is an implementation of Highlight that uses the ANSI OSC 8 sequence to emit a hyperlink.
type hyperlinkHighlight struct {
	value interface{}
	href  string
	id    string
}

func (h hyperlinkHighlight) Value() interface{} {
	return h.value
}

func (c hyperlinkHighlight) Apply(str string) string {
	return fmt.Sprintf("\x1B]8;%s;%s\x1B\\%v\x1B]8;;\x1B\\", c.id, c.href, str)
}
