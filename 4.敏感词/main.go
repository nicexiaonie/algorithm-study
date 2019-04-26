package main

import (
	"algorithm/4.敏感词/sensitive"
)

func main() {

	dict := []string{"傻逼", "我们", "我们遇到"}
	sensitive.Init()
	for _, v := range dict {
		sensitive.AppendWord(v, v)
	}

	//fmt.Println(sensitive.Sensitive.Root.Node[37027].Node[20123])
	content := "那些年我们遇到了好多傻逼."
	sensitive.Search(content)

}
