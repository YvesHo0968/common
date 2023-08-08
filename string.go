package common

import (
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"html"
	"os"
	"strings"
)

// AddCSlashes 返回在指定的字符前添加反斜杠的字符串
func AddCSlashes(str string, characters string) string {
	escapedStr := ""
	for _, ch := range str {
		chStr := string(ch)
		if strings.ContainsRune(characters, ch) {
			chStr = "\\" + chStr
		}
		escapedStr += chStr
	}
	return escapedStr
}

// AddSlashes 预定义的字符前添加反斜杠的字符串
func AddSlashes(str string) string {
	str = strings.ReplaceAll(str, `\`, `\\`)    // 转义反斜杠
	str = strings.ReplaceAll(str, `'`, `\'`)    // 转义单引号
	str = strings.ReplaceAll(str, `"`, `\"`)    // 转义双引号
	str = strings.ReplaceAll(str, "\x00", `\0`) // 转义 NULL 字符
	return str
}

// Bin2Hex 把 ASCII 字符的字符串转换为十六进制值
func Bin2Hex(str string) string {
	hexStr := ""
	for _, ch := range str {
		hexStr += fmt.Sprintf("%x", ch)
	}
	return hexStr
}

// Chop 移除字符串右侧的空白字符或其他预定义字符
func Chop(str string, charList ...string) string {
	if len(charList) == 0 {
		str = strings.TrimRight(str, "\x00")
		str = strings.TrimRight(str, "\t")
		str = strings.TrimRight(str, "\n")
		str = strings.TrimRight(str, "\x0B")
		str = strings.TrimRight(str, "\r")
		str = strings.TrimRight(str, " ")
		return str
	}

	return strings.TrimRight(str, charList[0])
}

// Chr 从指定 ASCII 值返回字符
func Chr(i int) string {
	return string(rune(i))
}

// ChunkSplit 把字符串分割为一连串更小的部分
func ChunkSplit(s string, length int, end string) string {
	var chunks []string
	for i := 0; i < len(s); i += length {
		endIndex := i + length
		if endIndex > len(s) {
			endIndex = len(s)
		}
		chunk := s[i:endIndex]
		chunks = append(chunks, chunk)
	}
	return strings.Join(chunks, end)
}

// Crc32 计算一个字符串的 32 位 CRC
func Crc32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

func Echo(values ...any) {
	str := make([]string, 0)
	for _, v := range values {
		str = append(str, StrVal(v))
	}

	fmt.Println(strings.Join(str, " "))
}

// Explode 字符转数组
func Explode(sep string, str string) []string {
	return strings.Split(str, sep)
}

// Fprintf 把格式化的字符串写入到指定的输出流
func Fprintf(w *os.File, format string, a ...interface{}) int {
	n, _ := fmt.Fprintf(w, format, a...)
	return n
}

// Hex2Bin 把十六进制值的字符串转换为 ASCII 字符
func Hex2Bin(str string) string {
	binaryData, err := hex.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(binaryData)
}

// HtmlEntityDecode 把 HTML 实体转换为字符。
func HtmlEntityDecode(str string) string {
	return html.UnescapeString(str)
}

// HtmlEntities 把字符转换为 HTML 实体
func HtmlEntities(text string) string {
	return html.EscapeString(text)
}

// HtmlSpecialCharsDecode 把预定义的 HTML 实体 "&lt;"（小于）和 "&gt;"（大于）转换为字符
func HtmlSpecialCharsDecode(str string) string {
	decodedString := strings.ReplaceAll(str, "&lt;", "<")
	decodedString = strings.ReplaceAll(decodedString, "&gt;", ">")
	decodedString = strings.ReplaceAll(decodedString, "&amp;", "&")
	decodedString = strings.ReplaceAll(decodedString, "&#039;", "'")
	decodedString = strings.ReplaceAll(decodedString, "&quot;", "\"")
	return decodedString
}

// HtmlSpecialChars 把预定义的字符 "<" （小于）和 ">" （大于）转换为 HTML 实体
func HtmlSpecialChars(str string) string {
	encodedString := strings.ReplaceAll(str, "&", "&amp;")
	encodedString = strings.ReplaceAll(encodedString, "<", "&lt;")
	encodedString = strings.ReplaceAll(encodedString, ">", "&gt;")
	encodedString = strings.ReplaceAll(encodedString, "'", "&#039;")
	encodedString = strings.ReplaceAll(encodedString, "\"", "&quot;")
	return encodedString
}

// Implode 数组转字符
func Implode(sep string, elems []string) string {
	return strings.Join(elems, sep)
}

// StrVal 任意类型转字符串
func StrVal(data any) string {
	return fmt.Sprintf("%v", data)
}
