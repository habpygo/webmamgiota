package main

import (
	"fmt"
	"os/exec"
	"runtime"

	"iota/webmamgiota/web"
	"iota/webmamgiota/web/controllers"
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
	//TODO: find out whether concurrency management would be appropriate here

	//open a new webpage...
	open("http://localhost:3000/")
	//...and serve it the MAMBoard
	web.Serve(msgwebpage)
}
