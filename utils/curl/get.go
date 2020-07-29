package curl

import (
	"bytes"
	"github.com/go-touch/mtype"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strconv"
)

type GetCaller struct {
	method string
	header map[string]string
}

// 获取一个GetCaller
func Get() *GetCaller {
	return &GetCaller{
		method: "GET",
		header: map[string]string{},
	}
}

// 设置头信息
func (gc *GetCaller) Header(header map[string]string) {
	gc.header = header
}

// 发送Get请求
func (gc *GetCaller) Call(url string, args ...map[string]interface{}) *mtype.AnyValue {
	// 参数处理
	param := url2.Values{}
	if args != nil {
		for key, value := range args[0] {
			if v, ok := value.(string); ok {
				param.Set(key, v)
			} else if v, ok := value.(int); ok {
				param.Set(key, strconv.Itoa(v))
			}
		}
	}
	parserUrl, err := url2.ParseRequestURI(url)
	if err != nil {
		return mtype.Eval(err)
	} else {
		parserUrl.RawQuery = param.Encode()
		url = parserUrl.String()
	}
	// 创建request
	request, err2 := http.NewRequest(gc.method, url, nil)
	if err2 != nil {
		return mtype.Eval(err2)
	}
	if len(gc.header) > 0 {
		for key, value := range gc.header {
			request.Header.Add(key, value)
		}
	}
	// 发送一个GET请求
	resp, err3 := http.DefaultClient.Do(request)
	if err3 != nil {
		return mtype.Eval(err3)
	}
	defer func() { _ = resp.Body.Close() }()
	// 响应数据处理
	respData, err4 := ioutil.ReadAll(resp.Body)
	if err4 != nil {
		return mtype.Eval(err4)
	} else {
		respData = bytes.TrimPrefix(respData, []byte("\xef\xbb\xbf"))
	}
	return mtype.Eval(respData)
}