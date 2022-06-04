package main

import (
	"fmt"
	"strings"
)

func main() {
	emails := []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}
	fmt.Println()
	fmt.Print(numUniqueEmails(emails))
}

// 独特的电子邮件地址
// @link https://leetcode.cn/problems/unique-email-addresses/
func uniqueEmailAddresses(emails []string) int {
	if len(emails) < 2 {
		return len(emails)
	}
	uniqueMap := make(map[string]int)
	sum := 0
	for _, v := range emails {
		tempEmail := ""
		strArray := strings.Split(v, "@")
		name := strArray[0]
		for _, ch := range name {
			if ch == '+' {
				break
			} else if ch != '.' {
				tempEmail += string(ch)
			}
		}
		key := tempEmail + "@" + strArray[1]
		if uniqueMap[key] < 1 {
			uniqueMap[key] = 1
			sum++
		}
	}
	return sum
}

// 优化
func numUniqueEmails(emails []string) int {
	// set集合 存储已处理的邮箱地址
	emailSet := map[string]bool{}
	for _, email := range emails {
		// 获取@符号对应下标
		index := strings.IndexByte(email, '@')
		// 去除+及后面字符
		localName := strings.SplitN(email[:index], "+", 2)[0]
		// 将.全部替换为空字符串
		localName = strings.ReplaceAll(localName, ".", "")
		// set 设置值
		emailSet[localName+email[index:]] = true
	}
	// 返回set大小 即不重复邮箱个数
	return len(emailSet)
}
