/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    parentVal := root.Val
    pVal := p.Val
    qVal := q.Val
    
    if pVal > parentVal && qVal > parentVal {
        return lowestCommonAncestor(root.Right, p, q)
    } else if pVal < parentVal && qVal < parentVal {
        return lowestCommonAncestor(root.Left, p, q)
    } 
    return root 
}
