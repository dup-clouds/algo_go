# [106. construct-binary-tree-from-inorder-and-postorder-traversal](https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)

## 题目

给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗二叉树。


![示例1](https://assets.leetcode.com/uploads/2021/02/19/tree.jpg)

~~~
输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
输出：[3,9,20,null,null,15,7]
~~~

示例 2：

~~~
输入：inorder = [-1], postorder = [-1]
输出：[-1]
~~~

## 题目大意

1. 给定一个二叉树的中序遍历和后序遍历，还原二叉树并返回root节点
2. 其中二叉树中每个节点的元素为不重复节点

## 解题思路

**方法一：**

1. 后序遍历的最后一个节点为root节点
2. 在中序遍历中找到root节点，即分隔左右子树，并计算左子树节点个数
3. 递归构建左右子树
4. 返回root节点