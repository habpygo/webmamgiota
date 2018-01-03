package web

import (
	"fmt"
	"net/http"

	"github.com/iotaledger/mamgoiota/web/controllers"
)

//Serve will listen for and serve up data on port 3000
func Serve(mammessageboard *controllers.MamSetup) {
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//login handlers
	http.HandleFunc("/index.html", mammessageboard.IndexHandler)
	http.HandleFunc("/login", mammessageboard.LoginHandler)

	http.HandleFunc("/tryout.html", mammessageboard.TryoutHandler)

	//inspection and query handlers
	http.HandleFunc("/queryallmessages.html", mammessageboard.AllMessagesHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/index.html", http.StatusTemporaryRedirect)
	})

	fmt.Println("Listening (http://localhost:3000/) ...")
	fmt.Println("=======================================================")
	http.ListenAndServe(":3000", nil)
}
