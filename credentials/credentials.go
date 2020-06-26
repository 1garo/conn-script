package credentials

import (
	"conn-script/types"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
)

func GetCredentials(hostname string) (*types.Credential, error) {
	var credentials types.Hostname
	jsonFile, err := os.Open("pass.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &credentials.Credentials)
	if err != nil {
		fmt.Println(err)
	}
	res := credentials.Credentials[hostname]
	return res, nil
}

func GetCredentialsTim(hostname string) (*types.CredentialTim, error) {
	var credentials types.HostnameTim
	jsonFile, err := os.Open("pass.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &credentials.Credentials)
	if err != nil {
		fmt.Println(err)
	}
	res := credentials.Credentials[hostname]
	return res, nil
}

func ChangeCredentials(credentials types.Credential, name string) error {
	// TODO: check if the field is empty, if yes, get it from the json that already exists, else you can change it
	// TODO: func CheckEmptyField(ex types.Credential)
	var hostname map[string]types.Credential
	jsonFile, _ := os.Open("pass.json")
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
	err = ioutil.WriteFile("pass.json", jsonString, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
func CreateCredentialVar(c *cli.Context) (types.Credential, error) {
	var credentials = types.Credential{
		User:        c.String("u"),
		Password:    c.String("p"),
		Description: c.String("d"),
		EnvType:     c.String("e"),
	}
	return credentials, nil
}
