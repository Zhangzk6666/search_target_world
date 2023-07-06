package main

import (
	"fmt"
	search_target_world "search_target_world"
)

func main() {
	// filter := search_target_world.NewFilter()
	// filter.LoadWordDict("dict/word.txt")
	// // model 1
	// found, res := filter.FindIn("肯德基麦辣鸡腿堡，买一送一")
	// found1, res1 := filter.FindIn("{汉}[堡](王) 狠霸王牛堡，美味无限")
	// fmt.Println(found, " ", res)
	// fmt.Println(found1, " ", res1)

	filter := search_target_world.NewFilterModelTwo()
	filter.LoadWordDict("dict/word.txt")
	// // TODO model 2
	// str := "麦 1 当 1  劳1112331 香浓咖啡，肯1德基3无限续杯"
	str := "肯1|fsdrjhgiudfjgb德基"
	// str := "肯德基麦辣鸡腿堡，买一送一"
	found3, res3 := filter.FindInWithoutStrict(str)
	fmt.Println(found3, " ", res3)
}
