# [226. invert-binary-tree](https://leetcode.cn/problems/invert-binary-tree/)

## 题目
**翻转二叉树**

给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

示例 1：

<img src="https://assets.leetcode.com/uploads/2021/03/14/invert1-tree.jpg">

~~~
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]
~~~

示例 2：

<img src="https://assets.leetcode.com/uploads/2021/03/14/invert2-tree.jpg">

~~~
输入：root = [2,1,3]
输出：[2,3,1]
~~~

示例 3：

~~~
输入：root = []
输出：[]
~~~


## 题目大意

给定一颗二叉树，将二叉树的所有左右节点进行翻转

## 解题思路

**方法一：**
1. 先进行根节点的左右节点翻转
2. 递归进行左节点翻转
3. 递归进行右节点翻转
4. 返回root节点即可
