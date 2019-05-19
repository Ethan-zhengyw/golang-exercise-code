package tree

import "fmt"

type TreeNode struct {
    Value int
    Left, Right *TreeNode
}

func CreateNode(Value int) *TreeNode {
    return &TreeNode{Value: Value}
}

func Print(node TreeNode) {
    fmt.Println(node.Value)
}

func (node TreeNode) Print2() {
    fmt.Println(node.Value)
}
