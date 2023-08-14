package arrayUtil

import (
	"reflect"
	"testing"
)

func TestDeDuplication(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal case",
			args: args{
				array: []int{1, 2, 3, 3, 3, 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Deduplicate(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deduplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		name  string
		left  int
		right int
		max   int
	}{
		{
			name:  "Left is max",
			left:  10,
			right: 9,
			max:   10,
		},
		{
			name:  "right is max",
			left:  1,
			right: 10,
			max:   10,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			returnedMax := MaxInt(test.left, test.right)
			if returnedMax != test.max {
				t.Errorf("Failed test %s\n\tleft: %v, right: %v, max: %v but received: %v",
					test.name, test.left, test.right, test.max, returnedMax)
			}
		})
	}
}

func TestMaxOfThree(t *testing.T) {
	testCases := []struct {
		name   string
		left   int
		middle int
		right  int
		max    int
	}{
		{
			name:   "right is max",
			left:   1,
			middle: 5,
			right:  10,
			max:    10,
		},
		{
			name:   "left is max",
			left:   10,
			middle: 5,
			right:  9,
			max:    10,
		},
		{
			name:   "left is max",
			left:   10,
			middle: 8,
			right:  6,
			max:    10,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actualMax := MaxInt(test.left, test.middle, test.right)
			if actualMax != test.max {
				t.Errorf("Failed test %s\n\tleft: %v, middle: %v, right: %v, max: %v but received: %v",
					test.name, test.left, test.middle, test.right, test.max, actualMax)
			}
		})
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name  string
		left  int
		right int
		min   int
	}{
		{
			name:  "left is min",
			left:  1,
			right: 10,
			min:   1,
		},
		{
			name:  "right is min",
			left:  10,
			right: 9,
			min:   9,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actualMin := MinInt(test.left, test.right)
			if actualMin != test.min {
				t.Errorf("Failed test %s\n\tleft: %v, right: %v, min: %v but received: %v",
					test.name, test.left, test.right, test.min, actualMin)
			}
		})
	}
}

func TestMinOfThree(t *testing.T) {
	testCases := []struct {
		name   string
		left   int
		middle int
		right  int
		min    int
	}{
		{
			name:   "left is min",
			left:   1,
			middle: 5,
			right:  10,
			min:    1,
		},
		{
			name:   "middle is min",
			left:   10,
			middle: 5,
			right:  9,
			min:    5,
		},
		{
			name:   "right is min",
			left:   10,
			middle: 8,
			right:  6,
			min:    6,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actualMin := MinInt(test.left, test.middle, test.right)
			if actualMin != test.min {
				t.Errorf("Failed test %s\n\tleft: %v, middle: %v, right: %v, min: %v but received: %v",
					test.name, test.left, test.middle, test.right, test.min, actualMin)
			}
		})
	}
}
