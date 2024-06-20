package highlight

import "testing"

func TestHyperlink(t *testing.T) {
	tests := map[string]struct {
		highlight Highlight
		expected  string
	}{
		"Emits OSC 8": {
			expected: "\x1B]8;;https://example.com\x1B\\Link text\x1B]8;;\x1B\\",
			highlight: hyperlinkHighlight{
				value: "Link text",
				href:  "https://example.com",
			},
		},

		"Emits ID": {
			expected: "\x1B]8;an_id;https://example.com\x1B\\Link text\x1B]8;;\x1B\\",
			highlight: hyperlinkHighlight{
				value: "Link text",
				href:  "https://example.com",
				id:    "an_id",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.highlight.Apply(tc.highlight.Value().(string))
			if tc.expected != got {
				t.Fatalf("expected: %s, got: %s", tc.expected, got)
			}
		})
	}
}
