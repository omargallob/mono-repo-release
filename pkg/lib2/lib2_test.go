package lib2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFarewell(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "basic farewell",
			want: "Good bye from lib2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Farewell()
			assert.Equal(t, tt.want, got, "Greet() = %v, want %v", got, tt.want)
		})
	}
}
