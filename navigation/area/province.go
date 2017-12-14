package area

type AreaProvince struct {
	AreaCounty
	Cities []*AreaCity `json:"cities"`
}

func (this *AreaProvince) GetAreaByCode(code string) (Arear, bool) {
	if this.Code == code {
		return this, true
	}

	for _, city := range this.Cities {
		if arear, ok := city.GetAreaByCode(code); ok {
			return arear, true
		}
	}

	return nil, false
}

func (this *AreaProvince) GetAreaByName(name string) (Arear, bool) {
	if this.Name == name {
		return this, true
	}

	for _, city := range this.Cities {
		if arear, ok := city.GetAreaByName(name); ok {
			return arear, true
		}
	}

	return nil, false
}

func (this *AreaProvince) Children() []string {
	children := []string{}
	for _, city := range this.Cities {
		if city.Code != "" {
			children = append(children, city.Code)
		}
		children = append(children, city.Children()...)
	}

	return children
}

func (this *AreaProvince) GetCode() string {
	return this.Code
}

func (this *AreaProvince) GetName() string {
	return this.Name
}
