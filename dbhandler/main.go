package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"project/github/comics/dbhandler/config"
	"project/github/comics/dbhandler/function"
	"project/github/comics/dbhandler/mongo"
)

type typeFunction func(*mongo.Mongo, string, []byte) error

var typeFunMapping map[string]typeFunction

func init() {
	typeFunMapping = make(map[string]typeFunction)
	typeFunMapping["characters"] = function.CharactersFunction
	typeFunMapping["character"] = function.CharactersFunction
}
func main() {
	var configFile string

	flag.StringVar(&configFile, "config", "", "Configuration file")
	flag.Parse()

	if configFile == "" {
		log.Println("Configuration file is required. Exiting.")
		os.Exit(1)
	}

	err, cfg := config.ParseConfig(configFile)

	if err != nil {
		log.Println("Error parsing configuration file. Exiting.")
		os.Exit(1)
	}

	m, err := mongo.Init(*cfg)

	if err != nil {
		log.Println("Error when connecting to db. Exiting.")
		os.Exit(1)
	}

	scheduleDBHandler(m)
}

func scheduleDBHandler(m *mongo.Mongo) {
	r := gin.Default()

	r.POST("/db/:type", func(c *gin.Context) {
		colType := c.Param("type")
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to read request body",
			})
			return
		}

		type MyData struct {
			Datatype string          `json:"datatype"`
			Data     json.RawMessage `json:"data"`
		}

		// extract 'data' portion
		var myData MyData
		if err := json.Unmarshal(body, &myData); err != nil {
			fmt.Println("Error:", err)
			return
		}
		dataJSON, err := json.Marshal(myData.Data)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		f := typeFunMapping[colType]
		err = f(m, colType, dataJSON)

		if err != nil {
			log.Println("Error inserting data ", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("Error - %s", err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Complete data for - %s", colType),
		})
	})

	r.Run(":8888")
}
