package main

import (
	"github.com/boybird/hello/pkg/models"
	"github.com/boybird/hello/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	setting.Setup()
	models.Setup()

}
