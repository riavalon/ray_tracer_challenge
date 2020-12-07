package matrix

import "testing"

func TestCreateMatrixAndGetValues(t *testing.T) {
	m := NewMatrix(
		1, 2, 3, 4,
		5.5, 6.5, 7.5, 8.5,
		9, 10, 11, 12,
		13.5, 14.5, 15.5, 16.5,
	)

	errorMessage := "Matrix should have correct values at given row and col. Got %v; Want %v"

	if m.Get(0, 0) != 1 {
		t.Errorf(errorMessage, m.Get(0, 0), 1)
	}

	if m.Get(0, 3) != 4 {
		t.Errorf(errorMessage, m.Get(0, 3), 4)
	}

	if m.Get(1, 0) != 5.5 {
		t.Errorf(errorMessage, m.Get(1, 0), 5.5)
	}

	if m.Get(1, 2) != 7.5 {
		t.Errorf(errorMessage, m.Get(1, 2), 7.5)
	}

	if m.Get(2, 2) != 11 {
		t.Errorf(errorMessage, m.Get(2, 2), 11)
	}

	if m.Get(3, 0) != 13.5 {
		t.Errorf(errorMessage, m.Get(3, 0), 13.5)
	}

	if m.Get(3, 2) != 15.5 {
		t.Errorf(errorMessage, m.Get(3, 2), 15.5)
	}
}

func TestMatrixSetsValueZeroIfFewerThan16ArgsAreGiven(t *testing.T) {
	m := NewMatrix(1, 2, 3, 4)

outerLoop:
	for y, row := range m.grid {
		if y > 0 {
			for _, cell := range row {
				if cell != 0 {
					t.Errorf("Expected default zero value when not enough args")
					break outerLoop
				}
			}
		}
	}
}

func TestSetSingleValueInMatrix(t *testing.T) {
	m := NewMatrix(1, 2, 3, 4)
	m.Set(0, 2, 5)

	if m.grid[0][2] != 5 {
		t.Errorf("Expected first row, third item to be updated by set method. Got %v, want %v", m.grid[0][2], 5)
	}
}
