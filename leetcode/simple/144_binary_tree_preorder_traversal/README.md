# [144. binary-tree-preorder-traversal](https://leetcode.cn/problems/binary-tree-preorder-traversal/)

## 题目
**二叉树的前序遍历**

给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

示例 1：

<img src="https://assets.leetcode.com/uploads/2020/09/15/inorder_1.jpg">

~~~
输入：root = [1,null,2,3]
输出：[1,2,3]
~~~

示例 2：
~~~
输入：root = []
输出：[]
~~~

示例 3：

~~~
输入：root = [1]
输出：[1]
~~~

## 题目大意

输出二叉树的前序遍历结果

## 解题思路

**方法一：**
1. 递归遍历
2. 先root节点，再左节点，再右节点
