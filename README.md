# Branchless

Branchless is a Go package that provides branchless operations on integers.

## Functions

### Min(x, y int) int
Returns the minimum of x and y.

### Max(x, y int) int
Returns the maximum of x and y.

### Min0(x int) int
Returns x if x > 0, otherwise 0.

### LessThan(x, y int) int
Returns 1 if x < y, otherwise 0.

### GreaterThan(x, y int) int
Returns 1 if x > y, otherwise 0.

### Equal(x, y int) int
Returns 1 if x == y, otherwise 0.

### NotEqual(x, y int) int
Returns 1 if x != y, otherwise 0.

### LessThanEqualTo(x, y int) int
Returns 1 if x <= y, otherwise 0.

### GreaterThanEqualTo(x, y int) int
Returns 1 if x >= y, otherwise 0.

### Abs(x int) int
Returns the absolute value of x.

### Diff(x, y int) int
Returns the absolute difference between x and y.

### Sign(x int) int
Returns -1 if x < 0, 0 if x == 0, and 1 if x > 0.

### IsPositive(x int) int
Returns 1 if x > 0, otherwise 0.

### IsNegative(x int) int
Returns 1 if x < 0, otherwise 0.

### IsZero(x int) int
Returns 1 if x == 0, otherwise 0.

### IsNotZero(x int) int
Returns 1 if x != 0, otherwise 0.

### Clamp(x, min, max int) int
Clamps the value x between the minimum (min) and maximum (max) values.

### IsPowerOfTwo(x int) int
Returns 1 if x is a power of 2 (or is zero), otherwise 0.

### IsEven(x int) int
Returns 1 if x is even, otherwise 0.

### IsOdd(x int) int
Returns 1 if x is odd, otherwise 0.

### Negate(x int) int
Returns the negation of x.

### IsDivisibleBy(x, y int) int
Returns 1 if x is divisible by y, otherwise 0.
