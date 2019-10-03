package arithmetic

func add(a int, b int) int      { return a + b }
func subtract(a int, b int) int { return a - b }
func multiply(a int, b int) int { return a * b }
func divide(a int, b int) int   { return a / b }

func Arithmetic(a int, b int, operator string) int {
	f := map[string]func(int, int) int{
		"add":      add,
		"subtract": subtract,
		"multiply": multiply,
		"divide":   divide,
	}

	return f[operator](a, b)
}
