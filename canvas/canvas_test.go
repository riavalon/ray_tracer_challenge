package canvas

import "testing"

func TestCreateCanvas(t *testing.T) {
	c := NewCanvas(10, 20)

	if c.Width != 10 {
		t.Errorf("Expected canvas to have width property.")
	}

	if c.Height != 20 {
		t.Errorf("Expected canvas to have height property.")
	}

	if len(c.Screen) != 20 {
		t.Errorf("Expected %v rows in canvas", 20)
	}

	if len(c.Screen[0]) != 10 {
		t.Errorf("Expected each row to have %v columns", 10)
	}
}

func TestCreateCanvasPixelsDefaultBlack(t *testing.T) {
	c := NewCanvas(2, 3)

	// Label outer loop in order to break out of it if
	// a bad value if found in a column
outerLoop:
	for _, row := range c.Screen {
		for _, col := range row {
			if col.Red != 0 || col.Green != 0 || col.Blue != 0 {
				t.Errorf("Found non-black pixel in new canvas")
				break outerLoop
			}
		}
	}
}

func TestGetPixelReturnsTrueForSecondParamIfBadCoords(t *testing.T) {
	c := NewCanvas(10, 20)

	if _, success := c.GetPixel(42, 101); success {
		t.Errorf("Should get falsy value for success if given bad coords")
	}
}

func TestWritePixelUpdatesInGrid(t *testing.T) {
	c := NewCanvas(10, 20)
	x, y := c.WritePixel(9, 19, NewColour(1, 1, 1))
	pixel, _ := c.GetPixel(x, y)

	if pixel.Red != 1 || pixel.Green != 1 || pixel.Blue != 1 {
		t.Errorf("Should be able to set a pixels colour on a canvas given a set of x y coords")
	}
}

func TestWritePixelReturnsNegativeOnesIfBadCoords(t *testing.T) {
	c := NewCanvas(10, 20)
	x, y := c.WritePixel(23, 141, NewColour(1, 1, 1))

	if x != -1 || y != -1 {
		t.Errorf("Should get -1 for both x, and y if invalid coords given")
	}
}

func TestGetPixelFromCanvas(t *testing.T) {
	c := NewCanvas(10, 20)
	x, y := c.WritePixel(5, 17, NewColour(1, 1, 1))
	pixel, _ := c.GetPixel(x, y)

	if pixel.Red != 1 || pixel.Green != 1 || pixel.Blue != 1 {
		t.Errorf("Expected to get pixel at x %v; y %v", x, y)
	}
}
