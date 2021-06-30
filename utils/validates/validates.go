package validates

import "github.com/go-playground/validator"

func GetClient() *validator.Validate {
	return validate
}

var validate *validator.Validate //定义

func InitValidate() {
	validate = validator.New() //初始化（赋值）
}
