package parser

import (
	"gorobot/engine"
	"gorobot/module"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents,
		"http://album.zhenai.com/u/1792597478",
		"1792597478",
		"Amy8月的雨")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1"+"element;but was %v", result.Items)
	}

	/*
		"basicInfo": ["未婚",
		"30岁",
		 "天蝎座(10.23-11.21)",
		 "160cm",
		 "49kg",
		 "工作地:成都青羊区",
		 "月收入:8千-1.2万",
		 "经理",
		 "大专"],
		 "detailInfo": ["汉族",
		  "籍贯:四川广元",
			"体型:一般",
			"不吸烟",
			"不喝酒",
			"已购房",
			"未买车",
			"没有小孩",
			"是否想要孩子:想要孩子",
			"何时结婚:时机成熟就结婚"],
			"educationString": "大专",
			"emotionStatus": 0,
		  "gender": 1,
			"genderString": "女士",
			"hasIntroduce": true,
			"heightString": "160cm",
			"hideVerifyModule": false,
			"introduceContent": "即便30.也依然坚持自己的想法，
			遵从自己的内心，希望能遇见三观一致、上进，能互相温
			暖的那个他，一起往后余生。愿往后的日子，有诗、
			有茶、有柔情！",
			"introducePraiseCount": 0,
			"isActive": true,
			"isFollowing": false,
			"isInBlackList": false,
			"isStar": false,
			"isZhenaiMail": true,
			"lastLoginTimeString": "当前在线",
			"liveAudienceCount": 0,
			"liveType": 0,
			"marriageString": "未婚",
			"memberID": 1792597478,
			"momentCount": 10,
			"nickname": "Amy8月的雨",
	*/
	actual := result.Items[0]
	expected := engine.Item{
		URL:  "https://album.zhenai.com/u/1792597478",
		Type: "zhenai",
		ID:   "1792597478",
		Payload: module.Profile{
			Name:       "Amy8月的雨",
			Gender:     "女",
			Height:     160,
			Age:        30,
			Weight:     49,
			Income:     "8001-12000元",
			Marriage:   "未婚",
			Education:  "大专",
			Occupation: "人事/行政",
			Hokou:      "四川广元",
			Xinzuo:     "天蝎座(10.23-11.21)",
			House:      "已购房",
			Car:        "未买车",
			WorkAddr:   "成都青羊区",
			WorkCity:   "成都",
		},
	}
	if actual != expected {
		t.Errorf("expected should have")
	}

}
