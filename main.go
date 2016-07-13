package main

import (
	_ "expvar"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli"
)

//
const MAJOR = "0.1"

// DO NOT UPDATE. This is populated by the build. See the Makefile.
var VERSION = "0"

//
func debugger() {
	func() {
		srv := &http.Server{
			Addr:         ":8008",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		fmt.Println("listening on 8008")
		fmt.Println(srv.ListenAndServe())
	}()
}

var settings struct {
	debugger  bool
	overwrite bool
	analyze   struct {
		overwrite bool
		this      struct {
			overwrite bool
		}
		that struct {
			overwrite bool
		}
	}
}

//
func main() {
	app := cli.NewApp()
	app.Name = "test"
	app.Usage = "Test."
	app.Version = MAJOR + "." + VERSION
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		cli.Command{
			Name:    "debug",
			Aliases: []string{"debugger", "debugging", "d"},
			Usage:   "Debug.",

			Action: func(ctx *cli.Context) (err error) {
				debugger()
				return nil
			},
		},
		cli.Command{
			Name:      "analyze",
			Aliases:   []string{"a"},
			Usage:     "Analyze.",
			ArgsUsage: "name",

			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "overwrite, force, f",
					Usage:       "Overwrite existing schema.",
					Destination: &settings.overwrite,
				},
				cli.BoolFlag{
					Name:        "debugging, debug, D",
					Usage:       "Enable debugging server.",
					Destination: &settings.debugger,
				},
			},

			Subcommands: []cli.Command{
				cli.Command{
					Name:      "this",
					Aliases:   []string{"t"},
					Usage:     "This.",
					ArgsUsage: "name",

					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:        "overwrite, force, f",
							Usage:       "Overwrite existing schema.",
							Destination: &settings.analyze.this.overwrite,
						},
					},

					Action: func(ctx *cli.Context) (err error) {
						log.Println("this", time.Now())
						return nil
					},
				},
				cli.Command{
					Name:      "that",
					Aliases:   []string{"a"},
					Usage:     "That.",
					ArgsUsage: "name",

					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:        "overwrite, force, f",
							Usage:       "Overwrite existing schema.",
							Destination: &settings.analyze.that.overwrite,
						},
					},

					Action: func(ctx *cli.Context) (err error) {
						log.Println("that", time.Now())
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
