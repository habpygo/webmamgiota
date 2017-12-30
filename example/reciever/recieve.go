package main

import (
	"fmt"
	"time"

	"github.com/iotaledger/mamgoiota"
)

func main() {
	address := "RQP9IFNFGZGFKRVVKUPMYMPZMAICIGX9SVMBPNASEBWJZZAVDCMNOFLMRMFRSQVOQGUVGEETKYFCUPNDDWEKYHSALY"
	provider := "http://node02.iotatoken.nl:14265"
	c, err := mamgoiota.NewConnection(provider, "")
	if err != nil {
		panic(err)
	}

	var lastTransactions []mamgoiota.Transaction
	doEvery(5*time.Second, func(t time.Time) {
		fmt.Println("Looking for new messages")

		newTransactions, err := mamgoiota.ReadTransactions(address, c)
		if err != nil {
			panic(err)
		}

		if len(lastTransactions) != 0 && len(lastTransactions) < len(newTransactions) {
			diff := len(newTransactions) - len(lastTransactions)

			fmt.Printf("Got %d new messages\n", diff)

			for i, m := range newTransactions[:diff] {
				fmt.Printf("%d. %v\n", i+1, m.Message)
			}

			lastTransactions = newTransactions
			return
		}

		lastTransactions = newTransactions
		fmt.Println("No new messages")
	})
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}
