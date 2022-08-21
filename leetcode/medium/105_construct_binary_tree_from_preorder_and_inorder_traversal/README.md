# [105. construct-binary-tree-from-preorder-and-inorder-traversal](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

## 题目

给定两个整数数组preorder 和 inorder，其中preorder 是二叉树的先序遍历， inorder是同一棵树的中序遍历，请构造二叉树并返回其根节点。

示例 1：

![示例1](https://assets.leetcode.com/uploads/2021/02/19/tree.jpg)

~~~
输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]
~~~

示例 2：

~~~
输入: preorder = [-1], inorder = [-1]
输出: [-1]
~~~

## 题目大意

1. 给定一个二叉树的前序遍历和中序遍历，还原二叉树并返回root节点
2. 其中二叉树中每个节点的元素为不重复节点

## 解题思路

**方法一：**

1. 前序遍历的第一个节点即为root节点
2. 在中序遍历中查找root节点，该节点的下标即为前序遍历左子树的最后一个节点
3. 分隔二叉树的左右子树，递归构造二叉树的左子树和右子树
4. 其中寻找中序遍历结果中的root节点可提前构造map进行查找

**注意点**

1. 构建左子树或右子树时取的是preStart位置元素，而非位置0的元素
2. 递归构建左右子树时，前序遍历中的左右临界点取法：不能使用中序遍历中的index，需要计算左子树节点个数，即通过中序遍历中的index-inStart即为左子节点个数
3. 前序遍历中左右临界点即为：index-inStart+preStart