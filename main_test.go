package main

import "testing"

func TestCalculatePinHeights(t *testing.T) {
	tests := []struct {
		name     string
		block    [][]rune
		expected []int
	}{
		{
			name: "First lock",
			block: [][]rune{
				{'#', '#', '#', '#', '#'},
				{'.', '#', '#', '#', '#'},
				{'.', '#', '#', '#', '#'},
				{'.', '#', '#', '#', '#'},
				{'.', '#', '.', '#', '.'},
				{'.', '#', '.', '.', '.'},
				{'.', '.', '.', '.', '.'},
			},
			expected: []int{0, 5, 3, 4, 3},
		},
		{
			name: "First key",
			block: [][]rune{
				{'.', '.', '.', '.', '.'},
				{'#', '.', '.', '.', '.'},
				{'#', '.', '.', '.', '.'},
				{'#', '.', '.', '.', '#'},
				{'#', '.', '#', '.', '#'},
				{'#', '.', '#', '#', '#'},
				{'#', '#', '#', '#', '#'},
			},
			expected: []int{5, 0, 2, 1, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := calculatePinHeights(test.block)
			if len(result) != len(test.expected) {
				t.Errorf("unexpected pin heights length, got %d, want %d", len(result), len(test.expected))
			}
			for i, v := range result {
				if v != test.expected[i] {
					t.Errorf("unexpected pin height at index %d, got %d, want %d", i, v, test.expected[i])
				}
			}
		})
	}
}

func TestFits(t *testing.T) {
	tests := []struct {
		name  string
		key   []int
		locks []int
		fits  bool
	}{
		{
			name:  "overlap",
			key:   []int{0, 5, 3, 4, 3},
			locks: []int{5, 0, 2, 1, 3},
			fits:  false,
		},
		{
			name:  "fits",
			key:   []int{0, 5, 3, 4, 3},
			locks: []int{3, 0, 2, 0, 1},
			fits:  true,
		},
		{
			name:  "fits",
			key:   []int{1, 2, 0, 5, 3},
			locks: []int{4, 3, 4, 0, 2},
			fits:  true,
		},
		{
			name:  "fits",
			key:   []int{1, 2, 0, 5, 3},
			locks: []int{3, 0, 2, 0, 1},
			fits:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := fits(test.key, test.locks)
			if result != test.fits {
				t.Errorf("unexpected result, got %t, want %t", result, test.fits)
			}
		})
	}
}
