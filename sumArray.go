package main

import (
    "fmt"
)

func main() {
   // 求数组所有元素之和
   var a=[3]int{1,5,10}
   sum:=0
   for i :=0;i<len(a);i++{
       sum+=a[i]
   }
   fmt.Println(sum)
}