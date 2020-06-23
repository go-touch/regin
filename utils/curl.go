package utils

import (
	netUrl "net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"bytes"
)

type CurlHandler struct {
}

var Curl *CurlHandler

func init() {
	Curl = &CurlHandler{}
}

// GET请求
func (ch *CurlHandler) Get(url string, args ...map[string]string) (result map[string]interface{}, err error) {
	// 参数处理
	param := netUrl.Values{}
	if args != nil {
		for key, value := range args[0] {
			param.Set(key, value)
		}
	}

	// 构造url
	reqUrl, err := netUrl.ParseRequestURI(url)
	if err != nil {
		return result, err
	}

	// 发送GET请求
	reqUrl.RawQuery = param.Encode()
	resp, err := http.Get(reqUrl.String())
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// 读取响应数据
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	return ch.ByteSliceToMap(respData)
}

// POST请求
func (ch *CurlHandler) Post(url string, args ...map[string]string) (result map[string]interface{}, err error) {
	// 参数处理
	var param []string
	if args != nil {
		for k, v := range args[0] {
			param = append(param, k+"="+v)
		}
	}

	// 发送请求
	paramStr := strings.Join(param, "&")
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(paramStr))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// 读取响应数据
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	// 发送带有字节顺序标记（BOM）的UTF-8文本字符串。 BOM标识文本是UTF-8编码的，但应在解码之前将其删除
	respData = bytes.TrimPrefix(respData, []byte("\xef\xbb\xbf"))
	return ch.ByteSliceToMap(respData)
}

// POST请求
func (ch *CurlHandler) JsonPost(url string, args ...map[string]interface{}) (result map[string]interface{}, err error) {
	// 参数处理
	var param string
	if args != nil {
		if jsonParam, err := ch.MapToJson(args[0]); err != nil {
			return result, err
		} else {
			param = jsonParam
		}
	}

	// 发送请求
	resp, err := http.Post(url, "application/json", strings.NewReader(param))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// 读取响应数据
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	// 发送带有字节顺序标记（BOM）的UTF-8文本字符串。 BOM标识文本是UTF-8编码的，但应在解码之前将其删除
	respData = bytes.TrimPrefix(respData, []byte("\xef\xbb\xbf"))
	return ch.ByteSliceToMap(respData)
}

// 转换请求结果 []byte 到 map[string]interface{}
func (ch *CurlHandler) ByteSliceToMap(respData []byte) (result map[string]interface{}, err error) {
	err = json.Unmarshal(respData, &result)
	return result, err
}

// 转换请求结果 []byte 到 map[string]interface{}
func (ch *CurlHandler) MapToJson(mapData map[string]interface{}) (jsonData string, err error) {
	result, err := json.Marshal(mapData)
	if err == nil {
		jsonData = string(result)
	}
	return jsonData, err
}
