package curl

import (
	"bytes"
	"github.com/go-touch/mtype"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type PostCaller struct {
	method string
	header map[string]string
}

// 获取一个PostCaller
func Post() *PostCaller {
	return &PostCaller{
		method: "POST",
		header: map[string]string{},
	}
}

// Set header.
func (pc *PostCaller) Header(header map[string]string) {
	pc.header = header
}

// Send a post request.
func (pc *PostCaller) Call(url string, args ...interface{}) *mtype.AnyValue {
	// 创建request
	request, err := http.NewRequest(pc.method, url, pc.IoReader(args...))
	if err != nil {
		return mtype.Eval(err)
	}
	if len(pc.header) > 0 {
		for key, value := range pc.header {
			request.Header.Add(key, value)
		}
	} else {
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	// 发送一个POST请求
	resp, err2 := http.DefaultClient.Do(request)
	if err2 != nil {
		return mtype.Eval(err2)
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

// 获取 io.Reader.
func (pc *PostCaller) IoReader(args ...interface{}) io.Reader {
	var ioReader io.Reader
	if args == nil {
		return nil
	}
	if mtype.GetType(args[0]) == mtype.TString {
		param := url.QueryEscape(args[0].(string))
		ioReader = strings.NewReader(param)
	} else {
		stringMap := mtype.Eval(args[0]).ToStringMap()
		var param []string
		for k, v := range stringMap {
			param = append(param, k+"="+url.QueryEscape(v))
		}
		ioReader = strings.NewReader(strings.Join(param, "&"))
	}
	return ioReader
}
