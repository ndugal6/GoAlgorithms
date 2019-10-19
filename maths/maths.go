//Package for maths
package maths

//computes the greatest common donimator
func GCD(x,y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
//computes the n-th Fibonacci term interatively
func FIB(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}