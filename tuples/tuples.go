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
	x float64
	y float64
	z float64
	w float64
}

func (t *Tuple) negate() *Tuple {
	zero := CreateVector(0, 0, 0)
	negated := zero.subtract(t)
	negated.w = t.w
	return negated
}

func (t *Tuple) add(b *Tuple) *Tuple {
	return CreatePoint(
		Calculate(t.x, b.x, Addition),
		Calculate(t.y, b.y, Addition),
		Calculate(t.z, b.z, Addition),
	)
}

func (t *Tuple) subtract(b *Tuple) *Tuple {
	return &Tuple{
		x: Calculate(t.x, b.x, Subtract),
		y: Calculate(t.y, b.y, Subtract),
		z: Calculate(t.z, b.z, Subtract),
		w: Calculate(t.w, b.w, Subtract),
	}
}

func (t *Tuple) multiplyByScalar(n float64) *Tuple {
	tup := CreateTuple(
		Calculate(n, t.x, Multiply),
		Calculate(n, t.y, Multiply),
		Calculate(n, t.z, Multiply),
		t.w,
	)
	return tup
}

func (t *Tuple) divideByScalar(n float64) *Tuple {
	tup := CreateTuple(
		Calculate(t.x, n, Divide),
		Calculate(t.y, n, Divide),
		Calculate(t.z, n, Divide),
		t.w,
	)
	return tup
}

func (t *Tuple) isPoint() bool {
	return t.w == 1.0
}

func (t *Tuple) isVector() bool {
	return t.w == 0
}

func (t *Tuple) isEquivalentTo(b *Tuple) bool {
	result := true

	x := Equals(t.x, b.x)
	y := Equals(t.y, b.y)
	z := Equals(t.z, b.z)
	w := Equals(t.w, b.w)

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
	return &Tuple{x: x, y: y, z: z, w: w}
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

	if vec.isPoint() {
		return aggregate
	}

	for _, val := range []float64{vec.x, vec.y, vec.z} {
		aggregate += Calculate(val, val, Multiply)
	}
	return math.Sqrt(aggregate)
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
