package main

import "fmt"

func main() {
	var f1 float64 = 0.123
	var f2 float32 = 4.56

	// Floating-point formatting
	_ = fmt.Sprintf("%.1f", f1) // want "potential floating-point precision issue in Sprintf format: .1f"
	_ = fmt.Sprintf("%.2f", f2) // want "potential floating-point precision issue in Sprintf format: .2f"

	// Floating-point division
	_ = f1 / 2.0 // want "potential floating-point division"
	_ = 3.0 / f2 // want "potential floating-point division"

	// Floating-point equality comparison
	if f1 == 0.123 { // want "potential floating-point comparison"
		println("equal")
	}
	if f2 != 4.56 { // want "potential floating-point comparison"
		println("not equal")
	}

	// Floating-point relational comparisons
	if f1 < 0.5 { // want "potential floating-point comparison"
		println("less than")
	}
	if f2 <= 5.0 { // want "potential floating-point comparison"
		println("less than or equal to")
	}
	if f1 > 0.1 { // want "potential floating-point comparison"
		println("greater than")
	}
	if f2 >= 4.0 { // want "potential floating-point comparison"
		println("greater than or equal to")
	}
	if f1 > 0.1 && f1 < 0.2 { // want "potential floating-point comparison"
		println("in range")
	}
	if f2 >= 4.0 && f2 <= 5.0 { // want "potential floating-point comparison"
		println("in range")
	}
}
