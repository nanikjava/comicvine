package main

import (
	"flag"
	"github.com/go-co-op/gocron"
	"log"
	"os"
	"project/github/comics/datacollector/configuration"
	"project/github/comics/datacollector/datajob/character"
	"project/github/comics/datacollector/datajob/characters"
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

	scheduleTask(cfg.Token)
}

func scheduleTask(tokenFlag string) {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Cron("*/1 * * * *").Do(func() { getCharacters(tokenFlag) }) // every minute

	if err != nil {
		log.Fatalf("Error scheduling")
	}
	s.StartBlocking()
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
