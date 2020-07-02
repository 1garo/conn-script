package main

import (
	cs "conn-script/credentials"
	hn "conn-script/hostname"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"text/tabwriter"
)

func main() {
	// TODO: make the pass.json reads with absolute path
	if _, err := os.Stat("pass.json"); os.IsNotExist(err) {
		fmt.Println("pass.json file doesn't exist!")
		return
	}
	app, err := appConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func appConfig() (*cli.App, error) {
	app := cli.NewApp()
	app.Name = "ssh-conn - Aliases to facilitate connect on servers based on a json file"
	app.Usage = "see the Usage below for more information."
	connFlags := []cli.Flag{
		&cli.StringFlag{
			Name:     "host",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "b",
			Required: false,
		},
	}
	addHostnameFlags := []cli.Flag{
		&cli.StringFlag{
			Name:     "n",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "u",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "p",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "d",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "e",
			Required: false,
		},
	}
	ChangeCredentialsFlags := []cli.Flag{
		&cli.StringFlag{
			Name:     "n",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "u",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "p",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "d",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "e",
			Required: false,
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:  "n",
			Usage: "Connect to the host that you pass",
			Action: func(c *cli.Context) error {
				bt := cs.CheckBalabit(c.String("b"))
				err := hn.Connect(c.String("host"), bt)
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
			Flags: connFlags,
		},
		{
			Name:  "l",
			Usage: "List all hostname unable to be connected",
			Action: func(c *cli.Context) error {
				res, err := hn.ListAllHostname()
				if err != nil {
					log.Fatal(err)
				}
				w := new(tabwriter.Writer)
				w.Init(os.Stdout, 8, 8, 0, '\t', 0)
				defer w.Flush()
				_, err = fmt.Fprintf(w, "\n %s\t\t%s\t\t", "Hostname", "Environment")
				if err != nil {
					log.Fatal(err)
				}
				_, err = fmt.Fprintf(w, "\n %s\t\t%s\t\t", "--------", "-----------")
				if err != nil {
					log.Fatal(err)
				}
				for i := range res {
					if i == "tim_config" {
						continue
					}
					fmt.Fprintf(w, "\n %s\t\t%s\t\t", i, res[i].EnvType)
				}
				return nil
			},
		},
		{
			Name:  "a",
			Usage: "Add a new hostname to the json file",
			Action: func(c *cli.Context) error {
				credentials := cs.CreateCredentialVar(c)
				err := hn.AddHostname(credentials, c.String("n"))
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
			Flags: addHostnameFlags,
		},
		{
			Name:  "c",
			Usage: "Change details of a hostname",
			Action: func(c *cli.Context) error {
				credentials := cs.CreateCredentialVar(c)
				err := cs.ChangeCredentials(credentials, c.String("n"))
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
			Flags: ChangeCredentialsFlags,
		},
	}
	return app, nil
}
