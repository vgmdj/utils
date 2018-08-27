package files

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

//File 文件
type File struct {
	fileName string
	mutex    sync.RWMutex
}

//NewFile 初始化
func NewFile(fileName string) *File {
	if fileName == "" {
		fmt.Println("need params filename")
		return nil
	}

	return &File{
		fileName: fileName,
	}

}

//SetFileName local file name set
func (lf *File) SetFileName(fileName string) {
	lf.fileName = fileName
}

//IsExist 判断文件是否存在，存在返回true，不存在返回false
func (lf *File) IsExist() bool {
	var exist = true
	if _, err := os.Stat(lf.fileName); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//ReadFile 读取文件内容
func (lf *File) ReadFile() (contents []byte, err error) {
	lf.mutex.RLock()
	defer lf.mutex.RUnlock()

	if !lf.IsExist() {
		return nil, fmt.Errorf("%s not exist", lf.fileName)
	}

	if contents, err = ioutil.ReadFile(lf.fileName); err != nil {
		fmt.Println("读取srcFile文件出错", lf.fileName)
		return
	}

	return
}

//CopyFileTo 从指定位置复制创建出一个新的文件
//需先调用IsExist判断文件是否存在,如果不存在则返回错误
func (lf *File) CopyFileTo(destFile string) (err error) {
	var (
		fileContent []byte //temp file content
	)

	if fileContent, err = lf.ReadFile(); err != nil {
		return
	}

	lf.mutex.Lock()
	defer lf.mutex.Unlock()

	if err = ioutil.WriteFile(destFile, fileContent, 0666); err != nil {
		fmt.Println("写入destFile文件出错", destFile)
		return
	}

	return
}

//ParseFileTo 将文件中的内容解析到给出结构体中去
func (lf *File) ParseFileTo(parseStruct interface{}) (err error) {
	var (
		fileContent []byte //temp file content
	)

	if fileContent, err = lf.ReadFile(); err != nil {
		return
	}

	if err = json.Unmarshal(fileContent, parseStruct); err != nil {
		fmt.Println("解析文件内容时出错", string(fileContent))
		return
	}

	return
}

//WriteContentTo 根据传来内容写入目标文件
func (lf *File) WriteContentTo(toWriteContent []byte) (err error) {
	lf.mutex.Lock()
	defer lf.mutex.Unlock()

	if err = ioutil.WriteFile(lf.fileName, toWriteContent, 0666); err != nil {
		fmt.Printf("write %s err  : %s ", lf.fileName, err.Error())
		return
	}
	return
}

//WriteJSONTo 根据传来内容写入目标文件
func (lf *File) WriteJSONTo(toWriteContent interface{}) (err error) {
	bts, err := json.Marshal(toWriteContent)
	if err != nil {
		fmt.Println(toWriteContent)
		return
	}

	lf.mutex.Lock()
	defer lf.mutex.Unlock()

	if err = ioutil.WriteFile(lf.fileName, bts, 0666); err != nil {
		fmt.Printf("write %s err  : %s ", lf.fileName, err.Error())
		return
	}
	return
}
