# Connection script 
Aliases to facilitate connect on servers based on a json file. 

Technologies that are being used:
- [go](https://golang.org/)
- [clion](https://www.jetbrains.com/pt-br/clion/)
- [go mod to see external libs](go.mod)
- [json](https://www.json.org/json-en.html) 
- [ssh/sftp protocol](https://pt.wikipedia.org/wiki/Secure_Shell)

## Installation

**Create a .env file with the following information:**

At least create the BT2 variable 'cause it's the default value (e.g: BT1, BT2, BT3 or BT3_VPN):

```BT2 = "<balabit ip>"```

SH variable with the current bash that you are using:

```SH = "<your bash path>" (e.g SH = "~/.zshrc")```


**Create a file credentials.json with the following structure:**

```json5
{
    "testHostname": {
        "user": "userTest",
        "pass": "passwordTest",
        "description": "descriptionTest",
        "env_type": "DEV"
    },
}
```

After it, exec the commands below on the root folder:

**build** - compile the binary to create the alias.  
**install** - create the alias to make your life easier.
```bash
make build 
```
```bash
make install 
```
## Usage 
**Add a hostname** 
```bash
ssh-conn add -n hostaname -u user -p password -d description -e DEV
```
**Connect to a hostname** 
```bash
ssh-conn conn -host <hostname> 
```
**Change a hostname** 
```bash
ssh-conn change -n testeHostname -u userTest -p passwordTest -d descriptionTest -e DEV
```
**List all hostname available** 
```bash
ssh-conn list
```
**Use help to lean more information about it** 
```bash
ssh-conn -h or --help
```
## License
[MIT](https://choosealicense.com/licenses/mit/)
