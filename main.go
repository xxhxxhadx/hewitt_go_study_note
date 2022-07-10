package main

import (
	"fmt"
	_ "hewitt_design_mode/design_mode"
	_ "hewitt_design_mode/design_mode/CreateMode"
)

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func main() {
	nums := make([]int, 0)
	nums = append(nums, 1, 2, 3, 4)
	fmt.Println(nums)
	fmt.Println(nums[2])
	nunMap := make(map[int]int)
	nunMap[0] = 1
	nunMap[1] = 2
	delete(nunMap, 1)
	fmt.Println(nunMap)
	fmt.Println(nunMap[1])

	i := 1
	i++ // i= i+1 and i += 1
	fmt.Println(i)

	for _, v := range nums {
		fmt.Println(v)
	}
}

//array 1,2,3,4,5
// map
// 1 -> 2
// 2 -> 3
