package canvas

import (
	"bufio"
	"strings"
	"testing"
)

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

func TestCanvasToPPMHeader(t *testing.T) {
	c := NewCanvas(5, 3)
	got := c.ToPPM()
	want := []string{"P3", "5 3", "255"}

	scanner := bufio.NewScanner(strings.NewReader(got))
	line := 0

	for scanner.Scan() {
		if line > 2 {
			break
		}
		text := scanner.Text()
		if text != want[line] {
			t.Errorf("should get valid plain PPM header. Got %q; Want %q", text, want[line])
		}
		line++
	}
}

func TestCanvasToPPMPixelBody(t *testing.T) {
	c := NewCanvas(5, 3)

	// colours we expect to find
	c1 := NewColour(1.5, 0, 0)
	c2 := NewColour(0, 0.5, 0)
	c3 := NewColour(-0.5, 0, 1)

	// Add coloured pixels to screen
	c.WritePixel(0, 0, c1)
	c.WritePixel(2, 1, c2)
	c.WritePixel(4, 2, c3)

	got := c.ToPPM()
	want := `P3
5 3
255
255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`

	if got != want {
		t.Errorf("Should build the pixel body according to colours specified.\nGot\n%+vWant\n%+v", got, want)
	}
}

// I feel very ashamed with how dumb this unit test is written
// but it does confirm the thing so I guess hoorah? I can refactor it later
func TestCanvasToPPMBodyLinesDontExceed70Chars(t *testing.T) {
	c := NewCanvas(10, 2)

	// set every pixel to colour(1, 0.8, 0.6)
	colour := NewColour(1, 0.8, 0.6)
	for y, row := range c.Screen {
		for x := range row {
			c.WritePixel(x, y, colour)
		}
	}

	got := c.ToPPM()
	want := []string{
		"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204",
		"153 255 204 153 255 204 153 255 204 153 255 204 153",
		"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204",
		"153 255 204 153 255 204 153 255 204 153 255 204 153",
	}

	scanner := bufio.NewScanner(strings.NewReader(got))
	iteration := 0
	line := 0

	for scanner.Scan() {
		if iteration < 3 {
			iteration++
			continue
		}

		if iteration > 6 {
			break
		}

		text := scanner.Text()
		if text != want[line] {
			t.Errorf("Line should move to the next line after 70 characters.\nGot\n%v\n\nWant\n%v", text, want[line])
		}
		line++
	}
}

func TestPPMFileEndsInNewLine(t *testing.T) {
	c := NewCanvas(5, 3)
	got := c.ToPPM()
	endsInNewline := strings.HasSuffix(got, "\n")

	if !endsInNewline {
		t.Errorf("PPM file should end in a newline character")
	}
}
