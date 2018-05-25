package controllers

import (
	"fmt"
	"iota/webmamgiota/connections"
	"iota/webmamgiota/web/metadata"
	"net/http"
	"time"

	"github.com/giota"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type MAMBoardSetup struct {
	Message   string
	Value     int64
	Timestamp time.Time
	Recipient string
	Number    int
	Balances  int64
}

//AllMessagesForAddressHandler will collect all the mesages and puts it into the messageCollection slice
func AllMessagesForAddressHandler(w http.ResponseWriter, r *http.Request) {
	//address := "TVWZVZZLWSMLXYTFQNVQSAGCQLRRCUXMUDDQWJILNQGOIFKMA9PKBRKORIWOOF9WQLJWGVGTWUXPNNKNYSRBAWUWQC"
	//seed := "THISISTHETESTSENTENCETOEXPERIMENTWITHIOTATANGLEFORPROGRAMMINGUSECASESASWELLASFUN9"
	trytesSeed, err := giota.ToTrytes(metadata.Seed)
	if err != nil {
		panic(err)
	}
	//provider := "http://node02.iotatoken.nl:14265"
	//provider := "http://nodes.spamnet.iota.org"

	c, err := connections.NewConnection(metadata.Provider, "")
	if err != nil {
		panic(err)
	}

	messageCollection, err := connections.ReadTransactions(metadata.Address, c)
	if err != nil {
		panic(err)
	}

	tempValue := MAMBoardSetup{}
	collectedMessages := []MAMBoardSetup{}

	myapi := giota.NewAPI(metadata.Provider, nil)

	for i, m := range messageCollection[:] {
		tempValue.Number = i
		tempValue.Message = m.Message
		tempValue.Value = m.Value
		tempValue.Timestamp = m.Timestamp

		collectedMessages = append(collectedMessages, tempValue)
	}
	tempValue.Recipient = metadata.Address

	balances, _ := giota.GetInputs(myapi, trytesSeed, 0, 10, 0, 2)
	tempValue.Balances = balances.Total()

	collectedMessages = append(collectedMessages, tempValue)

	//debug
	for i, m := range messageCollection[:] {
		fmt.Printf("%d. %v. Value is %v. Timestamp is %v. and recipient is %v\n", i+1, m.Message, m.Value, m.Timestamp, m.Recipient)
	}

	renderTemplate(w, r, "queryaddressformessages.html", collectedMessages)

}
