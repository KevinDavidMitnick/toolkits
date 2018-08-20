package core

import (
	"fmt"
)

func ReadableSize(raw float64) string {
	var t float64 = 1024
	var d float64 = 1

	if raw < t {
		return fmt.Sprintf("%.2fB", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.2fK", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.2fM", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.2fG", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.2fT", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.2fP", raw/d)
	}

	return "TooLarge"
}
