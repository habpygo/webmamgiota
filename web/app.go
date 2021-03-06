package web

import (
	"fmt"
	"net/http"

	"iota/webmamgiota/web/controllers"
)

//Serve will listen for and serve up data on port 3000
// func Serve(mammessageboard *controllers.Application)
func Serve(mamboard *controllers.MAMBoardSetup) {
	fmt.Println("We enter Serve")
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//send message|web-app starts here
	http.HandleFunc("/sendmessage.html", controllers.SendHandler)

	//show historic MAMs for particular address
	http.HandleFunc("/queryaddressformessages.html", controllers.AllMessagesForAddressHandler)

	//check messages for a particular transaction hash
	http.HandleFunc("/checkfortxid.html", controllers.CheckForTxIdHandler)

	//NOTE: if first page is changed, change it in controller.go as well
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Waiting for IOTA tangle response...")
		http.Redirect(w, r, "/sendmessage.html", http.StatusTemporaryRedirect)
	})

	fmt.Println("Listening (http://localhost:3000/) ...")
	fmt.Println("=======================================================")
	http.ListenAndServe(":3000", nil)
}
