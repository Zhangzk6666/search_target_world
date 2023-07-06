# test


## 对于`其中规则匹配过程中忽略特殊符号如[]、{}、()`的理解以及处理
- 目前设置的特殊符号有：`|-=_+!#^&%$@*(){}[]\`还有空格，如涉及到其他特殊字符，可添加在`common.go`中的`specialCharacter`里面。
- 为了处理效率，我首先处理掉字符串中的特殊字符再进行匹配，所以返回的结果中是`不包含`特殊字符的。

## 对于`规则一`的理解
- 字符串必须完全匹配，如：
>肯德基 => 可以匹配 `肯德基`;不能匹配到: `肯-这里是干扰项-德基`
- 忽略`特殊字符`对于规则一适用

## 对于`规则二`的理解:
#### 字符串的分割可以是多样的，如:
麦|当|劳
麦|当劳
麦当|劳
麦当劳
#### 对于`|`的理解
|表示字符前后可以存在其他干扰项，也可以没有干扰项。如:
麦|当|劳 =>可以匹配`麦当劳`、`麦1当劳`、`麦1当2劳`;可以匹配到只要按顺序出现‘麦当劳’的任何句子
#### 忽略`特殊字符`对于规则二适用

## 对于模式的说明
#### 本项目提供了两种模式
- 一种是仅应用规则一
- 另外一种应用规则一和规则二
- 对于不同模式的选择：如果不涉及到规则一请选择模式一，否则请选择模式二
#### 模式一的使用
```golang
// 仅适用于规则一
filter := search_target_world.NewFilter()
filter.LoadWordDict("dict/word.txt")
found, res := filter.FindIn("肯德基麦辣鸡腿堡，买一送一")
found1, res1 := filter.FindIn("{汉}[堡](王) 狠霸王牛堡，美味无限")
fmt.Println(found, " ", res)
fmt.Println(found1, " ", res1)
```

#### 模式二的使用
```golang
// 适用于规则一和规则二
filter := search_target_world.NewFilterModelTwo()
filter.LoadWordDict("dict/word.txt")
str := "肯德基"
str1 := "1肯2德3基4"
found3, res3 := filter.FindIn(str)
found4, res4 := filter.FindIn(str1)
fmt.Println(found3, " ", res3)
fmt.Println(found4, " ", res4)
```
