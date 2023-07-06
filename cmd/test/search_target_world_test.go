package test

import (
	"search_target_world"
	"testing"
)

// 仅应用规则一
func TestModelOne(t *testing.T) {
	filter := search_target_world.NewFilter()
	filter.AddWord("肯德基")
	str := "肯德基麦辣鸡腿堡，买一送一"
	if ok, res := filter.FindIn(str); ok != true {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(res)
	}

	str = `{汉}[堡](王) 狠霸王牛堡，美味无限`
	if ok, _ := filter.FindIn(str); ok != false {
		t.Errorf("%s not pass", str)
	} else {
		t.Log("没有匹配到，这是正常的")
	}
	filter.AddWord("汉堡王")
	str = `{汉\}[堡](王) 狠霸王牛堡，美味无限`
	if ok, res := filter.FindIn(str); ok != true {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(res)
	}
}

// 可以应用于规则一和规则二
func TestModelTwo(t *testing.T) {
	filter := search_target_world.NewFilterModelTwo()
	filter.AddWord("肯|德|基")
	filter.AddWord("汉|堡王")
	filter.AddWord("麦当|劳")
	filter.AddWord("沙县小吃")
	str := "肯啊德啊基啊"
	// 匹配规则一 和 规则二
	if found, res := filter.FindIn(str); found != true {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(res)
	}

	str = "1肯2德基"
	if found, res := filter.FindIn(str); found != true {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(res)
	}

	str = "肯德基啊"
	if found, res := filter.FindIn(str); found != true {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(res)
	}

	str = "汉12堡王"
	if found, res := filter.FindIn(str); found != true {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(res)
	}

	str = "麦1当劳"
	if found, res := filter.FindIn(str); found != false {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(res)
	}

	str = "麦当12劳"
	if found, res := filter.FindIn(str); found != true {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(res)
	}

	str = "沙)县小[吃"
	if found, res := filter.FindIn(str); found != true {
		t.Errorf("%s not pass", str)
	} else {
		t.Log(res)
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
}
