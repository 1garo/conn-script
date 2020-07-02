# Connection script - 
Aliases to facilitate connect on servers based on a json file. 

Technologies that are being used:
- [go](https://golang.org/)
- [clion](https://www.jetbrains.com/pt-br/clion/)
- [go mod to see external libs](go.mod)
- JSON 
- SSH/SFTP protocol
## Installation

**Create a .env file with the following information:**

At least create the BT2 variable 'cause it's the default value (e.g: BT1, BT2, BT3 or BT3_VPN):

```BT2 = "<your bash path>" (e.g SH = "~/.zshrc")```

SH variable with the current bash that you are using:

```SH = "<your bash path>" (e.g SH = "~/.zshrc")```


**Create a file pass.json with the following structure:**

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
ssh-conn a --n hostaname --u user --p password --d description -e DEV
```
**Connect to a hostname** 
```bash
ssh-conn n --host <hostname> 
```
**Change a hostname** 
```bash
ssh-conn c --n testeHostname --u userTeste --p passwordTest --d descriptionTeste -e DEV
```
**List all hostname available** 
```bash
ssh-conn l
```
**Use help to lean more information about it** 
```bash
ssh-conn -h or --help
```
## License
[MIT](https://choosealicense.com/licenses/mit/)