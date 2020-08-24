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
	//*http.Request
	*gin.Context
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
	request := &Request{
		Context:       c,
		//Request:       c.Request,
		Storage:       &mtype.AnyMap{},
		cookieMap:     map[string]string{},
		paramMap:      map[string]string{},
		getMap:        map[string]string{},
		postMap:       &mtype.AnyMap{},
		postFileSlice: []*multipart.FileHeader{},
		rawSlice:      make([]byte, 0),
	}
	_ = request.initParam(c.Params)
	_ = request.initGet()
	_ = request.initPost()
	return request
}

// Init Param data.
func (r *Request) initParam(params gin.Params) error {
	for _, element := range params {
		r.paramMap[element.Key] = element.Value
	}
	return nil
}

// Init GET data.
func (r *Request) initGet() error {
	params := r.Request.URL.Query()
	for key, value := range params {
		r.getMap[key] = value[len(value)-1]
	}
	return nil
}

// Init POST data.
func (r *Request) initPost() error {
	// When Request method is POST then parser data.
	if r.Request.Method == "POST" {
		// Handle data by Content-Type.
		ct := r.Request.Header.Get("Content-Type")
		if strings.Contains(ct, "/x-www-form-urlencoded") || strings.Contains(ct, "/form-data") {
			if r.error = r.Request.ParseMultipartForm(defaultMultipartMemory); r.error != nil && r.error != http.ErrNotMultipart {
				return r.error
			}
			for key, value := range r.Request.PostForm {
				if len(value) == 1 {
					(*r.postMap)[key] = value[0]
				} else {
					(*r.postMap)[key] = value
				}
			}
			// Parser postMap data when Contains [].
			*r.postMap = r.parserPostMap(*r.postMap)
			return nil
		} else { // Handle data by other Content-Type.
			if r.rawSlice, r.error = ioutil.ReadAll(r.Request.Body); r.error != nil {
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
	}
	return nil
}

// Get Http request method
func (r *Request) GetMethod() string {
	return r.Request.Method
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
func (r *Request) ParamAll() map[string]string {
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
func (r *Request) GetAll() map[string]string {
	return r.getMap
}

// POST param.
func (r *Request) Post() *mtype.AnyMap {
	return r.postMap
}

// POST file.
func (r *Request) PostFile(name string) []*multipart.FileHeader {
	if r.Request.MultipartForm != nil && r.Request.MultipartForm.File != nil {
		if fileHeaders, ok := r.Request.MultipartForm.File[name]; ok {
			r.postFileSlice = fileHeaders
		}
	}
	return r.postFileSlice
}

// 获取元数据
func (r *Request) Raw() []byte {
	return r.rawSlice
}

// Convert POST data to struct.
func (r *Request) ToStruct(object interface{}, method ...string) error {
	byteSlice := make([]byte, 0)

	// Handle data by method.
	m := "POST"
	if method != nil {
		m = strings.ToUpper(method[0])
	}
	if m == "GET" {
		if data, err := json.Marshal(r.getMap); err != nil {
			return err
		} else {
			byteSlice = data
		}
	} else if m == "POST" {
		if len(r.rawSlice) > 0 {
			byteSlice = r.rawSlice
		} else if data, err := json.Marshal(r.postMap); err != nil {
			return err
		} else {
			byteSlice = data
		}
	}

	// Parse []byte to struct.
	err := json.Unmarshal(byteSlice, object)
	if err != nil {
		return err
	}
	return nil
}

// Parser postMap data.
func (r *Request) parserPostMap(anyMap map[string]interface{}) map[string]interface{} {
	tmpMap := map[string]interface{}{}
	for key, value := range anyMap {
		i := strings.IndexByte(key, '[')
		j := strings.IndexByte(key, ']')

		if i >= 1 && j >= 1 {
			trimKey := strings.Trim(key, "]")
			k := trimKey[0:i]
			subK := trimKey[i+1:]
			if _, ok := tmpMap[k]; !ok {
				tmpMap[k] = map[string]interface{}{
					subK: value,
				}
			} else {
				tmpMap[k].(map[string]interface{})[subK] = value
			}
			delete(anyMap, key)
		}
	}
	for key, value := range tmpMap {
		anyMap[key] = value
	}
	return anyMap
}
