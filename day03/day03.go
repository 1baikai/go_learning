package day03

import "fmt"

//示例 1:
//
//输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
//示例 2:
//
//输入: [-1,-100,3,99] 和 k = 2
//输出: [3,99,-1,-100]
//解释:
//
//向右旋转 1 步: [99,-1,-100,3]
//向右旋转 2 步: [3,99,-1,-100]

func reversal(list []rune, k int) []rune {
	a := list[:len(list)-k]
	fmt.Println(a)
	b := list[len(list)-k:]
	fmt.Println(b)
	return append(b, a...)
}


func removeElement(num []int, val int) int {
	for i := 0; i < len(num); {
		if val == num[i] {
			num = append(num[:i], num[i+1:]...)
		}else {
			i++
		}

	}
	fmt.Println(num)
	return len(num)
}
