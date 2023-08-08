package common

import (
	"fmt"
	"os"
	"testing"
)

func TestAddCSlashes(t *testing.T) {
	fmt.Println(AddCSlashes("test sa", "sa"))
}

func TestAddSlashes(t *testing.T) {
	fmt.Println(AddSlashes("Who's Peter Griffin?"))
}

func TestBin2Hex(t *testing.T) {
	fmt.Println(Bin2Hex("123"))
}

func TestChop(t *testing.T) {
	str := "Hello, World!      \t\n"
	charList := "\t\n"
	choppedStr := Chop(str, charList)
	fmt.Println(choppedStr) //
}

func TestChr(t *testing.T) {
	fmt.Println(Chr(52))
	fmt.Println(Chr(052))
	fmt.Println(Chr(0x52))
}

func TestChunkSplit(t *testing.T) {
	fmt.Println(ChunkSplit("Hello world!", 1, ","))
}

func TestCrc32(t *testing.T) {
	fmt.Println(Crc32("Hello World!"))
}

func TestEcho(t *testing.T) {
	Echo("Hello World!", "")
}

func TestExplode(t *testing.T) {
	fmt.Println(Explode(",", "hello,word"))
}

func TestFprintf(t *testing.T) {
	//file, err := os.Create("output.txt")
	//
	//if err != nil {
	//	fmt.Println("Failed to open file:", err)
	//	return
	//}
	//defer file.Close()
	//
	//fmt.Println(Fprintf(file, "Name: %s\n", "ddd"))
	fmt.Println(Fprintf(os.Stdout, "Name: %s\n", "ddd"))
}

func TestHex2Bin(t *testing.T) {
	fmt.Println(Hex2Bin("313233"))
}

func TestHtmlEntityDecode(t *testing.T) {
	fmt.Println(HtmlEntityDecode("&lt;&copy; W3CS&ccedil;h&deg;&deg;&brvbar;&sect;&gt;"))
	fmt.Println(HtmlEntityDecode("This is some &lt;b&gt;bold&lt;/b&gt; text."))
}

func TestHtmlEntities(t *testing.T) {
	fmt.Println(HtmlEntities("<div>Hello & World</div>"))
}

func TestHtmlSpecialCharsDecode(t *testing.T) {
	fmt.Println(HtmlSpecialCharsDecode("This is some &lt;b&gt;bold&lt;/b&gt; text."))
}

func TestHtmlSpecialChars(t *testing.T) {
	fmt.Println(HtmlSpecialChars("This is some <b>bold</b> text."))
}

func TestImplode(t *testing.T) {
	fmt.Println(Implode(",", []string{"hello", "word"}))
}

func TestStrVal(t *testing.T) {
	fmt.Println(StrVal("test"))
	fmt.Println(StrVal(1))
	fmt.Println(StrVal(true))
}
