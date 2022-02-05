package main

import (
	"log"
	"os"

	httpdump "github.com/maruware/http-dump"
	"github.com/urfave/cli/v2"
)

func main() {
	var (
		port    int
		ip      string
		cert    string
		key     string
		oformat string
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
			&cli.StringFlag{
				Name:        "cert",
				Usage:       "TLS cert file path",
				Destination: &cert,
			},
			&cli.StringFlag{
				Name:        "key",
				Usage:       "TLS key file path",
				Destination: &key,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "output format (simple or json or simple_color)",
				Value:       "simple_color",
				Destination: &oformat,
			},
		},
		Action: func(c *cli.Context) error {
			opts := httpdump.ServeOpts{
				Port:   port,
				Ip:     ip,
				Cert:   cert,
				Key:    key,
				Format: oformat,
			}
			return httpdump.Serve(opts)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
