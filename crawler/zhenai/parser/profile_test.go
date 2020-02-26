package parser

import (
	"io/ioutil"
	"testing"

	"imooc.com/sb/learngo/crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "西西")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result)
	}
	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Name:      "西西",
		Age:       "26岁",
		Height:    "165cm",
		Income:    "5001-8000元",
		Marriage:  "未婚",
		Education: "中专",
		House:     "杭州",
	}
	if profile.Age != expected.Age {
		t.Errorf("expected Age %s; but was %s", expected.Age, profile.Age)
	}
	if profile.Income != expected.Income {
		t.Errorf("expected Income %s; but was %s", expected.Income, profile.Income)
	}
	if profile.Height != expected.Height {
		t.Errorf("expected Height %s; but was %s", expected.Height, profile.Height)
	}
	if profile.Income != expected.Income {
		t.Errorf("expected Income %s; but was %s", expected.Income, profile.Income)
	}
	if profile.Marriage != expected.Marriage {
		t.Errorf("expected Marriage %s; but was %s", expected.Marriage, profile.Marriage)
	}
	if profile.Education != expected.Education {
		t.Errorf("expected Education %s; but was %s", expected.Education, profile.Education)
	}
	if profile.House != expected.House {
		t.Errorf("expected House %s; but was %s", expected.House, profile.House)
	}

	//t.Errorf("profile: %v", profile)

}
