## Blind Spot

Record the bug because of ignorance.

Record the unfamiliar basic usage of golang.

### Unfamiliar basic usage of golang

Some unfamiliar basic usage of golang.

#### 1. Capitals in struct fields

`Struct内字段大小写区别`

`Only fields starting with a capital letter are exported, or in other words visible outside the curent package`

> References:
>
> [Capitals in struct fields](https://stackoverflow.com/questions/24837432/capitals-in-struct-fields)
>
> [Exported identifiers](https://go.dev/ref/spec#Exported_identifiers)

#### 2. You cannot define constant struct in golang

`在golang中不支持常量结构体的定义`

> References:
>
> There are boolean constants, rune constants, integer constants, floating-point constants, complex constants, and string constants. Rune, integer, floating-point, and complex constants are collectively called numeric constants.
> 
> A constant value is represented by a rune, integer, floating-point, imaginary, or string literal, an identifier denoting a constant, a constant expression, a conversion with a result that is a constant, or the result value of some built-in functions such as unsafe.Sizeof applied to any value, cap or len applied to some expressions, real and imag applied to a complex constant and complex applied to numeric constants. The boolean truth values are represented by the predeclared constants true and false. The predeclared identifier iota denotes an integer constant.
> 
> https://golang.org/ref/spec#Constants

#### 3. Convert interface{} to int

`将interface字段转换为int`

```go
iAreaId := val.(int)
```

> References:
>
> [Convert interface{} to int](https://stackoverflow.com/questions/18041334/convert-interface-to-int)
>
> [Golang Convert](http://golang.org/ref/spec#Conversions)
>
> Instead of
>
> ```golang
> iAreaId := int(val)
> ```
>
> you want a [type assertion](http://golang.org/ref/spec#Type_assertions):
>
> ```golang
> iAreaId := val.(int)
> iAreaId, ok := val.(int) // Alt. non panicking version 
> ```
>
> The reason why you cannot [convert](http://golang.org/ref/spec#Conversions) an interface typed value are these rules in the referenced specs parts:
>
> > Conversions are expressions of the form `T(x)` where `T` is a type and `x` is an expression that can be converted to type T.
>
> ...
>
> > A non-constant value x can be converted to type T in any of these cases:
> >
> > 1. x is assignable to T.
> > 2. x's type and T have identical underlying types.
> > 3. x's type and T are unnamed pointer types and their pointer base types have identical underlying types.
> > 4. x's type and T are both integer or floating point types.
> > 5. x's type and T are both complex types.
> > 6. x is an integer or a slice of bytes or runes and T is a string type.
> > 7. x is a string and T is a slice of bytes or runes.

#### 4. Assignment to the method receiver propagates only to callees but not to callers 

`在method receiver中修改caller本身不会生效`

> When you assign a value to the method receiver, the value will not be reflected outside of the method itself. Values will be reflected in subsequent calls from the same method.

#### 5. Golang Initialize Struct

```go
type Tmp struct{
	Val int
	Name string
} 

// t is nil
var t *Tmp
// t is not nil, t.Val = 0, t.Name = ""
t := &Tmp{}
```

#### 6. Timestamp之间比较的时候，如果采用time_a<time_b，需要注意两端相差不会超过一秒


#### 7. golang 对正则表达式的识别问题

```go
// 匹配整数或者一、二位小数
const RegexpGuidePurchasePrice = `^(\d+\.{1}\d{1,2})|(\d+)$`
// 此时match可以通过，因为正则表达式被解析为
// ^(\d+\.{1}\d{1,2}) 或者 (\d+)$
match, _ := regexp.MatchString(RegexpGuidePurchasePrice, "dd123.1212312")
// 此时可通过 match = true
// 加上括号后即可精准匹配
const RegexpGuidePurchasePrice = `^((\d+\.{1}\d{1,2})|(\d+))$`
```

#### 8. omitempty

> https://www.sohamkamani.com/golang/omitempty/

<span style="color:yellow">TODO</span>