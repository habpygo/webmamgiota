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

package controllers

import (
	"fmt"
	"iota/webmamgiota/connections"
	"iota/webmamgiota/web/metadata"
	"net/http"
	"time"
)

//ReceiveHandler checks for new messages every n seconds
func ReceiveHandler(w http.ResponseWriter, r *http.Request) {
	//address := "TVWZVZZLWSMLXYTFQNVQSAGCQLRRCUXMUDDQWJILNQGOIFKMA9PKBRKORIWOOF9WQLJWGVGTWUXPNNKNYSRBAWUWQC"
	//provider := "http://node02.iotatoken.nl:14265"

	c, err := connections.NewConnection(metadata.Provider, "")
	if err != nil {
		panic(err)
	}

	var n time.Duration
	n = 15

	var lastTransactions []connections.Transaction

	doEvery(n*time.Second, func(t time.Time) {
		fmt.Println("Looking for new messages")

		newTransactions, err := connections.ReadTransactions(metadata.Address, c)
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
