package controllers

import (
	"fmt"
	"net/http"
	"time"

	//"github.com/iotaledger/mamgoiota"

	"github.com/iotaledger/giota"
	"github.com/iotaledger/webmamgiota/connections"
)

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

//AllMessagesHandler will collect all the mesages and puts it into the messageCollection slice
func AllMessagesForAddressHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AllMessagesHandler is entered")
	//address := "RQP9IFNFGZGFKRVVKUPMYMPZMAICIGX9SVMBPNASEBWJZZAVDCMNOFLMRMFRSQVOQGUVGEETKYFCUPNDDWEKYHSALY"
	//address := "UOKSEHAQCBPTCYGLQHUFLGJLQVSGMF9EPITW9QFDVPPXXDINMTLCYYSYTSGSUHP9YBGYKDZBKSAGBVULZPOWXNDHPX"
	address := "TVWZVZZLWSMLXYTFQNVQSAGCQLRRCUXMUDDQWJILNQGOIFKMA9PKBRKORIWOOF9WQLJWGVGTWUXPNNKNYSRBAWUWQC"
	seed := "THISISTHETESTSENTENCETOEXPERIMENTWITHIOTATANGLEFORPROGRAMMINGUSECASESASWELLASFUN9"
	trytesSeed, err := giota.ToTrytes(seed)
	if err != nil {
		panic(err)
	}
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
	renderTemplate(w, r, "queryaddressformessages.html", collectedMessages)

}
