# [380. find-all-anagrams-in-a-string](https://leetcode.cn/problems/insert-delete-getrandom-o1/)

## 题目
实现RandomizedSet 类：
- RandomizedSet() 初始化 RandomizedSet 对象
- bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
- bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
- int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。<br/>

你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。


示例 1：
~~~
输入
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
输出
[null, true, false, true, 2, true, false, 2]

解释
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2 。
randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
~~~

## 题目大意

1. 实现的数据结构，使其满足在O(1)时间复杂度快速插入、删除、随机获取一个元素

## 解题思路

**方法一：**
1. 使用数组+散列表实现
2. 插入操作：数组尾部增加元素，散列表对应存储值对应的数组下标
3. 删除操作：通过值从散列表中获取数据的索引，将该值与数据最后一个元素进行交互完成删除
4. 随机获取：即随机获取数据元素-随机生成一个数组下标