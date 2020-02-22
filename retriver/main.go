package main

import (
	"fmt"
	"time"

	"imooc.com/sb/learngo/retriver/mock"
	"imooc.com/sb/learngo/retriver/real"
)

const url = "http://www.imooc.com"

// 接口里面保存两个值，一个实现者的类型，一个实现者的指针/值
// 指针接收者实现只能以指针方式使用，值接收者都可
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "jojo",
			"course": "golang",
		})
}

type RetriverPoster interface {
	Retriever
	Poster
}

func session(s RetriverPoster) string {

	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

// Type switch
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Println(" > Type switch:")
	switch v := r.(type) {
	case *mock.MockRetriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r Retriever
	retriever := &mock.MockRetriever{Contents: "this is a fake imooc.com"}
	inspect(retriever)

	//fmt.Println(download(r))
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	//fmt.Println(download(r))

	// Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(*mock.MockRetriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(retriever))
}
