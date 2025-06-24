package e2e

import (
	"qwant/internal/core"
	"testing"
)

func TestDirection_String(t *testing.T) {
	tests := []struct {
		direction core.Direction
		expected  string
	}{
		{core.North, "N"},
		{core.East, "E"},
		{core.South, "S"},
		{core.West, "W"},
	}

	for _, test := range tests {
		if result := test.direction.String(); result != test.expected {
			t.Errorf("Direction.String() = %s, expected %s", result, test.expected)
		}
	}
}

func TestParseDirection(t *testing.T) {
	tests := []struct {
		input     string
		expected  core.Direction
		shouldErr bool
	}{
		{"N", core.North, false},
		{"E", core.East, false},
		{"S", core.South, false},
		{"W", core.West, false},
		{"n", core.North, false}, // Case insensitive
		{"X", core.North, true},  // Invalid direction
	}

	for _, test := range tests {
		result, err := core.ParseDirection(test.input)
		if test.shouldErr {
			if err == nil {
				t.Errorf("ParseDirection(%s) expected error, got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ParseDirection(%s) unexpected error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("ParseDirection(%s) = %v, expected %v", test.input, result, test.expected)
			}
		}
	}
}

func TestLawn_IsValidPosition(t *testing.T) {
	lawn := core.InitLawn(5, 5)

	tests := []struct {
		pos      core.Position
		expected bool
	}{
		{core.Position{X: 0, Y: 0}, true},   // Bottom-left corner
		{core.Position{X: 5, Y: 5}, true},   // Top-right corner
		{core.Position{X: 3, Y: 3}, true},   // Middle
		{core.Position{X: -1, Y: 0}, false}, // Outside left
		{core.Position{X: 0, Y: -1}, false}, // Outside bottom
		{core.Position{X: 6, Y: 5}, false},  // Outside right
		{core.Position{X: 5, Y: 6}, false},  // Outside top
	}

	for _, test := range tests {
		result := lawn.IsValidPosition(test.pos)
		if result != test.expected {
			t.Errorf("IsValidPosition(%v) = %t, expected %t", test.pos, result, test.expected)
		}
	}
}
