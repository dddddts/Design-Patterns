package Singleton

import (
	"fmt"
	"sync"
)

var(
	s *simpleSingletonStruct
	lock = &sync.Mutex{}
	once = sync.Once{}
)

type simpleSingletonStruct struct {

}

func GetSimpleInstance()*simpleSingletonStruct{
	if s==nil{
		s = &simpleSingletonStruct{}
	}
	return s
}
func GetSimpleInstanceLock()*simpleSingletonStruct {
	lock.Lock()
	defer lock.Unlock()
	if s==nil{
		s = &simpleSingletonStruct{}
	}
	return s
}
func GetSimpleInstanceDoubleLock()*simpleSingletonStruct {
	if s==nil{
		lock.Lock()
		defer lock.Unlock()
		if s==nil{
			s = &simpleSingletonStruct{}
		}
	}

	return s
}
func GetSimpleInstanceOnce()*simpleSingletonStruct {
	once.Do(func(){
		s = &simpleSingletonStruct{}
	})
	return s
}
func (s *simpleSingletonStruct)Show(){
	fmt.Println("SingletonStruct")
}