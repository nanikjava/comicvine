package main

import (
	"flag"
	"log"
	"os"
	"project/github/comics/client/datajob/character"
)

func main() {
	var tokenFlag string

	flag.StringVar(&tokenFlag, "token", "", "api authentication token")
	flag.Parse()

	if tokenFlag == "" {
		log.Println("Token cannot be found. Exiting.")
		os.Exit(1)
	}

	getCharacters(tokenFlag)
}

func getCharacters(tokenFlag string) {
	c := character.New(tokenFlag)
	_, err := c.GetData()

	if err != nil {
		log.Println("error : ", err)
		os.Exit(1)
	}
}
