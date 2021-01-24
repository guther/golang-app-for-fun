package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Breed is the type of parameters. I did not put the binding:"required" because I wanna get the same behaviour of TheCatAPI
type Breed struct {
	Name string `form:"name"`
}

// GetFromWeb access THECATAPI api
func GetFromWeb(c *gin.Context, breed Breed) (dataJSON string) {

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
	dataBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// result is cached
	dataJSON = string(dataBytes[:])
	return
}

// DataToOutput is the function in charge of response request
func DataToOutput(c *gin.Context, dataJSON string) {
	var contentLength int64
	var contentType string

	if len(dataJSON) == 0 {
		c.JSON(http.StatusServiceUnavailable, []string{})
	}

	// JSON to []string
	var dataJSONRecover []string
	json.Unmarshal([]byte(dataJSON), &dataJSONRecover)
	dataBytesRecover, err := json.Marshal(dataJSONRecover)
	if err != nil {
		panic(err)
	}

	var dataRecover []string
	json.Unmarshal(dataBytesRecover, &dataRecover)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Network error?", r)
		}
	}()

	reader := ioutil.NopCloser(strings.NewReader(dataRecover[0]))

	contentLength, err = strconv.ParseInt(dataRecover[1], 10, 64)
	if err != nil {
		panic(err)
	}

	contentType = dataRecover[2]

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, nil)
}
