package main

import (
	"DesignPatterns/Singleton"
)
func main(){
	s := Singleton.GetSimpleInstance()
	s.Show()
}
