/*
Implement the following operations of a queue using stacks.

push(x) -- Push element x to the back of queue.
pop() -- Removes the element from in front of queue.
peek() -- Get the front element.
empty() -- Return whether the queue is empty.

链接：https://leetcode-cn.com/problems/implement-queue-using-stacks
*/
package implementQueueUsingStacks

type BaseStack struct {
	data []int
}

func (s *BaseStack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *BaseStack) Pop() int {
	result := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return result
}

func (s *BaseStack) Peek() int {
	return s.data[len(s.data)-1]
}

func (s *BaseStack) Empty() bool {
	return len(s.data) == 0
}

func (s *BaseStack) Size() int {
	return len(s.data)
}

type MyQueue struct {
	in  BaseStack
	out BaseStack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	q.in.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	if q.out.Size() == 0 {
		inSize := q.in.Size()
		for i := 0; i < inSize; i++ {
			q.out.Push(q.in.Pop())
		}
	}
	return q.out.Pop()
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	if q.out.Size() == 0 {
		inSize := q.in.Size()
		for i := 0; i < inSize; i++ {
			q.out.Push(q.in.Pop())
		}
	}
	return q.out.Peek()
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return q.in.Empty() && q.out.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
