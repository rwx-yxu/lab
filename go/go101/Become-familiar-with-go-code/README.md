# Notes
## Built in basic types
* Go supports 17 built in types
  * One boolean type: `bool`.
  * Eleven integer numeric types: `int8`, `uint8`, `uint16`, `int32`,
    `uint32`, `int`, `uint` and `uintptr`.
  * Two floating-point numeric types: `float32`, `float64`.
  * Two complex numeric types: `complex64`, `complex 128`
  * One built in string: `string`
* Each of these built in types can be used without importing any
  packages
* Go also support two built-in type aliases.
  * `byte` is an alias of `uint8`
  * `rune` is an alias of `int32`
* Integer types that start with u are unsigned types. This means that
  they are always non-negative. The number in the name of the type
  denotes the number of bits that will be occupied in memory during
  runtime. E.g `int8` largest value is 127 and smallest is -127. The
  first bit
  is used for the sign. Whereas the largest value for `uint8` is 255 and
  smallest is 0.
* The sizes of all values of a type are always the same so value sizes
  are often called type sizes where one byte contains eight bits
* The size of `uintptr`, `int`, and `uint` are implementation specific.
  Generally, the size of `int` and `uint` are 4 bytes on 32 bit
  architectures and 8 bytes on 64 bit architectures. The size of
  `uintptr` must be large enough to store the uninterpreted bits of any
  memory address.
* Each type has a default value.
  * The zero value of a boolean is false.
  * The zero value of a numeric type is 0.
  * The zero value of a string is an empty string.
## Basic literal values
* A literal is a text representation of the value in code. The
  literals denoting values of basic types are called basic value
  literals.
* Go does not define boolean literals but generally, we view the
  predeclared identifiers `false` and `true` as boolean
  literals.
* There are four integer literals
```
//Representation of 15 in decimal.
0xF // the hex form (starts with a "0x" or "0X")
0XF

017 // the octal form (starts with a "0", "0o" or "0O")
0o17
0O17

0b1111 // the binary form (starts with a "0b" or "0B")
0B1111

15 // the decimal form (starts without a "0")
```
* The binary and octal representation have been introduced since go 1.13
* Floating point value literals may contain a decimal integer part, a
  decimal fractional part and an integer exponent (10-based).
* The integer exponent part starts with a letter `e` or `E` along with a
  decimal integer suffix. `+ve` shift the decimal point to the right and
  `-ve` shift the decimal to the left.

```
1.23
01.23 // == 1.23
.23
1.
// An "e" or "E" starts the exponent part (10-based).
1.23e2 // == 123.0
123E2 // == 12300.0
123.E+2 // == 12300.0
1e-1 // == 0.1
.1e0 // == 0.1
0010e-2 // == 0.1
0e+5 // == 0.0
```

* Go 1.13 has also added support for hexadecimal floating point literal
  form.
  * This is denoted by the letters `p` or `P`. Followed by `n` where `n`
    is an integer that evaluated to `2^n`. This would form the
    expression of `yPn` where `y * 2^n`. Or `yP-n` where `y / 2^n`. `y`
    can an also be represented as `i.f` where `i` is an integer and `f`
    is `f/16`.
  * This means that `yPn = i.fPn = (i+f/16) * 2^n` as well as `yP-n =
    i.fP-n = (i+f/16) / 2^n`
  * The hexadecimal floating points must start with `0X` or `0x` like
    normal hexadecimal literals.

```
0x1p-2 // 1.0 / 2^2 = 1.0 / 4 = 0.25
0X2.p10 // 2.0 * 2^10 = 2.0 * 1024 = 2048.0
0x1.FP+0 // (1+15/16) * 2^0 = 1.9375 * 1 = 1.9375
0X.8P1 // (0+8/16) * 2^1 = 0.5 * 2 = 1
0x1FFFP-16 = 8191 / 2^16 = 8191 / 65536 = 0.1249847412109375

//Illegal declarations
0x.p1 // mantissa has no digits
1p-2 // p exponent requires hexadecimal mantissa
0x1.5e-2 // hexadecimal mantissa requires p exponent

```
* The zero value for floating point types are 0.0.
  * This can have different variations such as: `0., .0, 0e0, 0x0p0`
* A rune value is intended to store a Unicode code point. A Unicode code
  point is a Unicode character but some Unicode characters are composed
  of more than one code points each.
* A rune literal is expressed as one or more characters enclosed in a
  pair of quotes. The enclosed characters denote one Unicode point
  value.
* The most common rune literals is just to enclose the characters
  denoted by rune values between two single quotes.
* Rune integer literals start with a \.
  * Octal representation must be followed by three octal digits to
    represent a byte value where the first bit is the sign.
  * `\x` must be followed by exactly two hex digits.
  * `\U` must be followed by exactly eight hex digits.
  * `\u` must be followed by exactly four hex digits.
  * Each octal or hex sequence must represent a legal Unicode point or
    else, Go will fail to compile.

```
//Rune literals enclosed by single quotes
'a' // an English character
'π'
'众' // a Chinese character

//The unicode character of 'a' is 97.
// 141 is the octal representation of decimal number 97.
'\141'
// 61 is the hex representation of decimal number 97.
'\x61'
'\u0061'
'\U00000061'

fmt.Println('a')
\\Output 97
fmt.Println('ab')
\\more than one character in rune literal
```

* In practice, the four rune literal forms are rarely used for rune
  values. They are mainly used for interpreted string literals.
* If a run literal is composed by two characters (excluding the two
  quotes), the first character is \ and the second is not a digital
  digital character x, u and U, then the successive characters will be
  escaped as one special character.

```
//The possible character pairs to be escaped

\a (Unicode value 0x07) alert or bell
\b (Unicode value 0x08) backspace
\f (Unicode value 0x0C) form feed - advance downward to the next "page".
This is mainly used as section separators
\n (Unicode value 0x0A) line feed or newline
\r (Unicode value 0x0D) carriage return - Return to the beginning of the
current line without advancing downward.
\t (Unicode value 0x09) horizontal tab
\v (Unicode value 0x0b) vertical tab
\\ (Unicode value 0x5c) backslash
\' (Unicode value 0x27) single quote
```

* String value literals in Go are UTF-8 encoded.
* All Go source files must be UTF-8 encoding compatible.
* There are two forms of string value literals.
  * Interpreted - double quotes form.
  * Raw - back tick form.

```
// The interpreted form.
"Hello\nworld!\n\"你好世界\""

// The raw form.
`Hello
world!
"你好世界"`
```

* \" is only legal in interpreted string literals and \` is only legal
  in rune literals.
* Unicode string literals can also be used in interpreted string
  literals.

```
// The following interpreted string literals are equivalent.
"\141\142\143"
"\x61\x62\x63"
"\x61b\x63"
"abc"

// The following interpreted string literals are equivalent.
"\u4f17 = \xe4\xba\xba"
 // The Unicode of 众 is 4f17, which is
 // UTF-8 encoded as three bytes: e4 bc 97.
"\xe4\xbc\x97 = \u4eba"
 // The Unicode of 人 is 4eba, which is
  // UTF-8 encoded as three bytes: e4 ba ba.
"\xe4\xbc\x97\xe4\xba\xba"
"众人"
```

* English characters will be represented as one byte where as Chinese
  characters will be represented as three bytes.
* In a raw string literal, no character sequences will be escaped. The
  back quote character is not allowed to appear in a raw string literal.

## Untyped values and typed values
* An untyped value manes the type of the value has not been confirmed
  yet. Whereas a typed value is determined.
* For untyped values, each of them has one default type.
  * The predeclared `nil` is the only untyped value which has no default
    type.
* All literal constants are untyped values.
  * Most untyped values are literal constants and named constants.

## Explicit conversions of untyped constants.
* We can use the form T(v) to convert a value to the type denoted by T.
* for an untyped constant value v, there are two scenarios where T(v) is
  legal.
  1. v(or the literal denoted by v) is representable as a value of basic
     type T. The result is a typed constant of type T.
  1. The default type of v is an integer type (int or rune) and T is a
     string type. The result of T(v) is a string of type T and contains
     the UTF-8 representation of the integer as Unicode code point. The
     result string of a conversion from an integer always contains one
     and only one rune.

```
// Rounding happens in the following 3 lines.
complex128(1 + -1e-1000i) // 1.0+0.0i
float32(0.49999999) // 0.5
float32(17000000000000000)
// No rounding in the these lines.
float32(123)
uint(1.0)
int8(-123)
int16(6+0i)
complex128(789)
string(65) // "A"
string('A') // "A"
string('\u68ee') // "森"
string(-1) // "\uFFFD"
string(0xFFFD) // "\uFFFD"
string(0x2FFFFFFFF) // "\uFFFD"

// 1.23 is not representable as a value of int.
int(1.23)
// -1 is not representable as a value of uint8. uint should only be
positive integers
uint8(-1)
// 1+2i is not representable as a value of float64.
float64(1+2i)
// Constant -1e+1000 overflows float64.
float64(-1e1000)
// Constant 0x10000000000000000 overflows int.
int(0x10000000000000000)
// The default type of 65.0 is float64,
// which is not an integer type.
string(65.0)
// The default type of 66+0i is complex128,
// which is not an integer type.
string(66+0i)
```

* Sometimes, the form of explicit conversions must be written as (T)(v)
  to avoid ambiguities. This often happens if I is not an identifier.

## Introductions to type deductions
* Programmers do not need to explicitly specify the types of some values
  in Go.
* Go compilers will deduce the types for these values by context. This
  is also known as type inference.

## Constant declarations
* The keyword `const` is used to declare named constants.

```
package main

// Declare two individual constants. Yes,
// non-ASCII letters can be used in identifiers.
const π = 3.1416
const Pi = π // <=> const Pi = 3.1416

Declare multiple constants in a group.
const (
	No = !Yes
	Yes = true
	MaxDegrees = 360
	Unit = "radian"
)

func main() {
	// Declare multiple constants in one line.
	const TwoPi, HalfPi, Unit2 = π * 2, π * 0.5, "degree"
}
```

* The = symbol means "bind" instead of "assign".
* Constants can be declared both at the package level and in function
  bodies.
* The constants are untyped and the type are the same as the literals
  bound to it.

## Typed named constants
* Typed constants are all named.

```
const X float32 = 3.14

const (
 A, B int64 = -3, 5
 Y float32 = 2.718
)
```

* If multiple typed constants are declared, then their types must be the
  same.
* We can also use explicit conversion to provide enough information for
  Go compilers to deduce the types.
* If a basic value literal is bound to a typed constant, the basic value
  literal must be representable as a value of the type.

```
// error: 256 overflows uint8
const a uint8 = 256
// error: 256 overflows uint8
const b = uint8(255) + uint8(1)
// error: 128 overflows int8
const c = int8(-128) / int8(-1)
// error: -1 overflows uint
const MaxUint_a = uint(^0)
// error: -1 overflows uint. ^ is a bitwise-not operator
const MaxUint_b uint = ^0
```

* The following examples are valid on 64-bit OSes but invalid for 32-bit
  OSes because `unit` has only 64 buts on 64-bit OSes and same for
  32-bit OSes.

```
// bit-wise left operator doubles the value by 2^n in the form y=y *
2^n. For binary, this means to shift by one bit to the right. Since we
have a starting bit, we can shift by 64 which will form 1 and 64 zeros.
This will cause an overflow so we need to subtract 1 to get the max
value for unint. This would only be valid for 64-bit Oses
const Max unit = (1 << 64) -1

// We can also declare a uint and bind the largest value to it. The not
operator will flip the all bits to the opersite value. If we have a
value of 0 for uint, we can use bit-wise not to change all the bits to 1
by using the not operator. This operator can be used on the type or
variable only.
const MaxUint = ^uint(0)

// This means we can also declare a typed int constant and bind the largest
vlaue to it using the bitwise right operator. We can get the max size of
uint and shift by one bit to the right for int because the first bit
is the sign.
const MaxInt = int(^uint(0)>>1)

//By using the bitwise operators, we can get the number of bits of a
native word and check the current OS.

// NativeWordBits is 64 or 32.
const NativeWordBits = 32 << (^uint(0) >> 63)
const Is64bitOS = ^uint(0) >> 63 != 0
const Is32bitOS = ^uint(0) >> 32 == 0

```

## Autocomplete in constant declarations
* Group constant declarations except the first constant can be
  incomplete. An incomplete constant does not contain the = symbol. The
  remaining constants will inherit the first constant that is defined
  above.

```
const (
	A float32 = 5342
	B float32 = 218
	C
)
as
const (
	A float32 = 5342
	B float32 = 218
	C float32 = 218
)
```

## iota in constant declarations
* `iota` is a predeclared constant which can only be used in other
  constant declarations. It is declared as `const iota = 0`. This
  constant will increment by 1 for each constant declaration.
* `iota` is mostly useful in a `const` group whereby multiple constants
  are declared.
* `iota` can only be used in constant declarations
```
Package main

func main() {
 const (
 k = 3 // now, iota == 0

 m float32 = iota + .5 // m float32 = 1 + .5
 n // n float32 = 2 + .5

 p = 9 // now, iota == 3
 q = iota * 2 // q = 4 * 2
 _ // _ = 5 * 2
 r // r = 6 * 2
 s, t = iota, iota // s, t = 7, 7
 u, v // u, v = 8, 8
 _, w // _, w = 9, 9
 )

 const x = iota // x = 0
 const (
 y = iota // y = 0
 z // z = 1
 )

 println(m) // +1.500000e+000
 println(n) // +2.500000e+000
 println(q, r) // 8 12
 println(s, t, u, v, w) // 7 7 8 8 9
 println(x, y, z) // 0 0 1
}

const (
 Failed = iota - 1 // == -1
 Unknown // == 0
 Succeeded // == 1
)

const (
 Readable = 1 << iota // == 1
 Writable // == 2
 Executable // == 4
)
```

## Standard variable declaration forms
* Each standard declaration starts with the `var` keyword which is
  followed by the declared variable name. Variable names must be valid
  identifiers.

```
//When declaring multiple full standard variables at once, the type must be the same
for all the variables.
var lang, website string = "Go", "https://golang.org"
var compiled, dynamic bool = true, false
var announceYear int = 2009

// The types of the lang and dynamic variables
// will be deduced as built-in types "string"
// and "bool" by compilers, respectively.
var lang, dynamic = "Go", false

// The types of the compiled and announceYear
// variables will be deduced as built-in
// types "bool" and "int", respectively.
var compiled, announceYear = true, 2009

// The types of the website variable will be
// deduced as the built-in type "string".
var website = "https://golang.org"

// Both are initialized as blank strings.
var lang, website string
// Both are initialized as false.
var interpreted, dynamic bool
// n is initialized as 0.
var n int

var (
	lang, bornYear, compiled = "Go", 2007, true
	announceAt, releaseAt int = 2009, 2012
)
```

## Pure value assignments
* The expression items on the left of the = symbol are called
  destination or target values. The right side of the = symbol are
  called the source values.
* Constants are immutable (unchangeable) so constant cannot show up on
  the left side of a pure assignment as a destination value.
  * Constants can be used as the right side source value.
* Blank identifiers can also appear at the left side of pure value
  assignments. However, they cannot be used as source values for
  assignments.

```
const N = 123
var x int
var y, z float32

N = 9 // error: constant N is not modifiable
y = N // ok: N is deduced as a float32 value
x = y // error: type mismatch
x = N // ok: N is deduced as an int value
y = x // error: type mismatch
z = y // ok
_ = y // ok

x, y = y, x // error: type mismatch
x, y = int(y), float32(x) // ok
z, y = y, z // ok
 _, y = y, z // ok
//Go does not support assignment chains.
var a, b int
a = b = 123 // syntax error
```

## Short variable declarations forms
* We can also declare variables in a short form without the `var`
  keyword.
* These declarations can only be used for local variables. Not package
  level variables.

```
package main

 func main() {
  // Both lang and year are newly declared.
  lang, year := "Go language", 2007

  // Only createdBy is a new declared variable.
  // The year variable has already been
  // declared before, so here its value is just
  // modified, or we can say it is redeclared.
  year, createdBy := 2009, "Google Research"

  // This is a pure assignment.
  lang, year = "Go", 2012

  print(lang, " is created by ", createdBy)
  println(", and released at year", year)
 }
```

* In short variable declaration, there must be at least one new variable
  on the left of `:=`
* In a short variable declaration, all items t the left of the `:=` sign
  must be pure identifiers. This means that some other items cannot be
  the destination value.
  * Such as: container elements, pointer dereferences and struct field selectors. Pure assignments have no such limit.

## Each local declared variable must be used at lease once
* The go compiler do not allow local variables declared but not used.
  Package level variables have no such limit.
* If a local variable is only ever used as destination values, it will
  be viewed as unused.

## Dependency relations of package level variable initialization order
```
var x, y = a+1, 5 // 8 5
var a, b, c = b+1, c+1, y // 7 6 5
//Order of initalization are  y = 5, c = y, b = c+1, a = b+1, and
x = a+1.

//Package level variables cannot be used circularly.
var x, y = y, x // fail to compile
```

## Explicit conversions on non-Constant numeric values
* In Go, two typed values of two different basic types can't be assigned
  to each other. The types of the destination and source values in an
  assignment must be identical.
* If the type of the source and destination are not the same, then the
  source value must be explicitly converted to the type of the
  destination value.
* Non-constant integer values can be converted to strings.
  * Non-constant floating point and integer values can be explicitly
    converted to any other floating-point and integer types.
* Overflows are allowed in non-constant number conversions.
* When converting non-constant floating point values to an integer,
  rounding is allowed.

```
const a = -1.23
// The type of b is deduced as float64.
var b = a
// error: constant 1.23 truncated to integer.
var x = int32(a)
// error: cannot assign float64 to int32.
var y int32 = b
// okay: z == -1, and the type of z is int32.
// The fraction part of b is discarded.
var z = int32(b)
const k int16 = 255
// The type of n is deduced as int16.
var n = k
// error: constant 256 overflows uint8.
var f = uint8(k + 1)
// error: cannot assign int16 to uint8.
var g uint8 = n + 1
// okay: h == 0, and the type of h is uint8.
// n+1 overflows uint8 and is truncated.
var h = uint8(n + 1)
```

* Named constant declarations will be replaced with the literal it
  represents at compile time.
* If two operands of an operator operations are both constants, then the
  operation will be evaluated at compile time.

```
package main
const X = 3
const Y = X + X
var a = X

func main() {
  b := Y
  println(a, b, X, Y)
}

// will be viewed as at compile time. The constants have been evaluated.
package main

var a = X

func main() {
  b := 6
  println(a, b, 3, 6)
}
```

## Common Operators
* Like in C and C++, the multiplication binary operator can also be used
  as pointer dereference operator and the bitwise and operator `&` can
  also be used as pointer address.
* There is no power operator in Go. The `Pow` function in the math
  standard package will need to be used instead.
* The bit-wise clear operator `&^` is unique in Go. `m &6` is equivalent
  to `m &(n)`
* The rules for the result of a bitwise shift operator operation is
  always an integer value.
* Avoid cases that some bitwise shift operations return different result
  on different architectures on 64 bit and 32 bit OSes.

```
package main

const n = uint(8)
var m = uint(8)
func main() {
 println(a, b) // 2 0
}

var a byte = 1 << n / 128
var b byte = 1 << m / 128

//Above prints 2 and 0 because the two lines are equivalent to
var a = byte(int(1) << n / 128)
var b = byte(1) << m / 128

```

## Function Declarations
```
func SquaresOfSumAndDiff(a int64, b int64) (s int64, d int64) {
 x, y := a + b, a - b
 s = x * x
 d = y * y
 return // <=> return s, d
}
```
* Parameters and results are treated as local variables.
* The names in the result declaration list of a function can be present
  or absent all together. If a result is defined with a name,
  then the result is called a named result; otherwise it is called an
  anonymous result.
* When the results in a function declaration are anonymous; the return
  keyword must be followed by a sequence of return values.
```
func SquaresOfSumAndDiff(a int64, b int64) (int64, int64) {
 return (a+b) * (a+b), (a-b) * (a-b)
}
```
* If all the parameters are never used within the function body, the
  names in the parameter declaration list can be omitted.
* Local variables within a function body must be used.
* Go doesn't support default parameter values. The initial value of each
  result is the zero value of its type.
```
func f() (x int, y bool) {
 println(x, y) // 0 false
 return
}

func SquaresOfSumAndDiff(a, b int64) (s, d int64) {
 return (a+b) * (a+b), (a-b) * (a-b)
 // The above line is equivalent
 // to the following line.
 /*
 s = (a+b) * (a+b); d = (a-b) * (a-b); return
 */
}
```
* Functions must be directly declared at the package level.
* A function can't be declared within the body with another
  function.
* The type of an argument is not required to be identical with the
  corresponding parameter type. The only requirement for the
  argument is it must be assignable to the corresponding parameter
  type.
```
package main

func SquaresOfSumAndDiff(a int64, b int64) (int64, int64) {
 return (a+b) * (a+b), (a-b) * (a-b)
}

func CompareLower4bits(m, n uint32) (r bool) {
 r = m&0xF > n&0xF
 return
}

// Initialize a package-level variable
// with a function call.
var v = VersionString()

func main() {
 println(v) // v1.0
 x, y := SquaresOfSumAndDiff(3, 6)
 println(x, y) // 81 9
 b := CompareLower4bits(uint32(x), uint32(y))
 println(b) // false
 // "Go" is deduced as a string, 
 // and 1 is deduced as an int32.
 doNothing("Go", 1)
}

func VersionString() string {}
```

* Function calls can be deferred or invoked in new goroutines
  (green threads) in Go.

## Anonymous Functions
* Anonymous functions have no function name portion of the function
  declaration
```
package main

func main() {
	// This anonymous function has no parameters
	// but has two results.
	x, y := func() (int, int) {
		println("This function has no parameters.")
		return 3, 4
	}() // Call it. No arguments are needed.

	// The following anonymous function have no results.

	func(a, b int) {
		// The following line prints: a*a + b*b = 25
		println("a*a + b*b =", a*a+b*b)
	}(x, y) // pass argument x and y to parameter a and b.

	func(x int) {
		// The parameter x shadows the outer x.
		// The following line prints: x*x + y*y = 32
		println("x*x + y*y =", x*x+y*y)
	}(y) // pass argument y to parameter x.

	func() {
    // This function is in the scope of x and y variables declared
    above so it can use the twp variables directly.
		// The following line prints: x*x + y*y = 25
		println("x*x + y*y =", x*x+y*y)
	}() // no arguments are needed.
}
```

















