package main

import (
	"log"
	"os"

	h "github.com/RyuseiNomi/menu/src/handler"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "menu",
		Usage: "simple menu works on CLI",
		Action: func(c *cli.Context) error {
			h.Handle()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
