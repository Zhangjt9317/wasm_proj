package main

func main() {
	println("adding two numbers: ", add(2, 3))
	println("adding multiply numbers: ", multiply(2, 3))

}

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}
