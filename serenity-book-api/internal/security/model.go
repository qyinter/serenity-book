package security

type LoginModel struct {
	device          string // 设备
	isLastingCookie bool   // 是否持久化cookie // 默认为true
}
