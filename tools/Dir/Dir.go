package Dir

import (
	"os"
	"path/filepath"
)

// GetAbsDirPath 根据当前的工作目录来解析相对路径，可以使用../test
func GetAbsDirPath(filePath string) string {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return ""
	}
	return absPath
}

// DelDir 删除目录
func DelDir(dirPath string) bool {
	err := os.RemoveAll(dirPath)
	return err == nil
}

// MkDirAll 父目录不存在一同创建
func MkDirAll(dirPath string) bool {
	err := os.MkdirAll(dirPath, os.ModePerm)
	return err == nil
}

// IsExist 判断目录是否存在
func IsExist(dirPath string) bool {
	// 检查目录是否存在
	_, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			return false
		}
	} else {
		return true
	}
}
