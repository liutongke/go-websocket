package utils

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

/*
*
@保存图片到指定文件夹
@file 要保存的文件字节流
@dir 文件将要保存的文件夹
return 文件名称
return 文件的sid，唯一值
*/
func SaveImg(file []byte, dir string) (string, string) {
	dirPath, _ := os.Getwd()
	savePath := dirPath + "/" + dir
	_, err := os.Stat(savePath)
	if err != nil { //文件夹不存在，创建
		os.Mkdir(dir, 777)
	}
	var newFileName string
	suffixName := GetImgExt(file) //获取图片后缀
	uuidName := uuid.NewString()
	newFileName = savePath + "/" + uuidName + suffixName
	os.WriteFile(newFileName, file, 0644)
	return uuidName + suffixName, uuid.NewString()
}

type FileResult struct {
	FileName string
	UUID     string
}

func saveFile(file []byte, dir string) (*FileResult, error) {
	dirPath, err := getAbsPath(dir)
	if err != nil {
		return nil, err
	}

	// 确保目录存在，如果不存在则创建
	if err := ensureDirExists(dirPath); err != nil {
		return nil, err
	}

	// 生成新的文件名
	newFileName := generateFileName(dirPath)

	// 写入文件
	err = os.WriteFile(newFileName, file, 0644)
	if err != nil {
		return nil, err
	}

	// 返回自定义结构体
	return &FileResult{
		FileName: filepath.Base(newFileName),
		UUID:     uuid.NewString(),
	}, nil
}

func getAbsPath(dir string) (string, error) {
	if filepath.IsAbs(dir) {
		return dir, nil
	}

	dirPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(dirPath, dir), nil
}

func ensureDirExists(dirPath string) error {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return os.Mkdir(dirPath, 0755)
	}
	return err
}

func generateFileName(dirPath string) string {
	suffixName := ".jpg" // 你的文件后缀名，可以根据需要修改
	uuidName := uuid.NewString()
	return filepath.Join(dirPath, fmt.Sprintf("%s%s", uuidName, suffixName))
}
