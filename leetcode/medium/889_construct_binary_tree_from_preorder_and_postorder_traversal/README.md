# [889. construct-binary-tree-from-preorder-and-postorder-traversal](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/)

## 题目
**根据前序和后序遍历构造二叉树**

给定两个整数数组，preorder和 postorder ，其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，postorder 是同一棵树的后序遍历，重构并返回二叉树。
如果存在多个答案，您可以返回其中 任何 一个。


示例 1：

![](https://assets.leetcode.com/uploads/2021/07/24/lc-prepost.jpg)

~~~
输入：preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
输出：[1,2,3,4,5,6,7]
~~~

示例 2：

~~~
输入: preorder = [1], postorder = [1]
输出: [1]
~~~

## 题目大意

给定二叉树的前序遍历和后序遍历，构建二叉树

## 解题思路

**方法一：**
1. 先确定root节点即前序遍历的第一个元素
2. 确定前序遍历的第二个元素，即二叉树的第一个左子树root节点
3. 递归构建左子树
4. 递归构建右子树
5. 递归终止条件：startIndex>endIndex，其中当startIndex==endStart时终止

要点：
1. 确定两个数组左右子树范围
2. 动态计算左子树节点个数-->推导前序遍历左子树结束索引位置
3. 最后一个节点需要特殊判断