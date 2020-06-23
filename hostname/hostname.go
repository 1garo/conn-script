package hostname

import (
	"conn-script/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func AddHostname(credentials types.Credential, name string) error {
	var hostname map[string]types.Credential
	jsonFile, _ := os.Open("pass.json")
	file, _ := ioutil.ReadAll(jsonFile)
	fmt.Printf("%v\n", string(file))
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

func ListAllHostname() (map[string]types.Credential, error) {
	var hostname map[string]types.Credential
	jsonFile, err := os.Open("pass.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &hostname)
	res := hostname
	return res, nil
}
