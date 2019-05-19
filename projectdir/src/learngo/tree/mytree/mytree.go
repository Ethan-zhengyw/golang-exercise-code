package mytree

import "learngo/tree"

type Node struct {
    Node *tree.TreeNode
}

func PostOrder(root *Node) {
    if root == nil || root.Node == nil {
        return
    }

    left := &Node{Node: root.Node.Left}
    right := &Node{Node: root.Node.Right}

    PostOrder(left)
    PostOrder(right)

    root.Node.Print2()
}
