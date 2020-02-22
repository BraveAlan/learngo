package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"

	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

// 将错误抛出去，让外面的对象去做错误处理
func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start" +
			"with " + prefix)
	}
	path := request.URL.Path[len(prefix):] // /list/fib.txt
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
