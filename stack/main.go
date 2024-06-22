package main

import "fmt"

type Stack []string

func (s Stack) isEmpty() bool {
	return len(s) == 0
}

func (s *Stack) push(str string) {
	*s = append(*s, str)
}

func (s *Stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func main() {
	var stack Stack
	stack.push("Everyboody")
	stack.push("Loves")
	stack.push("Val")
	for !stack.isEmpty() {
		val, ok := stack.pop()
		if ok == true {
			fmt.Println(val)
		}
	}
}
