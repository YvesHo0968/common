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

func TestLcFirst(t *testing.T) {
	s := "Hello word"

	first := string(s[0])
	rest := s[1:]
	firstToLower := first + rest
	fmt.Println(firstToLower)
	fmt.Println(LcFirst("Hello word"))
}

func TestLtrim(t *testing.T) {
	fmt.Println(Ltrim("Hello World!", "Hello"))
}

func TestMd5(t *testing.T) {
	fmt.Println(Md5("123"))
}

func TestMd5File(t *testing.T) {
	fmt.Println(Md5File("./string.go"))
}

func TestNl2Br(t *testing.T) {
	fmt.Println(Nl2Br("One line.\nAnother line."))
}

func TestNumberFormat(t *testing.T) {
	fmt.Println(NumberFormat(1000000.1, 2, ".", ","))
}

func TestOrd(t *testing.T) {
	fmt.Println(Ord("我"))
	fmt.Println(Ord("h"))
}

func TestMbOrd(t *testing.T) {
	fmt.Println(MbOrd("我"))
	fmt.Println(MbOrd("h"))
}

func TestPrint(t *testing.T) {
	Print("hello word!")
}

func TestPrintf(t *testing.T) {
	Printf("hello %s", "word")
}

func TestQuotedPrintableEncode(t *testing.T) {
	fmt.Println(QuotedPrintableEncode("Hello world 我们"))
}

func TestQuotedPrintableDecode(t *testing.T) {
	fmt.Println(QuotedPrintableDecode("Hello world =E6=88=91=E4=BB=AC"))
}

func TestRtrim(t *testing.T) {
	fmt.Println(Rtrim("Hello World!", "World!"))
}

func TestSha1(t *testing.T) {
	fmt.Println(Sha1("123"))
}

func TestSha1File(t *testing.T) {
	fmt.Println(Sha1File("./string.go"))
}

func TestSimilarText(t *testing.T) {
	fmt.Println(SimilarText("Hello World", "Hello Peter"))
}

func TestSoundex(t *testing.T) {
	fmt.Println(Soundex("Apple"))
}

func TestSprintf(t *testing.T) {
	fmt.Println(Sprintf("Hello %s", "word!"))
}

func TestStrGetCsv(t *testing.T) {
	fmt.Println(StrGetCsv(`"John Doe",25,"New York"`))
}

func TestStrIReplace(t *testing.T) {
	fmt.Println(StrIReplace("WORLD", "Peter", "Hello world!"))
}

func TestStrPad(t *testing.T) {
	fmt.Println(StrPad("hello", 10, "*", "STR_PAD_LEFT"))
	fmt.Println(StrPad("hello", 10, "*", "STR_PAD_RIGHT"))
	fmt.Println(StrPad("hello", 10, "*", "STR_PAD_BOTH"))
}

func TestStrRepeat(t *testing.T) {
	fmt.Println(StrRepeat(".", 10))
}

func TestStrReplace(t *testing.T) {
	fmt.Println(StrReplace("body", "black", "<text text='body'>", 1))
}

func TestStrVal(t *testing.T) {
	fmt.Println(StrVal("test"))
	fmt.Println(StrVal(1))
	fmt.Println(StrVal(true))
}
