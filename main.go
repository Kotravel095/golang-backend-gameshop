package main

import (
	"github.com/Kotravel095/golang-backend-gameshop/config"
	"github.com/Kotravel095/golang-backend-gameshop/databases"
	"github.com/Kotravel095/golang-backend-gameshop/server"
)

func main(){
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db.ConnectionGetting())
	server.Start()
}

