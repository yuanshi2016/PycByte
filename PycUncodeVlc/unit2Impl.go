package PycUncodeVlc

import (
	"PycUncode/Utils"
	"fmt"
	"github.com/ying32/govcl/vcl"
	"os"
)

//::private::
type TMainFields struct {
}

func (f *TMain) OnFormCreate(sender vcl.IObject) {

}

func (f *TMain) OnPycSelectButtonClick(sender vcl.IObject) {
	if ok, dir := vcl.SelectDirectory2("", "./", true); ok {
		f.PycPathEdit.SetText(dir)
	}
}

func (f *TMain) OnSaveSelectButtonClick(sender vcl.IObject) {
	if ok, dir := vcl.SelectDirectory2("", "./", true); ok {
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
	StrUctByte := Utils.OpenFile(Path.Struct)
	for _, TempPyc := range Path.Pyc {
		var NewPyc []byte
		TempPycType := Utils.OpenFile(TempPyc.FullPath)
		NewPyc = append(NewPyc, StrUctByte[0:16]...)
		NewPyc = append(NewPyc, TempPycType[12:len(TempPycType)]...)
		var NewFullPath = SavePath + string(os.PathSeparator) + TempPyc.Name + ".py"
		var LogText = f.LogEdit.Text()
		f.LogEdit.SetText(fmt.Sprintf("%s Success \n", NewFullPath) + LogText)
		Utils.CreateFileWithDir(NewFullPath, NewPyc)
	}
}
