package main

func validateStackSequences(pushed []int, popped []int) bool {

	stack := make([]int, 0)
	index := 0
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for stack[len(stack)-1] == popped[index] {
			stack = stack[:len(stack)-1]
			index++
		}
	}

	return len(stack) == 0
}
