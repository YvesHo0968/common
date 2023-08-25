package common

import (
	"fmt"
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	fmt.Println(Abs(-18.9))
}

func TestACos(t *testing.T) {
	fmt.Println(ACos(0.64))
}

func TestACosh(t *testing.T) {
	fmt.Println(ACosh(7))
}

func TestASin(t *testing.T) {
	fmt.Println(ASin(0.64))
}

func TestASinH(t *testing.T) {
	fmt.Println(ASinH(7))
}

func TestATan(t *testing.T) {
	fmt.Println(ATan(0.50))
}

func TestATan2(t *testing.T) {
	fmt.Println(ATan2(0.50, 0.50))
}

func TestAtanh(t *testing.T) {
	fmt.Println(Atanh(0.5))
}

func TestBaseConvert(t *testing.T) {
	fmt.Println(BaseConvert("E196", 16, 8))
}

func TestBinDec(t *testing.T) {
	fmt.Println(BinDec("0011"))
}

func TestCeil(t *testing.T) {
	fmt.Println(Ceil(1.2))
}

func TestCos(t *testing.T) {
	fmt.Println(Cos(3))
}

func TestCosh(t *testing.T) {
	fmt.Println(Cosh(3))
}

func TestDecBin(t *testing.T) {
	fmt.Println(DecBin("3"))
}

func TestDecHex(t *testing.T) {
	fmt.Println(DecHex("30"))
}

func TestDecOct(t *testing.T) {
	fmt.Println(DecOct("30"))
}

func TestDeg2Rad(t *testing.T) {
	fmt.Println(Deg2Rad(45))
}

func TestExp(t *testing.T) {
	fmt.Println(Exp(1))
}

func TestExpm1(t *testing.T) {
	fmt.Println(Expm1(1))
}

func TestFloor(t *testing.T) {
	fmt.Println(Floor(1.9))
}

func TestFmod(t *testing.T) {
	fmt.Println(Fmod(7, 4.5))
}

func TestGetRandMax(t *testing.T) {
	fmt.Println(GetRandMax())
}

func TestHexDec(t *testing.T) {
	fmt.Println(HexDec("1e"))
}

func TestHypot(t *testing.T) {
	fmt.Println(Hypot(3, 4))
}

func TestIsFinite(t *testing.T) {
	fmt.Println(IsFinite(1))
	fmt.Println(IsFinite(math.Inf(1)))
}

func TestIsInFinite(t *testing.T) {
	fmt.Println(IsInFinite(2))
}

func TestIsNaN(t *testing.T) {
	fmt.Println(IsNaN(2))
	fmt.Println(IsNaN(math.NaN()))
}

func TestLcgValue(t *testing.T) {
	fmt.Println(LcgValue())
}

func TestLog(t *testing.T) {
	fmt.Println(Log(2))
}

func TestLog10(t *testing.T) {
	fmt.Println(Log10(2))
}

func TestLog1p(t *testing.T) {
	fmt.Println(Log1p(2))
}

func TestMtGetRandMax(t *testing.T) {
	fmt.Println(MtGetRandMax())
}

func TestMtRand(t *testing.T) {
	fmt.Println(MtRand(1000, 9999))
}

func TestOctDec(t *testing.T) {
	fmt.Println(OctDec("36"))
}

func TestPi(t *testing.T) {
	fmt.Println(Pi())
}

func TestPow(t *testing.T) {
	fmt.Println(Pow(2, 4))
}

func TestRad2Deg(t *testing.T) {
	fmt.Println(Rad2Deg(Pi()))
}

func TestRand(t *testing.T) {
	fmt.Println(Rand(1000, 9999))
}

func TestRound(t *testing.T) {
	fmt.Println(Round(1.5))
}

func TestSin(t *testing.T) {
	fmt.Println(Deg2Rad(30))
	fmt.Println(Sin(3))
}

func TestSinh(t *testing.T) {
	fmt.Println(Sinh(3))
}

func TestSqrt(t *testing.T) {
	fmt.Println(Sqrt(9))
}

func TestTanh(t *testing.T) {
	fmt.Println(Tanh(10))
}
