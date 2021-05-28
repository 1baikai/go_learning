package day02

import "strings"

//编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""

func Find(stringList []string) string {
	defaultString :=stringList[0]
	for _,value :=range stringList{
		num:=strings.Index(value, defaultString)
		if num==0{
			continue
		}else {
			defaultString = defaultString[:len(defaultString)-1]
		}
	}
	return defaultString
}
