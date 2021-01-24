package main

import (
	"config"
	"controller"
	"db"
	"fmt"
	"log"
	"middleware"
	"net/http"
	"service"
	"strings"
	"utils"
	. "utils"

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

	// Performs migration in database
	db.Migrate()

	authorized := router.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthorizeJWT() middleware just in the "authorized" group.
	authorized.Use(middleware.AuthorizeJWT())
	{
		authorized.GET("/breeds", breeds)
	}

	router.POST("/login", login)

	// By default it serves on :6060 unless a
	// PORT environment variable was defined.
	router.Run(fmt.Sprintf(":%d", config.PORT))
}

// breeds function performs performs the breeds endpoint
func breeds(c *gin.Context) {
	// Query parameter validation
	var breed utils.Breed
	if err := c.ShouldBindQuery(&breed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check cache
	query := "?name=" + c.Query("name")

	var cached = false
	var dataCache = []db.BreedCache{}

	defer func() {
		if r := recover(); r != nil {
			log.Println("Database error!", fmt.Sprint(r), strings.Contains(fmt.Sprint(r), "Error 1146"))
			if strings.Contains(fmt.Sprint(r), "Error 1146") {
				db.Migrate()
			}
			DataToOutput(c, GetFromWeb(c, breed))
		}
	}()

	cached, dataCache = db.CheckCacheResult(query)

	var dataJSON string

	if !cached {
		dataJSON = GetFromWeb(c, breed)
		db.Insert(query, dataJSON)
	} else {
		var dataStored db.BreedCache = dataCache[0]
		dataJSON = dataStored.Data
	}
	DataToOutput(c, dataJSON)
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
