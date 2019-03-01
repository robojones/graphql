//go:generate prisma generate
//go:generate go run gqlgen/cmd.go -c gqlgen/gqlgen.yml
//go:generate wire

package main

import (
	"log"
)

func main() {
	server, err := Initialize()
	if err != nil {
		panic(err)
	}

	log.Printf("Server is running on http://localhost:%s", server.Config.Port)
	err = server.Listen()

	if err != nil {
		log.Fatal(err)
	}
}
