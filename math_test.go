package common

import (
	"fmt"
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
