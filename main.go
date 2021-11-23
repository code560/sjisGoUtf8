package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	enableBom := false

	app := &cli.App{
		Usage: "Change charset shift-jis -> utf-8",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "b",
				Usage:       "global b-option",
				Destination: &enableBom,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Hello World!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
