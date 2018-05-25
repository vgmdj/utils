package gbdata

import "github.com/vgmdj/utils/area"

func init() {
	s := new(area.Selector)
	s.Register("2018", rev2018)

}

var rev2018 = map[string]string{
	"": "",
}
