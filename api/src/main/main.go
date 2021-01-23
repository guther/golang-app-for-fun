package main

import (
	"config"
	"controller"
	"fmt"
	"middleware"
	"net/http"
	"service"

	"github.com/gin-gonic/gin"
)

func main() {

	// Force log's color
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	config.Load()

	authorized := router.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(middleware.AuthorizeJWT())
	{
		authorized.GET("/breeds", getting)
	}

	router.POST("/login", login)

	// By default it serves on :6060 unless a
	// PORT environment variable was defined.
	router.Run(fmt.Sprintf(":%d", config.PORT))
}

// I did not put the binding:"required" because I wanna get the same behaviour of TheCatAPI
type Breed struct {
	Name string `form:"name"`
}

func getting(c *gin.Context) {
	// var loginService service.LoginService = service.StaticLoginService()
	// var jwtService service.JWTService = service.JWTAuthService()
	// var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	// Query parameter validation
	var breed Breed
	if err := c.ShouldBindQuery(&breed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get the content in the API
	response, err := http.Get("https://api.thecatapi.com/v1/breeds/search?q=" + breed.Name)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	defer reader.Close()
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, nil)
}

func login(c *gin.Context) {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	// Login data validation
	token := loginController.Login(c)
	if token != "" {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	}
}
