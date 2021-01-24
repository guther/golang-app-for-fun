package main

import (
	"bytes"
	"config"
	"controller"
	"db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"middleware"
	"net/http"
	"service"
	"strconv"
	"strings"

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

// I did not put the binding:"required" because I wanna get the same behaviour of TheCatAPI
type Breed struct {
	Name string `form:"name"`
}

// breeds function performs performs the breeds endpoint
func breeds(c *gin.Context) {
	// Query parameter validation
	var breed Breed
	if err := c.ShouldBindQuery(&breed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check cache
	query := "?name=" + c.Query("name")

	var cached = false
	var data_cache = []db.BreedCache{}

	defer func() {
		if r := recover(); r != nil {
			log.Println("Database error!", fmt.Sprint(r), strings.Contains(fmt.Sprint(r), "Error 1146"))
			if strings.Contains(fmt.Sprint(r), "Error 1146") {
				db.Migrate()
			}
			DataToOutput(c, getFromWeb(c, breed))
		}
	}()

	cached, data_cache = db.CheckCacheResult(query)

	var data_json string

	if !cached {
		data_json = getFromWeb(c, breed)
		db.Insert(query, data_json)
	} else {
		var data_stored db.BreedCache = data_cache[0]
		data_json = data_stored.Data
	}
	DataToOutput(c, data_json)
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

func getFromWeb(c *gin.Context, breed Breed) (data_json string) {

	var contentLength int64
	var contentType string

	// get the content in the API
	response, err := http.Get("https://api.thecatapi.com/v1/breeds/search?q=" + breed.Name)

	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	defer reader.Close()
	contentLength = response.ContentLength
	contentType = response.Header.Get("Content-Type")

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	bodyString := buf.String()

	// []string to JSON
	data := []string{bodyString, strconv.FormatInt(contentLength, 10), contentType}
	data_bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// result is cached
	data_json = string(data_bytes[:])
	return
}

func DataToOutput(c *gin.Context, data_json string) {
	var contentLength int64
	var contentType string

	if len(data_json) == 0 {
		c.JSON(http.StatusServiceUnavailable, []string{})
	}

	// JSON to []string
	var data_json_recover []string
	json.Unmarshal([]byte(data_json), &data_json_recover)
	data_bytes_recover, err := json.Marshal(data_json_recover)
	if err != nil {
		panic(err)
	}

	var data_recover []string
	json.Unmarshal(data_bytes_recover, &data_recover)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Network error?", r)
		}
	}()

	reader := ioutil.NopCloser(strings.NewReader(data_recover[0]))

	contentLength, err = strconv.ParseInt(data_recover[1], 10, 64)
	if err != nil {
		panic(err)
	}

	contentType = data_recover[2]

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, nil)
}
