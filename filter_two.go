package search_target_world

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

// FilterModelTwo 关键字过滤器
type FilterModelTwo struct {
	TrieModelTwo *TrieModelTwo
	noise        *regexp.Regexp
}

// New 返回一个关键字过滤器
func NewFilterModelTwo() *FilterModelTwo {
	return &FilterModelTwo{
		TrieModelTwo: NewTrieModelTwo(),
		noise:        regexp.MustCompile(`[\|\s-=_+!#^&%$@*(){}\[\]]+`),
	}
}

// LoadWordDict 加载关键字字典
func (FilterModelTwo *FilterModelTwo) LoadWordDict(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return FilterModelTwo.Load(f)
}

// Load common method to add words
func (FilterModelTwo *FilterModelTwo) Load(rd io.Reader) error {
	buf := bufio.NewReader(rd)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		FilterModelTwo.TrieModelTwo.Add(string(line))
	}
	return nil
}

// AddWord 添加关键字
func (FilterModelTwo *FilterModelTwo) AddWord(words ...string) {
	FilterModelTwo.TrieModelTwo.Add(words...)
}

// DelWord 删除关键字
// func (FilterModelTwo *FilterModelTwo) DelWord(words ...string) {
// 	FilterModelTwo.TrieModelTwo.Del(words...)
// }

// FindIn 检测关键字 -> 不连续 规则2
// 规则 2.【麦|当|劳】匹配：麦 1 当 1 劳 1112331 香浓咖啡，无限续杯
func (FilterModelTwo *FilterModelTwo) FindInWithoutStrict(text string) (bool, string) {
	text = FilterModelTwo.RemoveNoise(text)
	return FilterModelTwo.TrieModelTwo.FindInWithoutStrict(text)
}

// RemoveNoise 去除无效特殊字符
// 其中规则匹配过程中忽略特殊符号如[]、{}、()
func (FilterModelTwo *FilterModelTwo) RemoveNoise(text string) string {
	return FilterModelTwo.noise.ReplaceAllString(text, "")
}
