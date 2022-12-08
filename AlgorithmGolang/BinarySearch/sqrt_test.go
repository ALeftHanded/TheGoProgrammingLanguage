package BinarySearch

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	randNum := rand.Int31()
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "random test",
			args: args{
				x: int(randNum),
			},
			want: MathSqrt(int(randNum)),
		},
		{
			name: "zero test",
			args: args{
				x: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, mySqrt(tt.args.x), "mySqrt(%v)", tt.args.x)
		})
	}
}
