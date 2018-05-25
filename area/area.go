package area

import "fmt"

type CodeType int

const (
	GB2260 CodeType = iota
	POST
)

type Area interface {
	SetRevision(revision string)
	Get(code string) *AreaInfo
}

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

type AreaInfo struct {
	Province string
	City     string
	County   string
}

func (ai *AreaInfo) FullName() string {
	return ai.Province
}

func (ai *AreaInfo) GetGB2260() string {
	return ai.Province
}

func (ai *AreaInfo) GetPostCode() string {
	return ai.Province
}
