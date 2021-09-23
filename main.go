package main

import "fmt"

func firstFunc(x, y int) (a, b int) {
	return x + y, x - y
}
func learnMemory() (q, p *int) {
	p = new(int)         //allocate memory
	s := make([]int, 20) //allocate 20 integers as a single block of memory
	s[3] = 7
	return &s[3], p
}
func namedReturn(x, y int) (z int) {
	z = x + y
	return // z is implicit coz we already named it
}

type pair struct {
	x, y int
}

func (p pair) String() string {
	return fmt.Sprintf("%d, %d", p.x, p.y)
}
func funcWithVariadicParams(params ...interface{}) {
	for _, param := range params { // underscore is to ignore index value of the parameter
		fmt.Println(param)
	}
}

func main() {
	fmt.Println("Hello there!")
	fmt.Println(firstFunc(3, 2))
	str := "string"
	var x = "long declaration"
	fmt.Println(str)
	fmt.Println(x)
	slice := []int{3, 5, 6}
	fmt.Println(slice)
	var zeroArray [4]int
	fmt.Println(zeroArray)
	fixedSize := [...]int{0, 1, 2}
	fmt.Println(fixedSize)
	var data interface{}
	data = ""
	switch c := data.(type) {
	case string:
		fmt.Println(c, "It is a string")
	default:
		fmt.Println("idk")
	}
	for key, value := range map[string]int{"one": 1} {
		fmt.Println(key, value)
	}
	if x := 8; x > 0 {
		fmt.Println("yes")
	}
	xBig := func() (x int) {
		return 3
	}
	fmt.Println(xBig())

	m := map[string]int{"x": 1}
	if _, ok := m["y"]; !ok {
		fmt.Println("Not ok")
	}
}
