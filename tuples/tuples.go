package tuples

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

// EPSILON constant used to check if floating point values are
// equivalent. If the abs value of the difference between two
// floating points is less than EPSILON, they will be considered
// equivalent.
const EPSILON = 0.0001

// Tuple struct defines the shape of a point or vector
type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

// Negate inverts a tuple's values by subtracting from a x0 y0 z0 tuple
func (t *Tuple) Negate() *Tuple {
	zero := CreateVector(0, 0, 0)
	negated := zero.Subtract(t)
	negated.W = t.W
	return negated
}

// Add sums the corresponding values of both tuples and returns a new tuple
// with the total values
func (t *Tuple) Add(b *Tuple) *Tuple {
	return CreatePoint(
		Calculate(t.X, b.X, Addition),
		Calculate(t.Y, b.Y, Addition),
		Calculate(t.Z, b.Z, Addition),
	)
}

// Subtract takes the corresponding values of both tuples and subtracts them
// returning a new tuple with the resulting values.
func (t *Tuple) Subtract(b *Tuple) *Tuple {
	return &Tuple{
		X: Calculate(t.X, b.X, Subtract),
		Y: Calculate(t.Y, b.Y, Subtract),
		Z: Calculate(t.Z, b.Z, Subtract),
		W: Calculate(t.W, b.W, Subtract),
	}
}

// MultiplyByScalar multiplies each component of a tuple by a given scalar
// returning a new tuple with the product of each component.
func (t *Tuple) MultiplyByScalar(n float64) *Tuple {
	tup := CreateTuple(
		Calculate(n, t.X, Multiply),
		Calculate(n, t.Y, Multiply),
		Calculate(n, t.Z, Multiply),
		t.W,
	)
	return tup
}

// DivideByScalar works just like (*Tuple).MultiplyByScalar but with a
// division operation instead. Returns a new tuple with the resulting values.
func (t *Tuple) DivideByScalar(n float64) *Tuple {
	tup := CreateTuple(
		Calculate(t.X, n, Divide),
		Calculate(t.Y, n, Divide),
		Calculate(t.Z, n, Divide),
		t.W,
	)
	return tup
}

// IsPoint checks if tuple is a static point.
func (t *Tuple) IsPoint() bool {
	return t.W == 1.0
}

// IsVector checks if tuple is a vector.
func (t *Tuple) IsVector() bool {
	return t.W == 0
}

// IsEquivalentTo checks to see if the floating point values of each component
// in the tuple are equalivalent to each other. If all components are equal,
// the overall tuples are considered equal as well.
func (t *Tuple) IsEquivalentTo(b *Tuple) bool {
	result := true

	x := Equals(t.X, b.X)
	y := Equals(t.Y, b.Y)
	z := Equals(t.Z, b.Z)
	w := Equals(t.W, b.W)

	// Check if any of the values are not equivalent
	for _, value := range []bool{x, y, z, w} {
		if value == false {
			result = false
			break
		}
	}

	return result
}

// CreateTuple handles creating a tuple. Will eventually handle point vs vector
func CreateTuple(x, y, z, w float64) *Tuple {
	return &Tuple{X: x, Y: y, Z: z, W: w}
}

// CreatePoint creates a tuple designated as a point in order to talk
// about a specific location in space
func CreatePoint(x, y, z float64) *Tuple {
	return CreateTuple(x, y, z, 1.0)
}

// CreateVector creates a tuple designated as a vector. Same as point,
// but this one refers to a point _and_ a direction.
func CreateVector(x, y, z float64) *Tuple {
	return CreateTuple(x, y, z, 0)
}

// Equals checks if two floating point numbers are equivalent by
// seeing if the absolute difference between is less than EPSILON const
func Equals(a, b float64) bool {
	return math.Abs(a-b) < EPSILON
}

// Magnitude calculates the total distance traveled by a vector.
// if a non-vector is passed in, it returns a magnitude of zero.
// uses pythagoras theorem: Sqrt(x^2 + y^2 + z^2)
func Magnitude(vec *Tuple) float64 {
	var aggregate float64

	if vec.IsPoint() {
		return aggregate
	}

	for _, val := range []float64{vec.X, vec.Y, vec.Z} {
		aggregate += Calculate(val, val, Multiply)
	}
	return math.Sqrt(aggregate)
}

// NormalizeVector converts an arbitrary vector into a unit vector using
// the vectors magnitude.
func NormalizeVector(vec *Tuple) *Tuple {
	vectorMagnitude := Magnitude(vec)
	return CreateVector(
		Calculate(vec.X, vectorMagnitude, Divide),
		Calculate(vec.Y, vectorMagnitude, Divide),
		Calculate(vec.Z, vectorMagnitude, Divide),
	)
}

// DotProduct gets the dot product of two tuples.
func DotProduct(a, b *Tuple) float64 {
	return Calculate(a.X, b.X, Multiply) + Calculate(a.Y, b.Y, Multiply) + Calculate(a.Z, b.Z, Multiply) + Calculate(a.W, b.W, Multiply)
}

// CrossProduct gets the cross product of two vectors. Does not work with points.
// Second return value will be true if a non-vector is passed to this function.
func CrossProduct(a, b *Tuple) (*Tuple, bool) {
	if a.IsPoint() || b.IsPoint() {
		return CreateVector(0, 0, 0), true
	}

	return CreateVector(
		Calculate(a.Y, b.Z, Multiply)-Calculate(a.Z, b.Y, Multiply),
		Calculate(a.Z, b.X, Multiply)-Calculate(a.X, b.Z, Multiply),
		Calculate(a.X, b.Y, Multiply)-Calculate(a.Y, b.X, Multiply),
	), false
}

// Subtract function that holds onto the floating point substract function. Used
// as an argument to the Calculate function.
var Subtract func(a, b *big.Float) *big.Float = new(big.Float).Sub

// Addition function that holds onto the floating point addition function. Used
// as an argument to the Calculate function.
var Addition func(a, b *big.Float) *big.Float = new(big.Float).Add

// Multiply function that holds onto the floating point multiple function. Used
// as an argument to the calculate function
var Multiply func(a, b *big.Float) *big.Float = new(big.Float).Mul

// Divide function holds onto the math/big Float division function. Used
// as an argument to the calculate function
var Divide func(a, b *big.Float) *big.Float = new(big.Float).Quo

// Calculate used for doing floating point arithmetic. Takes two operands
// and an arithmetic function, returning the result as a parsed float64
func Calculate(a, b float64, arithmeticFn func(*big.Float, *big.Float) *big.Float) float64 {
	left, success := new(big.Float).SetPrec(200).SetString(strconv.FormatFloat(a, 'f', -1, 64))
	if success == false {
		panic(fmt.Sprintf("Failed to convert float into *big.Float for %v", a))
	}

	right, success := new(big.Float).SetPrec(200).SetString(strconv.FormatFloat(b, 'f', -1, 64))
	if success == false {
		panic(fmt.Sprintf("Failed to convert float into *big.Float for %v", b))
	}

	result := arithmeticFn(left, right)
	parsed, err := strconv.ParseFloat(result.String(), 64)
	if err != nil {
		panic(err)
	}

	return parsed
}
