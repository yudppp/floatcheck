package main

func main() {
	var f1 float64 = 0.123
	var f2 float32 = 4.56
	var i int = 10

	// Formatting that doesn't directly imply precision issues in the format string
	println("%.f", f1)
	println("%g", f2)
	println("%s", "hello")
	println("%d", i)

	// Integer division
	_ = i / 2
	_ = 3 / i
}
