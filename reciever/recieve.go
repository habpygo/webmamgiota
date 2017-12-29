package main

import (
	"fmt"
	"sort"
	"time"

	"iota/mamgoiota/mamutils"

	"github.com/iotaledger/giota"
)

func main() {
	address := "RQP9IFNFGZGFKRVVKUPMYMPZMAICIGX9SVMBPNASEBWJZZAVDCMNOFLMRMFRSQVOQGUVGEETKYFCUPNDDWEKYHSALY"
	provider := "http://node02.iotatoken.nl:14265"

	var lastMessages []string
	doEvery(5*time.Second, func(t time.Time) {
		fmt.Println("Looking for new messages")
		newMessages := readMessages(provider, address)

		if len(lastMessages) != 0 && len(lastMessages) < len(newMessages) {
			diff := len(newMessages) - len(lastMessages)

			fmt.Printf("Got %d new messages\n", diff)

			for i, m := range newMessages[len(newMessages)-diff:] {
				fmt.Printf("%d. %v\n", i+1, m)
			}

			lastMessages = newMessages
			return
		}

		lastMessages = newMessages
		fmt.Println("No new messages")
	})
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func readMessages(provider, address string) []string {

	api := giota.NewAPI(provider, nil)

	iotaAdress, err := giota.ToAddress(address)
	if err != nil {
		panic("error in address: " + address)
	}

	req := giota.FindTransactionsRequest{
		Command:   "findTransactions",
		Addresses: []giota.Address{iotaAdress},
	}

	res, err := api.FindTransactions(&req)

	if err != nil {
		panic(err)
	}

	response := *res

	trytesResp, err := api.GetTrytes(response.Hashes)
	if err != nil {
		panic(err)
	}

	sort.Slice(trytesResp.Trytes, func(i, j int) bool {
		return trytesResp.Trytes[i].Timestamp.Unix() < trytesResp.Trytes[j].Timestamp.Unix()
	})
	messages := make([]string, len(trytesResp.Trytes))
	for i, t := range trytesResp.Trytes {
		//Have to cut off one 9-padding at the end so it can be parsed
		message, err := mamutils.FromMAMTrytes(t.SignatureMessageFragment[:len(t.SignatureMessageFragment)-1])
		if err != nil {
			panic(err)
		}
		messages[i] = message
	}

	return messages
}
