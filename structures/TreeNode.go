package structures

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var NULL = -1 << 63

// Nums2TreeNode 数组转二叉树
func Nums2TreeNode(nums []int) *TreeNode {
	n := len(nums)
	if n < 1 {
		return nil
	}
	// 确定根节点
	root := &TreeNode{Val: nums[0]}
	// 存储按层遍历的节点 生成一个容量为10 数组大小为1的的切片充当队列
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root
	i := 1
	// 遍历原始数组中的每一个元素
	for i < n {
		// 取出当前队列的第一个节点
		currNode := queue[0]
		// 队列更新
		queue = queue[1:]
		// 设置左节点 同时将左节点放入队列中
		if i < n && nums[i] != NULL {
			leftNode := &TreeNode{Val: nums[i]}
			currNode.Left = leftNode
			queue = append(queue, leftNode)
		}
		i++
		// 设置右节点 同时将右节点放入队列中
		if i < n && nums[i] != NULL {
			rightNode := &TreeNode{Val: nums[i]}
			currNode.Right = rightNode
			queue = append(queue, rightNode)
		}
		i++
	}
	return root
}

// GetTargetNode 在二叉树中寻找Val等于target的节点并返回
func GetTargetNode(root *TreeNode, target int) *TreeNode {
	if root == nil || root.Val == target {
		return root
	}
	leftRoot := GetTargetNode(root.Left, target)
	if nil != leftRoot {
		return leftRoot
	}
	return GetTargetNode(root.Right, target)
}

// Equal 判断二个二叉树是否相等
func (tn *TreeNode) Equal(a *TreeNode) bool {
	if nil == tn && a == nil {
		return true
	}
	if nil == tn || nil == a || tn.Val != a.Val {
		return false
	}
	return tn.Left.Equal(a.Left) && tn.Right.Equal(a.Right)
}

// Tree2Nums 把 *TreeNode 按照层还原层数组[]int
func Tree2Nums(tn *TreeNode) []int {
	ret := make([]int, 0, 1024)
	if tn == nil {
		return ret
	}
	queue := []*TreeNode{tn}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node == nil {
				ret = append(ret, NULL)
			} else {
				ret = append(ret, node.Val)
				queue = append(queue, node.Left, node.Right)
			}
		}
		queue = queue[size:]
	}
	endIndex := len(ret) - 1
	for endIndex >= 0 && ret[endIndex] == NULL {
		endIndex--
	}
	return ret[:endIndex+1]
}
