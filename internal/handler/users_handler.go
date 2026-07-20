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
			Message: errForm.Error(),
		})
		return
	}

	res, error := h.svc.Create(&form)

	if error != nil {
		ctx.JSON(http.StatusBadRequest, &lib.Response{
			Success: false,
			Message: error.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &lib.Response{
			Success: true,
			Message: "Success Create Account",
			Results: &model.Users{
				Id:    res.Id,
				Email: res.Email,
			},
		})
	}
}
