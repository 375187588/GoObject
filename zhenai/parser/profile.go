package parser

import (
	"gorobot/engine"
	"gorobot/module"
	"gorobot/util"
	"math/rand"
	"regexp"
)

var ageRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">*(\d+)岁</div>`)
var marrgeRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`)
var heightRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">(\d+)cm</div>`)
var weightRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">(\d+)kg</div>`)
var incomeRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`)
var educationRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`)
var workCityRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">工作地:([^<]+)</div>`)
var workAddrRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`)
var xinzuoRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`)
var hukouRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn pink">([^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn pink">([^<]+)</div>`)
var carRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn pink">([^<]+)</div>`)
var guessRe = regexp.MustCompile(`<a class = "exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^>])`)

//ParseProfile function
func ParseProfile(
	contents []byte,
	surl string,
	id string,
	name string) engine.ParseResult {
	profile := module.Profile{}
	profile.Name = name
	/*
		match := ageRe.FindSubmatch(contents)
		if match != nil {
			age, err := strconv.Atoi(
				extractString(contents, ageRe))
			if err != nil {
				profile.Age = age
			}
		}

		match = heightRe.FindSubmatch(contents)
		if match != nil {
			height, err := strconv.Atoi(
				extractString(contents, heightRe))
			if err != nil {
				profile.Height = height
			}
		}
		match = weightRe.FindSubmatch(contents)
		if match != nil {
			weight, err := strconv.Atoi(
				extractString(contents, weightRe))
			if err != nil {
				profile.Weight = weight
			}
		}
	*/

	profile.Marriage = extractString(contents, marrgeRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Education = extractString(contents, educationRe)
	profile.WorkCity = extractString(contents, workCityRe)
	profile.WorkAddr = extractString(contents, workAddrRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.Hokou = extractString(contents, hukouRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	//模拟数据包
	profile.Age = util.RandInt(20, 55)
	profile.Height = util.RandInt(100, 230)
	profile.Weight = util.RandInt(55, 100)
	marr := [2]string{"已婚", "未婚"}
	profile.Marriage = marr[rand.Intn(1)]
	inc := [12]string{
		"500-1500元",
		"1501-2500元",
		"2501-3500元",
		"3501-4500元",
		"4501-5500元",
		"5501-6500元",
		"6501-7500元",
		"7501-9500元",
		"9501-15000元",
		"15001-25000元",
		"25001-30000元",
		"30000元以上",
	}
	profile.Income = inc[rand.Intn(11)]
	educa := [8]string{
		"小学",
		"中学",
		"高中",
		"大学本科",
		"大学专科",
		"硕士",
		"博士",
		"院士",
	}
	profile.Education = educa[rand.Intn(7)]
	workCity := [5]string{
		"北京",
		"上海",
		"杭州",
		"成都",
		"厦门",
	}
	profile.WorkCity = workCity[rand.Intn(4)]
	profile.WorkAddr = profile.WorkCity
	profile.Xinzuo = "白羊"
	profile.Hokou = "地球村"
	profile.House = "有房"
	profile.Car = "有车"

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				URL:     surl,
				Type:    "zhenai",
				ID:      id,
				Payload: profile,
			},
		},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//fmt.Printf("got profile:%s:%s:%s\n", m[1], m[2], m[3])
		result.Requests = append(result.Requests,
			engine.Request{GUrl: string(m[1]), ParserFunc: ProfileParser(string(m[2]), string(m[2]))})
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := marrgeRe.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

// ProfileParser name string
func ProfileParser(id string, name string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, url, id, name)
	}
}
