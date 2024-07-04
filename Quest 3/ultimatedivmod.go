package piscine

func UltimateDivMod(a *int, b *int) {
	DividedValue := *a / *b
	RemainderValue := *a % *b
	*a = DividedValue
	*b = RemainderValue
}
