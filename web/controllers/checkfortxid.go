package controllers

import (
	"fmt"
	"net/http"

	"github.com/iotaledger/webmamgiota/connections"
)

//TransactionVars contain the most important variables in the struct
// type TransactionVars struct {
// 	Timestamp time.Time
// 	Value     int64
// 	Message   string
// 	Recipient string
// }

//Check for Transaction IDs TxIds
//TxIds are for practice
/*
KEWIJJVTLQNSXJXV9BTQKYMXQRCDFTOXJVBVEJKGCELCPGAN9YOTZ9EESFGFKDG9R9XORHFCKIUE99999
UHGTSYS9DGKBERKLZELMAQVVOCH9PBXVV9KETWEXNKJKZO9CCBQWLASPKGUBMWIORPRYHYRYNUQM99999
*/

func CheckForTxIdHandler(w http.ResponseWriter, r *http.Request) {

	TxResult := &struct {
		Value     int64
		Message   string
		Recipient string
	}{
		Value:     0,
		Message:   "Just the same old story",
		Recipient: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
	}
	//address := "RQP9IFNFGZGFKRVVKUPMYMPZMAICIGX9SVMBPNASEBWJZZAVDCMNOFLMRMFRSQVOQGUVGEETKYFCUPNDDWEKYHSALY"
	//address := "UOKSEHAQCBPTCYGLQHUFLGJLQVSGMF9EPITW9QFDVPPXXDINMTLCYYSYTSGSUHP9YBGYKDZBKSAGBVULZPOWXNDHPX"
	//address := "TVWZVZZLWSMLXYTFQNVQSAGCQLRRCUXMUDDQWJILNQGOIFKMA9PKBRKORIWOOF9WQLJWGVGTWUXPNNKNYSRBAWUWQC"
	seed := "THISISTHETESTSENTENCETOEXPERIMENTWITHIOTATANGLEFORPROGRAMMINGUSECASESASWELLASFUN9"

	provider := "http://node02.iotatoken.nl:14265"
	c, err := connections.NewConnection(provider, seed)
	if err != nil {
		panic(err)
	}

	if r.FormValue("submitted") == "true" {
		RequiredTxID := r.FormValue("txid")

		Result, err := connections.ReadTransaction(RequiredTxID, c)
		if err != nil {
			fmt.Println("Error while reading TxResult")
			panic(err)
		}

		TxResult.Message = Result.Message
		TxResult.Value = Result.Value
		TxResult.Recipient = Result.Recipient

		//renderTemplate(w, r, "checkfortxid.html", TxResult)

	}

	renderTemplate(w, r, "checkfortxid.html", TxResult)

}
