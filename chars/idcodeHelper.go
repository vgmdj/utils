package chars

import (
	"fmt"
	"log"
	"time"

	_ "github.com/vgmdj/gb2260/gbdata" // gb2260 data
	"github.com/vgmdj/utils/area"
)

//Sex 性别
type Sex int

const (
	//Male 男
	Male Sex = 1
	//Female 女
	Female Sex = 2
)

func (s Sex) String() string {
	if s == Male {
		return "male"
	}

	return "female"
}

//CNString 中文输出
func (s Sex) CNString() string {
	if s == Male {
		return "男"
	}

	return "女"
}

var idCardAlpha = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2, 0}
var idCardCheckSum = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

//IDCard 身份证
type IDCard struct {
	CardNum string
	idNum   [18]byte
}

//ParseIDCard 身份证号解析
func ParseIDCard(idcard string) (*IDCard, error) {
	if len(idcard) == 15 {
		return ConvertOldIDCard(idcard), nil
	}

	res := new(IDCard)
	copy(res.idNum[:], idcard)
	if len(idcard) != 18 || res.idNum[17] != checkSum(res.idNum) {
		return res, fmt.Errorf("invalid id code")
	}

	res.CardNum = idcard
	return res, nil

}

//GetAge 年龄计算
func (i *IDCard) GetAge() int {
	bri := i.GetBirthday()
	now := time.Now()
	age := now.Year() - bri.Year()

	if (bri.Month() == now.Month() && bri.Day() <= now.Day()) ||
		bri.Month() < now.Month() || age == 0 {
		return age
	}

	return age - 1
}

//GetSex 性别
func (i *IDCard) GetSex() Sex {
	sex := ToInt(i.idNum[16:17])
	if sex%2 == 0 {
		return Female
	}

	return Male
}

//GetBirthday 生日
func (i *IDCard) GetBirthday() time.Time {
	year := ToInt(i.idNum[6:10])
	month := ToInt(i.idNum[10:12])
	day := ToInt(i.idNum[12:14])

	china, _ := time.LoadLocation("PRC")
	bri := time.Date(year, time.Month(month), day, 0, 0, 0, 0, china)

	return bri
}

//GetLastFour 后四位识别码
func (i *IDCard) GetLastFour() string {
	return i.CardNum[18-4:]
}

//GetPlaceOfBirth 出生地
func (i *IDCard) GetPlaceOfBirth() *area.Info {
	gb2260 := area.NewArea(area.GB2260)
	gb2260.SetRevision(i.CardNum[6:10])
	return gb2260.Get(i.CardNum[:6])
}

//ConvertOldIDCard 旧15位转成18位
func ConvertOldIDCard(idcard string) *IDCard {
	res := new(IDCard)
	copy(res.idNum[0:6], idcard[0:6])
	copy(res.idNum[8:17], idcard[6:15])
	if res.idNum[8] >= 2 {
		copy(res.idNum[6:8], []byte{49, 57})
	}
	res.idNum[17] = checkSum(res.idNum)
	res.CardNum = string(res.idNum[:])
	return res
}

//CheckIDCode 判断身份证号是否合法，返回true合法，返回false不合法
func CheckIDCode(idcode string) bool {
	length := len(idcode)
	if length != 15 && length != 18 {
		return false
	}
	if length == 18 {
		var bytes [18]byte
		copy(bytes[:], idcode)
		chk := checkSum(bytes)
		return (bytes[17] != 'x' && bytes[17] == chk) || (bytes[17] == 'x' && chk == 'X')
	}
	//默认返回错误
	return false
}

//GetIDCodeCheckSum 计算最后一位校验位
func GetIDCodeCheckSum(idcode string) byte {
	if len(idcode) != 18 && len(idcode) != 17 {
		log.Println("invalid length of idcode")
		return ' '
	}

	sum := 0
	for k, v := range idcode {
		sum += (int(v) - 48) * idCardAlpha[k]
	}
	i := sum % 11
	return idCardCheckSum[i]
}

func checkSum(idnum [18]byte) byte {
	sum := 0
	for k, v := range idnum {
		sum += (int(v) - 48) * idCardAlpha[k]
	}
	i := sum % 11
	return idCardCheckSum[i]
}
