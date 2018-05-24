package files

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

//CheckFileIsExist 判断文件是否存在，存在返回true，不存在返回false
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//CopyFileTo 从指定位置复制创建出一个新的文件
//需先调用CheckFileIsExist判断文件是否存在
func CopyFileTo(srcFile string, destFile string) (err error) {
	var (
		fileContent []byte //temp file content
	)

	if fileContent, err = ioutil.ReadFile(srcFile); err != nil {
		log.Println("读取srcFile文件出错", srcFile)
		return
	}

	if err = ioutil.WriteFile(destFile, fileContent, 0666); err != nil {
		log.Println("写入destFile文件出错", destFile)
		return
	}

	return
}

//ParseFileTo 将文件中的内容解析到给出结构体中去
func ParseFileTo(filename string, parseStruct interface{}, mutex *sync.Mutex) (err error) {
	mutex.Lock()
	defer mutex.Unlock()

	var (
		fileContent []byte //temp file content
	)

	if fileContent, err = ioutil.ReadFile(filename); err != nil {
		log.Printf("读取%s文件时出错\n", filename)
		return
	}

	if err = json.Unmarshal(fileContent, parseStruct); err != nil {
		log.Println("解析文件内容时出错", string(fileContent))
		return
	}

	return
}

//WriteContentTo 根据传来内容写入目标文件
func WriteContentTo(destFilePath string, toWriteContent []byte, mutex *sync.Mutex) (err error) {
	mutex.Lock()
	defer mutex.Unlock()
	if err = ioutil.WriteFile(destFilePath, toWriteContent, 0666); err != nil {
		log.Printf("write %s err  : "+err.Error(), destFilePath)
		return
	}
	return
}
