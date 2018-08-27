package area

import (
	"fmt"
	"sort"
	"sync"
)

//LatestRevision 最近年份
const LatestRevision = "2018"

var (
	gb2260Selector = make(map[string]map[string]string)
	mutex          sync.RWMutex
)

//Selector 选择器
type Selector struct{}

//Register 注册器
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

	for revision := range gb2260Selector {
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

//SetRevision 设置年份
func (gb *gb2260) SetRevision(revision string) {
	gb.gb = gb2260Selector[revision]
}

//Get 获取地域信息
func (gb *gb2260) Get(code string) *Info {
	if len(code) != 6 {
		return nil
	}

	return &Info{
		Province: gb.gb[code[:2]+"0000"],
		City:     gb.gb[code[:4]+"00"],
		County:   gb.gb[code],
	}
}

//Search 搜索匹配项
func (gb *gb2260) Search(code string) *Info {
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

//AllProvinces TODO
func (gb *gb2260) AllProvinces() []Info {

	return nil
}

//AllCities TODO
func (gb *gb2260) AllCities() []Info {

	return nil
}

//AllCounties TODO
func (gb *gb2260) AllCounties() []Info {

	return nil
}
