package controllers

import (
	"fmt"
	"net/http"

	"iota/webmamgiota/connections"
)

//CheckForTxIdHandler returns the message, value(when applicable) and address of recipient for a given transaction
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
	seed := "THISISTHETESTSENTENCETOEXPERIMENTWITHIOTATANGLEFORPROGRAMMINGUSECASESASWELLASFUN9"

	provider := "http://node02.iotatoken.nl:14265"
	//provider := "http://nodes.spamnet.iota.org"

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
