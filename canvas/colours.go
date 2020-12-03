package canvas

import (
	"math"

	tup "github.com/riavalon/ray_tracer/tuples"
)

// Colour struct that represents a tuple of red, green, blue
// TODO: In the future, try to leverage the tuples package for
// colours. Perhaps when we include alpha into the tuple?
type Colour struct {
	Red   float64
	Green float64
	Blue  float64
}

// ScaleWithMaxRange takes an integer representing the max range
// of the new scale. Will convert the colour components to conform
// with that scale and return the new converted colour.
func (c Colour) ScaleWithMaxRange(m float64) Colour {
	return NewColour(
		math.Max(0, math.Min(m, math.Round(tup.Calculate(m, c.Red, tup.Multiply)))),
		math.Max(0, math.Min(m, math.Round(tup.Calculate(m, c.Green, tup.Multiply)))),
		math.Max(0, math.Min(m, math.Round(tup.Calculate(m, c.Blue, tup.Multiply)))),
	)
}

// Subtract takes a passed in colour struct and subtracts it from
// the calling colour struct, returning the difference as a new
// colour struct.
func (c Colour) Subtract(c2 Colour) Colour {
	return NewColour(
		tup.Calculate(c.Red, c2.Red, tup.Subtract),
		tup.Calculate(c.Green, c2.Green, tup.Subtract),
		tup.Calculate(c.Blue, c2.Blue, tup.Subtract),
	)
}

// Add takes a passed in colour struct and adds it to the calling
// colour struct. Returns sum total colour
func (c Colour) Add(c2 Colour) Colour {
	return NewColour(
		tup.Calculate(c.Red, c2.Red, tup.Addition),
		tup.Calculate(c.Green, c2.Green, tup.Addition),
		tup.Calculate(c.Blue, c2.Blue, tup.Addition),
	)
}

// MultiplyByScalar multiplies each component of a colour by a given
// scalar value, returning the resulting colour.
func (c Colour) MultiplyByScalar(scalar float64) Colour {
	return NewColour(
		tup.Calculate(c.Red, scalar, tup.Multiply),
		tup.Calculate(c.Green, scalar, tup.Multiply),
		tup.Calculate(c.Blue, scalar, tup.Multiply),
	)
}

// DivideByScalar divides every component of the colour tuple by the
// given scalar value, returning the resulting colour.
func (c Colour) DivideByScalar(scalar float64) Colour {
	return NewColour(
		tup.Calculate(c.Red, scalar, tup.Divide),
		tup.Calculate(c.Green, scalar, tup.Divide),
		tup.Calculate(c.Blue, scalar, tup.Divide),
	)
}

// IsEquivalentTo checks if the passed in colour is equivalent to the
// calling colour.
func (c Colour) IsEquivalentTo(c2 Colour) bool {
	result := true

	red := tup.Equals(c.Red, c2.Red)
	green := tup.Equals(c.Green, c2.Green)
	blue := tup.Equals(c.Blue, c2.Blue)

	for _, val := range []bool{red, green, blue} {
		if !val {
			result = false
			break
		}
	}

	return result
}

// NewColour creates a colour tuple and returns it
func NewColour(r, g, b float64) Colour {
	return Colour{
		Red:   r,
		Green: g,
		Blue:  b,
	}
}

// MultiplyColours takes two colour tuples and returns the product
// colour of multiplying each component together.
func MultiplyColours(c1, c2 Colour) Colour {
	return NewColour(
		tup.Calculate(c1.Red, c2.Red, tup.Multiply),
		tup.Calculate(c1.Green, c2.Green, tup.Multiply),
		tup.Calculate(c1.Blue, c2.Blue, tup.Multiply),
	)
}
