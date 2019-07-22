package common

import (
	"github.com/BagusAK95/zaun/config"
	"github.com/jinzhu/gorm"
)

//NewMysqlConnection : initialized new mysql connection
func NewMysqlConnection(c *config.Configuration) (*gorm.DB, error) {
	dbType := c.Database.DbType

	dsn := c.Database.ConnectionURI
	dbConn, err := gorm.Open(dbType, dsn)
	if err != nil {
		return dbConn, err
	}

	debug := false
	if c.Server.Mode == "debug" {
		debug = true
	}

	dbConn.LogMode(debug)

	return dbConn, nil
}
