package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// 打印tree的值序列
// 思路: 使用 slice 存储二叉树的值 根据列表的长度 打印值序列
func (t *tree) String() string {
	var biT []int
	biT = appendValues(biT, t)
	return fmt.Sprintf("%v", biT)
}

func InitTree(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	//b := values[:0]
	//fmt.Printf("%v", b)
	appendValues(values[:0], root)
	return root
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	//b := values[:0]
	//fmt.Printf("%v", b)
	appendValues(values[:0], root)
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
