package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	deadends := []string{"0000"}
	target := "8888"
	fmt.Println(openLock(deadends, target))
}

/**
@link https://leetcode.cn/problems/open-the-lock/
752. 打开转盘锁
*/
func openLock(deadends []string, target string) int {
	// 等于初始密码则响应0
	if target == "0000" {
		return 0
	}
	// 队列 记录每层的节点
	var queue []string
	// 记录已遍历过的密码，避免重复死循环
	visited := make(map[string]int)
	for _, v := range deadends {
		visited[v] = 1
	}
	// 初始化密码再deadends中则直接返回
	if visited["0000"] > 0 {
		return -1
	}
	// 初始密码如队列
	queue = append(queue, "0000")
	visited["0000"] = 1
	// 初始化步数
	step := 0
	// 队列有元素时执行
	for len(queue) > 0 {
		// 遍历队列
		for _, v := range queue {
			curr := v
			queue = queue[1:]
			if curr == target {
				return step
			}
			// 对每个密码的四位进行遍历往上滑动或往下滑动
			// 滑动过程中如遇重复则跳过 否则入队列 入已遍历元素map
			for i := 0; i < 4; i++ {
				upStr := up(curr, i)
				if visited[upStr] < 1 {
					queue = append(queue, upStr)
					visited[upStr] = 1
				}
				downStr := down(curr, i)
				if visited[downStr] < 1 {
					queue = append(queue, downStr)
					visited[downStr] = 1
				}
			}
		}
		step++
	}
	return -1
}

// 向上滑动
func up(str string, i int) string {
	curr := strings.Split(str, "")
	if curr[i] == "9" {
		curr[i] = "0"
	} else {
		v, _ := strconv.Atoi(curr[i])
		curr[i] = strconv.Itoa(v + 1)
	}
	return strings.Join(curr, "")
}

// 向下滑动
func down(str string, i int) string {
	curr := strings.Split(str, "")
	if curr[i] == "0" {
		curr[i] = "9"
	} else {
		v, _ := strconv.Atoi(curr[i])
		curr[i] = strconv.Itoa(v - 1)
	}
	return strings.Join(curr, "")
}
