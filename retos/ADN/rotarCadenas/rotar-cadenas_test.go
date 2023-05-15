package rotarcadenas

import "testing"

func TestRotarCadenasDerecha(t *testing.T) {
	type Test struct {
		input    string
		expected string
	}

	var tests []Test = []Test{
		{
			input:    "ABC",
			expected: "CAB",
		},
		{
			input:    "FGH",
			expected: "HFG",
		},
	}

	for _, test := range tests {
		generated := RotarDerecha(test.input)
		if generated != test.expected {
			t.Errorf("No son iguales: \ngenerated: %s, \nexpected: %s", generated, test.expected)

		}
	}
}
func TestRotarCadenasIzquierda(t *testing.T) {
	type Test struct {
		input    string
		expected string
	}

	var tests []Test = []Test{
		{
			input:    "ABC",
			expected: "BCA",
		},
		{
			input:    "FGH",
			expected: "GHF",
		},
	}

	for _, test := range tests {
		generated := RotarIzquierda(test.input)
		if generated != test.expected {
			t.Errorf("No son iguales: \ngenerated: %s, \nexpected: %s", generated, test.expected)

		}
	}
}
