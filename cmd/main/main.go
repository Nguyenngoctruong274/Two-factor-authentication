package main

import (
	"authentication/cmd/authentication"
	"authentication/cmd/dataApplication"
	"authentication/source/utils"
	"log"
)

const (
	AuthenURI string = "mongodb://localhost:27017/authentication"
	AuthenDB  string = "authentication"
)

// var (
// 	g      errgroup.Group
// 	server *gin.Engine
// )

func main() {

	go func() {
		if err := utils.StartServer(":8081", dataApplication.Server02()); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := utils.StartServer(":8080", authentication.Server01()); err != nil {
			log.Fatal(err)
		}
	}()

	select {}
}
