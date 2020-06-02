package lesson

func Swap2() {
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)
}

func swap2(n, m *int) (*int, *int) {
	return n, m
}