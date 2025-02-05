package main

import (
	"log"

	starter "github.com/Dhairya3124/couchbase-go-starter/starter"
)

func main() {
	if err := starter.Run(); err != nil {
		log.Fatal(err)
	}
}
