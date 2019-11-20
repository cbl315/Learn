/*
Implement the following operations of a stack using queues.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
empty() -- Return whether the stack is empty.

链接：https://leetcode-cn.com/problems/implement-stack-using-queues
*/
package implementStackUsingQueue

type BaseQueue struct {
	data []int
}

func (q *BaseQueue) Push(x int) {
	q.data = append(q.data, x)
}

func (q *BaseQueue) Pop() (result int) {
	if !q.Empty() {
		result = q.data[len(q.data)-1]
		q.data = q.data[:len(q.data)-1]
		return
	} else {
		return
	}
}

func (q *BaseQueue) Top() (result int) {
	if !q.Empty() {
		result = q.data[len(q.data)-1]
		return
	} else {
		return
	}
}

func (q *BaseQueue) Size() int {
	return len(q.data)
}

func (q *BaseQueue) Empty() bool {
	return len(q.data) == 0
}

type MyStack struct {
	data BaseQueue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.data.Push(x)
	for i := 0; i < this.data.Size()-1; i++ {
		this.data.Push(this.data.Top())
		this.data.Pop()
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.data.Pop()
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.data.Top()
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.data.Empty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
