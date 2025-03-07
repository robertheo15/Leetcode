/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums)==0{
        return nil
    }
    index:=findMax(nums)
    node:=&TreeNode{Val:nums[index]}
    node.Left=constructMaximumBinaryTree(nums[:index])
    node.Right=constructMaximumBinaryTree(nums[index+1:])
    return node
}

func findMax(nums []int)int{
    index:=-1; max:=-1
    for i:=0; i<len(nums); i++ {
        if max < nums[i]{
            index=i
            max=nums[i]
        }
    }
    return index
}
