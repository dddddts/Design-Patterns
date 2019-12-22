package Factory

type FactoryInterface interface{
	CreatePigMeatBuns()ProductInterface
	CreateSXBuns()ProductInterface
}
type ProductInterface interface {
	Intro()
}
type HRBBunsFactory struct {

}
type BJBunsFactory struct {

}

func (h *HRBBunsFactory)CreatePigMeatBuns()ProductInterface{
	return &HRBPigMeatBunsProduct{}
}
func (h *HRBBunsFactory)CreateSXBuns()ProductInterface{
	return &HRBSXBunsProduct{}
}

func (b *BJBunsFactory)CreatePigMeatBuns()ProductInterface{
	return &BJPigMeatBunsProduct{}
}
func (b *BJBunsFactory)CreateSXBuns()ProductInterface{
	return &BJSXBunsProduct{}
}