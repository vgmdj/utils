package area

const LatestRevision = "2018"

var gb2260Selector map[string]map[string]string

type gb2260 struct {
	gb map[string]string
}

func newGB2260(revision ...string) *gb2260 {
	rev := LatestRevision
	if len(revision) != 0 {
		if _, ok := gb2260Selector[revision[0]]; !ok {
			return nil
		}

		rev = revision[0]
	}

	return &gb2260{gb2260Selector[rev]}
}

func (gb *gb2260) SetRevision(revision string) {
	gb.gb = gb2260Selector[revision]
}

func (gb *gb2260) Get(code string) *AreaInfo {
	if len(code) != 6 {
		return nil
	}

	return &AreaInfo{
		Province: gb.gb[code[:2]+"0000"],
		City:     gb.gb[code[:4]+"00"],
		County:   gb.gb[code],
	}
}
