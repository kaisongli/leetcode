package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Stu struct {
	P      Person
	School string
}
type Person struct {
	Name string
	Age  int
}

func main() {
	test()
}

//绿灯20s， 黄灯3，红23，一秒倒计时减1，绿黄红循环跳
func test() {
	m := map[string]int{
		"green":  20,
		"yellow": 3,
		"red":    23,
	}
	for true {
		for k, v := range m {
			for v > 0 {
				time.Sleep(time.Second)
				fmt.Println(fmt.Sprintf("%s: %d", k, v))
				v--
			}
		}
	}
}

func tuoFeng(s string) string {
	if len(s) < 2 {
		return s
	}
	ans := []byte{}
	ans = append(ans, s[0], s[1])
	left := 2
	for i := 2; i < len(s); i++ {
		if s[i] != s[i-2] {
			ans = append(ans, s[i])
			left++
		} else {
			if left > 2 {
				left = 2
			}
			for left > 0 {
				ans = ans[:len(ans)-1]
				left--
			}
		}
	}
	return string(ans)
}

func send(ch chan int) {
	defer fmt.Println("子协程结束....")
	ch <- 555
	fmt.Println("子协程开始....")

}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	ans := make([][]string, 0)
	m := make(map[string]int)
	idx := 0
	for _, str := range strs {
		chars := strings.Split(str, " ")
		sort.Strings(chars)
		toKey := strings.Join(chars, " ")
		if _, ok := m[toKey]; ok {
			ans[m[toKey]] = append(ans[m[toKey]], str)
		} else {
			m[toKey] = idx
			idx++
			ans = append(ans, []string{str})
		}
	}
	return ans
}
