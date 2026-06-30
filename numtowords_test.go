package numtowords

import "testing"

func TestConvert(t *testing.T) {

	tests := []struct {
		name     string
		input    int
		expected string
		hasError bool
	}{
		{
			name:     "zero",
			input:    0,
			expected: "zero",
			hasError: false,
		},
		{
			name:     "single digit",
			input:    7,
			expected: "seven",
			hasError: false,
		},
		{
			name:     "teen number",
			input:    15,
			expected: "fifteen",
			hasError: false,
		},
		{
			name:     "exact tens",
			input:    40,
			expected: "forty",
			hasError: false,
		},
		{
			name:     "two digit number",
			input:    42,
			expected: "forty two",
			hasError: false,
		},
		{
			name:     "maximum value",
			input:    99,
			expected: "ninety nine",
			hasError: false,
		},
		{
			name:     "negative number",
			input:    -5,
			expected: "minus five",
			hasError: false,
		},
		{
			name:     "negative two digit",
			input:    -42,
			expected: "minus forty two",
			hasError: false,
		},
		{
			name:     "above range",
			input:    100,
			expected: "",
			hasError: true,
		},
		{
			name:     "below range",
			input:    -100,
			expected: "",
			hasError: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got, err := Convert(tt.input)

			if tt.hasError {

				if err == nil {
					t.Errorf(
						"expected error but got %s",
						got,
					)
				}

				return
			}

			if err != nil {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			if got != tt.expected {
				t.Errorf(
					"expected %q, got %q",
					tt.expected,
					got,
				)
			}

		})
	}
}
