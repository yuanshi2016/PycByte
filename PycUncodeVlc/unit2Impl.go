package PycUncodeVlc

import (
	"PycUncode/Utils"
	"encoding/hex"
	"fmt"
	"github.com/ying32/govcl/vcl"
	"log"
	"os"
	"strings"
)

//::private::
type TMainFields struct {
}

func (f *TMain) OnFormCreate(sender vcl.IObject) {
	//f.SavePathEdit.SetText("/Volumes/D/GolangTool/PycUncode/test/save")
	//f.PycPathEdit.SetText("/Volumes/D/GolangTool/PycUncode/test/aaaa")
	f.LogEdit.SetText("")
}

func (f *TMain) OnPycSelectButtonClick(sender vcl.IObject) {
	if ok, dir := vcl.SelectDirectory2("", "./", false); ok {
		f.PycPathEdit.SetText(dir)
	}
}

func (f *TMain) OnSaveSelectButtonClick(sender vcl.IObject) {
	if ok, dir := vcl.SelectDirectory2("", "./", false); ok {
		f.SavePathEdit.SetText(dir)
	}
}

func (f *TMain) OnRunButtonClick(sender vcl.IObject) {
	f.LogEdit.SetText("")
	var SavePath = f.SavePathEdit.Text()
	var PycPath = f.PycPathEdit.Text()
	if len(SavePath) <= 0 || len(PycPath) <= 0 {
		vcl.ShowMessage("路径错误")
		return
	}
	var Path, _ = Utils.ListDir(PycPath)
	if len(Path.Struct) <= 0 {
		vcl.ShowMessage("Struct文件不存在")
	}
	//wp := workpool.New(10)
	StrUctByte := Utils.OpenFile(Path.Struct)
	for _, TempPyc := range Path.Pyc {
		var NewPyc []byte
		TempPycType := Utils.OpenFile(TempPyc.FullPath)
		//TODO 检测文件头
		if hex.EncodeToString(StrUctByte[0:16]) != hex.EncodeToString(TempPycType[0:16]) {
			NewPyc = append(NewPyc, StrUctByte[0:16]...)
			NewPyc = append(NewPyc, TempPycType[12:len(TempPycType)]...)
			Utils.CreateFileWithDir(TempPyc.FullPath, NewPyc)
		}
		var NewFullPath = SavePath + string(os.PathSeparator) + TempPyc.Name + ".py"
		var LogText = f.LogEdit.Text()
		f.LogEdit.SetText(fmt.Sprintf("\n %s Success \n", NewFullPath) + LogText)
		//保存目录是否存在
		os.MkdirAll(NewFullPath[:strings.LastIndex(NewFullPath, string(os.PathSeparator))], os.ModePerm)
		err, _, _ := Utils.Shellout(fmt.Sprintf("uncompyle6 -o %s %s", NewFullPath, TempPyc.FullPath))
		//fmt.Println(fmt.Sprintf("uncompyle6 -o %s %s",NewFullPath,TempPyc.FullPath))
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		//wp.Do(func() error {
		//	return nil
		//})
	}
	//wp.Wait()

}
