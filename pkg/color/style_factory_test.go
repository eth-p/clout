package color

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFunctions(t *testing.T) {
	TEST_COLOR := Color{
		kind:  colorNone,
		value: 999,
	}

	tests := map[string]struct {
		expected Style
		got      Style
	}{
		"Foreground": {
			got: Plain().Foreground(TEST_COLOR),
			expected: Style{
				foreground: TEST_COLOR,
				background: None,
				bold:       false,
			},
		},
		"Background": {
			got: Plain().Background(TEST_COLOR),
			expected: Style{
				foreground: None,
				background: TEST_COLOR,
				bold:       false,
			},
		},
		"Bold": {
			got: Plain().Bold(true),
			expected: Style{
				foreground: None,
				background: None,
				bold:       true,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			diff := cmp.Diff(tc.expected, tc.got, cmp.AllowUnexported(Style{}, Color{}))
			if diff != "" {
				t.Log("did not find expected Style; want -> -, got -> +")
				t.Fatalf(diff)
			}
		})
	}
}
