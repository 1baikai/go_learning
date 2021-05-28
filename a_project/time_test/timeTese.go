package main

import (
	"fmt"
	"time"
)

func main() {
	s:="00:00"
	e:="13:00"
	t ,err:= time.Parse("15:04",s)
	if err!=nil{
		fmt.Printf("Err %v",err)
	}
	t1 ,err:= time.Parse("15:04",e)
	if err!=nil{
		fmt.Printf("Err %v",err)
	}
	fmt.Println(t1.Sub(t))

	fmt.Println(t.Add(t1.Sub(t)).Format("15:04"))
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Unix(time.Now().Unix(),0).Format("2006-01-02 15:04:05"))


	dayList:=[]int{1,2,3,4}
	weekend:=[]time.Weekday{}
	for _,item:=range dayList{
		switch item {
		case 0:
			weekend=append(weekend,time.Sunday)
		case 1:
			weekend=append(weekend,time.Monday)
		case 2:
			weekend=append(weekend,time.Tuesday)
		case 3:
			weekend=append(weekend,time.Wednesday)
		case 4:
			weekend=append(weekend,time.Thursday)
		case 5:
			weekend=append(weekend,time.Friday)
		case 6:
			weekend=append(weekend,time.Saturday)
		}
	}
	fmt.Println(weekend)
}





