package BinarySearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowXN(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "normal case",
			args: args{
				x: 2.0,
				n: 17,
			},
			want: 131072.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PowXN(tt.args.x, tt.args.n), "PowXN(%v, %v)", tt.args.x, tt.args.n)
		})
	}
}
