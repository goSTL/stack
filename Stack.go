package stack

import (
	"container/list"
	"errors"
)

//Stack represents the LIFO data structure.
type Stack struct {
	stack *list.List
}

//Constructor for Stack.
//If params is empty, it will new a empty Stack.
//Or you can pass a List to initialize the Stack.
func New(params ...*list.List) (*Stack, error) {
	sta := new(Stack)
	var err error = nil
	switch len(params) {
	case 0:
		sta.stack = list.New()
	case 1:
		sta.stack = params[0]
	default:
		err = errors.New("The parameter is incorrect")
	}
	return sta, err
}

//Get the Stack size.
func (sta *Stack) Size() int {
	return sta.stack.Len()
}

//Check whether tthe Stack is empty or not.
func (sta *Stack) Empty() bool {
	return sta.stack.Len() == 0
}

//Insert the value above the top.
func (sta *Stack) Push(val interface{}) {
	sta.stack.PushBack(val)
}

//Delete the top value in Stack.
//If the Stack is empty, it will do nothing.
func (sta *Stack) Pop() {
	if !sta.Empty() {
		sta.stack.Remove(sta.stack.Back())
	}
}

//Get the top value in Stack.
//If the Stack is empty, you will get nil.
func (sta *Stack) Top() interface{} {
	if sta.Empty() {
		return nil
	} else {
		return sta.stack.Back().Value
	}
}

//Reverse the contents in the Stack.
func (sta *Stack) Reverse() {
	tmp := list.New()
	for !sta.Empty() {
		tmp.PushBack(sta.Top())
		sta.Pop()
	}
	sta.stack = tmp
}

//Swap the contents of two Stacks.
func (sta1 *Stack) Swap(sta2 *Stack) {
	tmp := sta1.stack
	sta1.stack = sta2.stack
	sta2.stack = tmp
}
