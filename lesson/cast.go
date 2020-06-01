package lesson

func Cast() {
	var sum int
	sum = 5 + 6 + 3
	// avg := sum / 3
	avg := float32(sum) / 3
	if avg > 4.5 {
		println("good")
	}
}