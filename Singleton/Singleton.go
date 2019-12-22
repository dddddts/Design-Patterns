package Singleton

import "fmt"

var s *SingletonStruct

type SingletonStruct struct {

}

func GetInstance()*SingletonStruct{
	if s==nil{
		s = &SingletonStruct{}
	}
	return s
}

func (s *SingletonStruct)Show(){
	fmt.Println("SingletonStruct")
}