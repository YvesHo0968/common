package common

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"github.com/jackpal/gateway"
	"github.com/leeqvip/gophp/serialize"
	goCache "github.com/patrickmn/go-cache"
	"github.com/sony/sonyflake"
	"io"
	r "math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Uuid 获取uuid
func Uuid() string {
	return uuid.New().String()
}

// Path 获取进程工作目录
func Path() string {
	dir, _ := os.Getwd()

	return dir
}

// FilePath 获取运行的目录
func FilePath() string {
	_, filename, _, _ := runtime.Caller(0)

	root := path.Dir(path.Dir(filename))

	return root
}

// Time 获取系统时间戳
func Time() int64 {
	return time.Now().Unix()
}

// Date 时间戳转日期 Date("Y-m-d H:i:s")
func Date(format string, times ...int64) string {
	// DateFormat pattern rules.
	var datePatterns = []string{
		// year
		"Y", "2006", // A full numeric representation of a year, 4 digits   Examples: 1999 or 2003
		"y", "06", // A two digit representation of a year   Examples: 99 or 03

		// month
		"m", "01", // Numeric representation of a month, with leading zeros 01 through 12
		"n", "1", // Numeric representation of a month, without leading zeros   1 through 12
		"M", "Jan", // A short textual representation of a month, three letters Jan through Dec
		"F", "January", // A full textual representation of a month, such as January or March   January through December

		// day
		"d", "02", // Day of the month, 2 digits with leading zeros 01 to 31
		"j", "2", // Day of the month without leading zeros 1 to 31

		// week
		"D", "Mon", // A textual representation of a day, three letters Mon through Sun
		"l", "Monday", // A full textual representation of the day of the week  Sunday through Saturday

		// time
		"g", "3", // 12-hour format of an hour without leading zeros    1 through 12
		"G", "15", // 24-hour format of an hour without leading zeros   0 through 23
		"h", "03", // 12-hour format of an hour with leading zeros  01 through 12
		"H", "15", // 24-hour format of an hour with leading zeros  00 through 23

		"a", "pm", // Lowercase Ante meridiem and Post meridiem am or pm
		"A", "PM", // Uppercase Ante meridiem and Post meridiem AM or PM

		"i", "04", // Minutes with leading zeros    00 to 59
		"s", "05", // Seconds, with leading zeros   00 through 59

		// time zone
		"T", "MST",
		"P", "-07:00",
		"O", "-0700",

		// RFC 2822
		"r", time.RFC1123Z,
	}

	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)

	var t int64
	if len(times) > 0 {
		t = times[0]
	} else {
		t = Time()
	}

	return time.Unix(t, 0).Format(format)
}

// DateToTime 日期转时间戳 DateToTime("2006-01-02 15:04:05", "2022-01-01 11:00:00")
func DateToTime(format, date string) (int64, error) {
	t, err := time.ParseInLocation(format, date, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// StrToTime 文本日期时间解析为 Unix 时间戳
func StrToTime(str string) int64 {
	uintToSeconds := map[string]int64{"minute": 60, "hour": 3600, "day": 86400, "week": 604800, "year": (365 * 86400) + 86400}

	accumulator := time.Now().Unix()

	var delta int64
	plus := true
	str = strings.TrimSpace(str)

	if strings.HasPrefix(str, "in ") {
		str = strings.Replace(str, "in ", "", 1)
	}

	if strings.Index(str, " ago") > 0 {
		str = strings.Replace(str, " ago", "", 1)
		plus = false
	}

	if strings.Index(str, "+") >= 0 {
		str = strings.Replace(str, "+", "", 1)
	}

	if strings.Index(str, "-") >= 0 {
		str = strings.Replace(str, "-", "", 1)
		plus = false
	}

	noteValMap := make(map[string]int64, 10)

	mustCompileStr := `\d+\s+(minute|hour|day|week|year)`
	re := regexp.MustCompile(mustCompileStr)

	parts := re.FindAllStringSubmatch(str, -1)

	for i := range parts {
		strArray := strings.Split(parts[i][0], " ")
		v, _ := strconv.Atoi(strArray[0])
		noteValMap[parts[i][1]] = int64(v)
	}

	delta = 0
	for k, v := range noteValMap {

		delta += uintToSeconds[k] * v
	}

	if plus {
		accumulator += delta
	} else {
		accumulator -= delta
	}

	return accumulator
}

// UniqueId 生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(Base64Encoded(string(b)))
}

// Base64Encoded Base64加密 Base64Encoded("hello")
func Base64Encoded(str string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	return encoded
}

// Base64Decode Base64解密 Base64Decode("aGVsbG8=")
func Base64Decode(str string) string {
	decode, _ := base64.StdEncoding.DecodeString(str)
	return string(decode)
}

// Sleep 延迟执行秒数
func Sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

// Usleep 延迟执行微秒
func Usleep(microseconds int) {
	time.Sleep(time.Duration(microseconds) * time.Microsecond)
}

// JsonEncode 结构体转json
func JsonEncode(data interface{}) string {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Map转化为byte数组失败,异常:%s\n", err)
		return ""
	}
	return string(jsonByte)
}

// JsonDecode json转结构体
func JsonDecode(data string, val interface{}) error {
	return json.Unmarshal([]byte(data), val)
}

// Serialize 结构体转字符
func Serialize(data interface{}) string {
	jsonByte, err := serialize.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonByte)
}

// UnSerialize 字符转结构体
func UnSerialize(str string) (interface{}, error) {
	return serialize.UnMarshal([]byte(str))
}

// UrlEncode url序列化
func UrlEncode(str string) string {
	return url.QueryEscape(str)
}

// UrlDecode url反序列化
func UrlDecode(str string) string {
	decodeStr, err := url.QueryUnescape(str)
	if err != nil {
		return ""
	}
	return decodeStr
}

var sonyFlakeData = sonyflake.NewSonyflake(sonyflake.Settings{})

// SonyFlakeId 雪花分布式id
func SonyFlakeId() int {
	//t, _ := time.Parse("2006-01-02", "2021-01-01")
	//settings := sonyflake.Settings{
	//	//StartTime:      t,              // 起始时间，默认值为2014-09-01 00:00:00 +0000 UTC
	//	//MachineID:      getMachineID,   // 是一个返回实例 ID 的函数，如果不定义此函外，默认用本机ip 的低16位
	//	//CheckMachineID: checkMachineID, // 验证实例 ID / 计算机ID 的唯一性，返回true时才创建
	//}

	//sf := sonyflake.NewSonyflake(settings)

	id, _ := sonyFlakeData.NextID()

	return int(id)
}

var snowflakeData, _ = snowflake.NewNode(time.Now().UnixMilli() % 1024)

// SnowflakeId 推特雪花id
func SnowflakeId() int {
	id := snowflakeData.Generate()

	return int(id)
}

// Mail 发送邮箱
func Mail(user, password, userName, host, port, to, subject, body, mailType string, isTls bool) error {
	auth := smtp.PlainAuth("", user, password, host)
	var contentType = "Content-Type: text/plain; charset=UTF-8"
	
	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	}

	//msg := []byte("To: " + to + "\r\nFrom: " + userName + "<" + user + ">" + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	msg := []byte("From: " + userName + "<" + user + ">" + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(Trim(to), ";")

	if isTls {
		// 配置TLS连接
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true, // 可能需要根据你的邮件服务器配置进行更改
			ServerName:         host,
		}
		// 连接到邮件服务器
		conn, err := tls.Dial("tcp", host+":"+port, tlsConfig)
		if err != nil {
			return err
		}

		defer func(conn *tls.Conn) {
			_ = conn.Close()
		}(conn)

		client, err := smtp.NewClient(conn, host)
		if err != nil {
			return err
		}

		defer func(client *smtp.Client) {
			_ = client.Close()
		}(client)

		// 进行SMTP认证
		err = client.Auth(auth)
		if err != nil {
			return err
		}

		// 设置发件人和收件人
		err = client.Mail(user)
		if err != nil {
			return err
		}

		for _, toEmail := range sendTo {
			_ = client.Rcpt(toEmail)
		}

		// 发送邮件内容
		wc, err := client.Data()
		if err != nil {
			return err
		}

		defer func(wc io.WriteCloser) {
			_ = wc.Close()
		}(wc)

		_, err = wc.Write(msg)
		return err
	} else {
		return smtp.SendMail(host, auth, user, sendTo, msg)
	}
}

// GetType 获取遍历类型
func GetType(v interface{}) string {
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		return reflect.TypeOf(v).Elem().Kind().String()
	}
	return reflect.TypeOf(v).Kind().String()
}

// GetPid 获取进程id
func GetPid() int {
	return os.Getpid()
}

// GetPpid 获取父级进程id
func GetPpid() int {
	return os.Getppid()
}

// GetGatewayIp 获取网关ip
func GetGatewayIp() (string, error) {
	gw, err := gateway.DiscoverGateway()
	if err != nil {
		return "", err
	}

	return gw.String(), nil
}

// IsAdmin 是否admin用户
func IsAdmin() bool {
	return os.Getuid() == 0
}

type reducetype func(interface{}) interface{}
type filtertype func(interface{}) bool

// InSlice checks given string in string slice or not.
func InSlice(v string, sl []string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// InSliceIface checks given interface in interface slice.
func InSliceIface(v interface{}, sl []interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// SliceRandList generate an int slice from min to max.
func SliceRandList(min, max int) []int {
	if max < min {
		min, max = max, min
	}
	length := max - min + 1
	t0 := time.Now()
	r.New(r.NewSource(int64(t0.Nanosecond())))
	list := r.Perm(length)
	for index := range list {
		list[index] += min
	}
	return list
}

// SliceMerge merges interface slices to one slice.
func SliceMerge(slice1, slice2 []interface{}) (c []interface{}) {
	c = append(slice1, slice2...)
	return
}

// SliceReduce generates a new slice after parsing every value by reduce function
func SliceReduce(slice []interface{}, a reducetype) (dsLice []interface{}) {
	for _, v := range slice {
		dsLice = append(dsLice, a(v))
	}
	return
}

// SliceRand returns random one from slice.
func SliceRand(a []interface{}) (b interface{}) {
	randnum := r.Intn(len(a))
	b = a[randnum]
	return
}

// SliceSum sums all values in int64 slice.
func SliceSum(intslice []int64) (sum int64) {
	for _, v := range intslice {
		sum += v
	}
	return
}

// SliceFilter generates a new slice after filter function.
func SliceFilter(slice []interface{}, a filtertype) (ftslice []interface{}) {
	for _, v := range slice {
		if a(v) {
			ftslice = append(ftslice, v)
		}
	}
	return
}

// SliceDiff returns diff slice of slice1 - slice2.
func SliceDiff(slice1, slice2 []interface{}) (diffslice []interface{}) {
	for _, v := range slice1 {
		if !InSliceIface(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

// SliceIntersect returns slice that are present in all the slice1 and slice2.
func SliceIntersect(slice1, slice2 []interface{}) (diffslice []interface{}) {
	for _, v := range slice1 {
		if InSliceIface(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

// SliceChunk separates one slice to some sized slice.
func SliceChunk(slice []interface{}, size int) (chunkslice [][]interface{}) {
	if size >= len(slice) {
		chunkslice = append(chunkslice, slice)
		return
	}
	end := size
	for i := 0; i <= (len(slice) - size); i += size {
		chunkslice = append(chunkslice, slice[i:end])
		end += size
	}
	return
}

// SliceRange generates a new slice from begin to end with step duration of int64 number.
func SliceRange(start, end, step int64) (intslice []int64) {
	for i := start; i <= end; i += step {
		intslice = append(intslice, i)
	}
	return
}

// SlicePad prepends size number of val into slice.
func SlicePad(slice []interface{}, size int, val interface{}) []interface{} {
	if size <= len(slice) {
		return slice
	}
	for i := 0; i < (size - len(slice)); i++ {
		slice = append(slice, val)
	}
	return slice
}

// SliceUnique cleans repeated values in slice.
func SliceUnique(slice []interface{}) (uniqueslice []interface{}) {
	for _, v := range slice {
		if !InSliceIface(v, uniqueslice) {
			uniqueslice = append(uniqueslice, v)
		}
	}
	return
}

// SliceShuffle shuffles a slice.
func SliceShuffle(slice []interface{}) []interface{} {
	for i := 0; i < len(slice); i++ {
		a := r.Intn(len(slice))
		b := r.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}

var (
	Cache = goCache.New(5*time.Second, 5*time.Second)
)

// SetCache 设置缓存
func SetCache(key string, data interface{}, ct int) bool {
	Cache.Set(key, data, time.Duration(ct)*time.Second)

	return true
}

// GetCache 获取缓存
func GetCache(key string) interface{} {
	var data interface{}

	cacheData, found := Cache.Get(key)

	if found {
		data = cacheData
	}

	return data
}

// DeleteCache 删除缓存
func DeleteCache(key string) bool {
	Cache.Delete(key)

	return true
}

// CurlPost 请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func CurlPost(url string, data interface{}, contentType string) string {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr := JsonEncode(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := io.ReadAll(resp.Body)
	return string(result)
}

// Pkcs7Padding PKCS#7 Padding：对于要填充的字节数n，将n个字节都填充为n
func Pkcs7Padding(data []byte, blockSize int) []byte {
	pad := blockSize - len(data)%blockSize
	b := bytes.Repeat([]byte{byte(pad)}, pad)
	return append(data, b...)
}

// UnPkcs7Padding PKCS#7去除Padding
func UnPkcs7Padding(data []byte) []byte {
	pad := int(data[len(data)-1])

	if pad >= len(data) {
		return []byte{}
	}
	return data[:len(data)-pad]
}
