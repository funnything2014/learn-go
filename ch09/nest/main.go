package main

import "fmt"

type MyReader interface {
	Read()
}

type MyWriter interface {
	Write(string)
}

type MyReadWriter interface {
	MyReader
	MyWriter
	ReadWrite()
}

type SreadWriter struct{}

func (s *SreadWriter) Read() {
	fmt.Println("read me")
}

func (s *SreadWriter) Write(s2 string) {
	fmt.Println("write me")
}

func (s *SreadWriter) ReadWrite() {
	fmt.Println("read write me")
}

func main() {
	var mrw MyReadWriter = &SreadWriter{}
	mrw.Read()
}
