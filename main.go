package main

import (
	"io"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	app := &cli.App{
		Usage: "Change charset shift-jis -> utf-8",
		Action: func(c *cli.Context) error {
			reader := transform.NewReader(os.Stdin, japanese.ShiftJIS.NewDecoder())
			if _, err := io.Copy(os.Stdout, reader); err != nil {
				return err
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
