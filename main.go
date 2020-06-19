package main

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"encoding/json"
	gp "github.com/keybase/gexpect"
	"os"
)

//type bt struct{
//    bt11 := ""
//    bt1 := "10.168.16.84"
//    bt2 := "10.112.16.84"
//    bt3 := "10.168.16.87"
//    bt3_vpn := "10.174.225.14"
//}
var bt string
var connectionType = "ssh"
var skipSshFingerprint = "StrictHostKeyChecking=no"

// TODO: move this code to the credentials module
type CredentialTim struct {
	User     string `json:"user"`
	Password string `json:"pass"`
}
type HostnameTim struct {
	Credentials map[string]CredentialTim
}
type Credential struct {
	User        string `json:"user"`
	Password    string `json:"pass"`
	Description string `json:"description"`
	EnvType     string `json:"env_type"`
}
type Hostname struct {
	Credentials map[string]Credential
}

//func main() {
//	app := cli.NewApp()
//	app.Name = "conn - aliases to facilitate navigation on servers"
//	app.Usage = "conn -u -n -d -b -p -e -c | need parameters (e.g, -e PROD)"
//	connFlags := []cli.Flag{
//		&cli.StringFlag{
//			Name: "host",
//		},
//		&cli.StringFlag{
//			Name: "test",
//		},
//	}
//	app.Commands = []*cli.Command{
//		{
//			Name:  "conn",
//			Usage: "Connect to the host that you pass",
//			Action: func(c *cli.Context) error {
//				command := []string{c.String("host"), c.String("test")}
//				//err := connect(command)
//
//				fmt.Printf("%v", command)
//			//	if err != nil {
//			//		log.Fatal(err)
//			//	}
//				return nil
//			},
//			Flags: connFlags,
//		},
//	}
//
//	err := app.Run(os.Args)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func main() {
	var err error
	var data Credential
	data, err = GetCredentials("opcsst")
	if err != nil {
		fmt.Printf("%s\nA error happened while trying to spawn the ssh connection.", err)
	}
	fmt.Print(data)
}
func connect(hostname string, bt string) error {
	child, err := gp.Spawn(fmt.Sprintf(`%s -o %s %s@%s`, connectionType, skipSshFingerprint, hostname, bt))
	if err != nil {
		fmt.Printf("%s\nA error happened while trying to spawn the ssh connection.", err)
	}
	r, _ := regexp.Compile(".*[P-p]assword.*")
	//TODO: Call  Hostname to return user and pass bt
	userData, err := GetCredentials(hostname)
	if err != nil {
		fmt.Printf("%s\nProblem while getting your credentials.", err)
	}
	timData, err := GetCredentialsTim("tim_config")
	if err != nil {
		fmt.Printf("%s\nProblems while getting your tim credentials", err)
	}
	child.Expect("Gateway username:")
	child.SendLine(timData.User)
	child.Expect("Gateway password:")
	child.SendLine(timData.Password)
	child.Expect(fmt.Sprintf("%s", r))
	child.SendLine(userData.Password)
	child.Interact()
	return nil
}

func GetCredentials(hostname string) (Credential, error) {
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

func GetCredentialsTim(hostname string) (CredentialTim, error) {
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
