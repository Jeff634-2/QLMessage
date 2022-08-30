package main

import (
	"QLProject/commands"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "qlSpider",
		Usage: "make an task",
	}
	app.Commands = commands.Commons()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	//app := cli.NewApp()
	//app.Name = "getQLMessage"
	//app.Usage = "make an task"
	//app.Commands = commands.Commons()
}
