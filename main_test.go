package main

import (
	"testing"
)

// func TestCard(t *testing.T) {
// 	card := newCard()
// 	if card != "Five of Diamonds" {
// 		t.Fatalf("error")
// 	}
// }

func Test_newCard(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Five of Diamonds",
			want: "Five of Diamonds",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newCard(); got != tt.want {
				t.Errorf("newCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
