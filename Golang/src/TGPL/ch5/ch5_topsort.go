package main

import (
	"fmt"
	"log"
	"sort"
)

// 拓扑排序
// 前置条件： 可以构成有向图
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	//"linear algebra": {"calculus",},
	"compliers": {
		"data structures",
		"formal language",
		"computer organization",
	},
	"data strctures":       {"discrete math"},
	"datebases":            {"data structures"},
	"discrete math":        {"intro to programming"},
	"formal language":      {"discrete math"},
	"networks":             {"operating system"},
	"operating systems":    {"data structures", "computer organization"},
	"programming language": {"data structures", "computer organization"},
}

func testTopSort() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// 深度优先 拓扑排序
func topoSort(m map[string][]string) []string {
	var order []string
	var seen = make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func sortMapIntKey(m map[int]string) (sortedIndex []int) {
	for index, _ := range m {
		sortedIndex = append(sortedIndex, index)
	}
	sort.Ints(sortedIndex)
	return sortedIndex
}

func testTopSort2() {
	sortedResult := topoSort2(prereqs)
	sortedIndex := sortMapIntKey(sortedResult)
	for _, index := range sortedIndex {
		fmt.Printf("%d:\t%s\n", index, sortedResult[index])
	}

}

func topoSort2(m map[string][]string) map[int]string {
	var order = make(map[int]string)
	currIndex := 1
	var seen = make(map[string]bool)
	var currSeen = make(map[string]int)
	var visitAll func(preCoutse string, items []string)
	visitAll = func(preCoutse string, items []string) {
		for _, item := range items {
			// 检测有向图是否有环
			currSeen[item]++
			if currSeen[item] >= 2 {
				log.Fatalf("exist cycle %s", item)
			}
			if !seen[item] {
				seen[item] = true
				visitAll(item, m[item])
				order[currIndex] = item
				currIndex++
			}
			currSeen[item]--
		}

	}
	var keys []string
	for key, _ := range m {
		keys = append(keys, key)
	}
	visitAll("", keys)
	return order
}
