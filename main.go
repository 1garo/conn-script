package main

import (
	cs "conn-script/credentials"
	hn "conn-script/hostname"
	"fmt"
	gp "github.com/keybase/gexpect"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"regexp"
	"text/tabwriter"
)

const bt2 = "10.112.16.84"
const connectionType = "ssh"
const skipSshFingerprint = "StrictHostKeyChecking=no"

func main() {
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
			Name: "host",
		},
	}
	addHostnameFlags := []cli.Flag{
		&cli.StringFlag{
			Name: "n",
		},
		&cli.StringFlag{
			Name: "u",
		},
		&cli.StringFlag{
			Name: "p",
		},
		&cli.StringFlag{
			Name: "d",
		},
		&cli.StringFlag{
			Name: "e",
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
				err := connect(c.String("host"), bt2)
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
				credentials, err := cs.CreateCredentialVar(c)
				if err != nil {
					log.Fatal(err)
				}
				err = hn.AddHostname(credentials, c.String("n"))
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
				credentials, err := cs.CreateCredentialVar(c)
				if err != nil {
					log.Fatal(err)
				}
				err = cs.ChangeCredentials(credentials, c.String("n"))
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

func connect(hostname string, bt string) error {
	userData, err := cs.GetCredentials(hostname)
	if err != nil {
		fmt.Printf("%s\nProblem while getting your credentials.", err)
	}
	timData, err := cs.GetCredentialsTim("tim_config")
	if err != nil {
		fmt.Printf("%s\nProblems while getting your tim credentials", err)
	}
	child, err := gp.Spawn(fmt.Sprintf(`%s -o %s %s@%s`, connectionType, skipSshFingerprint, userData.User, bt))
	if err != nil {
		fmt.Printf("%s\nA error happened while trying to spawn the ssh connection.", err)
	}
	r, _ := regexp.Compile(".*[P-p]assword.*")
	fmt.Println(r)
	child.Expect("Gateway username:")
	child.SendLine(timData.User)
	child.Expect("Gateway password:")
	child.SendLine(timData.Password)
	child.ExpectRegex(fmt.Sprintf("%s", r))
	child.SendLine(userData.Password)
	child.Interact()
	return nil
}
