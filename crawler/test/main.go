package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/text/transform"

	"golang.org/x/text/encoding"

	"golang.org/x/net/html/charset"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	e, _, _, respBodyReader := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(respBodyReader, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	//file, err := os.Open("zhenghun.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//e, _, _, fileReader := determineEncoding(file)
	//utf8Reader := transform.NewReader(fileReader, e.NewDecoder())
	//all, err := ioutil.ReadAll(utf8Reader)
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println("当前打印内容：", string(all))
	printCityList(all)

}

func determineEncoding(r io.Reader) (encoding.Encoding, string, bool, *bufio.Reader) {

	reader := bufio.NewReader(r)
	bytes, err := reader.Peek(1024) // 对bytes这个对象来说，Peek之后不会有指针偏移，但是对r这个对象来说，会偏移
	//bytes := []byte(`<!DOCTYPE html><html lang="zh-cn"><head><meta charset="utf-8"><title>珍爱征婚网_真实征婚交友信息_免费发布征婚启事</title><meta name="keywords" content="征婚网站,征婚信息,征婚启事"><meta name="description" content="珍爱征婚网提供网上最新最全的真实征婚信息，大龄男女同城征婚快速速配在线聊天，所有征婚信息都经过严格审核，精准适配自己的有缘人，找对象快捷安全！"><link rel="alternate" media="only screen and (max-width: 640px)" href="http://m.zhenai.com/zhenghun"><meta name="mobile-agent" content="format=xhtml; url=http://m.zhenai.com/zhenghun"><meta name="mobile-agent" content="format=html5; url=http://m.zhenai.com/zhenghun"><meta name="mobile-agent" content="format=wml; url=http://m.zhenai.com/zhenghun"><meta http-equiv="Cache-Control" content="no-transform"><meta http-equiv="Cache-Control" content="no-siteapp"><!--[if lt IE 9]>
	//<meta http-equiv="refresh" content="0;url='https://www.zhenai.com/n/update_guidan`)
	fmt.Println("第一个1024个字节：", string(bytes))
	if err != nil {
		panic(err)
	}
	bytes, err = reader.Peek(1024)
	fmt.Println("第二个1024个字节：", string(bytes))
	if err != nil {
		panic(err)
	}
	e, name, certain := charset.DetermineEncoding(bytes, "")
	return e, name, certain, reader
}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(matches))
}
