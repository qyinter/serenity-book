package security

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"serenity-book-api/internal/repository"
	"serenity-book-api/internal/service"
	"serenity-book-api/pkg/contents"
	"serenity-book-api/pkg/global"
	"serenity-book-api/pkg/helper/resp"
	"time"
)

type SecurityUtils interface {
	Login(loginId interface{}, loginModel LoginModel, ctx *gin.Context)
}

type securityUtils struct {
	*service.Service
	*repository.Repository
}

func NewSecurityUtils(service *service.Service, repository *repository.Repository) SecurityUtils {
	return &securityUtils{
		Service:    service,
		Repository: repository,
	}
}

func (su *securityUtils) Login(loginId interface{}, loginModel LoginModel, ctx *gin.Context) {
	// 1. 生成token
	token, err := GenerateToken(loginId)
	if err != nil {
		// 生成token失败
		resp.HandleError(ctx, http.StatusInternalServerError, http.StatusInternalServerError, contents.ServerError, nil)
		return
	}
	// 2. 保存token
	set := su.Repository.Rdb.HSet(ctx, "Authorization:login:token:"+token, token, token)
	if set.Err() != nil {
		// 保存token失败
		resp.HandleError(ctx, http.StatusInternalServerError, http.StatusInternalServerError, contents.ServerError, nil)
		return
	}
	expire := su.Repository.Rdb.Expire(ctx, "Authorization:login:token:"+token, time.Duration(global.GConfig.Security.Expire)*time.Hour)
	if expire.Err() != nil {
		// 设置过期时间失败
		resp.HandleError(ctx, http.StatusInternalServerError, http.StatusInternalServerError, contents.ServerError, nil)
		return
	}
	// 3. 保存session
	//TODO 获取菜单
	//TODO 获取权限字符串
	//TODO 获取自定义储存信息
	//TODO 获取token list ua 地址

	// 4. 保存cookie
	// 把cookie 到cookie里面
	return
}

// 登出方法 // 踢人下线

// 校验当前token 是否有效

// 获取session 中用户信息: 里面包含用户token ,用户菜单,用户角色字符串, token list ua 地址, 其他的用户信息

// 设置session 信息: 里面包含用户token ,用户菜单,用户角色字符串, 其他的用户信息

// 设置token

// 获取菜单

// 获取权限字符串 perm

// 获取自定义储存信息

/**
 *获取当前用户的login id
 */
// 如果获取不到token，则抛出: 无token
// 查找此token对应loginId, 如果找不到则抛出：无效token
// 如果是已经过期，则抛出已经过期
// 如果是已经被顶替下去了, 则抛出：已被顶下线
// 如果是已经被踢下线了, 则抛出：已被踢下线
