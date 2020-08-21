package main

import (
	"log"
	"os"

	"github.com/rivo/tview"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "menu",
		Usage: "simple menu works on CLI",
		Action: func(c *cli.Context) error {
			handle()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func handle() {
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("ずかん", "", 'd', nil).
		AddItem("コンテナ", "", 'c', nil).
		AddItem("アプリ", "", 'a', nil).
		AddItem("レポート", "", 'r', nil).
		AddItem("せってい", "", 's', nil).
		AddItem("とじる", "", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).Run(); err != nil {
		panic(err)
	}
}
