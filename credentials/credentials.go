package credentials

import (
	"conn-script/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
