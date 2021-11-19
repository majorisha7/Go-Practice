package main

func main() {
	strArray := []string{"apple", "pie", "apple", "red", "red", "red"}
	getMaxRepeatingElement(strArray)
}

func getMaxRepeatingElement(array []string) {
	elementMap := make(map[string]int)
	for _, element := range array {
		_, ok := elementMap[element]
		if ok {
			elementMap[element] = elementMap[element] + 1
		} else {
			elementMap[element] = 1
		}
	}
	var temp int
	var tempString string
	for k, v := range elementMap {
		if v > temp {
			temp = v
			tempString = k
		}
	}
	println(tempString)

}
