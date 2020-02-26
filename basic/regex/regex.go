package main

import (
	"fmt"
	"regexp"
	"strings"
)

const text = "My email is bravealan.adv@gmail.com.cn"
const infoRe = `<div class="des f-cl".+?>(.+?)</div>`
const textRe = `[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`
const info = `<div class="des f-cl" data-v-3c42fade>杭州 | 26岁 | 中专 | 未婚 | 165cm | 5001-8000元</div> <div class="actions" data-v-3c42fade><div class="item sayHi" data-v-3c42fade>打招呼</div> <div class="item sendMsg" data-v-3c42fade>发消息</div> <div class="item hongliang" data-v-3c42fade>红娘牵线</div></div></div></div></div></div> <div class="CONTAINER f-topIndex" `

func main() {
	re := regexp.MustCompile(textRe) // 正则表达式
	match := re.FindString(text)
	fmt.Println(match)

	re = regexp.MustCompile(infoRe) // 正则表达式
	submatch := re.FindSubmatch([]byte(info))
	basicInfo := string(submatch[1])
	fmt.Println(basicInfo)
	labels := strings.Split(basicInfo, "|")
	fmt.Println(labels)
	TrimSplice(labels)
	fmt.Println(labels)
}

func TrimSplice(arr []string) {
	for i, str := range arr {
		arr[i] = strings.Trim(str, " ")
	}
}
