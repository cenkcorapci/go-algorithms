package main

import (
	"fmt"
	"strconv"
	"strings"
)

type stack []int32

func (s *stack) Push(x int32) {
	*s = append(*s, x)
}

func (s *stack) Pop() (int32, bool) {
	if len(*s) == 0 {
		return 0, false
	} else {
		i := len(*s) - 1
		e := (*s)[i]
		*s = (*s)[:i]
		return e, true
	}
}
func (s *stack) Max() (int32, bool) {
	if len(*s) == 0 {
		return 0, false
	} else {
		x := (*s)[0]
		for i := 1; i < len(*s); i++ {
			if (*s)[i] > x {
				x = (*s)[i]
			}
		}
		return x, true
	}
}

func getMax(operations []string) (result []int32) {
	s := stack{}
	for _, op := range operations {
		cmd := strings.Fields(op)
		switch cmd[0] {
		case "1":
			x, err := strconv.Atoi(cmd[1])
			if err == nil {
				s.Push(int32(x))
			} else {
				fmt.Println(err)
			}
		case "2":
			s.Pop()
		case "3":
			if x, ok := s.Max(); ok {
				result = append(result, x)
			}

		}
	}
	return
}

func main() {
	inp := []string{"1 97", "2", "1 20", "2", "1 26", "1 20", "2", "3", "1 91", "3"}
	fmt.Println(getMax(inp))
}
