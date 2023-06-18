package Mgck

import (
	"github.com/bean-du/dfa"
)

var (
	fda       *dfa.DFA
	sensitive []string
)

type Msgck struct {
	Words string
}

// 初始化一下敏感词
func init() {
	//var tbWords []Msgck
	//
	//filePath := dir.GetAbsolutePath(config.GetConfClient().CommonConf.MgCk) // 打开json文件
	//
	//jsonFile, err := os.Open(filePath)
	//if err != nil {
	//	return
	//}
	//defer jsonFile.Close() // 要记得关闭
	//byteValue, err := ioutil.ReadAll(jsonFile)
	//if err != nil {
	//	return
	//}
	//json.Unmarshal(byteValue, &tbWords)
	//for _, v := range tbWords {
	//	sensitive = append(sensitive, v.Words)
	//}
	//fda = dfa.New()
	//fda.AddBadWords(sensitive)

}

func NewWords() *dfa.DFA {
	return fda
}

// 检查敏感词 true敏感词 false非敏感词
func CheckWord(words string) bool {
	_, _, res := fda.Check(words)
	return res
}
