package common

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecCmd(name string, params ...string) {
	process := exec.Command(name, params...)
	bs, err := process.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Oooops...] Execute command [%s] error: %v, %s", process.String(), err, string(bs))
		os.Exit(1)
	}

}
