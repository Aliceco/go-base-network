package main

import (
	"fmt"
	"sort"
)
type ServerSlice []Server
func (s ServerSlice) Len() int           { return len(s) }
func (s ServerSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ServerSlice) Less(i, j int) bool { return s[i].Weight < s[j].Weight }


type Server struct {
	Weight int
}
func main() {
	//nums:=[]int{2, 5, 1, 7, 4}
	//sort.Ints(nums)
	//fmt.Println(nums)
	//
	//strs:=[]string{"an", "lisi", "zhangsan"}
	//sort.Strings(strs)
	//fmt.Println(strs)

	ss:=ServerSlice{
		Server{Weight:4},
		Server{Weight:1},
		Server{Weight:2},
	}
	sort.Sort(ss)
	fmt.Println(ss) // [{1}, {2}, {4}]
}
