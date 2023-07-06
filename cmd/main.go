package main

import (
	"fmt"
	"search_target_world"
)

func main() {
	onlyRuleOneApplie()
	ruleOneAndTwoApplies()
}

func onlyRuleOneApplie() {
	// 仅适用于规则一
	fmt.Println("仅适用于规则一")
	filter := search_target_world.NewFilter()
	filter.LoadWordDict("dict/onlyRuleOneApplie.txt")
	found, res := filter.FindIn("肯德基麦辣鸡腿堡，买一送一")
	found1, res1 := filter.FindIn("{汉}[堡](王) 狠霸王牛堡，美味无限")
	fmt.Println(found, " ", res)
	fmt.Println(found1, " ", res1)
}

func ruleOneAndTwoApplies() {
	// 适用于规则一和规则二
	fmt.Println("适用于规则一和规则二")
	filter := search_target_world.NewFilterModelTwo()
	filter.LoadWordDict("dict/ruleOneAndTwoApplies.txt")
	str := "沙县小吃"
	//TODO bug
	//啊888 错误 但是后面  啊88888 是可以匹配的
	str1 := "肯德基啊888啊88888"
	found3, res3 := filter.FindIn(str)
	fmt.Println(found3, " ", res3)
	found4, res4 := filter.FindIn(str1)
	fmt.Println(found4, " ", res4)
	str2 := "麦当劳啊888当劳啊88888"
	found5, res5 := filter.FindIn(str2)
	fmt.Println(found5, " ", res5)

}
