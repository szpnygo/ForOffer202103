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
