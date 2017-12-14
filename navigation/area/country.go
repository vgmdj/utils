package area

type AreaCountry struct {
	AreaCounty
	Provinces []*AreaProvince `json:"provinces"`
}

func (this *AreaCountry) GetAreaByCode(code string) (Arear, bool) {
	if this.Code == code {
		return this, true
	}

	for _, province := range this.Provinces {
		if arear, ok := province.GetAreaByCode(code); ok {
			return arear, true
		}
	}

	return nil, false
}

func (this *AreaCountry) GetAreaByName(name string) (Arear, bool) {
	if this.Name == name {
		return this, true
	}

	for _, province := range this.Provinces {
		if arear, ok := province.GetAreaByName(name); ok {
			return arear, true
		}
	}

	return nil, false
}

func (this *AreaCountry) Children() []string {
	children := []string{}
	for _, province := range this.Provinces {
		if province.Code != "" {
			children = append(children, province.Code)
		}
		children = append(children, province.Children()...)
	}

	return children
}

func (this *AreaCountry) GetCode() string {
	return this.Code
}

func (this *AreaCountry) GetName() string {
	return this.Name
}
