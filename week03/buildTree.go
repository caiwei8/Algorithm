/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder, postorder []int) *TreeNode {
    // 后序数组的最后一个元素的指针，每递归一次减一
	postorderLastIndex:=len(postorder)-1 
	var build func(inorder []int) *TreeNode
	build = func(inorder []int) *TreeNode {
		if len(inorder) <= 0 {
			return nil
		}
		rootVal := postorder[postorderLastIndex]
		postorderLastIndex--
		var root = &TreeNode{Val: rootVal}
        // 根据根节点将中序数组一分为二，继续递归处理
		leftInorder, rightInorder := getInorder(inorder, rootVal)
		root.Right = build(rightInorder)
		root.Left = build(leftInorder)
		return root
	}
	return build(inorder)
}


func getInorder(inorder []in rootVal int) (left, right []int) {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			return inorder[:i], inorder[i+1:]
		}
	}
	return
}
