package main

import (
	"car.rental/cmd"
	Init "car.rental/init"
)

func init() {
	Init.MysqlInit()
	Init.RedisInit()
}
func main() {
	cmd.Execute()
}
