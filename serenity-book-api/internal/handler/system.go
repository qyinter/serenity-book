package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"net/http"
	v1 "serenity-book-api/api/v1"
	"serenity-book-api/internal/service"
	"serenity-book-api/pkg/global"
)

type SystemHandler interface {
	GetCaptcha(ctx *gin.Context)
}

type systemHandler struct {
	*Handler
	systemService service.SystemService
}

func NewSystemHandler(
	handler *Handler,
	systemService service.SystemService,
) SystemHandler {
	return &systemHandler{
		Handler:       handler,
		systemService: systemService,
	}
}

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

// GetCaptcha 获取验证码
func (h *systemHandler) GetCaptcha(ctx *gin.Context) {
	// 判断验证码是否开启
	openCaptcha := global.GConfig.Captcha.OpenCaptcha // 是否开启防爆次数
	//openCaptchaTimeOut := global.GConfig.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.GConfig.Captcha.ImgHeight, global.GConfig.Captcha.ImgWidth, global.GConfig.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, answer, err := cp.Generate()
	// 生成验证码失败
	if err != nil {
		global.GLog.Error("验证码获取失败!", zap.Any("err", err))
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrCaptchaGet, nil)
		return
	}
	// 生成验证码返回给前端
	v1.HandleSuccess(ctx, map[string]interface{}{
		"captchaId":     id,
		"picPath":       b64s,
		"answer":        answer,
		"CaptchaLength": global.GConfig.Captcha.KeyLong,
		//"openCaptcha":   oc,
		"openCaptcha": openCaptcha,
	})
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}
