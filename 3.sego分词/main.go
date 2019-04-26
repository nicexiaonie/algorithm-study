package main

import (
	"fmt"
	"github.com/huichen/sego"
)

func main()  {
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("3.分词/dictionary.txt")

	// 分词
	text := []byte("支持普通和搜索引擎两种分词，支持用户词典、词性标注，可运行JSON RPC服务。")
	segments := segmenter.InternalSegment(text, true)

	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
	//fmt.Println(sego.SegmentsToString(segments, true))
	output := sego.SegmentsToSlice(segments, false)

	for _, v := range output {
			fmt.Println(v)
	}
}
