package main

import "testing"

func TestFindLongestWord(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "empty",
			input: "",
			want:  "",
		},
		{
			name:  "single longest",
			input: "The old clock in the hallway chimed twelve times, but no one was there to hear it.",
			want:  "hallway",
		},
		{
			name:  "multiple longest",
			input: "The elephant wandered across the mountain with silent panthers nearby.",
			want:  "panthers",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := findLongestWord(tc.input)
			if got != tc.want {
				t.Fatalf("got: %s, want: %s", got, tc.want)
			}
		})
	}
}
