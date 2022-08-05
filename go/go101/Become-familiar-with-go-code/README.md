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



