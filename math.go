package common

import (
	"math"
	"strconv"
)

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

// ASin 返回一个数的反正弦
func ASin(num float64) float64 {
	return math.Asin(num)
}

// ASinH 返回一个数的反双曲正弦
func ASinH(num float64) float64 {
	return math.Asinh(num)
}

// ATan 返回一个数的反正切
func ATan(num float64) float64 {
	return math.Atan(num)
}

// ATan2 返回两个变量 x 和 y 的反正切
func ATan2(y, x float64) float64 {
	return math.Atan2(y, x)
}

// Atanh 返回一个数的反双曲正切
func Atanh(num float64) float64 {
	return math.Atanh(num)
}

// BaseConvert 在任意进制之间转换数字
func BaseConvert(num string, fromBase, toBase int) string {
	decimal, err := strconv.ParseInt(num, fromBase, 0)
	if err != nil {
		return ""
	}
	return strconv.FormatInt(decimal, toBase)
}

// BinDec 把二进制数转换为十进制数
func BinDec(binary string) int64 {
	decimal, err := strconv.ParseInt(binary, 2, 0)
	if err != nil {
		return 0
	}
	return decimal
}

// Ceil 向上舍入为最接近的整数
func Ceil(num float64) float64 {
	return math.Ceil(num)
}

// Cos 返回一个数的余弦
func Cos(num float64) float64 {
	return math.Cos(num)
}

// Cosh 返回一个数的双曲余弦
func Cosh(num float64) float64 {
	return math.Cosh(num)
}

// DecBin 把十进制数转换为二进制数
func DecBin(num string) string {
	return BaseConvert(num, 10, 2)
}

// DecHex 把十进制数转换为十六进制数
func DecHex(num string) string {
	return BaseConvert(num, 10, 16)
}

// DecOct 把十进制数转换为八进制数
func DecOct(num string) string {
	return BaseConvert(num, 10, 8)
}

// Deg2Rad 将角度值转换为弧度值
func Deg2Rad(num float64) float64 {
	radians := num * math.Pi / 180
	return radians
}

// Exp 返回E的x次方的值
func Exp(x float64) float64 {
	return math.Exp(x)
}

// Expm1 返回Exp(x) - 1
func Expm1(x float64) float64 {
	return Exp(x) - 1
}

// Floor 向下取整
func Floor(num float64) float64 {
	return math.Floor(num)
}

// Fmod 返回 x/y 的浮点数余数
func Fmod(x, y float64) float64 {
	return math.Mod(x, y)
}

// GetRandMax 返回通过调用 Rand() 函数显示的随机数的最大可能值
func GetRandMax() int64 {
	return int64(^uint64(0) >> 1)
}

// HexDec 把十六进制转换为十进制
func HexDec(num string) string {
	return BaseConvert(num, 16, 10)
}
