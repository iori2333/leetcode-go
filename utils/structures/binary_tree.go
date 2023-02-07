package structures

import (
	"math"
	"strconv"
	"strings"
)

const NIL = math.MinInt64

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func NewTree(sequence []int) *TreeNode {
	n := len(sequence)
	if n == 0 {
		return nil
	}

	var queue []*TreeNode
	root := NewTreeNode(sequence[0])
	queue = append(queue, root)

	for i := 1; i < n; {
		node := queue[0]
		queue = queue[1:]

		if i < n && sequence[i] != NIL {
			node.Left = NewTreeNode(sequence[i])
			queue = append(queue, node.Left)
		}
		i++

		if i < n && sequence[i] != NIL {
			node.Right = NewTreeNode(sequence[i])
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func (root *TreeNode) Sequence() []int {
	var res []int
	var queue []*TreeNode

	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			nd := queue[i]
			if nd == nil {
				res = append(res, NIL)
			} else {
				res = append(res, nd.Val)
				queue = append(queue, nd.Left, nd.Right)
			}
		}
		queue = queue[size:]
	}

	i := len(res)
	for i > 0 && res[i-1] == NIL {
		i--
	}

	return res[:i]
}

func (root *TreeNode) String() string {
	builder := strings.Builder{}
	builder.WriteString("[")
	for i, v := range root.Sequence() {
		if i > 0 {
			builder.WriteString(", ")
		}
		if v == NIL {
			builder.WriteString("null")
		} else {
			builder.WriteString(strconv.Itoa(v))
		}
	}
	builder.WriteString("]")
	return builder.String()
}
