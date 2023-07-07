package test

import (
	"search_target_world"
	"testing"
)

// 仅应用规则一
// func TestModelOne(t *testing.T) {
// 	filter := search_target_world.NewFilter()
// 	filter.AddWord("肯德基")
// 	str := "肯德基麦辣鸡腿堡，买一送一"
// 	if ok, res := filter.FindIn(str); ok != true {
// 		t.Errorf("%s not pass", str)
// 	} else {
// 		t.Log(str, "  ==>  ", res)
// 	}

// 	str = `{汉}[堡](王) 狠霸王牛堡，美味无限`
// 	if ok, _ := filter.FindIn(str); ok != false {
// 		t.Errorf("%s not pass", str)
// 	}
// 	filter.AddWord("汉堡王")
// 	str = `{汉\}[堡](王) 狠霸王牛堡，美味无限`
// 	if ok, res := filter.FindIn(str); ok != true {
// 		t.Errorf("%s not pass", str)
// 	} else {
// 		t.Log(str, "  ==>  ", res)
// 	}
// }

// 可以应用于规则一和规则二
func TestModelTwo(t *testing.T) {
	var str string
	filter := search_target_world.NewFilterModelTwo()
	filter.AddWord("肯|德|基")
	filter.AddWord("汉|堡王")
	filter.AddWord("麦当|劳")
	filter.AddWord("沙县小吃")
	str = "肯啊德啊基啊"
	//匹配规则一 和 规则二
	if found, res := filter.FindIn(str); found != true {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(str, "  ==>  ", res)
	}

	str = "1肯2德基"
	if found, res := filter.FindIn(str); found != true {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(str, "  ==>  ", res)
	}

	str = "肯德基啊"
	if found, res := filter.FindIn(str); found != true {
		t.Log(str, "  ==>  ", res, " ", found)

		t.Errorf("%s not pass", str)
	} else {
		t.Log(str, "  ==>  ", res)
	}

	str = "汉12堡王"
	if found, res := filter.FindIn(str); found != true {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(str, "  ==>  ", res)
	}

	str = "麦1当劳"
	if found, _ := filter.FindIn(str); found != false {
		t.Errorf("%s not pass", str)
	}

	str = "麦当12劳"
	if found, res := filter.FindIn(str); found != true {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(str, "  ==>  ", res)
	}

	str = "沙)县小[吃"
	if found, res := filter.FindIn(str); found != true {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(str, "  ==>  ", res)
	}

	str = "沙县好吃"
	if found, res := filter.FindIn(str); found != false {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(res)
	}

	str = "沙县小吃"
	if found, res := filter.FindIn(str); found != true {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(res)
	}

	filter.AddWord("沙县小,sw吃ssssssssssssssssss")
	str = "沙县小,sw吃ssssssssssssssssss"
	if found, res := filter.FindIn(str); found != true {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(res)
	}

	// 错误的将  将每一个rune认为是被|分割
	filter.AddWord("沙sdasfa吃ssssssss")
	str = "沙sdasfa吃ss吃ssssssss"
	if found, res := filter.FindIn(str); found != false {
		t.Log(found, " ", res)
		t.Errorf("%s not pass", str)
	}
}

func TestModelTwo_NewBug(t *testing.T) {
	filter := search_target_world.NewFilterModelTwo()
	// filter.AddWord("肯德基")
	filter.AddWord("肯德基|啊88888")
	str := "肯德基啊8啊88888"
	if found, res := filter.FindIn(str); found != true {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(str, "  ==>  ", res)
	}

	str = "肯s德基啊88888"
	if found, res := filter.FindIn(str); found != false {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(found, "  ==>  ", res)
	}

	filter.AddWord("沙sdasfa吃ssssssss")
	str = "沙sssssssssssssssssssssssss"
	if found, _ := filter.FindIn(str); found != false {
		t.Errorf("%s not pass", str)
	}
}
