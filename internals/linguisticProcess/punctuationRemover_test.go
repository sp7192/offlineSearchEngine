package linguisticprocess

import "testing"

func TestPunctuationRemoverTest(t *testing.T) {
	converter := PunctuationRemover{}
	tests := map[string]struct {
		input    string
		expected string
	}{
		`empty`: {
			input:    "",
			expected: "",
		},
		`noChangeCase1`: {
			input:    "amir",
			expected: "amir",
		},
		`noChangeCase2`: {
			input:    "am(ir",
			expected: "am(ir",
		},
		`changeCase1`: {
			input:    "amir.",
			expected: "amir",
		},
		`changeCase2`: {
			input:    "amir!!",
			expected: "amir",
		},
		`changeCase3`: {
			input:    "{amir}",
			expected: "amir",
		},
		`changeCase4`: {
			input:    `"amir."`,
			expected: "amir",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := converter.Convert(test.input)
			if got != test.expected {
				t.Errorf("got : %s, expected : %s\n", got, test.expected)
			}
		})
	}
}
