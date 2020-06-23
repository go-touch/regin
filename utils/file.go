package utils

import (
	"os"
	"path/filepath"
	"strings"
)

type FileHandler struct {
}

// 定义FileHandler结构体
var File *FileHandler

func init() {
	File = &FileHandler{}
}

// 遍历目录
func (fh *FileHandler) ScanDir(dirPath string) ([]os.FileInfo) {
	// 判断目录是否存在
	if _, err := fh.IsExist(dirPath); err != nil {
		panic(err)
	}

	// 打开目录
	dir, err := os.Open(dirPath)

	if err != nil {
		panic(err)
	}
	defer dir.Close()

	// 读取目录文件
	files, err := dir.Readdir(-1)

	if err != nil {
		panic(err)
	}
	return files
}

// 判断文件夹是否存在
func (fh *FileHandler) IsExist(path string) (bool, error) {
	// 获取文件的信息
	_, err := os.Stat(path)

	if err == nil {
		return true, err
	}
	return false, err
}

// 拼接路径
func (fh *FileHandler) JoinPath(path ...string) string {
	return strings.Join(path, string(filepath.Separator))
}
