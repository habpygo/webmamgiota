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
	"net/http"
	"strconv"
	"time"

	mamgoiota "github.com/giota/mamgoiota/connections"
)

//var address = "RQP9IFNFGZGFKRVVKUPMYMPZMAICIGX9SVMBPNASEBWJZZAVDCMNOFLMRMFRSQVOQGUVGEETKYFCUPNDDWEKYHSALY"
//var address = "UOKSEHAQCBPTCYGLQHUFLGJLQVSGMF9EPITW9QFDVPPXXDINMTLCYYSYTSGSUHP9YBGYKDZBKSAGBVULZPOWXNDHPX"
var address = "TVWZVZZLWSMLXYTFQNVQSAGCQLRRCUXMUDDQWJILNQGOIFKMA9PKBRKORIWOOF9WQLJWGVGTWUXPNNKNYSRBAWUWQC"

//var seed = "SIERTBRUINSISBEZIGOMEENRONDJESAMENMETWIMAMENTTEMAKENOMZODESUBSIDIERONDTEKRIJGENH9"
var seed = "THISISTHETESTSENTENCETOEXPERIMENTWITHIOTATANGLEFORPROGRAMMINGUSECASESASWELLASFUN9"

//MAMBoardSetup is the data to be transferred by bundles and exposed in the web-app
//It's different from Transfer type
type MAMBoardSetup struct {
	Message   string
	Value     int64
	Timestamp time.Time
	Recipient string
	Number    int
}

//SendHandler retrieves the message values from the webpage and sends it to the address given
func SendHandler(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		TransactionID string
		TimeStamp     string
		Success       bool
		Response      bool
	}{
		TransactionID: "",
		Success:       false,
		Response:      false,
	}
	//"https://testnet140.tangle.works"
	c, err := mamgoiota.NewConnection("http://node02.iotatoken.nl:14265", seed)
	//c, err := connections.NewConnection("http://eugene.iota.community:14265", seed)
	if err != nil {
		panic(err)
	}

	if r.FormValue("submitted") == "true" {
		newMamMessage := MAMBoardSetup{
			Message: r.FormValue("message"),
		}
		value, err := strconv.ParseInt(r.FormValue("value"), 10, 64)
		if err != nil {
			panic(err)
		}

		newMamMessage.Value = value

		/* WRITE YOUR MESSAGE HERE */
		//message := "Test text message from web-app to assert that all is good" //+ msgTime

		//we use mamgoiota here to test whether the mamgoiota folder in the library is still working
		//otherwise use connections.Send() like connections.NewConnection() above
		id, err := mamgoiota.Send(address, newMamMessage.Value, newMamMessage.Message, c)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Sent transaction: %v\n", id)

	}
	renderTemplate(w, r, "sendmessage.html", data)
}
