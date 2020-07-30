package terminal

import (
	//"fmt"
	//"fmt"
	//"log"
	//"strconv"

	"conn-script/types"
	//"fmt"

	//"log"
	"os"
	"os/exec"
	//"strconv"
	"strings"
)

func Size(terminalData *types.Terminal) {
	var a []string
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	a = make([]string, len(string(out)))
	a = strings.Split(string(out), " ")
	terminalData.Rows = a[0]
	columns := strings.Trim(a[1], "\n")
	terminalData.Columns = columns
}