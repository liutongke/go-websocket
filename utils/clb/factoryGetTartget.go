package clb

type GetTargetHost interface {
	GetTargetIp(serverTarget string, ip string) string
}

//工厂统一入口 https://juejin.cn/post/6871169933150486542#heading-3
func ProxyInfo(ObtainMode int) GetTargetHost {
	return &RandTarget{}
}

type RandTarget struct {
}

//随机轮询
func (rt RandTarget) GetTargetIp(serverTarget string, ip string) string {
	return "ip"
}
