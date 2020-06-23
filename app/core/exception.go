package core

import (
	"runtime"
	"strings"
	"fmt"
	"strconv"
)

type Exception struct {
	callback func(string) error
}

func (e *Exception) Init(stackString string) error {
	return e.callback(stackString)
}

// 捕获异常信息
func (e *Exception) Catch() error {
	if r := recover(); r != nil {
		var arr [4096]byte
		var stackString string
		buf := arr[:]
		n := runtime.Stack(buf, false)

		if n > 0 {
			stackString = e.Stack(r, buf)
		}
		_ = stackString
	}
	return nil
}

// Handle Stack message.
func (e *Exception) Stack(err interface{}, buf []byte) string {
	var stackSlice []string
	var msg string

	// Filename、line
	if _, srcName, line, ok := runtime.Caller(3); ok {
		msg = fmt.Sprintf("[Error]File:%s Line:%s Msg:%s \nMethod Stack meassage:", srcName, strconv.Itoa(line), err)
		stackSlice = append(stackSlice, msg)
	} else {
		msg = fmt.Sprintf("[Error]Msg:%s \nMethod Stack meassage:", err)
		stackSlice = append(stackSlice, msg)
	}

	// Handle data.
	stringStack := string(buf)                   // Convert to string
	tmpStack := strings.Split(stringStack, "\n") // Split to []string by \n
	var receiveStack []string

	for _, v := range tmpStack {
		receiveStack = append(receiveStack, strings.TrimSpace(v))
	}
	for i, j, k := 0, 0, len(receiveStack)-1; i < k; i += 2 {
		stackSlice = append(stackSlice, "["+strconv.Itoa(j)+"]"+receiveStack[i]+" "+receiveStack[i+1])

		if j == 10 {
			stackSlice = append(stackSlice, "...")
			break
		}
		j ++
	}
	return strings.Join(stackSlice, "\n")
}