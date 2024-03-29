package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents, "")
	const resultSize = 540

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCites := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}
	fmt.Printf("%v\n", expectedCites)

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d"+
			"requests: but had %d",
			resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].GUrl != url {
			t.Errorf("expected url#%d:%s; but "+
				"was %s", i, url, result.Requests[i].GUrl)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d"+
			"requests: but had %d",
			resultSize, len(result.Items))
	}

	// for i, city := range expectedCites {
	// 	if result.Items[i].(string) != city {
	// 		t.Errorf("expected city#%d:%s; but "+
	// 			"was %s", i, city, result.Items[i].(string))
	// 	}
	// }
}
