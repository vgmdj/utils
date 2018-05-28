package area

import (
	"fmt"
	"sort"
	"sync"
)

const LatestRevision = "2018"

var (
	gb2260Selector = make(map[string]map[string]string)
	mutex          sync.RWMutex
)

type Selector struct{}

func (s *Selector) Register(revision string, gb2260 map[string]string) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := gb2260Selector[revision]; !ok {
		gb2260Selector[revision] = gb2260
	}
}

//Revisions 列出所有支持的revisions，并按最近到远的时间排序
func (s Selector) Revisions() (list []string) {
	mutex.Lock()
	defer mutex.Unlock()

	for revision, _ := range gb2260Selector {
		list = append(list, revision)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(list)))

	return
}

type gb2260 struct {
	gb map[string]string
}

func newGB2260(revision ...string) Area {
	rev := LatestRevision
	if len(revision) != 0 {
		if _, ok := gb2260Selector[revision[0]]; !ok {
			fmt.Printf("no such revision: %s\n", revision[0])
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

//Search
func (gb *gb2260) Search(code string) *AreaInfo {
	revisions := new(Selector).Revisions()

	for _, revision := range revisions {
		gb := newGB2260(revision)
		area := gb.Get(code)
		if area != nil {
			return area
		}
	}

	return nil
}

//TODO AllProvinces
func (gb *gb2260) AllProvinces() []AreaInfo {

	return nil
}

//TODO AllCities
func (gb *gb2260) AllCities() []AreaInfo {

	return nil
}

//TODO AllCounties
func (gb *gb2260) AllCounties() []AreaInfo {

	return nil
}
