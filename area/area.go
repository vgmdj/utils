package area

import "fmt"

//CodeType 代码类型
type CodeType int

const (
	//GB2260 区域代码
	GB2260 CodeType = iota
	//POST 邮编
	POST
)

//Area 地址
type Area interface {
	SetRevision(revision string)
	Get(code string) *Info
}

//NewArea 初始化
func NewArea(codeType CodeType) Area {
	switch codeType {
	default:
		fmt.Println("invalid code type")
		return nil

	case GB2260:
		return newGB2260(LatestRevision)

	case POST:
		return nil

	}
}

//Info 地区信息
type Info struct {
	Province string
	City     string
	County   string
}

//FullName 全称
func (ai *Info) FullName() string {
	return fmt.Sprintf("%s%s%s", ai.Province, ai.City, ai.County)
}

//GetGB2260 反推区域代码
//TODO
func (ai *Info) GetGB2260() string {
	return ai.Province
}

//
//GetPostCode 反推邮编
//TODO
func (ai *Info) GetPostCode() string {
	return ai.Province
}
