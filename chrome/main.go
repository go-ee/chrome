package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/go-ee/chrome"
	"github.com/urfave/cli"
	"os"
)

const flagDebug = "debug"
const flagPattern = "pattern"

func main() {
	app := cli.NewApp()
	app.Usage = "Chrome Utils"
	app.Version = "1.0"

	app.Flags = []cli.Flag{}

	app.Commands = []*cli.Command{
		{
			Name:  "deleteExtension",
			Usage: "Delete extension by file pattern inside of extension",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  flagPattern,
					Usage: "File pattern, e.g. walkme*",
				},
			},
			Action: func(c *cli.Context) (err error) {
				if err = chrome.DeleteExtension(c.String(flagPattern)); err != nil {
					logrus.Errorf("error %v by %v to %v", err, c.Command.Name)
					return
				}
				return
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Warn("exit because of error.")
	}
}
