package canvas

// Canvas struct that represents the screen where elements
// are to be rendered.
type Canvas struct {
	Width  int
	Height int
	Screen [][]Colour
}

// WritePixel allows setting a pixel to a certain colour
// given a canvas and a set of valid co-ordinates.
func (c *Canvas) WritePixel(x, y int, colour Colour) (int, int) {
	if x >= c.Width || y >= c.Height {
		return -1, -1
	}

	c.Screen[y][x] = colour
	return x, y
}

// GetPixel takes an x and y coordinate to search for in the canvas screen.
// Returns the pixel colour if found. Second argument is a bool representing
// success. False if there was a failure, and true if a pixel was found.
func (c *Canvas) GetPixel(x, y int) (Colour, bool) {
	if x >= c.Width || y >= c.Height {
		return NewColour(0, 0, 0), false
	}
	colour := c.Screen[y][x]
	return colour, true
}

// NewCanvas creates a new canvas with specified width and height,
// creating a blank black colour pixel in each and every grid position.
func NewCanvas(w, h int) Canvas {
	screen := make([][]Colour, h)

	// make empty 2d slice
	for i := range screen {
		screen[i] = make([]Colour, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			screen[i][j] = NewColour(0, 0, 0)
		}
	}

	return Canvas{
		Width:  w,
		Height: h,
		Screen: screen,
	}
}
