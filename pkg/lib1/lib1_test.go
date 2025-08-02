package lib1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "basic greeting",
			want: "Hello from lib1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Greet()
			assert.Equal(t, tt.want, got, "Greet() = %v, want %v", got, tt.want)
		})
	}
}
