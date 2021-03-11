package tree

func main() {

}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var preoder func(*TreeNode)
	result := []int{}
	preoder = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		preoder(node.Left)
		preoder(node.Right)
	}
	preoder(root)
	return result
}

func inorderTraversal(root *TreeNode) []int {
	var inorder func(*TreeNode)
	result := []int{}
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return result
}

func postorderTraversal(root *TreeNode) []int {
	var postorder func(*TreeNode)
	result := []int{}
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		result = append(result, node.Val)
	}
	postorder(root)
	return result
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

// func maxDepth(root *TreeNode) int {
// 	maxValue := 0
// 	var max func(*TreeNode, int)
// 	max = func(node *TreeNode, i int) {
// 		if node == nil {
// 			return
// 		}
// 		if node.Left == nil && node.Right == nil {
// 			if i > maxValue {
// 				maxValue = i
// 			}
// 		}
// 		max(node.Left, i+1)
// 		max(node.Right, i+1)
// 	}
// 	max(root, 1)
// 	return maxValue
// }

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func isSymmetric(root *TreeNode) bool {
	return isSame(root.Left, root.Right)
}

func isSame(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSame(p.Left, q.Right) && isSame(p.Right, q.Left)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
