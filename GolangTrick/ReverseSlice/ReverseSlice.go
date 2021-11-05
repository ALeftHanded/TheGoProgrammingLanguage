package ReverseSlice

func Reverse(arr []int) []int {
	for l,r := 0, len(arr) - 1; l<r; l,r = l+1, r-1{
		arr[l], arr[r] = arr[r], arr[l]
	}
	return arr
}
