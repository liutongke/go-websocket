package tools

import "errors"

const (
	ErrProxyUrlAppListIsEmpty string = "应用服务器地址为空"
)

func MyErr(errMsg string) error {
	return errors.New(errMsg)
}
