package main

import (
	"fmt"
	"regexp"
)
var  RegexpSpecial = regexp.MustCompile(`[~!#$^&*()+=<>?:"{}|,\;'\\[\]·~#￥%……&*（）——+={}|《》？：“”【】、；‘'，。、]`)
var RegexpNum = regexp.MustCompile(`^(\-|\+)?\d+(\.\d+)?$`)
var regexVersion  = regexp.MustCompile(`\d+\_\d+\_\d+`)
func main()  {
	//if !RegexpNum.MatchString("+12312123123123") {
	//	fmt.Println("cuowu")
	//}else {
	//	fmt.Println("yes")
	//}
	a:=regexVersion.FindStringSubmatch("openGauss-2_0——0-CentOS_x86-64bit-all.tar.gz")
	fmt.Println(a)
	fmt.Println(regexVersion.FindString("openGauss-2_0_0-CentOS_x86-64bit-all.tar.gz"))
}
