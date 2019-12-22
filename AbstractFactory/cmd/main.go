package main

import (
	"DesignPatterns/AbstractFactory"
)
func main(){
	hrbFact := &Factory.HRBBunsFactory{}
	hrbPigMeatBuns := hrbFact.CreatePigMeatBuns()
	hrbPigMeatBuns.Intro()
	hrbSXBuns := hrbFact.CreateSXBuns()
	hrbSXBuns.Intro()

	bjFact := &Factory.BJBunsFactory{}
	bjPigMeatBuns := bjFact.CreatePigMeatBuns()
	bjPigMeatBuns.Intro()
	bjSxBuns := bjFact.CreateSXBuns()
	bjSxBuns.Intro()

}