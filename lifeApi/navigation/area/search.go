package area

import "fmt"

var (
	defaultArea Arear
)

func init() {
	defaultArea = NewAreaByJson(areaJson)
}

func GetAreaNameByCode(code string) string {
	arear, f := defaultArea.GetAreaByCode(code)
	if !f {
		return ""
	}
	return arear.GetName()
}

func GetAreaCodeByName(name string) string {
	arear, f := defaultArea.GetAreaByName(name)
	if !f {
		return ""
	}

	return arear.GetCode()
}

func IsAreaByName(name string) bool {
	if GetAreaCodeByName(name) == "" {
		return false
	}

	return true
}

func IsAreaByCode(code string) bool {
	if GetAreaNameByCode(code) == "" {
		return false
	}

	return true
}

func Children(code string) []string {
	arear, f := defaultArea.GetAreaByCode(code)
	if !f {
		return []string{}
	}

	return arear.Children()
}

func AreaDetailByCode(code string) string {
	if code == "" {
		return "未知"
	}

	province := fmt.Sprintf("%s000", code[:3])
	pDetail := GetAreaNameByCode(province)
	if province == code {
		return pDetail
	}

	city := fmt.Sprintf("%s00", code[:4])
	cDetail := GetAreaNameByCode(city)
	if city == code {
		return fmt.Sprintf("%s%s", pDetail, cDetail)
	}

	return fmt.Sprintf("%s%s%s", pDetail, cDetail, GetAreaNameByCode(code))
}
