package Factory

import "fmt"

type HRBPigMeatBunsProduct struct {

}
func (h *HRBPigMeatBunsProduct)Intro(){
	fmt.Println("哈尔滨猪肉包子")
}
type HRBSXBunsProduct struct {

}

func (h *HRBSXBunsProduct)Intro(){
	fmt.Println("哈尔滨三鲜包子")
}
type BJPigMeatBunsProduct struct {

}
func (b *BJPigMeatBunsProduct)Intro(){
	fmt.Println("北京猪肉包子")
}
type BJSXBunsProduct struct {

}
func (b *BJSXBunsProduct)Intro(){
	fmt.Println("北京三鲜包子")
}