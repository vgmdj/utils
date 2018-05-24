package area

import "github.com/vgmdj/utils/chars"

type area interface {
	GetProvince() string
	GetCity() string
	GetFullName() string
}

type idArea struct {
	code int
	area string
}

func NewArea(code int) area {
	if code > 659001 || code < 110000 {
		return nil
	}

	return &idArea{code: code}
}

func (ia *idArea) GetProvince() string {
	c := chars.TakeLeftInt(ia.code, 2) * 10000
	return GB2260[c]
}

func (ia *idArea) GetCity() string {
	c := chars.TakeLeftInt(ia.code, 4) * 100
	return GB2260[c]
}

func (ia *idArea) GetFullName() string {
	return GB2260[ia.code]
}
