package main

import "fmt"

func main() {
	var f1 float64 = 0.123
	var f2 float32 = 4.56

	// Floating-point formatting
	_ = fmt.Sprintf("%.1f", f1)
	_ = fmt.Sprintf("%.2f", f2)

	// Floating-point division
	_ = f1 / 2.0
	_ = 3.0 / f2

	// Floating-point equality comparison
	if f1 == 0.123 { // want "potential floating-point comparison, consider using a tolerance"
		println("equal")
	}
	if f2 != 4.56 { // want "potential floating-point comparison, consider using a tolerance"
		println("not equal")
	}

	// Floating-point relational comparisons
	if f1 < 0.5 { // want "potential floating-point comparison, consider using a tolerance"
		println("less than")
	}
	if f2 <= 5.0 { // want "potential floating-point comparison, consider using a tolerance"
		println("less than or equal to")
	}
	if f1 > 0.1 { // want "potential floating-point comparison, consider using a tolerance"
		println("greater than")
	}
	if f2 >= 4.0 { // want "potential floating-point comparison, consider using a tolerance"
		println("greater than or equal to")
	}
	if f1 > 0.1 && f1 < 0.2 { // want "potential floating-point comparison, consider using a tolerance"
		println("in range")
	}
	if f2 >= 4.0 && f2 <= 5.0 { // want "potential floating-point comparison, consider using a tolerance"
		println("in range")
	}
}
