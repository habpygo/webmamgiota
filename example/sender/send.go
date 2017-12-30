/*
MIT License
Copyright (c) 2017 Harry Boer, Jonah Polack

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
	"time"

	"github.com/iotaledger/mamgoiota"
)

//These nodes were working during the demo in Amsterdam 25 October, 2017. Currently only node01 is active
//var nodes = [4]string{"http://node01.iotameetup.nl:14265", "http://node02.iotameetup.nl:14265", "http://node03.iotameetup.nl:14265", "http://node04.iotatoken.nl:14265"}

//Below the address of the message board http://node01.iotameetup.nl:1337/ to where you can send the MAM
//Perhaps you should find another address
//var address = "XHBQNNJB9ESMBABXJVVRLXTKXTKOINIJCXOEHIMOJIGLOCPFXYCZGVTHK9RBQWECIXGOKLYFMOXRPYBPWVZG9B9LTZ"
var address = "RQP9IFNFGZGFKRVVKUPMYMPZMAICIGX9SVMBPNASEBWJZZAVDCMNOFLMRMFRSQVOQGUVGEETKYFCUPNDDWEKYHSALY"
var seed = "SIERTBRUINSISBEZIGOMEENRONDJESAMENMETWIMAMENTTEMAKENOMZODESUBSIDIERONDTEKRIJGENH9"

func main() {
	//"https://testnet140.tangle.works"
	//"http://node01.iotameetup.nl:14265" //This node worked @ December-1-2017; no guarantee that it will work in the future
	c, err := mamgoiota.NewConnection("http://node02.iotatoken.nl:14265", seed)
	if err != nil {
		panic(err)
	}

	msgTime := time.Now().UTC().String()
	message := "Testmessage By Harry during a rainy day on: " + msgTime

	id, err := mamgoiota.Send(address, 0, message, c)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Sent Transaction: %v\n", id)
}
