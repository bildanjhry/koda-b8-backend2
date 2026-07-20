package handler

import (
	"net/http"

	"github.com/bildanjhry/auth/internal/lib"
	"github.com/bildanjhry/auth/internal/model"
	"github.com/bildanjhry/auth/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) Create(ctx *gin.Context) {
	var form model.UserForm

	errForm := ctx.ShouldBind(&form)
	if errForm != nil {
		ctx.JSON(http.StatusBadRequest, &lib.Response{
			Success: false,
			Status:  400,
			Message: errForm.Error(),
		})
		return
	}

	res, errorCreate := h.svc.Create(&form)
	if errorCreate != nil {
		ctx.JSON(http.StatusBadRequest, &lib.Response{
			Success: false,
			Status:  400,
			Message: errorCreate.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &lib.Response{
		Success: true,
		Status:  200,
		Message: "Success Create Account",
		Results: &model.Users{
			Id:        res.Id,
			Email:     res.Email,
			CreatedAt: res.CreatedAt,
		},
	})
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var form model.UserForm

	errForm := ctx.ShouldBind(&form)
	if errForm != nil {
		ctx.JSON(http.StatusBadRequest, &lib.Response{
			Success: false,
			Status:  400,
			Message: errForm.Error(),
		})
		return
	}

	res, errorCreate := h.svc.Login(&form)
	if errorCreate != nil {
		ctx.JSON(http.StatusBadRequest, &lib.Response{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: errorCreate.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &lib.Response{
		Success: true,
		Status:  200,
		Message: "Login Success",
		Results: &model.Users{
			Id:        res.Id,
			Email:     res.Email,
			CreatedAt: res.CreatedAt,
		},
	})
}
