package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"serenity-book-api/internal/security"
	"serenity-book-api/internal/service"
	"serenity-book-api/pkg/helper/resp"
)

type UserHandler interface {
	GetUserById(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type userHandler struct {
	*Handler
	securityUtils security.SecurityUtils
	userService   service.UserService
}

func NewUserHandler(handler *Handler, securityUtils security.SecurityUtils, userService service.UserService) UserHandler {
	return &userHandler{
		Handler:       handler,
		securityUtils: securityUtils,
		userService:   userService,
	}
}

func (h *userHandler) Login(ctx *gin.Context) {
	h.securityUtils.Login(1, security.LoginModel{}, ctx)
	resp.HandleSuccess(ctx, nil)
	return
}

func (h *userHandler) GetUserById(ctx *gin.Context) {
	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	user, err := h.userService.GetUserById(params.Id)
	h.logger.Info("GetUserByID", zap.Any("user", user))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, user)
}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}
