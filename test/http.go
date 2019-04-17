package main

import (
	"bytes"
	"crypto/tls"
	"dial/base"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// httpclient 超时2s 支持http/https
func NewHttpClient() *http.Client {
	return &http.Client{
		Timeout:   2 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
}

// 检查请求结果 通过返回码检查或者返回内容检查 content:比对内容 code:http状态码
func CkeckResponse(resp *http.Response, code int, content string) bool {
	if code > 0 && resp.StatusCode != code {
		return false
	}

	if len(content) > 0 {
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		if err != nil {
			return false
		}
		if !bytes.Contains(body, []byte(content)) {
			fmt.Println("no")
			return false
		}
		fmt.Println(content)
		fmt.Println("ok")
	}
	return true
}

// http/https get请求 url如http://192.168.5.41:8080/index 返回值:延时
func HttpGet(url, content string, code int) int64 {
	c := NewHttpClient()

	t := time.Now()
	resp, err := c.Get(url)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer resp.Body.Close()

	fmt.Println(resp)
	if !CkeckResponse(resp, code, content) {
		return 0
	}

	return time.Since(t).Nanoseconds() / 1000
}

// http/https post请求 datatype:post类型 data:post数据 返回值:延时
func HttpPost(url, content, datatype, data string, code int) int64 {
	c := NewHttpClient()

	t := time.Now()
	resp, err := c.Post(url, datatype, bytes.NewReader([]byte(data)))
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer resp.Body.Close()

	fmt.Println(resp)
	if !CkeckResponse(resp, code, content) {
		return 0
	}

	return time.Since(t).Nanoseconds() / 1000
}

// method:HEAD/PUT/GET等，不支持post data:比如method为POST的提交数据 返回值:延时
func HttpRequest(method, url, content, data string, code int) int64 {

	req, err := http.NewRequest(method, url, bytes.NewReader([]byte(data)))
	if err != nil {
		fmt.Println(err)
		return 0
	}

	client := NewHttpClient()

	t := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer resp.Body.Close()

	fmt.Println(resp)
	if !CkeckResponse(resp, code, content) {
		return 0
	}

	return time.Since(t).Nanoseconds() / 1000
}

func main() {
	fmt.Println(base.HttpGet("http://192.168.6.190:13636/", "企业智能DDI系统", 200))
	fmt.Println(base.HttpGet("http://192.168.6.190:13636/", "企业智能DDI系统xxx", 200))
	//fmt.Println(HttpGet("http://192.168.6.190:13636/", "", 0))
	fmt.Println(base.HttpGet("http://192.168.6.190:13636/", "企业智能DDI系统", 0))
	//fmt.Println(HttpGet("https://192.168.6.53:8081/", "service.ws.yamutech.com", 200))
	//fmt.Println(HttpGet("https://192.168.6.53:8081/", "service.ws.yamutech.com.xxx", 200))
	//fmt.Println(HttpGet("http://192.168.6.195:13636/", "企业智能DDI系统", 0))
	//fmt.Println(HttpPost("http://192.168.5.41:8080/login", "log", "application/x-www-form-urlencoded", "UserName=hww&PassWord=666", 200))
	//fmt.Println(HttpRequest("GET", "http://192.168.6.190:13636/", "企业智能DDI系统", "", 200))
	//fmt.Println(Http("HEAD", "http://192.168.6.190:13636/", "", "", 200))

	//fmt.Println(base.HttpRequest("GET", "http://192.168.6.190:13636/", "企业智能DDI系统", "", 200))
	fmt.Println(base.HttpRequest("GET", "http://192.168.6.190:13636/", "企业智能DDI系统xxx", "", 200))
	fmt.Println(base.HttpPost("http://192.168.5.41:8080/login", "log", "application/x-www-form-urlencoded", "UserName=hww&PassWord=666", 200))
	//fmt.Println(base.HttpRequest("HEAD", "http://192.168.6.190:13636/", "", "", 200))
	//fmt.Println(base.HttpRequest("GET", "https://192.168.6.53:8081/", "service.ws.yamutech.com", "", 200))
	fmt.Println(base.HttpRequest("GET", "https://192.168.6.53:8081/", "service.ws.yamutech.com", "", 302))
}
