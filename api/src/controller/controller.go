package controller

import (
	"dto"
	"net/http"
	"service"

	"github.com/gin-gonic/gin"
)

// login controller interface
type LoginController interface {
	Login(c *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func LoginHandler(loginService service.LoginService, jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(c *gin.Context) string {
	var credential dto.LoginCredentials

	// Query parameter validation
	// Example for binding JSON ({"username": "alexandre", "password": "damasceno"})
	err := c.ShouldBindJSON(&credential)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return ""
	}

	// Login data validation
	isUserAuthenticated := controller.loginService.LoginUser(credential.User, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.User, true)

	}
	return ""
}
