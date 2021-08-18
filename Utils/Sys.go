package Utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type FileInfo struct {
	Name     string
	Path     string
	FullPath string
	Suffix   string
}
type GetPycInfo struct {
	Pyc    []FileInfo
	Struct string
}

func (this GetPycInfo) ToJson() string {
	b, err := json.Marshal(this)
	if err != nil {
		return ""
	}
	return string(b)
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
		default: //防止无文件名溢出
			if fi.IsDir() { // 忽略目录
				ListDir(dirPth + PthSep + fi.Name())
			}
		}
		if strings.LastIndex(fi.Name(), ".") <= 0 || Getsuffix(fi.Name()) != "pyc" {
			continue
		}
		Temp := FileInfo{
			Name:     GetFileName(fi.Name()),
			Path:     dirPth + PthSep,
			Suffix:   Getsuffix(fi.Name()),
			FullPath: dirPth + PthSep + fi.Name(),
		}
		files.Pyc = append(files.Pyc, Temp)
	}
	return files, nil
}
func GetFileName(Name string) string {
	var NewFilePathName = Name[0:strings.LastIndex(Name, ".")]
	NewFilePathName = strings.ReplaceAll(NewFilePathName, ".", string(os.PathSeparator))
	return NewFilePathName
}
func Getsuffix(Name string) string {
	var c = strings.Split(Name, ".")
	return strings.ToLower(c[len(c)-1])
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

	os.MkdirAll(fullpath[:strings.LastIndex(fullpath, string(os.PathSeparator))], os.ModePerm)
	file, _ := os.OpenFile(fullpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	file.Write(content)
}

func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", command)
	default:
		cmd = exec.Command("bash", "-c", command)
	}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
