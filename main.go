package main

import (
	"flag"
	"log"
	"os"
	"project/github/comics/client/datajob/character"
	"project/github/comics/client/datajob/characters"
)

func main() {
	var tokenFlag string

	flag.StringVar(&tokenFlag, "token", "", "api authentication token")
	flag.Parse()

	if tokenFlag == "" {
		log.Println("Token cannot be found. Exiting.")
		os.Exit(1)
	}

	scheduleTask(tokenFlag)
	for {
	}
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
