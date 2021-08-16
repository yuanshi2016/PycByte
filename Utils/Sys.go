package Utils

import (
	"io/ioutil"
	"os"
	"strings"
)

type FileInfo struct {
	Name     string
	Path     string
	FullPath string
	suffix   string
}
type GetPycInfo struct {
	Pyc    []FileInfo
	Struct string
}

var files GetPycInfo
var _struct string

//获取指定目录下的所有文件和目录
func ListDir(dirPth string) (GetPycInfo, error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return files, err
	}
	PthSep := string(os.PathSeparator)
	// suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		switch strings.ToLower(fi.Name()) {
		case ".ds_store":
			continue
		case "struct":
			files.Struct = dirPth + PthSep + fi.Name()
			//_struct = dirPth + PthSep + fi.Name()
			continue
		}
		if fi.IsDir() { // 忽略目录
			ListDir(dirPth + PthSep + fi.Name())
			//fmt.Println(dirPth + PthSep + fi.Name())
		} else {
			files.Pyc = append(files.Pyc, FileInfo{
				Name:     GetFileName(fi.Name()),
				Path:     dirPth + PthSep,
				suffix:   Getsuffix(fi.Name()),
				FullPath: dirPth + PthSep + fi.Name(),
			})
		}
	}
	return files, nil
}
func GetFileName(Name string) string {
	//var c = strings.Split(Name,".")
	var NewFilePathName = Name[0:strings.LastIndex(Name, ".")]
	NewFilePathName = strings.ReplaceAll(NewFilePathName, ".", "/")
	return NewFilePathName
}
func Getsuffix(Name string) string {
	var c = strings.Split(Name, ".")
	return c[len(c)-1]
}
func OpenFile(fileName string) []byte {
	var ContentByte []byte
	var f *os.File
	var err error
	defer f.Close()
	f, err = os.OpenFile(fileName, os.O_RDONLY, 0600)
	if err == nil {
		ContentByte, _ = ioutil.ReadAll(f)
	}
	return ContentByte
}

/**
fullpath 全路径
*/
func CreateFileWithDir(fullpath string, content []byte) {
	os.MkdirAll(fullpath[:strings.LastIndex(fullpath, "/")], os.ModePerm)
	file, _ := os.OpenFile(fullpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	file.Write(content)
}
