package version

import (
	"imooc.com/sb/learngo/trygomock/version/spider"
)

// GetGoVersion 获取Go的版本信息
func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
