package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
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

// ShowAllAccount godoc
//
//	@Summary		Show all users
//	@Description	get all users
//	@Tags			users
//	@Security		BearerAuth
//	@Param 		page query string false "page"
//	@Param 		limit query string false "limit"
//	@Param 		search[name] query string false "search users by name" search(name)
//	@Param 		search[email] query string false "search users by email" search(email)
//	@Param 		sort query string false "Sort by" Enums(id, name, email, created_at, updated_at) default(id)
//	@Param 		order query string false "Order" Enums(ASC, DESC) default(ASC)
//	@Success		200	{object}	lib.Response
//	@Failure		500	{object}	lib.Response
//	@Router			/user/all [get]
func (h *UserHandler) GetAll(ctx *gin.Context) {
	par := model.UserParams{
		PAGE:     "1",
		LIMIT:    "5",
		ORDER:    "ASC",
		NAME:     "",
		EMAIL:    "",
		ORDER_BY: "id",
	}
	search := ctx.QueryMap("search")
	name := search["name"]
	email := search["email"]

	if name != "" {
		par.NAME = name
	}
	if email != "" {
		par.EMAIL = email
	}
	if ctx.Query("page") != "" {
		par.PAGE = ctx.Query("page")
	}
	if ctx.Query("limit") != "" {
		par.LIMIT = ctx.Query("limit")
	}
	if ctx.Query("sort") != "" {
		par.ORDER_BY = ctx.Query("sort")
	}
	if ctx.Query("order") != "" {
		par.ORDER = ctx.Query("order")
	}

	res, err := h.svc.GetAll(&par)
	fmt.Println(res)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &lib.Response{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &lib.ResponseUsers{
		Success:     true,
		Status:      200,
		Message:     "Success Get All Users",
		Page:        par.PAGE,
		Data_length: par.LIMIT,
		Results:     res,
	})

}

// ShowUserById godoc
//
//	@Summary		Show an user by id
//	@Description	get user by id
//	@Tags			users
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Security		BearerAuth
//	@Param 		id path int true "user's Id"
//	@Success		200	{object}	lib.Response
//	@Failure		500	{object}	lib.Response
//	@Router			/user/detail/{id} [get]
func (h *UserHandler) GetById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	res, err := h.svc.GetById(&id)
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
		Message: "Success get data",
		Results: &model.Users{
			Id:        res.Id,
			Name:      res.Name,
			Email:     res.Email,
			CreatedAt: res.CreatedAt,
			Picture:   res.Picture,
		},
	})
}

// DeleteUserById godoc
//
//	@Summary		Delete user by id
//	@Description	Delete user by id
//	@Tags			users
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Security		BearerAuth
//	@Param 		id path int true "user's Id"
//	@Success		200	{object}	lib.Response
//	@Failure		500	{object}	lib.Response
//	@Router			/user/delete/{id} [delete]
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

// UploadProfilePicture godoc
//
//	@Summary		Upload profile picture
//	@Description	Upload profile picture by id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param 		id path int true "user's Id"
//	@Param 		picture formData file true "Profile Picture"
//	@Success		200	{object}	lib.Response
//	@Failure		500	{object}	lib.Response
//	@Router			/user/upload-pic/{id} [patch]
func (h *UserHandler) UploadPicture(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	file, err := ctx.FormFile("picture")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &lib.Response{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ext := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("user-picture-%s%s", idStr, ext)
	dst := filepath.Join("uploads", filepath.Base(fileName))
	ctx.SaveUploadedFile(file, dst)
	h.svc.UploadPicture(&id, &model.UserPicture{
		Picture: dst,
	})

	ctx.JSON(http.StatusOK, &lib.Response{
		Success: true,
		Status:  http.StatusOK,
		Message: "Success Update Profile",
	})
}

// EditUserByiD godoc
//
//	@Summary		Edit User Email
//	@Description	Edit User Email by id
//	@Tags			users
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Security		BearerAuth
//	@Param 		id path int true "user's Id"
//	@Success		200	{object}	lib.Response
//	@Failure		500	{object}	lib.Response
//	@Router			/user/edit/{id} [patch]
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
			Name:      res.Name,
			Email:     res.Email,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
		},
	})
}

// RegisterUser godoc
//
//	@Summary		Register User
//	@Description	Register User
//	@Tags			auth
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param 		name formData string true "Name"
//	@Param 		email formData string true "Email"
//	@Param 		password formData string true "Password"
//	@Success		200	{object}	lib.Response
//	@Failure		500	{object}	lib.Response
//	@Router			/auth/register [post]
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
		Results: res,
	})
}

// LoginUser godoc
//
//	@Summary		Login User
//	@Description	Login User
//	@Tags			auth
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param 		email formData string true "Email"
//	@Param 		password formData string true "Password"
//	@Success		200	{object}	lib.Response
//	@Failure		500	{object}	lib.Response
//	@Router			/auth/login [post]
func (h *UserHandler) Login(ctx *gin.Context) {
	var form model.LoginForm

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

	token := lib.GenerateToken(res.Id)

	ctx.JSON(http.StatusOK, &lib.Response{
		Success: true,
		Status:  200,
		Message: "Login Success",
		Results: &lib.LoginResponse{
			Id:    res.Id,
			Token: token,
		},
	})
}
