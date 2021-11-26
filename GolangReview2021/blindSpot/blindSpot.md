## Blind Spot

Record the bug because of ignorance.

Record the unfamiliar basic usage of golang.

### Unfamiliar basic usage of golang

Some unfamiliar basic usage of golang.

#### 1. Capitals in struct fields

*Struct内字段大小写区别*

`Only fields starting with a capital letter are exported, or in other words visible outside the curent package`

> References:
>
> [Capitals in struct fields](https://stackoverflow.com/questions/24837432/capitals-in-struct-fields)
>
> [Exported identifiers](https://go.dev/ref/spec#Exported_identifiers)

#### 2. You cannot define constant struct in golang

*在golang中不支持常量结构体的定义*

> References:
>
> There are boolean constants, rune constants, integer constants, floating-point constants, complex constants, and string constants. Rune, integer, floating-point, and complex constants are collectively called numeric constants.
> 
> A constant value is represented by a rune, integer, floating-point, imaginary, or string literal, an identifier denoting a constant, a constant expression, a conversion with a result that is a constant, or the result value of some built-in functions such as unsafe.Sizeof applied to any value, cap or len applied to some expressions, real and imag applied to a complex constant and complex applied to numeric constants. The boolean truth values are represented by the predeclared constants true and false. The predeclared identifier iota denotes an integer constant.
> 
> https://golang.org/ref/spec#Constants