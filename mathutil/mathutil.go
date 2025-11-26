package mathutil

import "errors"

var Age = 19

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

func Divide(x, y int) (float64, error) {
	xFloat := float64(x)
	yFloat := float64(y)
	if y <= 0 {
		return 0, errors.New("it can't be 0")
	}

	return xFloat / yFloat, nil
}
