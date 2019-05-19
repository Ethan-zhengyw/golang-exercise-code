package tree

func TraverseTree(treeNode *TreeNode) {
    if treeNode == nil {
        return
    }
    TraverseTree(treeNode.Left)
    treeNode.Print2()
    TraverseTree(treeNode.Right)
}
