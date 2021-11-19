package main

func main() {
	recursiveFunction(9)
}

func recursiveFunction(input int) {
	var result int
	if input/2 > 1 {
		result = input / 2
		recursiveFunction(result)
	}
	println(input)
}
