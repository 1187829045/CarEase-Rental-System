package main

import (
	"car.rental/dao/model"
	"car.rental/global"
	Init "car.rental/init"
)

func main() {
	Init.MysqlInit()
	_ = global.DB.AutoMigrate(&model.User{})
}
