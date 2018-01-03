package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/iotaledger/mamgoiota"
)

//MessageStruct is the data to be exposed in the web-app
//It's the same as Transaction type
type MessageStruct struct {
	Message   string
	Value     int64
	Timestamp time.Time
	Recipient string
}

func AllMessagesHandler(w http.ResponseWriter, r *http.Request) {
	address := "RQP9IFNFGZGFKRVVKUPMYMPZMAICIGX9SVMBPNASEBWJZZAVDCMNOFLMRMFRSQVOQGUVGEETKYFCUPNDDWEKYHSALY"
	provider := "http://node02.iotatoken.nl:14265"
	c, err := mamgoiota.NewConnection(provider, "")
	if err != nil {
		panic(err)
	}

	messageCollection := []MessageStruct{}

	//PROBABLY HAVE TO UNMARSHAL IT AND APPEND TO A TEMPVALUE
	messageCollection, err = mamgoiota.ReadTransactions(address, c)
	if err != nil {
		panic(err)
	}

	//debug
	fmt.Println("messageCollection", messageCollection)
	renderTemplate(w, r, "queryallmessages.html", messageCollection)

}
