package version

import (
	"testing"

	"github.com/golang/mock/gomock"
	"imooc.com/sb/learngo/trygomock/version/spider"
)

func TestGetGoVersion(t *testing.T) {
	mockctl := gomock.NewController(t)
	defer mockctl.Finish()
	mockSpider := spider.NewMockSpider(mockctl) // 接口的实现对象
	time := 2
	mockSpider.EXPECT().GetBody().Return("go.1.14").Times(time)
	for i := 0; i < time; i++ {
		goVer1 := GetGoVersion(mockSpider)
		if goVer1 != "go.1.13" {
			t.Errorf("Get wrong version %s", goVer1)
		}
	}
}
