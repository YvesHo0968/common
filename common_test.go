package common

import (
	"encoding/base64"
	"fmt"
	"github.com/rs/zerolog"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestUuid(t *testing.T) {
	fmt.Println(Uuid())
}

func TestPath(t *testing.T) {
	fmt.Println(Path())
}

func TestSha1(t *testing.T) {
	fmt.Println(Sha1("123"))
}

func TestUniqueId(t *testing.T) {
	fmt.Println(UniqueId())
}

func TestBase64Encoded(t *testing.T) {
	fmt.Println(Base64Encoded("hello"))
}

func TestBase64Decode(t *testing.T) {
	fmt.Println(Base64Decode("aGVsbG8="))
}

func TestMd5File(t *testing.T) {
	fmt.Println(Md5File("/Volumes/DATA/镜像/CentOS-7-x86_64-Minimal-2009.iso"))
}

func TestSha1File(t *testing.T) {
	fmt.Println(Sha1File("/Volumes/DATA/镜像/CentOS-7-x86_64-Minimal-2009.iso"))
}

func TestStrToLower(t *testing.T) {
	fmt.Println(StrToLower("Hello"))
}

func TestUcWords(t *testing.T) {
	fmt.Println(UcWords("hello word"))
}

func TestUcFirst(t *testing.T) {
	fmt.Println(UcFirst("hello word"))
}

func TestLcFirst(t *testing.T) {
	fmt.Println(LcFirst("Hello word"))
}

func TestMdStrLen(t *testing.T) {
	fmt.Println(MdStrLen("hello中国"))
}

func TestStrContains(t *testing.T) {
	fmt.Println(StrContains("aaddegg", "aa"))
}

func TestStrRepeat(t *testing.T) {
	fmt.Println(StrRepeat("-=", 10))
}

func TestStrReplace(t *testing.T) {
	fmt.Println(StrReplace("body", "black", "<text text='body'>", 1))
}

func TestStrShuffle(t *testing.T) {
	fmt.Println(StrShuffle("123456"))
}

func TestStrToUpper(t *testing.T) {
	fmt.Println(StrToUpper("Hello"))
}

func TestRand(t *testing.T) {
	fmt.Println(Rand(1000, 9999))
}

func TestCeil(t *testing.T) {
	fmt.Println(Ceil(1.2))
}

func TestFloor(t *testing.T) {
	fmt.Println(Floor(1.9))
}

func TestRound(t *testing.T) {
	fmt.Println(Round(1.5))
}

func TestSleep(t *testing.T) {
	Sleep(1)
}

func TestUsleep(t *testing.T) {
	Usleep(1000000)
}

func TestGetHostName(t *testing.T) {
	fmt.Println(GetHostName())
}

func TestGetOS(t *testing.T) {
	fmt.Println(GetOS())
}

func TestGetArch(t *testing.T) {
	fmt.Println(GetArch())
}

func TestGetArchBit(t *testing.T) {
	fmt.Println(GetArchBit())
}

func TestGetCpuCores(t *testing.T) {
	fmt.Println(GetCpuCores())
}

func TestSetGoMaxProcs(t *testing.T) {
	fmt.Println(SetGoMaxProcs(0))
}

func TestLog(t *testing.T) {
	//config.InitLog()
	//
	//log := config.Log
	//
	//log.Info().Str("foo", "bar").Msg("Hello World")
	//log.Error().Str("foo", "bar").Msg("Hello World")
	//
	//log.Info().
	//	Str("foo", "bar").
	//	Dict("dict", zerolog.Dict().
	//		Str("bar", "baz").
	//		Int("n", 1),
	//	).Msg("hello world")
	//
	//err := errors.New("A repo man spends his life getting into tense situations")
	//
	//fmt.Println(err)
	//service := "myservice"
	//
	//log.Fatal().
	//	Err(err).
	//	Str("service", service).
	//	Msgf("Cannot start %s", service)
}

func TestLogDemo(t *testing.T) {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, TimeFormat: time.Stamp}
	consoleWriter.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	consoleWriter.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	consoleWriter.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	consoleWriter.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s;", i)
	}

	multi := zerolog.MultiLevelWriter(consoleWriter)
	Logg := zerolog.New(multi).With().Timestamp().Caller().Logger().Level(zerolog.DebugLevel)

	Logg.Printf("ddddd")
	Logg.Info().Str("foo", "bar").Msg("Hello World")
}

func TestStructToJson(t *testing.T) {
	//DataMap := map[string]int{"a": 1, "b": 2, "c": 3}
	type S struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
		Sex  string `json:"sex,omitempty"`
	}
	sList := []S{{
		Name: "小红",
		Age:  18,
		Sex:  "女",
	}, {
		Name: "小李",
		Age:  20,
		Sex:  "男",
	}, {
		Name: "夏龙",
		Age:  18,
		Sex:  "女",
	}}
	for i := 0; i <= 1000; i++ {
		fmt.Println(JsonEncode(sList))
	}
}

func TestJsonToStruct(t *testing.T) {
	str := `[{"name":"小红","age":18,"sex":"女"},{"name":"小李","age":20,"sex":"男"},{"name":"夏龙","age":18,"sex":"女"}]`

	type S struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
		Sex  string `json:"sex,omitempty"`
	}

	for i := 0; i <= 10000; i++ {
		var data []S
		JsonDecode(str, &data)
		fmt.Println(data)
	}
}

func TestSerialize(t *testing.T) {
	data := map[string]interface{}{"php": "世界上最好的语言"}

	//data := []int{1, 3}

	fmt.Println(Serialize(data))
}

func TestUnSerialize(t *testing.T) {
	str := `a:1:{s:3:"php";s:24:"世界上最好的语言";}`
	out, _ := UnSerialize(str)
	m := out.(map[string]interface{})
	fmt.Println(m["php"])
}

func TestUrlEncode(t *testing.T) {
	fmt.Println(UrlEncode("中国"))
}

func TestUrlDecode(t *testing.T) {
	fmt.Println(UrlDecode("%E4%B8%AD%E5%9B%BD"))
}

func TestSonyFlakeId(t *testing.T) {
	fmt.Println(SonyFlakeId())
	fmt.Println(SonyFlakeId())
	fmt.Println(SonyFlakeId())
	fmt.Println(SonyFlakeId())
}

func TestSnowflakeId(t *testing.T) {
	fmt.Println(SnowflakeId())
	fmt.Println(SnowflakeId())
	fmt.Println(SnowflakeId())
}

func TestSendEmail(t *testing.T) {
	s := SendEmailData{
		FormName: "Go邮箱测试",
		ToEmail:  []string{"11111@qq.com"},
		Subject:  "测试第三方 email 库",
		Text:     "",
		HTML:     "<h1>HTML 正文</h1>",
	}

	c := SmtpConfig{
		Username: "xxx@aliyun.com",
		Password: "Password",
		Host:     "smtpdm.aliyun.com",
		Port:     465,
		Tls:      true,
	}
	fmt.Println(SendEmail(s, c))
}

func TestImplode(t *testing.T) {
	dd := reflect.TypeOf([]string{"hello", "word"})
	fmt.Println(dd.String())
	fmt.Println(Implode(",", []string{"hello", "word"}))
}

func TestExplode(t *testing.T) {
	fmt.Println(Explode(",", "hello,word"))
}

func TestSubStr(t *testing.T) {
	fmt.Println(SubStr("1212", 0, 3))
}

func TestGetPid(t *testing.T) {
	fmt.Println(GetPid())
}

func TestGetPpid(t *testing.T) {
	fmt.Println(GetPpid())
}

func TestIsAdmin(t *testing.T) {
	fmt.Println(IsAdmin())
}

func TestGetGatewayIp(t *testing.T) {
	fmt.Println(GetGatewayIp())
}

func TestPkcs7(t *testing.T) {
	pkcs7Pad := Pkcs7Padding([]byte("121212"), 16)
	fmt.Println(pkcs7Pad)
	pkcs7UnPad := UnPkcs7Padding(pkcs7Pad)
	fmt.Println(pkcs7UnPad)
}

func TestEncryptAES(t *testing.T) {
	// 原始数据
	text := []byte("Hello world!")

	// 统一使用32 字节的密钥和16 字节的iv向量
	key := []byte("12345678123456781234567812345678")
	iv := []byte("1234567812345678")

	// ECB 模式测试
	b1, err := AesEncryptECB(text, key[:16])
	if err != nil {
		panic(err)
	}

	fmt.Printf("AES-128-ECB 加密结果：%s\n", base64.StdEncoding.EncodeToString(b1))
	b1, err = AesDecryptECB(b1, key[:16])
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-128-ECB 解密结果：%s\n", string(b1))

	b2, err := AesEncryptECB(text, key[:24])
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-192-ECB 加密结果：%s\n", base64.StdEncoding.EncodeToString(b2))
	b2, err = AesDecryptECB(b2, key[:24])
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-192-ECB 解密结果：%s\n", string(b2))

	b3, err := AesEncryptECB(text, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-256-ECB 加密结果：%s\n", base64.StdEncoding.EncodeToString(b3))
	b3, err = AesDecryptECB(b3, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-256-ECB 解密结果：%s\n", string(b3))

	// CBC 模式测试
	b4, err := AesEncryptCBC(text, key[:16], iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-128-CBC 加密结果：%s\n", base64.StdEncoding.EncodeToString(b4))
	b4, err = AesDecryptCBC(b4, key[:16], iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-128-CBC 解密结果：%s\n", string(b4))

	b5, err := AesEncryptCBC(text, key[:24], iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-192-CBC 加密结果：%s\n", base64.StdEncoding.EncodeToString(b5))
	b5, err = AesDecryptCBC(b5, key[:24], iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-192-CBC 解密结果：%s\n", string(b5))

	b6, err := AesEncryptCBC(text, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-256-CBC 加密结果：%s\n", base64.StdEncoding.EncodeToString(b6))
	b6, err = AesDecryptCBC(b6, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-256-CBC 解密结果：%s\n", string(b6))

	// CTR 模式测试
	b7, err := AesEncryptCTR(text, key[:16], iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-128-CTR 加密结果：%s\n", base64.StdEncoding.EncodeToString(b7))
	b7, err = AesDecryptCTR(b7, key[:16], iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-128-CTR 解密结果：%s\n", string(b7))

	b8, err := AesEncryptCTR(text, key[:24], iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-192-CTR 加密结果：%s\n", base64.StdEncoding.EncodeToString(b8))
	b8, err = AesDecryptCTR(b8, key[:24], iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-192-CTR 解密结果：%s\n", string(b8))

	b9, err := AesEncryptCTR(text, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-256-CTR 加密结果：%s\n", base64.StdEncoding.EncodeToString(b9))
	b9, err = AesDecryptCTR(b9, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-256-CTR 解密结果：%s\n", string(b9))

	b10, err := AesEncryptCFB(text, key[:16], iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-128-CFB 加密结果：%s\n", base64.StdEncoding.EncodeToString(b10))
	b10, err = AesDecryptCFB(b10, key[:16], iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-128-CFB 解密结果：%s\n", string(b10))

	b11, err := AesEncryptCFB(text, key[:24], iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-192-CFB 加密结果：%s\n", base64.StdEncoding.EncodeToString(b11))
	b11, err = AesDecryptCFB(b11, key[:24], iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-192-CFB 解密结果：%s\n", string(b11))

	b12, err := AesEncryptCFB(text, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-256-CFB 加密结果：%s\n", base64.StdEncoding.EncodeToString(b12))
	b12, err = AesDecryptCFB(b12, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-256-CFB 解密结果：%s\n", string(b12))

	b13, err := AesEncryptOFB(text, key[:16], iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-128-OFB 加密结果：%s\n", base64.StdEncoding.EncodeToString(b13))
	b13, err = AesDecryptOFB(b13, key[:16], iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES-128-OFB 解密结果：%s\n", string(b13))
}

func TestGetType(t *testing.T) {
	var str string = "Hello, World!"
	var num int = 123
	var flag bool = true
	fmt.Println(GetType(str), GetType(num), GetType(flag))

	data := map[string]interface{}{
		"key1": "value1",
		"key2": 123,
		"key3": true,
		"key4": []int{1, 2, 3, 4, 5},
		"key5": []string{"e", "e"},
		"key6": map[string]any{
			"key1": "value1",
			"key2": 123,
			"key3": true,
		},
	}

	for key, value := range data {
		valueType := GetType(value)
		fmt.Printf("Key: %s, Value: %v, Type: %s\n", key, value, valueType)
	}

}

func TestName(t *testing.T) {
	cmd := exec.Command("ls", "-lah")
	//cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	r, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	r, _ = simplifiedchinese.GBK.NewDecoder().Bytes(r)
	fmt.Println(string(r))
}
