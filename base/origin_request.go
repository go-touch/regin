package base

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/go-touch/mtype"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

const defaultMultipartMemory = 32 << 20 // 32 MB

type Request struct {
	*http.Request
	Storage       *mtype.AnyMap // 运行期间的存储容器
	paramMap      StringMap
	getMap        StringMap
	cookieMap     StringMap
	postMap       *mtype.AnyMap
	postFileSlice []*multipart.FileHeader
	rawSlice      []byte
	error         error
}

// Define static Request
var RequestInvoker *Request

func init() {
	RequestInvoker = &Request{}
}

// Create an new Request.
func (r *Request) Factory(c *gin.Context) *Request {
	request := &Request{Request: c.Request}
	_ = request.init()
	_ = request.initParam(c.Params)
	_ = request.initGet()
	_ = request.initPost()
	return request
}

// Init data.
func (r *Request) init() error {
	r.Storage = &mtype.AnyMap{}
	return nil
}

// Init Param data.
func (r *Request) initParam(params gin.Params) error {
	r.paramMap = StringMap{}
	for _, entry := range params {
		r.paramMap[entry.Key] = entry.Value
	}
	return nil
}

// Init GET data.
func (r *Request) initGet() error {
	r.getMap = StringMap{}
	paramGroup := r.Request.URL.Query()

	for key, value := range paramGroup {
		r.getMap[key] = value[len(value)-1]
	}
	return nil
}

// Init POST data.
func (r *Request) initPost() error {
	r.postMap = &mtype.AnyMap{}
	if r.Method != "POST" {
		return r.error
	}

	// Handle data by Content-Type.
	ct := r.Header.Get("Content-Type")
	if strings.Contains(ct, "/x-www-form-urlencoded") || strings.Contains(ct, "/form-data") {
		if r.error = r.ParseMultipartForm(defaultMultipartMemory); r.error != nil && r.error != http.ErrNotMultipart {
			return r.error
		}
		for key, value := range r.PostForm {
			if len(value) == 1 {
				r.postMap.Set(key, value[0])
			} else {
				r.postMap.Set(key, value)
			}
		}
		return nil
	} else { // Handle data by other Content-Type.
		if r.rawSlice, r.error = ioutil.ReadAll(r.Body); r.error != nil {
			return r.error
		}
		if strings.Contains(ct, "/json") { // Content-Type is Json.
			if r.error = json.Unmarshal(r.rawSlice, r.postMap); r.error != nil {
				return r.error
			}
		} else if strings.Contains(ct, "/xml") { // Content-Type is Xml.
			if r.error = xml.Unmarshal(r.rawSlice, r.postMap); r.error != nil {
				return r.error
			}
		}
	}
	return nil
}

// Get Http request method
func (r *Request) GetMethod() string {
	return r.Method
}

// Get error info
func (r *Request) GetError() error {
	return r.error
}

//  Path data.
func (r *Request) Param(key string, defaultValue ...string) string {
	val := ""
	if defaultValue != nil {
		val = defaultValue[0]
	}
	if value, ok := r.paramMap[key]; ok {
		return value
	}
	return val
}

// Get param array.
func (r *Request) ParamAll() StringMap {
	return r.paramMap
}

// Get data.
func (r *Request) Get(key string, defaultValue ...string) string {
	val := ""
	if defaultValue != nil {
		val = defaultValue[0]
	}
	if value, ok := r.getMap[key]; ok {
		return value
	}
	return val
}

// Get param array.
func (r *Request) GetAll() StringMap {
	return r.getMap
}

// POST param.
func (r *Request) Post(key string, defaultValue ...interface{}) interface{} {
	var val interface{}
	if defaultValue != nil {
		val = defaultValue[0]
	}
	if value, ok := (*r.postMap)[key]; ok {
		return value
	}
	return val
}

// POST param array.
func (r *Request) PostAll() *mtype.AnyMap {
	return r.postMap
}

// POST file.
func (r *Request) PostFile(name string) []*multipart.FileHeader {
	if r.MultipartForm != nil && r.MultipartForm.File != nil {
		if fileHeaders, ok := r.MultipartForm.File[name]; ok {
			r.postFileSlice = fileHeaders
		}
	}
	return r.postFileSlice
}

// 获取元数据
func (r *Request) Raw() []byte {
	return r.rawSlice
}
