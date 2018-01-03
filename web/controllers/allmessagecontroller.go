package controllers

import (
	"fmt"
	"net/http"
)

func AllMessagesHandler(w http.ResponseWriter, r *http.Request) {

	var messageCollection []string

	fmt.Println("messageCollection", messageCollection)

}
