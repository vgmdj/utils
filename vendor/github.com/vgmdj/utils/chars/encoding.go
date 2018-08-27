package chars

import "github.com/axgle/mahonia"

//ConvertGbkToUtf8 gbk转utf8
func ConvertGbkToUtf8(src string) string {
	return convertToString(src, "gbk", "utf8")
}

func convertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
