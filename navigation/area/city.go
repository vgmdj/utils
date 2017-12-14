package area

type AreaCity struct {
	AreaCounty
	Counties []*AreaCounty `json:"counties"`
}

func (this *AreaCity) GetAreaByCode(code string) (Arear, bool) {
	if this.Code == code {
		return this, true
	}

	for _, county := range this.Counties {
		if arear, ok := county.GetAreaByCode(code); ok {
			return arear, true
		}
	}

	return nil, false
}

func (this *AreaCity) GetAreaByName(name string) (Arear, bool) {
	if this.Name == name {
		return this, true
	}

	for _, county := range this.Counties {
		if arear, ok := county.GetAreaByName(name); ok {
			return arear, true
		}
	}

	return nil, false
}

func (this *AreaCity) Children() []string {
	children := []string{}
	for _, county := range this.Counties {
		if county.Code != "" {
			children = append(children, county.Code)
		}
		children = append(children, county.Children()...)
	}

	return children
}

func (this *AreaCity) GetCode() string {
	return this.Code
}

func (this *AreaCity) GetName() string {
	return this.Name
}
