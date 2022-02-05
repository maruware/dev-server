package main

import (
	"log"
	"os"

	httpdump "github.com/maruware/http-dump"
	"github.com/urfave/cli/v2"
)

func main() {
	var (
		port int
		ip   string
	)

	app := &cli.App{
		Name:  "http-dump",
		Usage: "http dump dev server",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Value:       8080,
				Usage:       "listen port",
				Destination: &port,
			},
			&cli.StringFlag{
				Name:        "bind",
				Aliases:     []string{"b"},
				Value:       "0.0.0.0",
				Usage:       "listen ip",
				Destination: &ip,
			},
		},
		Action: func(c *cli.Context) error {
			opts := httpdump.ServeOpts{
				Port: port,
				Ip:   ip,
			}
			return httpdump.Serve(opts)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
