package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/spf13/viper"

	"github.com/BagusAK95/zaun/common"
	"github.com/BagusAK95/zaun/config"
	"github.com/BagusAK95/zaun/delivery/http"
	"github.com/BagusAK95/zaun/domain/generator"

	_ "database/sql"

	"github.com/labstack/echo"
)

//init : initialized
func init() {
	//CommandlineExecute()

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

}

//main : main function
func main() {
	app, err := initializeContainer()
	if err != nil {
		fmt.Printf("Cannot start app: %+v\n", err)
		os.Exit(1)
	}
	app.Start()
}

//Container : container configuration struct
type Container struct {
	Configuration *config.Configuration
	Database      *gorm.DB
	HTTP          *echo.Echo
}

//initializeContainer : initialized container
func initializeContainer() (Container, error) {
	config, errNewConfig := config.New()
	if errNewConfig != nil {
		return Container{}, errNewConfig
	}
	db, errNewConnMysql := common.NewMysqlConnection(config)
	if errNewConnMysql != nil {
		return Container{}, errNewConnMysql
	}
	redis, errNewConnRedis := common.NewRedisConnection(config)
	if errNewConnRedis != nil {
		return Container{}, errNewConnRedis
	}
	cache := common.NewCache(redis, config)

	generator := generator.Init(db, config, cache)
	httpHandler := http.New(*generator)

	return Container{
		Configuration: config,
		Database:      db,
		HTTP:          httpHandler,
	}, nil
}

//Start : start server
func (e Container) Start() {
	conf := e.Configuration.Server
	go func() {
		e.HTTP.Start(conf.Addr)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Shutting down server")

	duration := time.Duration(conf.ShutdownTimeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	if err := e.HTTP.Shutdown(ctx); err != nil {
		fmt.Printf("Failed to shut down server gracefully: %s", err)
	}
}
