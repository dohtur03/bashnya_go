package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Set the stack: Push, Pop, IsEmpty, Size, Clear
type Stack struct {
	items []int
}

func (stack *Stack) Push(val int) {
	stack.items = append(stack.items, val)
}

func (stack *Stack) Pop() int {
	if stack.IsEmpty() {
		return 0
	}
	popPosition := len(stack.items) - 1
	popItem := stack.items[popPosition]
	stack.items = stack.items[:len(stack.items)-1]
	return popItem
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.items) == 0
}

func (stack *Stack) Size() int {
	return len(stack.items)
}

func (stack *Stack) Clear() {
	stack.items = []int{}
}

// Convert a string of numbers into an array
func converter(inp string) []int {
	var nums []int

	ints := strings.Fields(inp)
	for _, p := range ints {
		num, err := strconv.Atoi(p)
		if err != nil {
			fmt.Println("Wrong input")
			nums = nil
			break
		} else {
			nums = append(nums, num)
		}
	}
	return nums
}

// Input a string of ints
func inp_str() string {
	text := bufio.NewReader(os.Stdin)
	inp, err := text.ReadString('\n')
	if err != nil {
		fmt.Println("Error in input")
		return ""
	}
	return inp
}

// Convert the array into a stack, process, output the result
func processor(arr []int) {
	myStack := Stack{}
	fmt.Println(myStack)

	for i := 0; i < len(arr); i++ {
		myStack.Push(arr[i])
		lnth := myStack.Size()
		fmt.Printf("Hi %d, stack: %v, size: %d\n", arr[i], myStack, lnth)
	}
	fmt.Println("full stack")
	for i := 0; i < len(arr)-1; i++ {
		num := myStack.Pop()
		lnth := myStack.Size()
		fmt.Printf("Bye %d, stack: %v, size: %d\n", num, myStack, lnth)
	}
	myStack.Clear()
	fmt.Println("empty inside", myStack)
}

// Call everything from here
func main() {
	text := inp_str()
	if text != "" {
		arr := converter(text)
		if len(arr) != 0 {
			processor(arr)
		}
	}
}
