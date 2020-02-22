package mock

import "fmt"

type MockRetriever struct {
	Contents string
}

func (r *MockRetriever) String() string {
	return fmt.Sprintf("Retriever: {Contents=%s}", r.Contents)
}

func (r *MockRetriever) Post(url string,
	form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *MockRetriever) Get(url string) string {
	return r.Contents
}
