package client

import (
	"fmt"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	//req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("req:", req)
			fmt.Println("via:", req)

			return nil
		},
	}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	//resp, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//s, err := httputil.DumpResponse(resp, true)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s\n", s)

	//resp, err = http.Get("http://www.imooc.com")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//
	//s, err := httputil.DumpResponse(resp, true)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s\n", s)

}
