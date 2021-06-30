package Dir

import "os"

//获取当前文件路径
func GetAbsolutePath(path string) string {
	tmpDir, err := os.Getwd()
	if err != nil {
		panic("path GetAbsolutePath error")
	}
	return tmpDir + path
}

//初始化创建配置文件
func Mkdir(dir string) string {
	dirPath, _ := os.Getwd()
	savePath := dirPath + "/" + dir
	_, err := os.Stat(savePath)
	if err != nil { //文件夹不存在，创建
		os.Mkdir(dir, 777)
		return savePath
	}
	return ""
}
