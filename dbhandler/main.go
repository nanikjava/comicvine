package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"project/github/comics/client/models"
	"project/github/comics/client/utils"
	"project/github/comics/dbhandler/dbconfig"
	"project/github/comics/dbhandler/mongo"
)

type typeFunction func(*mongo.Mongo, string, []byte) error

var typeFunMapping map[string]typeFunction

func init() {
	typeFunMapping = make(map[string]typeFunction)
	typeFunMapping["characters"] = charactersFunction
	typeFunMapping["character"] = charactersFunction
}
func main() {
	var configFile string

	flag.StringVar(&configFile, "config", "", "Configuration file")
	flag.Parse()

	if configFile == "" {
		log.Println("Configuration file is required. Exiting.")
		os.Exit(1)
	}

	err, cfg := dbconfig.ParseConfig(configFile)

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

		f := typeFunMapping[colType]
		err = f(m, colType, body)

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

func charactersFunction(m *mongo.Mongo, coltype string, body []byte) error {
	var bsonDocument bson.M
	err := bson.UnmarshalExtJSON(body, false, &bsonDocument)
	if err != nil {
		return err
	}

	_, err = m.Create(context.Background(), coltype, &models.CreateRequest{
		Document:  bsonDocument,
		Operation: utils.One,
		IsBatch:   false,
	})

	return err
}
