package fileutil

import (
	"os"
	"path/filepath"
)

// GetAbsDirPath 根据当前的工作目录来解析相对路径，可以使用../test
// Absolute Path: /var/www/html/log
func GetAbsDirPath(filePath string) string {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return ""
	}
	return absPath
}

// DelDir 删除目录
func DelDir(dirPath string) error {
	err := os.RemoveAll(dirPath)
	return err
}

// MkDirAll 父目录不存在一同创建
func MkDirAll(dirPath string) error {
	err := os.MkdirAll(dirPath, os.ModePerm)
	return err
}

// IsExist 判断目录是否存在
func IsExist(dirPath string) bool {
	// 检查目录是否存在
	_, err := os.Stat(dirPath)
	return err == nil
}
