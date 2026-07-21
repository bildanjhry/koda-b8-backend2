package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

func (h *UserHandler) GetAll(ctx *gin.Context) {
	res, err := h.svc.GetAll()
	fmt.Println(res)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &lib.Response{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &lib.Response{
		Success: true,
		Status:  200,
		Message: "Success Create Account",
		Results: res,
	})

}

func (h *UserHandler) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	res := h.svc.Delete(&id)
	if res != nil {
		ctx.JSON(http.StatusInternalServerError, &lib.Response{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: res.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &lib.Response{
		Success: true,
		Status:  http.StatusOK,
		Message: "Success Delete Data",
	})
}

func (h *UserHandler) Edit(ctx *gin.Context) {
	var form model.UserEmail

	errForm := ctx.ShouldBind(&form)
	if errForm != nil {
		ctx.JSON(http.StatusBadRequest, &lib.Response{
			Success: false,
			Status:  400,
			Message: errForm.Error(),
		})
		return
	}

	fmt.Println(&form)

	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	res, editErr := h.svc.Edit(&id, &form)

	if editErr != nil {
		ctx.JSON(http.StatusInternalServerError, &lib.Response{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: editErr.Error(),
		})
		return
	}

	fmt.Println(res)

	ctx.JSON(http.StatusOK, &lib.Response{
		Success: true,
		Status:  http.StatusOK,
		Message: "Success Update Data",
		Results: &model.Users{
			Id:        res.Id,
			Email:     res.Email,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
		},
	})
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
