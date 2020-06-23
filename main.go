package main

import (
	"encoding/json"
	"fmt"
	gp "github.com/keybase/gexpect"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"text/tabwriter"
)

//type bt struct{
//    bt11 := ""
//    bt1 := "10.168.16.84"
var bt2 = "10.112.16.84"

//    bt3 := "10.168.16.87"
//    bt3_vpn := "10.174.225.14"
//}
var connectionType = "ssh"
var skipSshFingerprint = "StrictHostKeyChecking=no"

type CredentialTim struct {
	User     string `json:"user"`
	Password string `json:"pass"`
}

type HostnameTim struct {
	Credentials map[string]*CredentialTim
}

type Credential struct {
	User        string `json:"user"`
	Password    string `json:"pass"`
	Description string `json:"description"`
	EnvType     string `json:"env_type"`
}
type Hostname struct {
	Credentials map[string]*Credential
}

func main() {
	app := cli.NewApp()
	app.Name = "conn - aliases to facilitate navigation on servers"
	app.Usage = "conn -u -n -d -b -p -e -c | need parameters (e.g, -e PROD)"
	connFlags := []cli.Flag{
		&cli.StringFlag{
			Name: "host",
		},
	}
	addFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "n",
			Value: "",
		},
		&cli.StringFlag{
			Name:  "u",
			Value: "",
		},
		&cli.StringFlag{
			Name:  "p",
			Value: "",
		},
		&cli.StringFlag{
			Name:  "d",
			Value: "",
		},
		&cli.StringFlag{
			Name:  "e",
			Value: "",
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
				res, err := listAllHostnames()
				if err != nil {
					log.Fatal(err)
				}
				w := new(tabwriter.Writer)
				w.Init(os.Stdout, 8, 8, 0, '\t', 0)
				defer w.Flush()
				_, err = fmt.Fprintf(w, "\n %s\t%s\t", "Hostname", "Environment")
				if err != nil {
					log.Fatal(err)
				}
				_, err = fmt.Fprintf(w, "\n %s\t%s\t", "--------", "-----------")
				if err != nil {
					log.Fatal(err)
				}
				for i := range res {
					if i == "tim_config" {
						continue
					}
					fmt.Fprintf(w, "\n %s\t%s\t", res[i].User, res[i].EnvType)
				}
				return nil
			},
		},
		{
			Name:  "a",
			Usage: "Add a new hostname to the json file",
			Action: func(c *cli.Context) error {
				credentials := Credential{
					User:        c.String("u"),
					Password:    c.String("p"),
					Description: c.String("d"),
					EnvType:     c.String("e"),
				}
				err := addHostname(credentials, c.String("n"))
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
			Flags: addFlags,
		},
		{
			Name:  "c",
			Usage: "Change details of a hostname",
			Action: func(c *cli.Context) error {
				fmt.Println("Change function")
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func addHostname(credentials Credential, name string) error {
	var hostname map[string]Credential
	jsonFile, _ := os.Open("pass.json")
	file, _ := ioutil.ReadAll(jsonFile)
	fmt.Printf("%v\n", string(file))
	json.Unmarshal(file, &hostname)
	hostname[name] = credentials
	jsonString, err := json.MarshalIndent(hostname, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("pass.json", jsonString, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func connect(hostname string, bt string) error {
	userData, err := GetCredentials(hostname)
	if err != nil {
		fmt.Printf("%s\nProblem while getting your credentials.", err)
	}
	timData, err := GetCredentialsTim("tim_config")
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

func GetCredentials(hostname string) (*Credential, error) {
	jsonFile, err := os.Open("pass.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var credentials Hostname
	err = json.Unmarshal(byteValue, &credentials.Credentials)
	if err != nil {
		fmt.Println(err)
	}
	res := credentials.Credentials[hostname]
	return res, nil
}

func GetCredentialsTim(hostname string) (*CredentialTim, error) {
	jsonFile, err := os.Open("pass.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var credentials HostnameTim
	err = json.Unmarshal(byteValue, &credentials.Credentials)
	if err != nil {
		fmt.Println(err)
	}
	res := credentials.Credentials[hostname]
	return res, nil
}

func listAllHostnames() (map[string]Credential, error) {
	jsonFile, err := os.Open("pass.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	var hostname map[string]Credential
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &hostname)
	res := hostname
	return res, nil
}
