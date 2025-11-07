package main

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "empty slice",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "no duplicates",
			input: []string{"a", "b", "c"},
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "with duplicates",
			input: []string{"a", "b", "a", "c", "b"},
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "all duplicates",
			input: []string{"x", "x", "x"},
			want:  []string{"x"},
		},
		{
			name:  "mixed case duplicates",
			input: []string{"Go", "go", "GO", "go"},
			want:  []string{"Go", "go", "GO"},
		},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			got := unique(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("input: %v, got: %v, want: %v", tc.input, got, tc.want)
			}
		})
	}
}
