package main

import (
	"car.rental/global"
	Init "car.rental/init"
	"car.rental/model"
)

func main() {
	Init.MysqlInit()
	_ = global.DB.AutoMigrate(&model.User{})
}
