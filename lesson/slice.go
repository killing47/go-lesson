package lesson

import "fmt"

func Slice() {
	n := []int{19, 86, 1, 12}
	var sum int
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	fmt.Println(sum)
}