package sorter

import "testing"

func TestSort(t *testing.T) {
	tests := []struct {
		width, height, length, mass int
		expected                    string
	}{
		{100, 100, 10, 5, "STANDARD"},
		{200, 160, 150, 10, "SPECIAL"},
		{200, 160, 150, 25, "REJECTED"},
		{140, 140, 140, 21, "REJECTED"},
	}
	for _, test := range tests {
		result := Sort(test.width, test.height, test.length, test.mass)
		if result != test.expected {
			t.Errorf("Sort(%d, %d, %d, %d) = %s; expected %s", test.width, test.height, test.length, test.mass, result, test.expected)
		}
	}
}