package main

import (
	"fmt"

	"github.com/iotaledger/mamgoiota/web"
	"github.com/iotaledger/mamgoiota/web/controllers"
)

func main() {

	msgwebpage := &controllers.MAMBoardSetup{}
	fmt.Println("We enter main()")
	web.Serve(msgwebpage)
}
