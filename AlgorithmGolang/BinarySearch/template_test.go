Generated Test_condition
Generated TestBinarySearch
package BinarySearch

import "testing"

func Test_condition(t *testing.T) {
	type args struct {
		num int
		n   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := condition(tt.args.num, tt.args.n); got != tt.want {
				t.Errorf("condition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr []int
		n   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.arr, tt.args.n); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
input.Files: os.Stat: stat /Users/xiaohan.lu/TheGoProgrammingLanguage/AlgorithmGolang/BinarySearch/-: no such file or directory

