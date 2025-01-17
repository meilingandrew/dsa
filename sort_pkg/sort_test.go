package sort_pkg

import (
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		width    int
		height   int
		length   int
		mass     int
		expected Stack
	}{
		{
			name:     "rejected",
			width:    500,
			height:   500,
			length:   500,
			mass:     500,
			expected: Stack(Rejected),
		},
		{
			name:     "dimension greater to 150",
			width:    10,
			height:   10,
			length:   150,
			mass:     19,
			expected: Stack(Special),
		},
		{
			name:     "heavy mass greater then 20",
			width:    10,
			height:   10,
			length:   10,
			mass:     21,
			expected: Stack(Special),
		},
		{
			name:     "standard",
			width:    10,
			height:   10,
			length:   10,
			mass:     10,
			expected: Stack(Standard),
		},
		{
			name:     "invalid",
			width:    0,
			height:   10,
			length:   10,
			mass:     10,
			expected: Stack(Rejected),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := sort(tt.width, tt.height, tt.length, tt.mass)

			if stack != tt.expected.String() {
				t.Errorf("assert.Equal(t, got: %v, expected: %v)", stack, tt.expected.String())
			}
		})
	}
}
