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

	"github.com/iotaledger/giota"
	"github.com/iotaledger/send-message/stringutils"
)

//These nodes were working during the demo in Amsterdam 25 October, 2017. Currently only node01 is active
//var nodes = [4]string{"http://node01.iotameetup.nl:14265", "http://node02.iotameetup.nl:14265", "http://node03.iotameetup.nl:14265", "http://node04.iotatoken.nl:14265"}

//Below the address of http://node01.iotameetup.nl:1337/ to where you could send the MAM
//fmt.Println() statements are for debugging purposes
var address = "XHBQNNJB9ESMBABXJVVRLXTKXTKOINIJCXOEHIMOJIGLOCPFXYCZGVTHK9RBQWECIXGOKLYFMOXRPYBPWVZG9B9LTZ"
var seed = ""

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // We need this to avoid ending up with the same sequence. Comment out for debugging.

	trits := "9ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	message := "IOTA World from Golang, do you read?"

	// Select a random seed
	for i := 0; i < 81; i++ {
		x := rand.Float64() * 27
		y := int(x)
		seed += string(trits[y])
	}
	seedTrytes, _ := giota.ToTrytes(seed)
	fmt.Println("The seed is: ", seedTrytes)

	// Select a random host when all the hosts are actively listening
	// x := rand.Float64() * 4
	// y := int(x)
	//host := nodes[y]
	provider := "http://node01.iotameetup.nl:14265"

	api := giota.NewAPI(provider, nil)
	fmt.Println("api is: ", api)

	// Transform the message to tryte values suitable to send MAM's
	msg := stringutils.ToMAMTrytes(message)

	// Convert the string address to giota address
	address, err := giota.ToAddress(address)
	if err != nil {
		fmt.Println("error in address: ", address)
	}
	//fmt.Println("The address is: ", address)

	// trs := []giota.Transfer{
	// 	giota.Transfer{
	// 		Address: address,
	// 		Value:   0,
	// 		Message: msg,
	// 		Tag:     "",
	// 	},
	// }

	trx := []giota.Transaction{
		giota.Transaction{
			Address:                  address,
			SignatureMessageFragment: msg,
			Value:     0,
			Timestamp: time.Now(),
		},
	}

	fmt.Println("trs is: ", trx)

	//Uncomment two lines below if you want to check the message
	//stringMessage := stringutils.FromMAMTrytes(msg)
	//fmt.Println("stringMessage is: ", stringMessage)

	// bdl, trsErr := giota.Send(api, seedTrytes, 9, trs, 14, nil)
	// if trsErr != nil {
	// 	fmt.Println("Error while sending Trytes: ", trsErr)
	// }
	// fmt.Println("The bundle is: ", bdl)

	trxErr := giota.SendTrytes(api, 3, trx, 14, nil)
	if trxErr != nil {
		fmt.Println("Error sending Trytes: ", trxErr)
	}

}
