package main

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/iotaledger/mamgoiota/web"
	"github.com/iotaledger/mamgoiota/web/controllers"
)

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func main() {

	msgwebpage := &controllers.MAMBoardSetup{}
	fmt.Println("We enter main()")
	open("http://localhost:3000/")
	web.Serve(msgwebpage)
}
