package _116_populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	ret := root
	var queue []*Node
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			ret = queue[i]
			if i+1 < size {
				ret.Next = queue[i+1]
			} else {
				ret.Next = nil
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
	}
	return root
}

func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}
	dfs(root)
	return root
}

func dfs(root *Node) {
	if root.Left == nil {
		return
	}
	root.Left.Next = root.Right
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	dfs(root.Left)
	dfs(root.Right)
}
