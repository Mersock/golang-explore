// Package mersockacdc test now
package mersockacdc

//Sum add an unlimited number of values of
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
