package main

import "fmt"

type MyWriter interface {
	Write(string) error
}

type MyCloser interface {
	Close(string) error
}

type writeCloser struct {
	MyWriter // interface也是一个类型
}

type fileWriter struct{}

func (fw *fileWriter) Write(s string) error {
	fmt.Println("file write:", s)
	return nil
}

type dbWriter struct{}

func (dw *dbWriter) Write(s string) error {
	fmt.Println("db write:", s)
	return nil
}

//func (w *writeCloser) Write(s string) error {
//	fmt.Println("writeCloser.Write", s)
//	return nil
//}

func (w *writeCloser) Close(s string) error {
	fmt.Println("writeCloser.Close")
	return nil
}

func main() {
	var mw MyWriter = &writeCloser{&dbWriter{}}
	mw.Write("bobby")
}
