package main

import (
	"github.com/boybird/hello/pkg/gredis"
	"github.com/boybird/hello/pkg/logging"
	"github.com/boybird/hello/pkg/models"
	"github.com/boybird/hello/pkg/setting"
	"github.com/boybird/hello/pkg/util"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

}
