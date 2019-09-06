package main

/**
 ____            _        _______
|  _ \          (_)      |__   __|
| |_) | __ _ ___ _  ___     | |_   _ _ __   ___  ___
|  _ < / _` / __| |/ __|    | | | | | '_ \ / _ \/ __|
| |_) | (_| \__ \ | (__     | | |_| | |_) |  __/\__ \
|____/ \__,_|___/_|\___|    |_|\__, | .__/ \___||___/
                                __/ | |
                               |___/|_|
**/
func main() {

	// boolean
	var isCool bool              // declare your variables
	isCool = true                // assign value
	println("Boolean: ", isCool) // must use variable - compilation will fail otherwise

	// string - "a string is in effect a read-only slice of bytes."
	// - They contain a list of arbitrary bytes
	// - it does not require unicode text or some other format
	var myName string
	myName = "Brian"
	println("String: ", myName)

	// integers
	// int    : int is a signed integer type that is at least 32 bits in size.
	//          Distinct type form int32
	// int8   : -128 to 127
	// int16  : -32768 to 32767
	// int32  : -2147483648 to 2147483647
	// int64  : -9223372036854775808 to 9223372036854775807
	var (
		// declare your variables in groups
		i   int
		i8  int8
		i16 int16
		i32 int32
		i64 int64
	)
	i = 1
	i8 = 127
	i16 = 32767
	i32 = 2147483647
	i64 = -9223372036854775808
	println("Integers: ", i, i8, i16, i32, i64)

	// uint8  : 0 to 255
	// uint16 : 0 to 65535
	// uint32 : 0 to 4294967295
	// uint64 : 0 to 18446744073709551615
	// uintptr: is an integer type that is large enough to hold the bit pattern of any pointer.
	var (
		iu   uint
		iu8  uint8
		iu16 uint16
		iu32 uint32
		iu64 uint64
	)
	iu = 1
	iu8 = 255
	iu16 = 65535
	iu32 = 4294967295
	iu64 = 18446744073709551615
	println("UIntegers: ", iu, iu8, iu16, iu32, iu64)

	// byte - alias for uint8 and is equivalent to uint8 in all ways.
	// Used to distinguish byte values from 8-bit unsigned integer values.
	var theByte = byte(255) // declare your variables by specifying type
	println("Byte: ", theByte)

	// rune - rune is an alias for int32 and is equivalent to int32 in all ways.
	// Used to distinguish character values from integer values.

	r := rune(61)
	println("Rune: ", string(r))

	// float32 - IEEE-754 32-bit floating-point numbers.
	f32 := float32(3.234888)
	// float64 -  IEEE-754 64-bit floating-point numbers
	f64 := float64(3.234888)
	println("Floats: ", f32, f64)

	// variables with initializers
	var a, b, c int = 1, 2, 3
	println("Init vars: ", a, b, c)

	// short declarations
	z, y, x := 4, 5, "hey!"
	println("Short declarations: ", z, y, x)
}

/*
Boolean:  true
String:  Brian
Integers:  1 127 32767 2147483647 -9223372036854775808
UIntegers:  1 255 65535 4294967295 18446744073709551615
Byte:  255
Rune:  =
Floats:  +3.234888e+000 +3.234888e+000
Init vars:  1 2 3
Short declarations:  4 5 hey!
*/
