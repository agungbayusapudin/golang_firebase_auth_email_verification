package controller

import (
	"crud_fire/model"
	"crud_fire/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authController struct {
	authService service.ServiceAuth
}

func NewAuthController(authService service.ServiceAuth) authController {
	return authController{authService: authService}
}

func (c *authController) Login(ctx echo.Context) error {
	var form model.ModelAuth

	// binding data
	if err := ctx.Bind(&form); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// login
	token, err := c.authService.Login(ctx.Request().Context(), form)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, token)
}

func (c *authController) Register(ctx echo.Context) error {
	var form model.ModelAuth

	// binding data
	if err := ctx.Bind(&form); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// register
	user, err := c.authService.Register(ctx.Request().Context(), form)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, user)
}

func (c *authController) VerifyEmail(ctx echo.Context) error {
	token := ctx.QueryParam("token")
	if token == "" {
		return ctx.JSON(http.StatusBadRequest, "token is required")
	}

	// verify email
	err := c.authService.VerifyEmail(ctx.Request().Context(), token)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "success verify email")
}

func (c *authController) LoginUsingGoogle(ctx echo.Context) error {
	idToken := ctx.QueryParam("idToken")
	if idToken == "" {
		return ctx.JSON(http.StatusBadRequest, "idToken is required")
	}

	// login using google
	token, err := c.authService.LoginUsingGoogle(ctx.Request().Context(), idToken)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, token)

}
