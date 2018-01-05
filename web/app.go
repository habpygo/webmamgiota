package web

import (
	"fmt"
	"net/http"

	"github.com/iotaledger/mamgoiota/web/controllers"
)

//Serve will listen for and serve up data on port 3000
// func Serve(mammessageboard *controllers.Application)
func Serve(mamboard *controllers.MAMBoardSetup) {
	fmt.Println("We enter Serve")
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/queryallmessages.html", controllers.AllMessagesHandler)
	//http.HandleFunc("/", controllers.AllMessagesHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HandleFunc will redirect the / to queryallmessages.html")
		http.Redirect(w, r, "/queryallmessages.html", http.StatusTemporaryRedirect)
	})

	fmt.Println("Listening (http://localhost:3000/) ...")
	fmt.Println("=======================================================")
	http.ListenAndServe(":3000", nil)
}
