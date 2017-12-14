package area

import (
	"bytes"
	"encoding/json"
)

type Arear interface {
	GetAreaByCode(code string) (Arear, bool)
	GetAreaByName(name string) (Arear, bool)
	Children() []string
	GetCode() string
	GetName() string
}

type Area struct {
	Countries []*AreaCountry `json:"countries"`
}

func (this *Area) GetAreaByCode(code string) (Arear, bool) {
	for _, country := range this.Countries {
		if arear, ok := country.GetAreaByCode(code); ok {
			return arear, true
		}
	}
	return nil, false
}

func (this *Area) GetAreaByName(name string) (Arear, bool) {
	for _, country := range this.Countries {
		if arear, ok := country.GetAreaByName(name); ok {
			return arear, true
		}
	}
	return nil, false
}

func (this *Area) Children() []string {
	children := []string{}
	for _, country := range this.Countries {
		if country.Code != "" {
			children = append(children, country.Code)
		}
		children = append(children, country.Children()...)
	}

	return children
}

func (this *Area) GetCode() string {
	return ""
}

func (this *Area) GetName() string {
	return ""
}

func (this *Area) Parent() Arear {
	return nil
}

func (this *Area) String() string {
	return ""
}

func NewAreaByJson(areaJson string) *Area {
	area := &Area{
		Countries: make([]*AreaCountry, 0),
	}
	if err := json.NewDecoder(bytes.NewReader([]byte(areaJson))).Decode(area); err != nil {
		return nil
	}

	return area
}

func NewArea(countries ...*AreaCountry) *Area {
	area := &Area{
		Countries: make([]*AreaCountry, 0),
	}

	for _, country := range countries {
		area.Countries = append(area.Countries, country)
	}

	return area
}
