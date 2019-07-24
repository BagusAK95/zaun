package main

import (
	"fmt"
	"os"

	"github.com/BagusAK95/zaun/common"
	"github.com/BagusAK95/zaun/config"
	"github.com/BagusAK95/zaun/domain/route"
	"github.com/BagusAK95/zaun/domain/target"
)

//CommandlineExecute : execute command line
func CommandlineExecute() {
	if len(os.Args) <= 1 {
		return
	}
	arg := os.Args[1]

	switch arg {
	case "mysql_migrate":
		mysqlMigrate()
		os.Exit(0)
	}
}

//mysqlMigrate : migrate mysql structure
func mysqlMigrate() {
	config, _ := config.New()
	db, err := common.NewMysqlConnection(config)
	if err != nil {
		fmt.Printf("Cannot connect to database %+v\n", err)
		os.Exit(1)
	}
	db.AutoMigrate(
		&route.Route{},
		&target.Target{},
	)
}
