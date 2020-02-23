package main

import (
	"fmt"
	"regexp"
)

const text = "My email is bravealan.adv@gmail.com.cn"

func main() {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`) // 正则表达式
	match := re.FindString(text)
	fmt.Println(match)

}
