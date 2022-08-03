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








