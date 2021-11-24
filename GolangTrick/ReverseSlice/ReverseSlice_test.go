package ReverseSlice

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
			[]int{1,2,3,4,5,6,7,8,9,10},
			},
			want: []int{10,9,8,7,6,5,4,3,2,1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
