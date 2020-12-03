package canvas

import (
	"testing"
)

func TestNewColour(t *testing.T) {
	col := NewColour(0.8, 0, 1)

	if col.Red != 0.8 {
		t.Errorf("NewColour func should have red value in tuple. Got %v; Want %v", col.Red, 0.8)
	}

	if col.Green != 0 {
		t.Errorf("NewColour func should have green value in tuple. Got %v; Want %v", col.Green, 0)
	}

	if col.Blue != 1 {
		t.Errorf("NewColour func should have blue value in tuple. Got %v; Want %v", col.Blue, 1)
	}
}

func TestCheckIfColoursAreEquivalent(t *testing.T) {
	col1 := NewColour(0.1, 0.2, 0.3)
	col2 := NewColour(0.1, 0.2, 0.3)

	if col1.IsEquivalentTo(col2) == false {
		t.Errorf("Should be able to check equivalence of two colours.")
	}
}

func TestAddTwoColours(t *testing.T) {
	col1 := NewColour(0.9, 0.6, 0.75)
	col2 := NewColour(0.7, 0.1, 0.25)
	got := col1.Add(col2)
	want := NewColour(1.6, 0.7, 1)

	if got.IsEquivalentTo(want) == false {
		t.Errorf("Failed to add two tuples together.\nGot  %v;\nWant %v;", got, want)
	}
}

func TestSubtractTwoColours(t *testing.T) {
	col1 := NewColour(0.8, 0.2, 0.93)
	col2 := NewColour(0.4, 0.1, 0.45)
	got := col1.Subtract(col2)
	want := NewColour(0.4, 0.1, 0.48)

	if got.IsEquivalentTo(want) == false {
		t.Errorf("Failed to subtract two colour tuples from each other.\nGot  %v;\nWant %v;", got, want)
	}
}

func TestMultiplyColourByScalar(t *testing.T) {
	col := NewColour(0.2, 0.5, 0.8)
	got := col.MultiplyByScalar(2)
	want := NewColour(0.4, 1, 1.6)

	if got.IsEquivalentTo(want) == false {
		t.Errorf("Colour should be result of multiplying by scalar value.\nGot  %v;\nWant %v;", got, want)
	}
}

func TestDivideColourByScalar(t *testing.T) {
	col := NewColour(0.8, 0.2, 0.6)
	got := col.DivideByScalar(2)
	want := NewColour(0.4, 0.1, 0.3)

	if got.IsEquivalentTo(want) == false {
		t.Errorf("Colour should be result of dividing by scalar value.\nGot  %v;\nWant %v;", got, want)
	}
}

func TestMultiplyingColours(t *testing.T) {
	c1 := NewColour(1, 0.2, 0.4)
	c2 := NewColour(0.9, 1, 0.1)
	got := MultiplyColours(c1, c2)
	want := NewColour(0.9, 0.2, 0.04)

	if got.IsEquivalentTo(want) == false {
		t.Errorf("Colours should be able to multiply correctly.\nGot  %v;Want %v;", got, want)
	}
}
