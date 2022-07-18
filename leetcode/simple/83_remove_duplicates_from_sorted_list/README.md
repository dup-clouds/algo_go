# [83. remove-duplicates-from-sorted-list](https://leetcode.cn/problems/remove-duplicates-from-sorted-list/)

## 题目

给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。

**示例1：**

<img src="https://assets.leetcode.com/uploads/2021/01/04/list1.jpg">

~~~
输入：head = [1,1,2]
输出：[1,2]
~~~

**示例2：**

<img src="https://assets.leetcode.com/uploads/2021/01/04/list2.jpg">

~~~
输入：head = [1,1,2,3,3]
输出：[1,2,3]
~~~

## 题目大意

同移除数组中的重复元素，链表为升序，移除完成后返回链表头节点

## 解题思路

**方法一：**
1. 定义slow+fast指针，dummy节点
2. fast每移动一步，如fast值等于slow值，则删除fast节点，同时slow.next=fast.next
3. 重复2操作，直至到链表尾结点
4. 返回dummy.next节点
