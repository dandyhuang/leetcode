package main

import "fmt"

func DeferLoopV20() {
	// i变量在每次调用中被保存了下来
	for i := 0; i < 10; i++ {
		defer func(i int) {
			println(i)
		}(i)
	}
}

func main() {
	var s string
	s = "dsdfdsfsd"
	fmt.Println(s[1])
	arr := []int{0, 1, 2, 3, 4, 6}
	fmt.Println(arr[0:2])
	DeferLoopV20()
	prerequisites := [][]int{{1, 3}, {2, 4}}
	graph := make(map[int][]int, 2)
	for _, prerequisite := range prerequisites {
		course, prereq := prerequisite[0], prerequisite[1]
		graph[prereq] = append(graph[prereq], course)
	}
	for k, v := range graph[3] {
		fmt.Println(k, v)
	}
}
