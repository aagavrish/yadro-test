package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestCheckBallsDistribution(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"case 1",
			`2
		1 2
		2 1
		`,
			"yes"},
		{"case 2",
			`3
		10 20 30
		1 1 1
		0 0 1
	`, "no"},
		{"case 3",
			`1
		1
		`,
			"yes"},
		{"case 4",
			`4
			1 1 1 1
			1 1 1 1
			1 1 1 1
			1 1 1 1
			`, "yes"},
		{"case 5",
			`2
			100000 100000
			100000 100000
			`, "yes"},
		{"case 6",
			`2
			1 2
			5 4
			`, "no"},
		{"case 7",
			`3
			1 0 0
			0 1 0
			0 0 1
			`, "yes"},
		{"case 8",
			`3
			1 2 3
			4 5 6
			7 8 9
			`, "no"},
		{"case 9",
			`3
			10 10 10
			10 0 10
			10 10 10
			`, "yes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			input := strings.NewReader(tt.input)
			in := bufio.NewReader(input)
			out := bufio.NewWriter(&buf)

			CheckBallsDistribution(in, out)
			out.Flush()

			if got := buf.String(); got != tt.expected {
				t.Errorf("CheckBallsDistribution() = %q, want %q", got, tt.expected)
			}
		})
	}
}
