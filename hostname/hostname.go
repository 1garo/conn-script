package hostname

import (
	cs "conn-script/credentials"
	expect "conn-script/expect"
	"conn-script/types"
	"encoding/json"
	"fmt"
	//"io"
	"io/ioutil"
	"log"
	"os"
	//	"regexp"
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
	// TODO: Write a spawn/expect feature, all that exist is shit!
}
