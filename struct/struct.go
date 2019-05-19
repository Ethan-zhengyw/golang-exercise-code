package main

import "fmt"

type TreeNode struct {
    value int
    left, right *TreeNode
}

func createNode(value int) *TreeNode {
    return &TreeNode{value: value}
}

func print(node TreeNode) {
    fmt.Println(node.value)
}

func (node TreeNode) print2() {
    fmt.Println(node.value)
}

func traverseTree(treeNode *TreeNode) {
    if treeNode == nil {
        return
    }
    traverseTree(treeNode.left)
    treeNode.print2()
    traverseTree(treeNode.right)
}

func main() {
    var root TreeNode
    fmt.Println(root)

    root = TreeNode{value: 5}
    root.right = new(TreeNode)
    root.right.left = &TreeNode{6, nil, nil}
    root.right.right = createNode(2)

    fmt.Println(root)
    
    print(root)
    root.print2()
    
    fmt.Println("Traversing...")
    traverseTree(&root)
}
