package app

import (
	cs "conn-script/credentials"
	hn "conn-script/hostname"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"text/tabwriter"
)

func Config() (*cli.App, error) {
	app := cli.NewApp()
	app.Name = "ssh-conn"
	app.Usage = "Navigate through server easily!"
	connFlags := []cli.Flag{
		&cli.StringFlag{
			Name:     "hostname",
			Aliases:  []string{"host"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "balabit",
			Aliases:  []string{"b"},
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "sftp",
			Required: false,
			Aliases:  []string{"f"},
		},
	}
	addHostnameFlags := []cli.Flag{
		&cli.StringFlag{
			Name:     "hostname",
			Aliases:  []string{"n"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "user",
			Aliases:  []string{"u"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "password",
			Aliases:  []string{"p"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "description",
			Aliases:  []string{"d"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "environment",
			Aliases:  []string{"e"},
			Required: true,
		},
	}
	ChangeCredentialsFlags := []cli.Flag{
		&cli.StringFlag{
			Name:     "hostname",
			Aliases:  []string{"n"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "user",
			Aliases:  []string{"u"},
			Required: false,
		},
		&cli.StringFlag{
			Name:     "password",
			Aliases:  []string{"p"},
			Required: false,
		},
		&cli.StringFlag{
			Name:     "description",
			Aliases:  []string{"d"},
			Required: false,
		},
		&cli.StringFlag{
			Name:     "environment",
			Aliases:  []string{"e"},
			Required: false,
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:    "connect",
			Aliases: []string{"conn"},
			Usage:   "Connect to the host",
			Action: func(c *cli.Context) error {
				bt := cs.CheckBalabit(c.String("b"))
				if !c.Bool("f") {
					err := hn.Connect(c.String("host"), bt, hn.ConnectionType)
					if err != nil {
						log.Fatal(err)
					}
					return nil
				} else {
					fmt.Println("ENTROU AQUI")
					hn.ConnectionType = "sftp"
					err := hn.Connect(c.String("host"), bt, hn.ConnectionType)
					if err != nil {
						log.Fatal(err)
					}
					return nil
				}
			},
			Flags: connFlags,
		},
		{
			Name:    "list",
			Aliases: []string{"list"},
			Usage:   "List all hostname enable to be connected",
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
					_, err := fmt.Fprintf(w, "\n %s\t\t%s\t\t", i, res[i].EnvType)
					if err != nil {
						log.Fatal(err)
					}
				}
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"add"},
			Usage:   "Add a new hostname to the json file",
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
			Name:    "change",
			Aliases: []string{"change"},
			Usage:   "Change details of a hostname",
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
