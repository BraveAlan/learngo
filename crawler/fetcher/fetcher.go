package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/text/encoding/unicode"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	e, respBodyReader := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(respBodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}

func determineEncoding(r io.Reader) (encoding.Encoding, *bufio.Reader) {

	reader := bufio.NewReader(r)
	bytes, err := reader.Peek(1024) // 对bytes这个对象来说，Peek之后不会有指针偏移，但是对r这个对象来说，会偏移
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8, reader
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e, reader
}
