package linguisticprocess

import "testing"

func TestCheckStopWordTest(t *testing.T) {
	converter := CheckStopWord{}
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
		`changeCase1`: {
			input:    "a",
			expected: "",
		},
		`changeCase2`: {
			input:    "in",
			expected: "",
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
