package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"NurtureNest/pkg/api/sdk/crypto"
	sdk_jwt "NurtureNest/pkg/api/sdk/jwt"
	"NurtureNest/pkg/api/sdk/response"
	domain "NurtureNest/pkg/domain"
	services "NurtureNest/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

type Response struct {
	ID      uint   `copier:"must"`
	Name    string `copier:"must"`
	Surname string `copier:"must"`
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (cr *UserHandler) UserRegister(c *gin.Context) {
	var user domain.UserRegister
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "BAD REQUEST", err)
		return
	}

	//validasi email
	if !strings.Contains(user.Email, "@gmail.com") {
		msg := "EMAIL NOT VALID"
		response.FailOrError(c, http.StatusBadRequest, msg, errors.New(msg))
		return
	}

	if len(user.Password) < 8 {
		msg := "PASSWORD IS TOO SHORT"
		response.FailOrError(c, http.StatusBadRequest, msg, errors.New(msg))
		return
	}

	_, err = cr.userUseCase.UserFindByUsernameOrEmail(c.Request.Context(), user.Username, user.Email)
	if err == nil {
		msg := "USERNAME OR EMAIL IS ALREADY IN USE"
		response.FailOrError(c, http.StatusBadRequest, msg, errors.New(msg))
		return
	}

	result, err := cr.userUseCase.UserRegister(c.Request.Context(), user)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "USERNAME OR EMAIL IS ALREADY IN USE", err)
		return
	}

	response.Success(c, http.StatusCreated, "Success create user", result)
}

func (cr *UserHandler) UserFindAll(c *gin.Context) {
	users, err := cr.userUseCase.FindAll(c.Request.Context())

	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "failed get all user", err)
	}

	response.Success(c, http.StatusCreated, "Success get all user", users)
}

func (cr *UserHandler) UserLogin(c *gin.Context) {
	var request domain.UserLogin
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}

	user, err := cr.userUseCase.UserFindByUsernameOrEmail(c.Request.Context(), request.UsernameOrEmail, request.UsernameOrEmail)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Username atau email tidak ditemukan", err)
		return
	}

	err = crypto.ValidateHash(request.Password, user.Password)
	if err != nil {
		msg := "PASSWORD INVALID"
		response.FailOrError(c, http.StatusBadRequest, msg, errors.New(msg))
		return
	}

	tokenJwt, err := sdk_jwt.GenerateToken(user)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "create token failed", err)
		return
	}

	response.Success(c, http.StatusOK, "Berhasil masuk", gin.H{
		"token": tokenJwt,
	})
}

func (cr *UserHandler) FindByID(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	user, err := cr.userUseCase.FindByID(c.Request.Context(), uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &user)

		c.JSON(http.StatusOK, response)
	}
}

func (cr *UserHandler) Save(c *gin.Context) {
	var user domain.Users

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := cr.userUseCase.Save(c.Request.Context(), user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &user)

		c.JSON(http.StatusOK, response)
	}
}

func (cr *UserHandler) Delete(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot parse id",
		})
		return
	}

	ctx := c.Request.Context()
	user, err := cr.userUseCase.FindByID(ctx, uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	cr.userUseCase.Delete(ctx, user)

	c.JSON(http.StatusOK, gin.H{"message": "User is deleted successfully"})
}
