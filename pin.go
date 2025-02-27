package drivers

type Pin interface {
	High()
	Low()
	Set(high bool)
}
