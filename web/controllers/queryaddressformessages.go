package controllers

import (
	"fmt"
	"net/http"
	"time"

	//"github.com/iotaledger/mamgoiota"

	"iota/webmamgiota/connections"

	"github.com/giota"
	"github.com/gorilla/websocket"
	//"github.com/iotaledger/giota"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//var upgrader = websocket.Upgrader{}

//MAMBoardSetup is the data to be transferred by bundles and exposed in the web-app
//It's different from Transfer type
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
	fmt.Println("AllMessagesHandler is entered")
	//address := "RQP9IFNFGZGFKRVVKUPMYMPZMAICIGX9SVMBPNASEBWJZZAVDCMNOFLMRMFRSQVOQGUVGEETKYFCUPNDDWEKYHSALY"
	//address := "UOKSEHAQCBPTCYGLQHUFLGJLQVSGMF9EPITW9QFDVPPXXDINMTLCYYSYTSGSUHP9YBGYKDZBKSAGBVULZPOWXNDHPX"
	//address := "RXMDCFKXYUW9NDUFNLAMEFEICJGUYEGVQDFSQERLPMSIRKEMXKAIARXJFXDHZBTAFCE9ZFWFNQKTIJLOXYFNJZIW9W" //Address Remko
	address := "TVWZVZZLWSMLXYTFQNVQSAGCQLRRCUXMUDDQWJILNQGOIFKMA9PKBRKORIWOOF9WQLJWGVGTWUXPNNKNYSRBAWUWQC"
	seed := "THISISTHETESTSENTENCETOEXPERIMENTWITHIOTATANGLEFORPROGRAMMINGUSECASESASWELLASFUN9"
	trytesSeed, err := giota.ToTrytes(seed)
	if err != nil {
		panic(err)
	}
	provider := "http://node02.iotatoken.nl:14265"
	//provider := "http://eugene.iota.community:14265"

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

	myapi := giota.NewAPI(provider, nil)

	for i, m := range messageCollection[:] {
		tempValue.Number = i
		tempValue.Message = m.Message
		tempValue.Value = m.Value
		tempValue.Timestamp = m.Timestamp

		collectedMessages = append(collectedMessages, tempValue)
	}
	tempValue.Recipient = address

	balances, _ := giota.GetInputs(myapi, trytesSeed, 0, 10, 0, 2)
	tempValue.Balances = balances.Total()

	collectedMessages = append(collectedMessages, tempValue)

	//debug
	for i, m := range messageCollection[:] {
		fmt.Printf("%d. %v. Value is %v. Timestamp is %v. and recipient is %v\n", i+1, m.Message, m.Value, m.Timestamp, m.Recipient)
	}

	// var n time.Duration
	// n = 10
	// var lastTransactions []mamgoiota.Transaction

	//TODO: doesn't work, should be achieved by using websockets
	// defer doEvery(n*time.Second, func(t time.Time) {
	// 	fmt.Println("Looking for new messages")

	// 	newTransactions, err := mamgoiota.ReadTransactions(address, c)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if len(lastTransactions) != 0 && len(lastTransactions) < len(newTransactions) {
	// 		diff := len(newTransactions) - len(lastTransactions)

	// 		fmt.Printf("Got %d new messages\n", diff)

	// 		for i, m := range newTransactions[:diff] {
	// 			tempValue.Number = i
	// 			tempValue.Message = m.Message
	// 			tempValue.Value = m.Value
	// 			tempValue.Timestamp = m.Timestamp

	// 			collectedMessages = append(collectedMessages, tempValue)
	// 			fmt.Printf("%d. %v\n", i+1, m.Message)
	// 		}

	// 		lastTransactions = newTransactions
	// 		renderTemplate(w, r, "queryaddressformessages.html", collectedMessages)
	// 		return
	// 	}

	//lastTransactions = newTransactions
	//fmt.Println("No new messages")
	//}

	renderTemplate(w, r, "queryaddressformessages.html", collectedMessages)

}
