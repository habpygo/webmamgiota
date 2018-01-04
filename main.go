package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/iotaledger/mamgoiota"
	"github.com/iotaledger/mamgoiota/web"
	"github.com/iotaledger/mamgoiota/web/controllers"
	"github.com/iotaledger/mamgoiota/mamutils"
)

func main() {

	myboard := &controllers.

	app := &controllers.Application{
		Mamboard: myboard,
	}
	web.Serve(app)
}