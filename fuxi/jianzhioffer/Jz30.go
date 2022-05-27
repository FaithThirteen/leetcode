package main

var stack = []int{}
var stackMin = []int{}

func Push1(node int) {
	// write code here
	stack = append(stack, node)
	if len(stackMin) == 0 {
		stackMin = append(stackMin, node)
	} else if node > stackMin[len(stackMin)-1] {
		stackMin = append(stackMin, stackMin[len(stackMin)-1])
	} else if node < stackMin[len(stackMin)-1] {
		stackMin = append(stackMin, node)
	} else {
		stackMin = append(stackMin, stackMin[len(stackMin)-1])
	}
}
func Pop1() {
	// write code here
	stack = stack[:len(stack)-1]
	stackMin = stackMin[:len(stackMin)-1]
}
func Top() int {
	// write code here
	return stack[len(stack)-1]
}
func Min() int {
	// write code here
	return stackMin[len(stackMin)-1]
}
