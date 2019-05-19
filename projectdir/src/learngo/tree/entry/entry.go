package main

import "fmt"
import "learngo/tree"
import "learngo/tree/mytree"

func main() {
    var root tree.TreeNode
    fmt.Println(root)

    root = tree.TreeNode{Value: 5}
    root.Right = new(tree.TreeNode)
    root.Right.Left = &tree.TreeNode{6, nil, nil}
    root.Right.Right = tree.CreateNode(2)

    fmt.Println(root)
    
    tree.Print(root)
    root.Print2()
    
    fmt.Println("Traversing...")
    tree.TraverseTree(&root)

    newroot := tree.TreeNode{Value: 5}
    myRoot := mytree.Node{Node: &newroot}
    fmt.Println("Traversing post order...")
    mytree.PostOrder(&myRoot)
}
