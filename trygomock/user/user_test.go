package user

import (
	"testing"

	"github.com/golang/mock/gomock"
	"imooc.com/sb/learngo/trygomock/mock"
)

func TestUser_GetUserInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var id int64 = 1
	mockMale := mock.NewMockMale(ctl) // 创建一个新的mock实例
	// EXPECT()返回一个允许调用者设置期望和返回值的对象
	// Get(id) 是设置入参并调用 mock 实例中的方法。Return(nil) 是设置先前调用的方法出参
	gomock.InOrder(
		mockMale.EXPECT().Get(id).Return(nil),
	)

	user := NewUser(mockMale)
	err := user.GetUserInfo(id)
	if err != nil {
		t.Errorf("user.GetUserInfo err: %v", err)
	}
}
