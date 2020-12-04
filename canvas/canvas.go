package canvas

import (
	"fmt"
	"strings"
)

// Canvas struct that represents the screen where elements are to be rendered
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

// ToPPM converts the calling canvas to a PPM Header string and returns it.
func (c *Canvas) ToPPM() string {
	var ppmFile strings.Builder
	// ppm header
	ppmFile.WriteString(fmt.Sprintf("P3\n%d %d\n%d\n", c.Width, c.Height, 255))

	// ppm pixel body
	for _, row := range c.Screen {
		var rowStringBuilder strings.Builder

		for i, cell := range row {
			scaled := cell.ScaleWithMaxRange(255)
			colourString := fmt.Sprintf("%v %v %v", scaled.Red, scaled.Green, scaled.Blue)
			if i != len(row)-1 {
				colourString += " "
			}
			rowStringBuilder.WriteString(colourString)
		}

		finalString := breakLineAt(70, rowStringBuilder.String())
		ppmFile.WriteString(fmt.Sprintf("%s\n", finalString))
	}
	return ppmFile.String()
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

func breakLineAt(end int, s string) string {
	copy := s
	var sb strings.Builder
	for {
		if len(copy) > end {
			substr := copy[0:70]
			lastSpaceIdx := findLastSpaceChar(substr)
			substr = substr[0:lastSpaceIdx]
			copy = copy[lastSpaceIdx+1:]
			sb.WriteString(fmt.Sprintf("%s\n", substr))
			continue
		}

		trimmedStr := strings.Trim(copy, " ")
		sb.WriteString(trimmedStr)
		break
	}
	return sb.String()
}

func reverseString(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteString(string(s[i]))
	}
	return sb.String()
}

func findLastSpaceChar(s string) int {
	reversed := reverseString(s)
	revIdx := strings.Index(reversed, " ")
	return (len(s) - 1) - revIdx
}
