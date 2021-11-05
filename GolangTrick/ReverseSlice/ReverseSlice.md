## How to reverse slice in golang?

References: [Reversing Tricks](https://github.com/golang/go/wiki/SliceTricks#reversing)

To replace the contents of a slice with the same elements but in reverse order:

```go
for i := len(a)/2-1; i >= 0; i-- {
	opp := len(a)-1-i
	a[i], a[opp] = a[opp], a[i]
}
```

The same thing, except with two indices:

```go
for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
	a[left], a[right] = a[right], a[left]
}
```

