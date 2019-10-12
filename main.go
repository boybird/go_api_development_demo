package main

import (
	"github.com/boybird/hello/pkg/gredis"
	"github.com/boybird/hello/pkg/logging"
	"github.com/boybird/hello/pkg/models"
	"github.com/boybird/hello/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()

}
