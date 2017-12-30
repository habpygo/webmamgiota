/*
MIT License
Copyright (c) 2017 Harry Boer

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/iotaledger/mamgoiota/mamutils" // in $GOPATH/src

	"github.com/iotaledger/giota"
)

//These nodes were working during the demo in Amsterdam 25 October, 2017. Currently only node01 is active
//var nodes = [4]string{"http://node01.iotameetup.nl:14265", "http://node02.iotameetup.nl:14265", "http://node03.iotameetup.nl:14265", "http://node04.iotatoken.nl:14265"}

//Below the address of the message board http://node01.iotameetup.nl:1337/ to where you can send the MAM
//Perhaps you should find another address
//var address = "XHBQNNJB9ESMBABXJVVRLXTKXTKOINIJCXOEHIMOJIGLOCPFXYCZGVTHK9RBQWECIXGOKLYFMOXRPYBPWVZG9B9LTZ"
var address = "RQP9IFNFGZGFKRVVKUPMYMPZMAICIGX9SVMBPNASEBWJZZAVDCMNOFLMRMFRSQVOQGUVGEETKYFCUPNDDWEKYHSALY"
var seed = "SIERTBRUINSISBEZIGOMEENRONDJESAMENMETWIMAMENTTEMAKENOMZODESUBSIDIERONDTEKRIJGENH9"

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // We need this to avoid ending up with the same sequence. Comment out for debugging.

	trits := "9ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Genereate a random number to distinguish the messages
	//number := rand.Float64() * 100
	msgTime := time.Now().UTC().String()
	message := "Testmessage By Harry during a rainy day on: "
	//message2 := strconv.FormatFloat(number, 'f', 0, 64)
	//message += message2
	message += msgTime

	// Select a random seed
	for i := 0; i < 81; i++ {
		x := rand.Float64() * 27
		y := int(x)
		seed += string(trits[y])
	}
	seedTrytes, _ := giota.ToTrytes(seed)
	//provider := "https://testnet140.tangle.works"
	//provider := "http://node01.iotameetup.nl:14265" //This node worked @ December-1-2017; no guarantee that it will work in the future
	provider := "http://node02.iotatoken.nl:14265"

	api := giota.NewAPI(provider, nil)

	// Transform the message to tryte values suitable to send MAM's
	msg, err := mamutils.ToMAMTrytes(message)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Convert the string address to giota address
	address, err := giota.ToAddress(address)
	if err != nil {
		fmt.Println("error in address: ", address)
	}

	trs := []giota.Transfer{
		giota.Transfer{
			Address: address,
			Value:   0,
			Message: msg,
			Tag:     "",
		},
	}

	//Uncomment two lines below if you want to check the message on the CLI
	stringMessage, err := mamutils.FromMAMTrytes(msg)
	if err != nil {
		//TODO add a proper error
		fmt.Println("Error message")
	}
	fmt.Println("The stringMessage is: ", stringMessage)

	_, bestPow := giota.GetBestPoW()

	mamBundle, txErr := giota.Send(api, seedTrytes, 9, trs, 15, bestPow)
	if txErr != nil {
		//TODO add a proper error
		fmt.Println("From send.go: Error while sending Trytes: ", txErr)
	}
	//Doesn't give tx hash, but bundle hash so it seems
	bdHash := mamBundle.Hash()

	fmt.Println("The Bundle Hash is: ", bdHash)

}
