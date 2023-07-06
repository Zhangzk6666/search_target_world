package search_target_world

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

// Filter 关键字过滤器
type Filter struct {
	trie  *Trie
	noise *regexp.Regexp
}

// New 返回一个关键字过滤器
func NewFilter() *Filter {
	return &Filter{
		trie:  NewTrie(),
		noise: regexp.MustCompile(specialCharacter),
	}
}

// LoadWordDict 加载关键字字典
func (filter *Filter) LoadWordDict(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return filter.Load(f)
}

// Load common method to add words
func (filter *Filter) Load(rd io.Reader) error {
	buf := bufio.NewReader(rd)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		filter.trie.Add(string(line))
	}
	return nil
}

// AddWord 添加关键字
func (filter *Filter) AddWord(words ...string) {
	filter.trie.Add(words...)
}

// DelWord 删除关键字
func (filter *Filter) DelWord(words ...string) {
	filter.trie.Del(words...)
}

// FindIn 检测关键字 -> 规则1
// 规则 1.【肯德基】匹配：肯德基麦辣鸡腿堡，买一送一
func (filter *Filter) FindIn(text string) (bool, string) {
	text = filter.RemoveNoise(text)
	return filter.trie.FindIn(text)
}

// RemoveNoise 去除无效特殊字符
// 其中规则匹配过程中忽略特殊符号如[]、{}、()
func (filter *Filter) RemoveNoise(text string) string {
	return filter.noise.ReplaceAllString(text, "")
}
