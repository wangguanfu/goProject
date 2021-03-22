package log

import "fmt"

type FileLog struct {
	
}

func NewFileLog(file string) *FileLog  {
	return &FileLog{}
}


func (f *FileLog) LogDebug(msg string){
	fmt.Println(msg)
}


func (f *FileLog) LogWarn(msg string){
	fmt.Println(msg)
}





