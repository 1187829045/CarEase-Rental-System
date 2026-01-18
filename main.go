package main

import (
	"car.rental/cmd"
	"car.rental/init"
)

func init() {
	Init.MysqlInit()
}
func main() {
	cmd.Execute()
}
