package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, s := range tokens {
		switch s {
		case "+":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		case "-":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b-a)
		case "*":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case "/":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b/a)
		default:
			n, _ := strconv.Atoi(s)
			stack = append(stack, n)
		}

	}

	return stack[0]
}

func main() {
	fmt.Println(evalRPN1([]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}))
}
func evalRPN1(tokens []string) int {
	stack := []int{}

	for _, s := range tokens {
		if s != "+" && s != "-"&& s != "*" && s != "/"{
			n, _ := strconv.Atoi(s)
			stack = append(stack, n)
			continue
		}

		a,b := stack[len(stack)-1],stack[len(stack)-2]
		stack = stack[:len(stack)-2]

		switch s {
		case "+":
			stack = append(stack, a+b)
		case "-":
			stack = append(stack, b-a)
		case "*":
			stack = append(stack, a*b)
		case "/":
			stack = append(stack, b/a)
		}
	}

	return stack[0]
}
