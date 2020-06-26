# Connection script [golang version]
Aliases to facilitate connect on servers based on a json file. 
## Installation

Create a .env file in the repo with a variable called SH:
 
```SH = "<your bash path>" (e.g SH = "~/.zshrc")```

Create a file pass.json with the following structure:

```json5
{
    "testeHostname": {
        "user": "TESTANDO",
        "pass": "passwordTest",
        "description": "descriptionTestt",
        "env_type": "DEV"
    },
}
```


After it, exec the commands below on the root folder:

*build* - compile the binary to create the alias.  
*install* - create the alias to make your life easier.
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