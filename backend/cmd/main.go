package main

import (
	"fmt"
	"github.com/KainoaGardner/webMahjongCalc/api"
	"github.com/KainoaGardner/webMahjongCalc/config"
	"log"
)

func main() {
	server := api.NewAPIServer(fmt.Sprintf("%s:%s", config.Envs.PublicHost, config.Envs.Port))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
