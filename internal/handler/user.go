package handler

import (
	"fmt"
	"net/http"
	"time"

	userDTO "github.com/fnxr21/voucher-manager/internal/dto/user"
	"github.com/fnxr21/voucher-manager/internal/model"
	repositories "github.com/fnxr21/voucher-manager/internal/repository"
	"github.com/fnxr21/voucher-manager/pkg/bcrypt"
	errorhandler "github.com/fnxr21/voucher-manager/pkg/error"
	jwtToken "github.com/fnxr21/voucher-manager/pkg/jwt"
	customLog "github.com/fnxr21/voucher-manager/pkg/log"
	typeResult "github.com/fnxr21/voucher-manager/pkg/type"
	validator "github.com/go-playground/validator/v10"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerUser struct {
	UserRepository repositories.User
}

func HandlerUser(UserRepository repositories.User) *handlerUser {
	return &handlerUser{UserRepository}
}

type Token struct {
	Token string
}

func (h *handlerUser) Login(c echo.Context) error {
	request := new(userDTO.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, typeResult.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})

	}

	err := c.Validate(request)
	if err != nil {
		customLog.Error(c.Request(), c, err.Error(), "200")
		for _, err := range err.(validator.ValidationErrors) {

			switch err.Tag() {
			case "required":
				return errorhandler.HandlerValidationError(c, http.StatusBadRequest, fmt.Sprintf("%s is required", err.Field()))
			case err.Tag():
				return errorhandler.HandlerValidationError(c, http.StatusBadRequest, fmt.Sprintf("%s is not valid "+err.Tag(), err.Field()))
			case "gte":
				return errorhandler.HandlerValidationError(c, http.StatusBadRequest, fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param()))
			case "lte":
				return errorhandler.HandlerValidationError(c, http.StatusBadRequest, fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param()))
			}
		}
	}

	user, err := h.UserRepository.Login(request.Username)
	if err != nil {
		return c.JSON(http.StatusNotFound, typeResult.ErrorResult{Status: http.StatusNotFound, Message: "User Not found"})
	}

	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, typeResult.ErrorResult{Status: http.StatusBadRequest, Message: "Password Incorect"})
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["name"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		return c.JSON(http.StatusBadRequest, typeResult.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})

	}
	Tokens := Token{Token: token}
	return c.JSON(http.StatusOK, typeResult.SuccessResult{Status: http.StatusOK, Data: Tokens})
}

func (h *handlerUser) Register(c echo.Context) error {
	request := new(userDTO.RequestRegister)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, typeResult.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})

	}
	err := c.Validate(request)
	if err != nil {
		customLog.Error(c.Request(), c, err.Error(), "200")

		for _, err := range err.(validator.ValidationErrors) {

			switch err.Tag() {
			case "required":
				return errorhandler.HandlerValidationError(c, http.StatusBadRequest, fmt.Sprintf("%s is required", err.Field()))
			case err.Tag():
				return errorhandler.HandlerValidationError(c, http.StatusBadRequest, fmt.Sprintf("%s is not valid "+err.Tag(), err.Field()))
			case "gte":
				return errorhandler.HandlerValidationError(c, http.StatusBadRequest, fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param()))
			case "lte":
				return errorhandler.HandlerValidationError(c, http.StatusBadRequest, fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param()))
			}
		}
	}

	pass, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, typeResult.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}
	modelUser := model.User{
		Username: request.Username,
		Password: pass,
		Email:    request.Email,
	}
	user, err := h.UserRepository.Register(modelUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, typeResult.ErrorResult{Status: http.StatusBadRequest, Message: "Register Failed"})
	}

	return c.JSON(http.StatusOK, typeResult.SuccessResult{Status: http.StatusOK, Data: user})
}

func (h *handlerUser) Reauth(c echo.Context) error {
	adminLogin := c.Get("userLogin")

	adminID := adminLogin.(jwt.MapClaims)["id"].(float64)

	_, err := h.UserRepository.Reauth(uint(adminID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, typeResult.ErrorResult{Status: http.StatusBadRequest, Message: "User Not Found"})
	}

	return c.JSON(http.StatusOK, typeResult.SuccessReauth{Status: http.StatusOK, Data: "Reauth Success"})

}
