package main

import "testing"

func TestBinChar(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{
			input: 0,
			want:  "0",
		}, {
			input: 1,
			want:  "1",
		}, {
			input: 4,
			want:  "100",
		}, {
			input: 5,
			want:  "101",
		},
	}
	for _, test := range tests {
		got := binChar(test.input)
		if got != test.want {
			t.Errorf("TestBinChar failed: input %v, got %v, want %v", test.input, got, test.want)
		}
	}
}
