## Sort in Golang

> References:
> 
> https://yourbasic.org/golang/how-to-sort-in-go/
> 
> https://pkg.go.dev/sort

### sort.Slice

sort.Slice

Keeping equal elements in their original order

```go
family := []struct {
    Name string
    Age  int
}{
    {"Alice", 23},
    {"David", 2},
    {"Eve", 2},
    {"Bob", 25},
}

// Sort by age, keeping original order or equal elements.
sort.SliceStable(family, func(i, j int) bool {
    return family[i].Age < family[j].Age
})
fmt.Println(family) // [{David 2} {Eve 2} {Alice 23} {Bob 25}]
```

if zero means max

```go
family := []struct {
    Name string
    Age  int
}{
    {"Alice", 23},
    {"David", 2},
    {"Eve", 2},
    {"Bob", 25},
	{"test", 0}
}

// Sort by age, keeping original order or equal elements.
sort.SliceStable(family, func(i, j int) bool {
	if family[i].Age == 0 {
		return false
    }
    return family[i].Age < family[j].Age
})
fmt.Println(family) // [{David 2} {Eve 2} {Alice 23} {Bob 25} {test 0}]
```