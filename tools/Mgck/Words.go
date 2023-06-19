package Mgck

import (
	"encoding/json"
	"github.com/bean-du/dfa"
	"go-websocket/config"
	"go-websocket/tools"
	"go-websocket/tools/Dir"
	"io/ioutil"
)

var (
	fda *dfa.DFA
)

type Data struct {
	Words interface{} `json:"words"`
}

// 初始化一下敏感词
func InitWord() {

	filePath := Dir.GetAbsDirPath(config.GetConf().CommonConf.MgCk)

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		tools.EchoError(err.Error())
	}

	var data []Data
	err = json.Unmarshal(content, &data)
	if err != nil {
		tools.EchoError(err.Error())
	}

	var sensitive []string

	for _, v := range data {
		// 将接口值断言为字符串类型
		str, ok := v.Words.(string)
		if ok {
			sensitive = append(sensitive, str)
		}

	}

	fda = dfa.New()
	fda.AddBadWords(sensitive)
}

// CheckWord https://pkg.go.dev/github.com/bean-du/dfa@v1.0.2#section-readme
// DFA 敏感词检测
// 检查敏感词 true敏感词 false非敏感词
func CheckWord(words string) ([]string, []string, bool) {
	w1, w2, res := fda.Check(words)
	return w1, w2, res
}
