package day01

import (
	"sort"

	"github.com/sirupsen/logrus"
)

/*

第350题：两个数组的交集
给定两个数组，编写一个函数来计算它们的交集。
示例 1:
输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2,2]
示例 2:
输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [4,9]
说明：
输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
我们可以不考虑输出结果的顺序。
进阶:
如果给定的数组已经排好序呢？将如何优化你的算法呢？
思路：设定两个为0的指针，比较两个指针的元素是否相等。如果指针的元素相等，我们将两个指针一起向后移动，并且将相等的元素放入空白数组
*/

func intersect(nums, nums1 []int) []int {
	numsMap := map[int]int{}

	for _, val := range nums {
		//if _,ok :=numsMap[val];ok{
		//	numsMap[val] = i+1
		//}else {
		//	numsMap[val] =i
		//}
		numsMap[val] += 1
	}
	logrus.Println(numsMap)
	newNum := []int{}
	for _, item := range nums1 {
		if _, ok := numsMap[item]; ok {
			numsMap[item] -= 1
			newNum = append(newNum, item)
		}
	}
	return newNum
}

func intersect_1(nums, num1 []int) []int {
	sort.Ints(nums)
	sort.Ints(num1)
	i, j := 0, 0
	newList:= []int{}
	for i<len(nums)&&j<len(num1){
		if nums[i]<num1[j]{
			i++
		}else if nums[i]>num1[j]{
			j++
		}else {
			newList =append(newList,nums[i])
			i++
			j++
		}
	}
	return newList
}

func intersect3(nums1 []int, nums2 []int) []int {
	m0 := map[int]int{}
	for _, v := range nums1 {
		//遍历nums1，初始化map
		m0[v] += 1
	}
	k := 0
	for _, v := range nums2 {
		//如果元素相同，将其存入nums2中，并将出现次数减1
		if m0[v] > 0 {
			m0[v] -=1
			nums2[k] = v
			k++
		}
	}
	return nums2[0:k]
}
//GO
func intersect5(nums1 []int, nums2 []int) []int {
	i, j, k := 0, 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}