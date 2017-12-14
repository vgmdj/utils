package area

import "fmt"

type AreaCounty struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func (this *AreaCounty) GetAreaByCode(code string) (Arear, bool) {
	if this.Code == code {
		return this, true
	}

	return nil, false
}

func (this *AreaCounty) GetAreaByName(name string) (Arear, bool) {
	if this.Name == name {
		return this, true
	}

	return nil, false
}

func (this *AreaCounty) Children() []string {
	return []string{}
}

func (this *AreaCounty) GetCode() string {
	return this.Code
}

func (this *AreaCounty) GetName() string {
	return this.Name
}

func (this *AreaCounty) String() string {
	return fmt.Sprintf("[%s]%s", this.Code, this.Name)
}

func NewAreaCounty(name string, code string) *AreaCounty {
	return &AreaCounty{
		Name: name,
		Code: code,
	}
}
