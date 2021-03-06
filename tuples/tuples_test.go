package tuples

import (
	"math"
	"testing"
)

func TestCreateTuple(t *testing.T) {
	point := CreateTuple(4.3, -4.2, 3.1, 1.0)

	if point.X != 4.3 {
		t.Errorf("point.x should have init value. Got %v; Want %v", point.X, 4.3)
	}

	if point.Y != -4.2 {
		t.Errorf("point.y should have init value. Got %v; Want %v", point.Y, -4.2)
	}

	if point.Z != 3.1 {
		t.Errorf("point.z should have init value. Got %v; Want %v", point.Z, 3.1)
	}

	if point.W != 1.0 {
		t.Errorf("point.w should have init value. Got %v; Want %v", point.W, 1.0)
	}
}

func TestCheckIsPoint(t *testing.T) {
	tup := CreateTuple(4.3, 4.2, 3.1, 1.0)

	if tup.IsPoint() == false {
		t.Errorf("Expected tuple with w 1.0 to be point.")
	}
}

func TestCheckIsVector(t *testing.T) {
	tup := CreateTuple(4.3, -4.2, 3.1, 0.0)

	if tup.IsVector() == false {
		t.Errorf("Expected tuple with w 0 to be vector.")
	}
}

func TestCreatePoint(t *testing.T) {
	point := CreatePoint(4.3, -4.2, 3.1)

	if point.W != 1.0 {
		t.Errorf("Expected w value to be for a point. Got %v; Want %v", point.W, 1.0)
	}
}

func TestCreateVector(t *testing.T) {
	vector := CreateVector(4.3, -4.2, 3.1)

	if vector.W != 0.0 {
		t.Errorf("Expected w value to be for vector. Got %v; Want %v", vector.W, 0.0)
	}
}

func TestFloatingPointNotEqual(t *testing.T) {
	a := 4.2
	b := 4.9

	if Equals(a, b) == true {
		t.Errorf("Expected values to be non-equivalent")
	}
}

func TestFloatingPointEqual(t *testing.T) {
	a := 4.2
	b := 4.2

	if Equals(a, b) == false {
		t.Errorf("Expected values to be equivalent")
	}
}

func TestCompareTuplesForEquivalency(t *testing.T) {
	a := CreatePoint(4.3, -4.2, 3.1)
	b := CreatePoint(4.3, -4.2, 3.1)

	if a.IsEquivalentTo(b) == false {
		t.Errorf("Expected tuples to be equal, but got falsy value")
	}
}

func TestCompareTuplesNotEquivalent(t *testing.T) {
	a := CreatePoint(4.3, -4.2, 3.1)
	b := CreatePoint(4.3, -2.2, 3.1)

	if a.IsEquivalentTo(b) == true {
		t.Errorf("Expected tuples to not be equivalent, but got truthy value.")
	}
}

func TestPointNotEqualToVector(t *testing.T) {
	a := CreateVector(4.3, 3.2, 1.7)
	b := CreatePoint(4.3, 3.2, 1.7)

	if a.IsEquivalentTo(b) {
		t.Errorf("Expected Point and Vector to not be equivalent")
	}
}

func TestAddTwoTuplesTogether(t *testing.T) {
	a := CreatePoint(4.2, 2.1, 1.9)
	b := CreatePoint(-2.3, 4.7, 10.2)
	want := CreatePoint(1.9, 6.8, 12.1)
	got := a.Add(b)

	if got.IsEquivalentTo(want) == false {
		t.Errorf("Expected points to add together.\nGot %v;\nWant %v", got, want)
	}
}

func TestAddTwoFloatsTogether(t *testing.T) {
	a := 4.3
	b := -2.5
	got := Calculate(a, b, Addition)
	want := 1.8

	if Equals(got, want) == false {
		t.Errorf("Expected properly rounded floating point number after addition.\nGot %v;\nWant %v", got, want)
	}
}

func TestSubtractTwoFloats(t *testing.T) {
	a := 4.3
	b := 4.1
	got := Calculate(a, b, Subtract)
	want := 0.2

	if Equals(got, want) == false {
		t.Errorf("Expected properly rounded floating point number after subtraction.\nGot %v;\nWant %v", got, want)
	}
}

func TestSubtractTwoPoints(t *testing.T) {
	a := CreatePoint(2.3, 3.1, 4.2)
	b := CreatePoint(2.1, 3.0, 4.1)
	got := a.Subtract(b)
	want := CreateVector(0.2, 0.1, 0.1)

	if got.IsEquivalentTo(want) == false {
		t.Errorf("Expected subtracting a point from point to be a vector with proper values.\nGot %v;\nWant %v", got, want)
	}
}

func TestSubstractTwoPointsEqualsVector(t *testing.T) {
	a := CreatePoint(5.3, -2.3, 8.2)
	b := CreatePoint(2.3, 1.2, 6.3)
	got := a.Subtract(b)

	if got.IsVector() == false {
		t.Errorf("Expected vector from point - point.\nGot %v", got)
	}
}

func TestSubtractVectorFromPoint(t *testing.T) {
	a := CreatePoint(3.2, 82.2, -23.2)
	b := CreateVector(7.2, 13.6, -8.2)
	got := a.Subtract(b)

	if got.IsPoint() == false {
		t.Errorf("Expected point from point - vector.\nGot %v", got)
	}
}

func TestSubtractVectorFromVector(t *testing.T) {
	a := CreateVector(3.2, 5.2, 8.32)
	b := CreateVector(3.8, 9.22, -82.1)
	got := a.Subtract(b)

	if got.IsVector() == false {
		t.Errorf("Expected vector from vector - vector.\nGot %v", got)
	}
}

func TestSubtractVectorFromZeroVector(t *testing.T) {
	a := CreateVector(3.2, 1.3, 5.4)
	b := CreateVector(0, 0, 0)
	got := b.Subtract(a)
	want := CreateVector(-3.2, -1.3, -5.4)

	if got.IsEquivalentTo(want) == false {
		t.Errorf("Expected negated version of vector.\nGot %v;\nWant %v", got, want)
	}
}

func TestNegateTupleFunction(t *testing.T) {
	orig := CreatePoint(3.2, -1.3, 5.4)
	got := orig.Negate()
	want := CreatePoint(-3.2, 1.3, -5.4)

	if got.IsEquivalentTo(want) == false {
		t.Errorf("Expected tuple to be negated.\nGot %v;\nWant %v", got, want)
	}
}

func TestMultiplyFloatingPointNumbers(t *testing.T) {
	got := Calculate(3, 7.5, Multiply)
	want := 22.5

	if Equals(got, want) == false {
		t.Errorf("Expected accurate floating point multiplication.\nGot %v;\nWant %v", got, want)
	}
}

func TestMultiplyTupleByScalar(t *testing.T) {
	point := CreatePoint(3.2, 2.5, 5.4)
	got := point.MultiplyByScalar(3)
	want := CreatePoint(9.6, 7.5, 16.2)

	if got.IsEquivalentTo(want) == false {
		t.Errorf("Expected correct scalar multiplication of tuple.\nGot %v;\nWant %v", got, want)
	}
}

func TestMultiplyTupleByScalarFraction(t *testing.T) {
	point := CreatePoint(5.0, 6.0, 8.0)
	got := point.MultiplyByScalar(0.5)
	want := CreatePoint(2.5, 3.0, 4.0)

	if got.IsEquivalentTo(want) == false {
		t.Errorf("Expected correct scalar multiplication by a fraction of a tuple.\nGot %v;\nWant %v", got, want)
	}
}

func TestDivideFloatingPointNumbers(t *testing.T) {
	got := Calculate(3.2, 2, Divide)
	want := 1.6

	if Equals(got, want) == false {
		t.Errorf("Expected proper floating point division.\nGot %v;\nWant %v", got, want)
	}
}

func TestDivideTupleByScalarValue(t *testing.T) {
	point := CreatePoint(4.0, 2.0, -10.0)
	got := point.DivideByScalar(2.0)
	want := CreatePoint(2.0, 1.0, -5.0)

	if got.IsEquivalentTo(want) == false {
		t.Errorf("Expected proper division of tuples by scalar value.\nGot %v;\nWant %v", got, want)
	}
}

func TestGetMagnitudeOfVector(t *testing.T) {
	vec := CreateVector(1, 0, 0)
	got := Magnitude(vec)
	want := 1.0

	if Equals(got, want) == false {
		t.Errorf("Expected to get magnitude one for vector. Instead got %v", got)
	}
}

func TestMagnitudeOfPointShouldBeZero(t *testing.T) {
	point := CreatePoint(1, 0, 0)
	got := Magnitude(point)
	want := 0.0

	if Equals(got, want) == false {
		t.Errorf("Expected to get %v for a non-vector passed to Magnitude. Got %v", want, got)
	}
}

func TestMagnitudeOfVectorY(t *testing.T) {
	vec := CreateVector(0, 1, 0)
	got := Magnitude(vec)
	want := math.Sqrt(1)

	if Equals(got, want) == false {
		t.Errorf("Expected to get %v for magnitude of vector with non-zero Y. Got %v", want, got)
	}
}

func TestMagnitudeEntireVector(t *testing.T) {
	vec := CreateVector(3, 8, 10)
	got := Magnitude(vec)
	want := math.Sqrt(173)

	if Equals(got, want) == false {
		t.Errorf("Expected total magnitude from vector. Got %v; Want %v", got, want)
	}
}

func TestMagnitudeWithNegativeValues(t *testing.T) {
	vec := CreateVector(-1, -2, -3)
	got := Magnitude(vec)
	want := math.Sqrt(14)

	if Equals(got, want) == false {
		t.Errorf("Failed to get magnitude with negative numbers. Got %v; Want %v", got, want)
	}
}

func TestNormalizeSimpleVector(t *testing.T) {
	vec := CreateVector(4, 0, 0)
	got := NormalizeVector(vec)
	want := CreateVector(1, 0, 0)

	if got.IsEquivalentTo(want) == false {
		t.Errorf("Expected normalized unit vector.\nGot %v;\nWant %v", got, want)
	}
}

func TestNormalizeVector(t *testing.T) {
	vec := CreateVector(1, 2, 3)
	got := NormalizeVector(vec)
	want := CreateVector(
		Calculate(vec.X, math.Sqrt(14), Divide),
		Calculate(vec.Y, math.Sqrt(14), Divide),
		Calculate(vec.Z, math.Sqrt(14), Divide),
	)

	if got.IsEquivalentTo(want) == false {
		t.Errorf("Expected vectors to match.\nGot  %v;\nWant %v;", got, want)
	}
}

func TestMagnitudeOfNormalizedVector(t *testing.T) {
	vec := CreateVector(1, 2, 3)
	normalized := NormalizeVector(vec)
	got := Magnitude(normalized)
	want := 1.0

	if Equals(got, want) == false {
		t.Errorf("Expected to get magnitude of one.\nGot  %v;\nWant %v;", got, want)
	}
}

func TestDotProductOfTwoTuples(t *testing.T) {
	a := CreateVector(1, 2, 3)
	b := CreateVector(2, 3, 4)

	got := DotProduct(a, b)
	want := 20.0

	if Equals(got, want) == false {
		t.Errorf("Expected dot product to be %v. Got %v", want, got)
	}
}

func TestCrossProductOfTwoVectors(t *testing.T) {
	a := CreateVector(1, 2, 3)
	b := CreateVector(2, 3, 4)
	aFirst, _ := CrossProduct(a, b)
	bFirst, _ := CrossProduct(b, a)
	aWant := CreateVector(-1, 2, -1)
	bWant := CreateVector(1, -2, 1)

	if aFirst.IsEquivalentTo(aWant) == false {
		t.Errorf("Expected correct cross product of `a` first vector.\nGot  %v;\nWant %v;", aFirst, aWant)
	}

	if bFirst.IsEquivalentTo(bWant) == false {
		t.Errorf("Expected correct cross product of `b` first vector.\nGot  %v;\nWant %v;", bFirst, bWant)
	}
}

func TestInfinityFromCrossProductForNonVector(t *testing.T) {
	a := CreatePoint(1, 2, 3)
	b := CreateVector(2, 3, 4)

	if vec, err := CrossProduct(a, b); !err {
		t.Errorf("Expected truthy error when supplying `a` with a point. Vec is %v", vec)
	}

	if vec, err := CrossProduct(b, a); !err {
		t.Errorf("Expected truthy error when supplying `b` with a point. Vec is %v", vec)
	}
}
