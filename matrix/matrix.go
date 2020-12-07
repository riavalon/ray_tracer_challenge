package matrix

// Matrix struct represents a 4x4 matrix of numbers.
type Matrix struct {
	grid [][]float64
}

// Get will take a row and col argument and return the value held at that
// position in the matrix.
func (m Matrix) Get(row, col int) float64 {
	return m.grid[row][col]
}

// Set will take a row and column argument as well as a new value and update
// the item at that coordinate in the matrix with the given value. If index
// is out of bounds, nothing happens.
func (m *Matrix) Set(row, col int, value float64) {
	if row < 0 || row > 3 || col < 0 || col > 3 {
		return
	}
	m.grid[row][col] = value
}

// SeedMatrix takes a variadic param of float64s to populate
// into the grid going from top to bottom, left to right. Default or
// existing values are left in pace if not enough arguments are given.
// If > 16 arguments are given, the rest after the 16th will be ignored.
func (m *Matrix) SeedMatrix(items ...float64) {
outerLoop:
	for y, row := range m.grid {
		for x := range row {
			if len(items) > 0 {
				m.grid[y][x] = items[0]
				items = items[1:]
				continue
			}
			break outerLoop
		}
	}
}

// NewMatrix creates a new matrix struct with provided values squentially.
// Values are set from top-down, left-right.
// e.g
// NewMatrix(1, 2, 3, 4, 5, 6, 7, 8 ...) =>
// [
//		[1, 2, 3, 4],
//		[5, 6, 7, 8],
//		...
// ]
func NewMatrix(values ...float64) Matrix {
	m := Matrix{}
	m.grid = buildEmptyMatrix()
	m.SeedMatrix(values...)
	return m
}

func buildEmptyMatrix() [][]float64 {
	matrix := make([][]float64, 4)
	for i := range matrix {
		row := make([]float64, 4)
		matrix[i] = row
	}
	return matrix
}
