package main

func mmadd(a int, b int) (c int, d int) {
	return a + b, a - b
}

func main() {
	a := 22
	c, d := mmadd(50, 30)
	println(a, c, d)
}
