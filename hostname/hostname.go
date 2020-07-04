package hostname

import (
	cs "conn-script/credentials"
	"conn-script/types"
	"encoding/json"
	"fmt"
	gp "github.com/keybase/gexpect"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

var ConnectionType = "ssh"

const skipSshFingerprint = "StrictHostKeyChecking=no"

func AddHostname(credentials types.Credential, name string) error {
	var hostname map[string]types.Credential
	filename, err := cs.GetCredentialsFile()
	if err != nil {
		log.Fatal(err)
	}
	jsonFile, _ := os.Open(filename)
	file, _ := ioutil.ReadAll(jsonFile)
	errMarshal := json.Unmarshal(file, &hostname)
	if errMarshal != nil {
		log.Fatal(errMarshal)
	}
	hostname[name] = credentials
	jsonString, err := json.MarshalIndent(hostname, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("credentials.json", jsonString, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func ListAllHostname() (map[string]types.Credential, error) {
	var hostname map[string]types.Credential
	filename, err := cs.GetCredentialsFile()
	if err != nil {
		log.Fatal(err)
	}
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &hostname)
	res := hostname
	return res, nil
}

func Connect(hostname string, bt string, connectionType string) error {
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
	child.Expect("Gateway username:")
	child.SendLine(timData.User)
	child.Expect("Gateway password:")
	child.SendLine(timData.Password)
	child.ExpectRegex(fmt.Sprintf("%s", r))
	child.SendLine(userData.Password)
	child.Interact()
	// TODO: create the color highlight based on the env and send the stty when conn-type were ssh
	return nil
}
