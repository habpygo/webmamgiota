package controllers

import (
	"fmt"
	"net/http"

	//"github.com/iotaledger/mamgoiota"

	"github.com/iotaledger/mamgoiota/connections"
)

//MAMBoardSetup is the data to be exposed in the web-app
//It's the same as Transaction type
// type MAMBoardSetup struct {
// 	Message   string
// 	Value     int64
// 	Timestamp time.Time
// 	Recipient string
// 	Number    int
// }

//AllMessagesHandler will collect all the mesages and puts it into the messageCollection slice
func AllMessagesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AllMessagesHandler is entered")
	//address := "RQP9IFNFGZGFKRVVKUPMYMPZMAICIGX9SVMBPNASEBWJZZAVDCMNOFLMRMFRSQVOQGUVGEETKYFCUPNDDWEKYHSALY"
	//address := "UOKSEHAQCBPTCYGLQHUFLGJLQVSGMF9EPITW9QFDVPPXXDINMTLCYYSYTSGSUHP9YBGYKDZBKSAGBVULZPOWXNDHPX"
	address := "TVWZVZZLWSMLXYTFQNVQSAGCQLRRCUXMUDDQWJILNQGOIFKMA9PKBRKORIWOOF9WQLJWGVGTWUXPNNKNYSRBAWUWQC"
	provider := "http://node02.iotatoken.nl:14265"
	c, err := connections.NewConnection(provider, "")
	if err != nil {
		panic(err)
	}

	messageCollection, err := connections.ReadTransactions(address, c)
	if err != nil {
		panic(err)
	}

	tempValue := MAMBoardSetup{}
	collectedMessages := []MAMBoardSetup{}

	for i, m := range messageCollection[:] {
		tempValue.Number = i
		tempValue.Message = m.Message
		tempValue.Value = m.Value
		tempValue.Timestamp = m.Timestamp

		collectedMessages = append(collectedMessages, tempValue)
	}
	tempValue.Recipient = address
	collectedMessages = append(collectedMessages, tempValue)
	//debug
	for i, m := range messageCollection[:] {
		fmt.Printf("%d. %v. Value is %v. Timestamp is %v. and recipient is %v\n", i+1, m.Message, m.Value, m.Timestamp, m.Recipient)
	}
	renderTemplate(w, r, "queryallmessages.html", collectedMessages)

}
