package lesson

type MyInt int

func (n *MyInt) Inc() { (*n)++ }

func Receiver() {
	var n MyInt
	println(n)
	n.Inc()
	println(n)
}