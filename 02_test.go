package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "empty",
			input: "",
			want:  true,
		},
		{"one word", "a", true},
		{
			name:  "level Test",
			input: "level",
			want:  true,
		},
		{
			name:  "madam Test",
			input: "madam",
			want:  true,
		},
		{
			name:  "negative",
			input: "carl",
			want:  false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := isPalindrome(tc.input)
			if got != tc.want {
				t.Fatalf("got: %t, want: %t", got, tc.want)
			}
		})
	}
}
