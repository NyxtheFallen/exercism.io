package collatzconjecture

import "errors"

func CollatzConjecture (n int) (int, error) {
	var steps int = 0

	if n < 0 {
		return n, errors.New("Argument n is negative.")
	} else if n == 0 {
		return n, errors.New("Argument n is 0.")
	}

	for n != 1 {
		switch remainder := n % 2; remainder {
		case 0: n /= 2
		default: n = (3 * n) + 1
		}
		steps++
	}
	return steps, nil
}