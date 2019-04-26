package sensitive

import (
	"sync"
	"fmt"
)

type Lexicon struct {
}

type Node struct {
	Node map[rune]*Node
	End  bool
	Info interface{}
}

type Trie struct {
	Root  *Node
	Mutex sync.RWMutex
}

type resultInfo struct {
	StartIndex int
	EndIndex   int
	Word       string
}

var Sensitive = new(Trie)

func Init() {
	Sensitive.Root = newNode()
}
func newNode() *Node {
	n := new(Node)
	n.End = false
	n.Node = make(map[rune]*Node)
	return n
}

func AppendWord(text string, info interface{}) {

	chars := []rune(text)
	node := Sensitive.Root
	for _, char := range chars {
		if _, ok := node.Node[char]; !ok {
			node.Node[char] = newNode()
			node.Node[char].Info = string(char)
		}
		node = node.Node[char]
	}
	node.End = true
}

func Search(text string) {
	chars := []rune(text)
	result := make([]*resultInfo, 0)
	for k := range chars {
		node := Sensitive.Root
		i := 0
		inWord := ""
		for {
			if k+i >= len(chars) {
				break
			}
			node = node.Node[chars[k+i]]
			inWord += string(chars[k+i])
			if node != nil {
				if node.End {
					resultInfo := new(resultInfo)
					resultInfo.StartIndex = k
					resultInfo.EndIndex = k + i
					resultInfo.Word = inWord
					result = append(result, resultInfo)
				}
				if len(node.Node) == 0 {
					break
				}
			} else {
				break
			}
			i++
		}
	}
	for _, v := range result {
		fmt.Println(v)
	}
}
