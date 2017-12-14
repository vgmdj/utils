package characters

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	_ "errors"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	_ "strconv"
	"strings"
	_ "strings"
	"sync"
	"time"
)

var (
	defaultRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Creates a random string based on a variety of options, using
// supplied source of randomness.
//
// If start and end are both 0, start and end are set
// to ' ' and 'z', the ASCII printable
// characters, will be used, unless letters and numbers are both
// false, in which case, start and end are set to 0 and math.MaxInt32.
//
// If set is not nil, characters between start and end are chosen.
//
// This method accepts a user-supplied rand.Rand
// instance to use as a source of randomness. By seeding a single
// rand.Rand instance with a fixed seed and using it for each call,
// the same random sequence of strings can be generated repeatedly
// and predictably.
func RandomSpec0(count uint, start, end int, letters, numbers bool,
	chars []rune, rand *rand.Rand) string {
	if count == 0 {
		return ""
	}
	if start == 0 && end == 0 {
		end = 'z' + 1
		start = ' '
		if !letters && !numbers {
			start = 0
			end = math.MaxInt32
		}
	}
	buffer := make([]rune, count)
	gap := end - start
	for count != 0 {
		count--
		var ch rune
		if len(chars) == 0 {
			ch = rune(rand.Intn(gap) + start)
		} else {
			ch = chars[rand.Intn(gap)+start]
		}
		if letters && ((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')) ||
			numbers && (ch >= '0' && ch <= '9') ||
			(!letters && !numbers) {
			if ch >= rune(56320) && ch <= rune(57343) {
				if count == 0 {
					count++
				} else {
					buffer[count] = ch
					count--
					buffer[count] = rune(55296 + rand.Intn(128))
				}
			} else if ch >= rune(55296) && ch <= rune(56191) {
				if count == 0 {
					count++
				} else {
					// high surrogate, insert low surrogate before putting it in
					buffer[count] = rune(56320 + rand.Intn(128))
					count--
					buffer[count] = ch
				}
			} else if ch >= rune(56192) && ch <= rune(56319) {
				// private high surrogate, no effing clue, so skip it
				count++
			} else {
				buffer[count] = ch
			}
		} else {
			count++
		}
	}
	return string(buffer)
}

// Creates a random string whose length is the number of characters specified.
//
// Characters will be chosen from the set of alpha-numeric
// characters as indicated by the arguments.
//
// Param count - the length of random string to create
// Param start - the position in set of chars to start at
// Param end   - the position in set of chars to end before
// Param letters - if true, generated string will include
//                 alphabetic characters
// Param numbers - if true, generated string will include
//                 numeric characters
func RandomSpec1(count uint, start, end int, letters, numbers bool) string {
	return RandomSpec0(count, start, end, letters, numbers, nil, defaultRand)
}

// Creates a random string whose length is the number of characters specified.
//
// Characters will be chosen from the set of alpha-numeric
// characters as indicated by the arguments.
//
// Param count - the length of random string to create
// Param letters - if true, generated string will include
//                 alphabetic characters
// Param numbers - if true, generated string will include
//                 numeric characters
func RandomAlphaOrNumeric(count uint, letters, numbers bool) string {
	return RandomSpec1(count, 0, 0, letters, numbers)
}

func RandomString(count uint) string {
	return RandomAlphaOrNumeric(count, false, false)
}

func RandomStringSpec0(count uint, set []rune) string {
	return RandomSpec0(count, 0, len(set)-1, false, false, set, defaultRand)
}

func RandomStringSpec1(count uint, set string) string {
	return RandomStringSpec0(count, []rune(set))
}

// Creates a random string whose length is the number of characters
// specified.
//
// Characters will be chosen from the set of characters whose
// ASCII value is between 32 and 126 (inclusive).
func RandomAscii(count uint) string {
	return RandomSpec1(count, 32, 127, false, false)
}

// Creates a random string whose length is the number of characters specified.
// Characters will be chosen from the set of alphabetic characters.
func RandomAlphabetic(count uint) string {
	return RandomAlphaOrNumeric(count, true, false)
}

// Creates a random string whose length is the number of characters specified.
// Characters will be chosen from the set of alpha-numeric characters.
func RandomAlphanumeric(count uint) string {
	return RandomAlphaOrNumeric(count, true, true)
}

// Creates a random string whose length is the number of characters specified.
// Characters will be chosen from the set of numeric characters.
func RandomNumeric(count uint) string {
	return RandomAlphaOrNumeric(count, false, true)
}

// 生成6位随机数
func RandomNumString() string {
	Count := uint(6)
	return RandomAlphaOrNumeric(Count, false, true)
}

//CheckFileIsExist 判断文件是否存在，存在返回true，不存在返回false
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//CopyFileTo 从指定位置复制创建出一个新的文件
//需先调用CheckFileIsExist判断文件是否存在
func CopyFileTo(srcFile string, destFile string) (err error) {
	var (
		fileContent []byte //temp file content
	)

	if fileContent, err = ioutil.ReadFile(srcFile); err != nil {
		log.Println("读取srcFile文件出错", srcFile)
		return
	}

	if err = ioutil.WriteFile(destFile, fileContent, 0666); err != nil {
		log.Println("写入destFile文件出错", destFile)
		return
	}

	return
}

//ParseFileTo 将文件中的内容解析到给出结构体中去
func ParseFileTo(filename string, parseStruct interface{}, mutex *sync.Mutex) (err error) {
	mutex.Lock()
	defer mutex.Unlock()

	var (
		fileContent []byte //temp file content
	)

	if fileContent, err = ioutil.ReadFile(filename); err != nil {
		log.Printf("读取%s文件时出错\n", filename)
		return
	}

	if err = json.Unmarshal(fileContent, parseStruct); err != nil {
		log.Println("解析文件内容时出错", string(fileContent))
		return
	}

	return
}

//WriteContentTo 根据传来内容写入目标文件
func WriteContentTo(destFilePath string, toWriteContent []byte, mutex *sync.Mutex) (err error) {
	mutex.Lock()
	defer mutex.Unlock()
	if err = ioutil.WriteFile(destFilePath, toWriteContent, 0666); err != nil {
		log.Printf("write %s err  : "+err.Error(), destFilePath)
		return
	}
	return
}

//将给定的str，从offset位开始，一直替换limit位的repStr
func HideString(str string, offset int, limit int, repStr string) (result string, err error) {
	var (
		temp     []byte
		strlen   = len(str)
		endPoint = offset + limit - 1
		repByte  byte
	)

	if strlen < 2 || strlen-1 < endPoint || len(repStr) != 1 {
		err = errors.New("字符长度不合法或函数使用错误")
		log.Println(str)
		return
	}

	repByte = repStr[0]
	for i := 0; i < strlen; i++ {
		if i < offset || i > endPoint {
			temp = append(temp, str[i])
			continue
		}
		temp = append(temp, repByte)
	}
	result = string(temp)

	return
}

func IsIntContain(char int, strArr []int) bool {
	if len(strArr) == 0 {
		return false
	}

	for _, v := range strArr {
		if char == v {
			return true
		}
	}

	return false

}

func IsStringContain(char string, strArr []string) bool {
	if len(strArr) == 0 {
		return false
	}

	for _, v := range strArr {
		if char == v {
			return true
		}
	}

	return false

}

func DBBytesToInt(b []byte) int {
	temp, err := strconv.Atoi(string(b))
	if err != nil {
		log.Println(err.Error(), string(b))
		return 0
	}

	return temp
}

func StringToBool(str string) bool {
	if str == "1" || str == "true" {
		return true
	}

	return false
}

func ConvertArrayMapBytesToString(dbMaps []map[string][]byte) (maps []map[string]string) {
	maps = make([]map[string]string, 0)
	for _, dbmap := range dbMaps {
		strMap := make(map[string]string)
		for k, v := range dbmap {
			strMap[k] = string(v)
		}
		maps = append(maps, strMap)
	}

	return

}

func ConvertMapBytesToString(dbMaps map[string][]byte) (maps map[string]string) {
	maps = make(map[string]string)

	for k, v := range dbMaps {
		maps[k] = string(v)
	}

	return

}

func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}

func PrepareData(strQuery string, strRep ...string) (data string) {
	strRep = append(strRep, " ")

	query := strings.Split(strQuery, "?")
	lenQuery := len(query)
	lenReq := len(strRep)
	if lenQuery == 0 || lenReq == 0 {
		return strQuery
	}

	if lenQuery != lenReq {
		log.Println("数量对应错误", lenQuery, lenReq)
		return strQuery
	}

	for k, v := range query {
		data += v
		data += strRep[k]
	}

	return
}

func CheckInt(strs ...string) bool {
	for _, v := range strs {
		match, err := regexp.Match(`^[0-9]*$`, []byte(v))
		if err != nil || !match {
			log.Println(err, match, v)
			return false
		}
	}

	return true
}

func ParseMoney(strm string) int64 {
	f32m, err := strconv.ParseFloat(strm, 32)
	if err != nil {
		log.Println(err.Error())
	}

	ff32m := f32m * 100

	return int64(ff32m)
}

func FormatMoney(i64m int64) string {
	ff32m := float64(i64m / 100)

	return strconv.FormatFloat(ff32m, 'f', 2, 32)

}

func FormatTime(t time.Time) string {
	china, _ := time.LoadLocation("PRC")
	cstTime := t.In(china).Format("20060102150405")

	return cstTime
}

func TranslateFee(strFee string) int {
	fee, err := strconv.Atoi(strFee)
	if err != nil {
		log.Println(err.Error())
		return -1
	}

	return fee

}

func StringInc(str string) string {
	i, err := strconv.Atoi(str)
	if err != nil {
		return "1"
	}

	return strconv.Itoa(i + 1)

}

func StoreStrMoney(str string) int {
	temp, err := strconv.Atoi(str)
	if err != nil {
		log.Println(err.Error())
		return -1
	}

	return temp * 100

}

func ShowStrMoney(str string) string {
	temp, _ := strconv.Atoi(str)

	temp = temp / 100

	return strconv.Itoa(temp)

}
