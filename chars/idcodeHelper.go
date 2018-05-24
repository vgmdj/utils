package chars

import (
	"fmt"
	"time"
)

type Sex int

const (
	Male   Sex = 1
	Female Sex = 2
)

func (s Sex) String() string {
	if s == Male {
		return "male"
	}

	return "female"
}

func (s Sex) CNString() string {
	if s == Male {
		return "男"
	}

	return "女"
}

var IdCardAlpha = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2, 0}
var IdCardCheckSum = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

type IdCard struct {
	CardNum string
	idNum   [18]byte
}

func ParseIdCard(idcard string) (*IdCard, error) {
	if len(idcard) == 15 {
		return ConvertOldIdCard(idcard), nil
	}

	res := new(IdCard)
	copy(res.idNum[:], idcard)
	if len(idcard) != 18 || res.idNum[17] != checkSum(res.idNum) {
		return res, fmt.Errorf("invalid id code")
	}

	res.CardNum = idcard
	return res, nil

}

func (i *IdCard) GetAge() int {
	bri := i.GetBirthday()
	now := time.Now()
	age := now.Year() - bri.Year()

	if (bri.Month() == now.Month() && bri.Day() <= now.Day()) ||
		bri.Month() < now.Month() || age == 0 {
		return age
	}

	return age - 1
}

func (i *IdCard) GetSex() Sex {
	sex := BytesToInt(i.idNum[16:17])
	if sex%2 == 0 {
		return Female
	}

	return Male
}

func (i *IdCard) GetBirthday() time.Time {
	year := BytesToInt(i.idNum[6:10])
	month := BytesToInt(i.idNum[10:12])
	day := BytesToInt(i.idNum[12:14])

	china, _ := time.LoadLocation("PRC")
	bri := time.Date(year, time.Month(month), day, 0, 0, 0, 0, china)

	return bri
}

func (i *IdCard) GetLastFour() string {
	return i.CardNum[18-4:]
}

func ConvertOldIdCard(idcard string) *IdCard {
	res := new(IdCard)
	copy(res.idNum[0:6], idcard[0:6])
	copy(res.idNum[8:17], idcard[6:15])
	if res.idNum[8] >= 2 {
		copy(res.idNum[6:8], []byte{49, 57})
	}
	res.idNum[17] = checkSum(res.idNum)
	res.CardNum = string(res.idNum[:])
	return res
}

//CheckIdCode 判断身份证号是否合法，返回true合法，返回false不合法
func CheckIdCode(idcode string) bool {
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

func checkSum(idnum [18]byte) byte {
	sum := 0
	for k, v := range idnum {
		sum += (int(v) - 48) * IdCardAlpha[k]
	}
	i := sum % 11
	return IdCardCheckSum[i]
}
