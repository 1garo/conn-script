package credentials

import (
	"conn-script/types"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func GetCredentials(hostname string) (*types.Credential, error) {
	var credentials types.Hostname
	filename, err := GetCredentialsFile()
	if err != nil {
		log.Fatal(err)
	}
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &credentials.Credentials)
	if err != nil {
		log.Fatal(err)
	}
	res := credentials.Credentials[hostname]
	return res, nil
}

func GetCredentialsTim(hostname string) (*types.CredentialTim, error) {
	var credentials types.HostnameTim
	filename, err := GetCredentialsFile()
	if err != nil {
		log.Fatal(err)
	}
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &credentials.Credentials)
	if err != nil {
		log.Fatal(err)
	}
	res := credentials.Credentials[hostname]
	return res, nil
}

func ChangeCredentials(credentials types.Credential, name string) error {
	var hostname map[string]types.Credential
	filename, err := GetCredentialsFile()
	if err != nil {
		log.Fatal(err)
	}
	jsonFile, _ := os.Open(filename)
	file, _ := ioutil.ReadAll(jsonFile)
	errMarshal := json.Unmarshal(file, &hostname)
	if errMarshal != nil {
		log.Fatal(errMarshal)
	}
	credentialsChecked, err := CheckEmptyField(credentials, name)
	if err != nil {
		log.Fatal(err)
	}
	hostname[name] = credentialsChecked
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

func CreateCredentialVar(c *cli.Context) types.Credential {
	var credentials = types.Credential{
		User:        c.String("u"),
		Password:    c.String("p"),
		Description: c.String("d"),
		EnvType:     c.String("e"),
	}
	return credentials
}

func CheckEmptyField(credentials types.Credential, name string) (types.Credential, error) {
	var hostname map[string]types.Credential
	filename, err := GetCredentialsFile()
	if err != nil {
		log.Fatal(err)
	}
	jsonFile, _ := os.Open(filename)
	file, _ := ioutil.ReadAll(jsonFile)
	errMarshal := json.Unmarshal(file, &hostname)
	if errMarshal != nil {
		log.Fatal(errMarshal)
	}
	if credentials.EnvType == "" {
		credentials.EnvType = hostname[name].EnvType
	}
	if credentials.Description == "" {
		credentials.Description = hostname[name].Description
	}
	if credentials.User == "" {
		credentials.User = hostname[name].User
	}
	if credentials.Password == "" {
		credentials.Password = hostname[name].Password
	}
	return credentials, nil
}

func CheckBalabit(bt string) string {
	var res string
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	if bt == "1" {
		res = os.Getenv("BT1")
	} else if bt == "2" {
		res = os.Getenv("BT2")
	} else if bt == "3" {
		res = os.Getenv("BT3")
	} else if bt == "3_vpn" {
		res = os.Getenv("BT3_VPN")
	} else if bt == "" {
		res = os.Getenv("BT2")
	}
	return res
}

func GetCredentialsFile() (string, error) {
	var path, err = filepath.Abs("credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	return path, nil
}
