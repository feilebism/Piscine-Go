package piscine

func DivMod(a int, b int, div *int, mod *int) {
	DividedValue := a / b
	RemainderValue := a % b
	*div = DividedValue
	*mod = RemainderValue
}
