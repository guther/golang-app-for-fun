package config

import (
	"log"
	"os"
	"strconv"
)

var (
	// PORT is the defalt webserver port number
	PORT = 0
)

// Load function gets the PORT variable from environment or it's is used 6060 port
func Load() {
	var err error
	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println(err)
		PORT = 6060
	}
}
