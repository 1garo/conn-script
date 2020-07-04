package main

import (
	"conn-script/app"
	cs "conn-script/credentials"
	"fmt"
	"log"
	"os"
)

func main() {
	// TODO: make the credentials.json reads with absolute path
	file, err := cs.GetCredentialsFile()
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println("credentials.json file doesn't exist!")
		return
	}
	application, err := app.Config()
	if err != nil {
		log.Fatal(err)
	}
	err = application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
