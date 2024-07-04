package piscine

func Swap(a *int, b *int) {
	AValue := *a
	BValue := *b

	*a = BValue
	*b = AValue
}
