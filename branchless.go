// Package branchless provides branchless operations on integers.
package branchless

// Min returns the minimum of x and y.
func Min(x, y int) int {
	return y ^ ((x ^ y) & ((x - y) >> 63))
}

// Max returns the maximum of x and y.
func Max(x, y int) int {
	return x ^ ((x ^ y) & ((x - y) >> 63))
}

// Min0 returns x if x > 0, otherwise 0.
func Min0(x int) int {
	return x & ^(x >> 63)
}

// LessThan returns 1 if x < y, otherwise 0.
func LessThan(x, y int) int {
	return ((x - y) >> 63) & 1
}

// GreaterThan returns 1 if x > y, otherwise 0.
func GreaterThan(x, y int) int {
	return ((y - x) >> 63) & 1
}

// Equal returns 1 if x == y, otherwise 0.
func Equal(x, y int) int {
	return (((x ^ y) | -(x ^ y)) >> 63 & 1) ^ 1
}

// NotEqual returns 1 if x != y, otherwise 0.
func NotEqual(x, y int) int {
	return (((x ^ y) | -(x ^ y)) >> 63 & 1)
}

// LessThanEqualTo returns 1 if x <= y, otherwise 0.
func LessThanEqualTo(x, y int) int {
	return ((y - x) >> 63 & 1) ^ 1
}

// GreaterThanEqualTo returns 1 if x >= y, otherwise 0.
func GreaterThanEqualTo(x, y int) int {
	return ((x - y) >> 63 & 1) ^ 1
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	return (x + (x >> 63)) ^ (x >> 63)
}

// Diff returns the absolute difference between x and y.
func Diff(x, y int) int {
	return ((x - y) ^ ((x - y) >> 63)) - ((x - y) >> 63)
}

// Sign returns -1 if x < 0, 0 if x == 0, and 1 if x > 0.
func Sign(x int) int {
	return (x >> 63) | (-(x >> 63))
}

// IsPositive returns 1 if x > 0, otherwise 0.
func IsPositive(x int) int {
	return ^((x >> 63) & 1)
}

// IsNegative returns 1 if x < 0, otherwise 0.
func IsNegative(x int) int {
	return (x >> 63) & 1
}

// IsZero returns 1 if x == 0, otherwise 0.
func IsZero(x int) int {
	return (((x | -x) >> 63) & 1) ^ 1
}

// IsZero returns 1 if x == 0, otherwise 0.
func IsNotZero(x int) int {
	return (((x | -x) >> 63) & 1)
}

// Clamp clamps the value x between the minimum (min) and maximum (max) values.
func Clamp(x, min, max int) int {
	return Min(Max(x, min), max)
}

// IsPowerOfTwo returns 1 if x is a power of 2 (or is zero), otherwise 0.
func IsPowerOfTwo(x int) int {
	return ^x & (x - 1) >> 63
}

// IsEven returns 1 if x is even, otherwise 0.
func IsEven(x int) int {
	return (x & 1) ^ 1
}

// IsOdd returns 1 if x is odd, otherwise 0.
func IsOdd(x int) int {
	return x & 1
}

// Negate returns the negation of x.
func Negate(x int) int {
	return -x
}

// IsDivisibleBy returns 1 if x is divisible by y, otherwise 0.
func IsDivisibleBy(x, y int) int {
	return (x % y & 1) ^ 1
}