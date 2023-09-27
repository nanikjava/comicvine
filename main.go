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
	"project/github/comics/client/configuration"
	"project/github/comics/client/datajob/character"
	"project/github/comics/client/datajob/characters"
	"project/github/comics/client/models"
	"project/github/comics/client/mongo"
	"project/github/comics/client/utils"
	"time"
)

func main() {
	var configFile string

	flag.StringVar(&configFile, "config", "", "Configuration file")
	flag.Parse()

	if configFile == "" {
		log.Println("Configuration file is required. Exiting.")
		os.Exit(1)
	}

	err, cfg := configuration.ParseConfig(configFile)

	if err != nil {
		log.Println("Error parsing configuration file. Exiting.")
		os.Exit(1)
	}

	m, err := mongo.Init(*cfg)

	if err != nil {
		log.Println("Error when connecting to db. Exiting.")
		os.Exit(1)
	}

	go func() {
		time.Sleep(5 * time.Second)
		scheduleTask(cfg.Token)
	}()
	go func() {
		scheduleDBHandler(m)
	}()
	for {
	}
}

func scheduleDBHandler(m *mongo.Mongo) {
	r := gin.Default()

	// Define a route with a dynamic route parameter named "type"
	r.POST("/db/:type", func(c *gin.Context) {
		colType := c.Param("type")
		requestBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to read request body",
			})
			return
		}

		var bsonDocument bson.M
		err = bson.UnmarshalExtJSON(requestBody, false, &bsonDocument)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to parse request body as BSON",
			})
			return
		}

		rows, err := m.Create(context.Background(), colType, &models.CreateRequest{
			Document:  bsonDocument,
			Operation: utils.One,
			IsBatch:   false,
		})

		if err != nil {
			log.Println("Error inserting data ", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Total number of record inserted - %d for type - %s", rows, colType),
		})
	})

	r.Run(":8888")
}

func scheduleTask(tokenFlag string) {
	//s := gocron.NewScheduler(time.UTC)
	//_, err := s.Cron("*/1 * * * *").Do(func() { getCharacters(tokenFlag) }) // every minute
	//
	//if err != nil {
	//	log.Fatalf("Error scheduling")
	//}
	//s.StartAsync()
	getCharacters(tokenFlag)

}

func getCharacters(tokenFlag string) {
	c := characters.New(tokenFlag)
	err := c.GetData("")

	if err != nil {
		log.Println("error : ", err)
		os.Exit(1)
	}

	chrc := character.New(tokenFlag)
	for i := 0; i < c.Len(); i++ {
		r := c.Get(i)

		for j := 0; j < len(r.Results); j++ {
			l := r.Results[j]
			chrc.GetData(l.APIDetailURL)
		}
	}
}
