package common

import "math"

// Abs 返回一个数的绝对值
func Abs(num float64) float64 {
	return math.Abs(num)
}

// ACos 返回一个数的反余弦
func ACos(num float64) float64 {
	return math.Acos(num)
}

// ACosh 返回一个数的反双曲余弦
func ACosh(num float64) float64 {
	return math.Log(num + math.Sqrt(num*num-1.0))
}
