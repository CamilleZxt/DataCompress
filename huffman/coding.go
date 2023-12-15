package huffman

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
)

type TreeNode struct {
	Value      string  //1位长度字符
	Weight     int
	LeftChild  *TreeNode
	RightChild *TreeNode
	NodeParent *TreeNode
}

type TreeNodes []TreeNode

//type RootNode struct {
//	Root *TreeNode
//}

func (n TreeNodes) Len() int {
	return len(n)
}

func (n TreeNodes) Less(i, j int) bool {
	return n[i].Weight > n[j].Weight
}

func (n TreeNodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n TreeNode) IsLeaf() (bool) {
	return n.LeftChild == nil && n.RightChild == nil
}

func CalWeight(str string) map[string]int {
	tmpMap := make(map[string]int)
	for i:=0;i<len(str);i++ {
		tmpMap[string(str[i])] += 1
	}
	return tmpMap
}

func SortedNodes(maps map[string]int) []TreeNode {
	nodes := make(TreeNodes, len(maps))
	i := 0
	for value, weight := range maps {
		nodes[i] = TreeNode{Value: value, Weight: weight}
		i++
	}
	sort.Sort(sort.Reverse(nodes))
	return nodes
}

func CreateTree(nodes TreeNodes) *TreeNode {
	if len(nodes) < 2 {
		panic("Must contain 2 or more emlments")
	}
	tree :=&TreeNode{Weight: nodes[0].Weight + nodes[1].Weight, LeftChild: &nodes[0], RightChild: &nodes[1]}
	for i := 2; i < len(nodes); {
		if nodes[i].Weight == 0 {
			i++
			continue
		}
		oldRoot := tree
		if i+1 < len(nodes) && tree.Weight > nodes[i+1].Weight {
			newNode := TreeNode{Weight: nodes[i].Weight + nodes[i+1].Weight, LeftChild: &nodes[i], RightChild: &nodes[i+1]}
			tree = &TreeNode{Weight: newNode.Weight + oldRoot.Weight, LeftChild: oldRoot, RightChild: &newNode}
			i += 2
		} else {
			tree = &TreeNode{Weight: nodes[i].Weight + oldRoot.Weight, LeftChild: oldRoot, RightChild: &nodes[i]}
			i++
		}
	}
	return tree
}

var treeCode string
func (n TreeNode) traverse(codes string, visit func(string, string)) {
	if n.Value == ""{
		treeCode = treeCode + "0"
	} else {
		treeCode = treeCode + "1"
		treeCode = treeCode + BytesToBinaryString([]byte(n.Value))
	}

	if leftNode := n.LeftChild; leftNode != nil {
		leftNode.traverse(codes+"0", visit)
	} else {
		visit(n.Value, codes)
		return
	}
	n.RightChild.traverse(codes+"1", visit)
}

var index = 0
func convert2Map(parentNode *TreeNode) (*TreeNode) {
	length := len(treeCode)
	if length - 1 == index {
		return nil
	}
	if treeCode[index] == '0'{
		index++
		treeNode := &TreeNode{Value:"",NodeParent:parentNode}
		treeNode.LeftChild = convert2Map(treeNode)
		return treeNode
	} else if treeCode[index] == '1'{
		index = index + 9
		byteTmp := treeCode[index - 8:index]
		treeNode := &TreeNode{Value:string(BinaryStringToBytes(byteTmp)) }

		//if treeCode[index + 1] == '1' {
		//	index = index + 9
		//	nodeRight := &TreeNode{Value: string(BinaryStringToBytes(treeCode[index - 8:index]))}
		//}

		return treeNode
	}
	return nil
}

func Convert2Code(str string){

}

func Encode(str string) {//ABRACADABRA!
	maps := CalWeight(str)
	nodes := SortedNodes(maps)
	tree := CreateTree(nodes)

	encodeMap := make(map[string]string)
	tree.traverse("", func(value string, code string) {
		encodeMap[value] = code
	})

	var tmpStr string
	for i:=0;i<len(str);i++ {
		tmpStr = tmpStr + encodeMap[string(str[i])]
	}

	fmt.Println(treeCode) //字典
	fmt.Println(tmpStr)  //编码之后的字符串

	byte1 := BinaryStringToBytes(treeCode)
	byte2 := BinaryStringToBytes(tmpStr)
	fmt.Println(byte1)
	fmt.Println(byte2)
	fmt.Println(len(tmpStr)/8)
	fmt.Println(len(treeCode)/8)
	//treeMap := convert2Map(nil)

	//fmt.Println(treeMap)
}


const (
	zero  = byte('0')
	one   = byte('1')
	//space = byte(' ')
)
var uint8arr [8]uint8

// ErrBadStringFormat represents a error of input string's format is illegal .
var ErrBadStringFormat = errors.New("bad string format")

// ErrEmptyString represents a error of empty input string.
var ErrEmptyString = errors.New("empty string")

func init() {
	uint8arr[0] = 128
	uint8arr[1] = 64
	uint8arr[2] = 32
	uint8arr[3] = 16
	uint8arr[4] = 8
	uint8arr[5] = 4
	uint8arr[6] = 2
	uint8arr[7] = 1
}

// append bytes of string in binary format.
func appendBinaryString(bs []byte, b byte) []byte {
	var a byte
	for i := 0; i < 8; i++ {
		a = b
		b <<= 1
		b >>= 1
		switch a {
		case b:
			bs = append(bs, zero)
		default:
			bs = append(bs, one)
		}
		b <<= 1
	}
	return bs
}

// ByteToBinaryString get the string in binary format of a byte or uint8.
func ByteToBinaryString(b byte) string {
	buf := make([]byte, 0, 8)
	buf = appendBinaryString(buf, b)
	return string(buf)
}

// BytesToBinaryString get the string in binary format of a []byte or []int8.
func BytesToBinaryString(bs []byte) string {
	l := len(bs)
	bl := l*8 //+ l + 1
	buf := make([]byte, 0, bl)
	//buf = append(buf, lsb)
	for _, b := range bs {
		buf = appendBinaryString(buf, b)
		//buf = append(buf, space)
	}
	//buf[bl-1] = rsb
	return string(buf)
}

// regex for delete useless string which is going to be in binary format.
var rbDel = regexp.MustCompile(`[^01]`)

// BinaryStringToBytes get the binary bytes according to the
// input string which is in binary format.
func BinaryStringToBytes(s string) (bs []byte) {
	if len(s) == 0 {
		panic(ErrEmptyString)
	}
	s = rbDel.ReplaceAllString(s, "")
	l := len(s)
	if l == 0 {
		panic(ErrBadStringFormat)
	}
	mo := l % 8
	l /= 8
	if mo != 0 {
		l++
	}
	bs = make([]byte, 0, l)
	mo = 8 - mo
	var n uint8
	for i, b := range []byte(s) {
		m := (i + mo) % 8
		switch b {
		case one:
			n += uint8arr[m]
		}
		if m == 7 {
			bs = append(bs, n)
			n = 0
		}
	}
	return
}