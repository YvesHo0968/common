package common

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"hash/crc32"
	"html"
	"io"
	r "math/rand"
	"mime/quotedprintable"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// AddCSlashes 返回在指定的字符前添加反斜杠的字符串
func AddCSlashes(str, characters string) string {
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
func Explode(sep, str string) []string {
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

// LcFirst 首字母小写
func LcFirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToLower(v))
		return u + str[len(u):]
	}
	return ""
}

// Ltrim 移除字符串左侧的字符
func Ltrim(str, chars string) string {
	return strings.TrimLeft(str, chars)
}

// Md5 生成32位md5字串
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5File 文件MD5
func Md5File(fileName string) string {
	file, err := os.Open(fileName)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if err != nil {
		return ""
	}
	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

// Nl2Br 在字符串中的每个新行之前插入 HTML 换行符
func Nl2Br(s string) string {
	return strings.ReplaceAll(s, "\n", "<br>")
}

// NumberFormat 通过千位分组来格式化数字
func NumberFormat(number float64, decimals int, decimalSep, thousandsSep string) string {
	// 将浮点数转换为字符串，保留指定小数位数
	numberStr := strconv.FormatFloat(number, 'f', decimals, 64)
	// 分割整数和小数部分
	parts := strings.Split(numberStr, ".")
	// 添加千分位分隔符
	if thousandsSep != "" && len(parts[0]) > 3 {
		for i := len(parts[0]) - 3; i > 0; i -= 3 {
			parts[0] = parts[0][:i] + thousandsSep + parts[0][i:]
		}
	}
	// 组合整数和小数部分
	result := parts[0]
	if decimals > 0 {
		result += decimalSep + parts[1]
	}
	return result
}

// Ord 返回字符串中第一个字符的 ASCII 值
func Ord(char string) int {
	bytes := []byte(char)
	if len(bytes) > 0 {
		return int(bytes[0])
	}
	return -1
}

// MbOrd 返回字符串中第一个字符的 ASCII 值
func MbOrd(char string) int {
	runes := []rune(char)
	if len(runes) > 0 {
		return int(runes[0])
	}
	return -1
}

// Print 输出一个或多个变量
func Print(a ...any) {
	fmt.Print(a...)
}

// Printf 输出格式化的字符串
func Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

// QuotedPrintableEncode 把 8 位字符串转换为 quoted-printable 字符串
func QuotedPrintableEncode(str string) string {
	var builder strings.Builder
	writer := quotedprintable.NewWriter(&builder)
	_, _ = writer.Write([]byte(str))
	defer func(writer *quotedprintable.Writer) {
		_ = writer.Close()
	}(writer)
	return builder.String()
}

// QuotedPrintableDecode 把 quoted-printable 字符串转换为 8 位字符串
func QuotedPrintableDecode(encoded string) string {
	b, err := io.ReadAll(quotedprintable.NewReader(strings.NewReader(encoded)))
	if err != nil {
		return ""
	}
	return string(b)
}

// Rtrim 移除字符串左侧的字符
func Rtrim(str, chars string) string {
	return strings.TrimRight(str, chars)
}

// Sha1 生成sha1字串 sha1("123")
func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1File 文件sha1
func Sha1File(fileName string) string {
	file, err := os.Open(fileName)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if err != nil {
		return ""
	}
	h := sha1.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

// SimilarText 计算两个字符串的相似度
func SimilarText(str1, str2 string) int {
	m := len(str1)
	n := len(str2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 0
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = 0
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[m][n]
}

// Soundex 计算字符串的 soundex 键
func Soundex(word string) string {
	soundexMap := map[rune]int{
		'A': 0, 'E': 0, 'I': 0, 'O': 0, 'U': 0, 'Y': 0,
		'H': 0, 'W': 0,
		'B': 1, 'F': 1, 'P': 1, 'V': 1,
		'C': 2, 'G': 2, 'J': 2, 'K': 2, 'Q': 2, 'S': 2, 'X': 2, 'Z': 2,
		'D': 3, 'T': 3,
		'L': 4,
		'M': 5, 'N': 5,
		'R': 6,
	}
	word = strings.ToUpper(word)
	soundexCode := string(word[0])
	for i := 1; i < len(word); i++ {
		c := word[i]
		if val, ok := soundexMap[rune(c)]; ok {
			if val != soundexMap[rune(word[i-1])] {
				soundexCode += fmt.Sprintf("%d", val)
			}
		}
	}
	soundexCode = soundexCode + "000"
	soundexCode = soundexCode[:4]
	return soundexCode
}

// Sprintf 把格式化的字符串写入一个变量中
func Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

// StrGetCsv 把 CSV 字符串解析到数组中
func StrGetCsv(csvString string) []string {
	reader := strings.NewReader(csvString)
	var values []string
	var insideQuotes bool
	var currentVal string
	for {
		runes, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if runes == '"' { // Check for opening or closing quotes
			insideQuotes = !insideQuotes
		} else if runes == ',' && !insideQuotes { // Separator found outside quotes
			values = append(values, currentVal)
			currentVal = ""
		} else {
			currentVal += string(runes) // Add character to current value
		}
	}
	if len(currentVal) > 0 { // Add the last value
		values = append(values, currentVal)
	}
	return values
}

// StrIReplace 替换字符串中的一些字符（大小写不敏感）
func StrIReplace(search, replace, subject string) string {
	return regexp.MustCompile("(?i)"+search).ReplaceAllString(subject, replace)
}

// StrPad 把字符串填充为新的长度
func StrPad(input string, padLength int, padString, padType string) string {
	inputLength := len(input)
	if padLength <= inputLength {
		return input
	}
	padStringLen := len(padString)
	difference := padLength - inputLength
	switch padType {
	case "STR_PAD_LEFT":
		padCount := difference / padStringLen
		padRemainder := difference % padStringLen
		leftPad := strings.Repeat(padString, padCount) + padString[:padRemainder]
		return leftPad + input
	case "STR_PAD_RIGHT":
		padCount := difference / padStringLen
		padRemainder := difference % padStringLen
		rightPad := padString[:padRemainder] + strings.Repeat(padString, padCount)
		return input + rightPad
	case "STR_PAD_BOTH":
		padCount := difference / (padStringLen * 2)
		padRemainder := difference % (padStringLen * 2)
		leftPad := strings.Repeat(padString, padCount+1)
		rightPad := padString[:padRemainder] + strings.Repeat(padString, padCount)
		return leftPad[:difference/2] + input + rightPad[:difference/2]
	default:
		return input
	}
}

// StrRepeat 把字符串重复指定的次数
func StrRepeat(input string, repeatCount int) string {
	return strings.Repeat(input, repeatCount)
}

// StrReplace 替换字符串中的一些字符（大小写敏感）
func StrReplace(search, replace, subject string, count int) string {
	return strings.Replace(subject, search, replace, count)
}

// StrRot13 对字符串执行 ROT13 编码
func StrRot13(str string) string {
	var output strings.Builder
	for _, char := range str {
		switch {
		case char >= 'A' && char <= 'M', char >= 'a' && char <= 'm':
			output.WriteRune(char + 13)
		case char >= 'N' && char <= 'Z', char >= 'n' && char <= 'z':
			output.WriteRune(char - 13)
		default:
			output.WriteRune(char)
		}
	}
	return output.String()
}

// StrShuffle 随机地打乱字符串中的所有字符
func StrShuffle(str string) string {
	runes := []rune(str)
	randData := r.New(r.NewSource(time.Now().UnixNano()))
	s := make([]rune, len(runes))
	for i, v := range randData.Perm(len(runes)) {
		s[i] = runes[v]
	}
	return string(s)
}

// StrSplit 把字符串分割到数组中
func StrSplit(str string, length int) []string {
	var result []string
	for len(str) > length {
		result = append(result, str[:length])
		str = str[length:]
	}
	if len(str) > 0 {
		result = append(result, str)
	}
	return result
}

// StrWordCount 计算字符串中的单词数
func StrWordCount(s string) int {
	words := strings.Fields(s)
	count := len(words)
	return count
}

// StrCaseCmp 比较两个字符串（大小写不敏感）
func StrCaseCmp(s1, s2 string) int {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	switch {
	case s1 < s2:
		return -1
	case s1 > s2:
		return 1
	default:
		return 0
	}
}

// StrChr 查找字符串在另一字符串中的第一次出现
func StrChr(s, substr string) string {
	index := StrPos(s, substr)
	if index == -1 {
		return ""
	}
	return s[index:]
}

// StrCmp 比较两个字符串（大小写敏感）
func StrCmp(s1, s2 string) int {
	return strings.Compare(s1, s2)
}

// StrCSpn 返回在找到任何指定的字符之前，在字符串查找的字符数
func StrCSpn(str, chars string) int {
	return strings.IndexAny(str, chars)
}

// StripTags 剥去字符串中的 HTML 标签
func StripTags(html string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	stripped := re.ReplaceAllString(html, "")
	return stripped
}

// StripCSlashes 删除由 AddCSlashes() 函数添加的反斜杠
func StripCSlashes(str string) string {
	// 特殊转义字符映射
	escapeMap := map[string]string{
		`\\`: `\`,
		`\"`: `"`,
		`\'`: "'",
		`\n`: "\n",
		`\r`: "\r",
		`\t`: "\t",
		`\f`: "\f",
		`\b`: "\b",
	}
	// 替换特殊转义字符
	for k, v := range escapeMap {
		str = strings.ReplaceAll(str, k, v)
	}
	// 使用 strconv.Unquote 解析转义字符
	unquoted, err := strconv.Unquote(`"` + str + `"`)
	if err != nil {
		// 解析转义序列失败，返回原始字符串
		return str
	}
	return unquoted
}

// StrIpOs 返回字符串在另一字符串中第一次出现的位置（大小写不敏感）
func StrIpOs(haystack, needle string) int {
	haystack = strings.ToLower(haystack)
	needle = strings.ToLower(needle)
	return StrPos(haystack, needle)
}

// StrIStr 查找字符串在另一字符串中第一次出现的位置（大小写不敏感）
func StrIStr(haystack, needle string) string {
	lowerHaystack := strings.ToLower(haystack)
	lowerNeedle := strings.ToLower(needle)
	index := StrPos(lowerHaystack, lowerNeedle)
	if index == -1 {
		return ""
	}
	substring := haystack[index : index+len(needle)]
	return substring
}

// StrLen 返回字符串的长度
func StrLen(str string) int {
	return len(str)
}

// MdStrLen 字符串长度
func MdStrLen(str string) int {
	return len([]rune(str))
}

// StrNatCaseCmp 使用一种"自然排序"算法来比较两个字符串（大小写不敏感）
func StrNatCaseCmp(s1, s2 string) int {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	return StrNatCmp(s1, s2)
}

// StrNatCmp 使用一种"自然排序"算法来比较两个字符串（大小写敏感）
func StrNatCmp(s1, s2 string) int {
	i, j, n1, n2 := 0, 0, len(s1), len(s2)
	for i < n1 && j < n2 {
		// 获取当前字符
		c1 := s1[i]
		c2 := s2[j]
		// 如果是数字字符，则按数字进行比较
		if c1 >= '0' && c1 <= '9' && c2 >= '0' && c2 <= '9' {
			num1 := 0
			num2 := 0
			// 提取数字
			for i < n1 && s1[i] >= '0' && s1[i] <= '9' {
				num1 = num1*10 + int(s1[i]-'0')
				i++
			}
			for j < n2 && s2[j] >= '0' && s2[j] <= '9' {
				num2 = num2*10 + int(s2[j]-'0')
				j++
			}
			// 比较数字大小
			if num1 < num2 {
				return -1
			} else if num1 > num2 {
				return 1
			}
		} else if c1 != c2 { // 如果是字母字符，则按字母的ASCII码进行比较
			if c1 < c2 {
				return -1
			} else {
				return 1
			}
		}
		// 移动指针
		i++
		j++
	}
	// 如果两个字符串长度不同，则较长的字符串较大
	if n1 < n2 {
		return -1
	} else if n1 > n2 {
		return 1
	}
	// 如果两个字符串完全相同，则返回0
	return 0
}

// StrNCaseCmp 前 n 个字符的字符串比较（大小写不敏感）
func StrNCaseCmp(s1, s2 string, n int) int {
	s1 = strings.ToLower(s1[:n])
	s2 = strings.ToLower(s2[:n])

	return StrNatCmp(s1, s2)
}

// StrNCmp 前 n 个字符的字符串比较（大小写敏感）
func StrNCmp(s1, s2 string, n int) int {
	s1 = s1[:n]
	s2 = s2[:n]
	return StrNatCmp(s1, s2)
}

// StrPBrk 在字符串中搜索指定字符中的任意一个
func StrPBrk(str, char string) string {
	index := StrCSpn(str, char)
	if index >= 0 {
		return str[index:]
	} else {
		return ""
	}
}

// StrPos 字符串在另一字符串中第一次出现的位置（大小写敏感）
func StrPos(haystack, needle string) int {
	return strings.Index(haystack, needle)
}

// StrRChr 查找字符串在另一个字符串中最后一次出现
func StrRChr(str, char string) string {
	index := StrRPos(str, char)
	if index >= 0 {
		return str[index:]
	} else {
		return ""
	}
}

// StRRev 反转字符串
func StRRev(str string) string {
	runes := []rune(str)
	length := len(runes)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// StrRIPos 查找字符串在另一字符串中最后一次出现的位置(大小写不敏感)
func StrRIPos(haystack, needle string) int {
	haystack = strings.ToLower(haystack)
	needle = strings.ToLower(needle)
	return StrRPos(haystack, needle)
}

// StrRPos 查找字符串在另一字符串中最后一次出现的位置(大小写敏感)
func StrRPos(haystack, needle string) int {
	return strings.LastIndex(haystack, needle)
}

// StrSpn 返回在字符串中包含的特定字符的数目
func StrSpn(str, accept string) int {
	return strings.IndexFunc(str, func(r rune) bool {
		return !strings.ContainsRune(accept, r)
	})
}

// StrStr 查找字符串在另一字符串中的第一次出现（大小写敏感）
func StrStr(haystack, needle string) string {
	index := StrPos(haystack, needle)
	if index == -1 {
		return ""
	}
	return haystack[index:]
}

// StrTok 把字符串分割为更小的字符串
func StrTok(str, sep string) (string, string) {
	tokens := Explode(sep, str)
	if len(tokens) <= 1 {
		return "", ""
	}
	remaining := strings.Join(tokens[1:], sep)
	return tokens[0], remaining
}

// StrToLower 字符转小写
func StrToLower(str string) string {
	return strings.ToLower(str)
}

// StrToUpper 字符转大写
func StrToUpper(str string) string {
	return strings.ToUpper(str)
}

// StrTr 转换字符串中特定的字符
func StrTr(str string, replacements map[string]string) string {
	for old, newStr := range replacements {
		str = strings.ReplaceAll(str, old, newStr)
	}
	return str
}

// SubStr 字符串裁剪
func SubStr(str string, start int, l ...int) string {
	runes := []rune(str)
	length := len(runes)
	if len(l) > 0 {
		length = l[0]
	}

	if start < 0 {
		start = len(runes) + start
	} else if start > len(runes) {
		start = len(runes)
	}

	end := start + length
	if end < 0 {
		end = len(runes) + end
	} else if end > len(runes) {
		end = len(runes)
	}

	if length < 0 {
		length = 0
	} else if end > len(runes) {
		end = len(runes)
	}
	return string(runes[start:end])
}

// MdSubStr 返回中文字符串的一部分
func MdSubStr(str string, start int, length ...int) string {
	return SubStr(str, start, length...)
}

// SubstrCompare 从指定的开始位置比较两个字符串
func SubstrCompare(str1, str2 string, startPos, length int, b ...bool) int {
	substr1 := ""
	if startPos >= 0 && length > 0 {
		substr1 = SubStr(str1, startPos, length)
	}

	isTrue := true
	if len(b) > 0 {
		isTrue = b[0]
	}
	if isTrue {
		substr1 = StrToLower(substr1)
		str2 = StrToLower(str2)
	}
	return StrCmp(substr1, str2)
}

// SubstrCount 计算子串在字符串中出现的次数
func SubstrCount(str, substr string) int {
	count := strings.Count(str, substr)
	return count
}

// SubstrReplace 把字符串的一部分替换为另一个字符串
func SubstrReplace(str, repl string, start int, l ...int) string {
	runes := []rune(str)

	length := len(runes)
	if len(l) > 0 {
		length = l[0]
	}

	if start < 0 {
		start = len(runes) + start
	} else if start > len(runes) {
		start = len(runes)
	}

	if length < 0 {
		length = (len(runes) + length) - start
	}

	end := start + length

	if end < 0 {
		end = 0
	} else if end > len(runes) {
		end = len(runes)
	}
	return str[:start] + repl + str[end:]
}

func Trim(str string) string {
	start := 0
	end := len(str) - 1
	// 从左边开始找到第一个非空白字符的索引
	for start <= end && (str[start] == ' ' || str[start] == '\t' || str[start] == '\n' || str[start] == '\r') {
		start++
	}
	// 从右边开始找到第一个非空白字符的索引
	for start <= end && (str[end] == ' ' || str[end] == '\t' || str[end] == '\n' || str[end] == '\r') {
		end--
	}
	if start > end {
		return ""
	}
	return str[start : end+1]
}

// UcFirst 首字母大写
func UcFirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}

// UcWords 把每个单词的首字符转换为大写
func UcWords(str string) string {
	return cases.Title(language.Und, cases.NoLower).String(str)
}

// VFprintf 把格式化的字符串写入到指定的输出流
func VFprintf(w *os.File, format string, a ...interface{}) int {
	return Fprintf(w, format, a...)
}

// VSprintf 把格式化字符串写入变量中
func VSprintf(format string, a ...any) string {
	return Sprintf(format, a...)
}

// Wordwrap 按照指定长度对字符串进行折行处理
func Wordwrap(str string, width int, breakChar string) string {
	words := strings.Fields(str)
	result := ""
	lineLength := 0
	for _, word := range words {
		wordLength := len(word)
		if lineLength+wordLength > width {
			result += breakChar
			lineLength = 0
		}
		result += word + " "
		lineLength += wordLength + 1
	}
	return strings.TrimRight(result, " ")
}

// StrVal 任意类型转字符串
func StrVal(data any) string {
	return fmt.Sprintf("%v", data)
}
