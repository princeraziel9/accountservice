package main

import (
	"fmt"
	"github.com/callistaenterprise/goblog/accountservice/dbclient"
	"github.com/callistaenterprise/goblog/accountservice/service"
)

var appName = "Account Service"

func main() {
	fmt.Println( "Starting " + appName  )
	initBoltClient()
	service.StartWebServer( "6767" )
}

func initBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}