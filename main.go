package main

import (
	"log"
)

func main() {
	c, err := InitConfig()
	if err != nil {
		log.Fatal("Could not start app. Error while reading config")
	}

	r := InitRouter(c)

	r.Run(":8000")

}
