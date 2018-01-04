package controllers

import (
	"fmt"
	"net/http"

	"github.com/iotaledger/mamgoiota"
)

//MessageStruct is the data to be exposed in the web-app
//It's the same as Transaction type
// type MessageStruct struct {
// 	Message   string
// 	Value     int64
// 	Timestamp time.Time
// 	Recipient string
// }

//AllMessagesHandler will collect all the mesages and puts it into the messageCollection slice
func AllMessagesHandler(w http.ResponseWriter, r *http.Request) {
	address := "RQP9IFNFGZGFKRVVKUPMYMPZMAICIGX9SVMBPNASEBWJZZAVDCMNOFLMRMFRSQVOQGUVGEETKYFCUPNDDWEKYHSALY"
	provider := "http://node02.iotatoken.nl:14265"
	c, err := mamgoiota.NewConnection(provider, "")
	if err != nil {
		panic(err)
	}

	messageCollection := []mamgoiota.Transaction{}

	messageCollection, err = mamgoiota.ReadTransactions(address, c)
	if err != nil {
		panic(err)
	}

	//debug
	fmt.Println("messageCollection", messageCollection)
	renderTemplate(w, r, "queryallmessages.html", messageCollection)

}
