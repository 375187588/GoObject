package fetcher

//说明：Fetcher 从指定URL获取数据
//创建日期：2019/10/14
//Copyright:go工作者
//修改：last edit

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

//determineEncoding 是判断是UTF8还是GBK 是GBK解码成UTF8
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error:%v\n", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

//限制速率
var rateLimiter = time.Tick(100 * time.Millisecond)

//Fetch 从URL获得URL源码数据
func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	
	//构造http请求
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	//访问手机版本还是访问电脑版本
	//请填写：User-Agent
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")

	//http.Client创建
	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {
			//fmt.Println("Redirect1", req)
			return nil
		},
	}

	//执行http请求
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	//最后关闭
	defer resp.Body.Close()

	//返回不成功，返回失败状态代码
	if resp.StatusCode != http.StatusOK {

		fmt.Println("Error:status code", resp.StatusCode)
		return nil,
			fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	//GBK转UTF8编码
	bodyReader := bufio.NewReader(resp.Body)
	encode := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, encode.NewDecoder())
	fetchData, err := ioutil.ReadAll(utf8Reader)

	if err != nil {
		return nil, err
	}

	return fetchData, nil
}
