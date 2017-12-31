# mamgoiota

## Sending MAMs to the IOTA tangle with Go

### API

#### Create a new Connection
```go
import "github.com/iotaledger/mamgoiota"

func main(){
    c, err := mamgoiota.NewConnection("someNodeURL", "yourSeed")
    if c != nil && err == nil{
        fmt.Println("Connection is valid")
    }
}
```
If you don't have a nodeURL try out one from: http://iotasupport.com/lightwallet.shtml

If you don't have a seed yet, follow the description here: https://iota.readme.io/docs/securely-generating-a-seed

Please keep in mind that you may NEVER loose this seed nor give it to anybody else, because the seed is the connection to your funds!


#### Send a MAM to the IOTA tangle
```go
import "github.com/iotaledger/mamgoiota"

func main(){
    c, err := mamgoiota.NewConnection("someNodeURL", "yourSeed")
    if err != nil{
        panic(err)
    }
    id, err := Send("the receiving address", 0, "your stringified message", c)
    if err != nil{
        panic(err)
    }
    fmt.Printf("Send to the Tangle. TransactionId: %v\n", id)
}
```
After sending, you find your transaction here https://thetangle.org giving the TransactionId

If you want to transfer value aswell (here 100 IOTA) call the send method like this: ```Send("the receiving address", 100, "your stringified message", c)```.

#### Read data from the IOTA tangle
Reading all transaction received by a certain adress:
```go
import "github.com/iotaledger/mamgoiota"

func main(){
    c, err := NewConnection("someNodeURL", "")
    if err != nil{
        panic(err)
    }

    ts, err := ReadTransactions("Receiving Adress", c)
    if err != nil{
        panic(err)
    }
    for i, tr := range ts {
        t.Logf("%d. %v: %d IOTA, %v to %v\n", i+1, tr.Timestamp, tr.Value, tr.Message, tr.Recipient)
    }
}
```
The seed can be ommitted here, since reading does not require an account



Reading a special transaction by transactionID:
```go
import "github.com/iotaledger/mamgoiota"

func main(){
    c, err := NewConnection("someNodeURL", "")
    if err != nil{
        panic(err)
    }

    tx, err := ReadTransaction("Some transactionID", c)
    if err != nil{
        panic(err)
    }
    t.Logf("%v: %d IOTA, %v to %v\n", tx.Timestamp, tx.Value, tx.Message, tx.Recipient)
}
```


#### Examples
Check out our [example folder](/example) for a send and a receive example.

To run this edit the `sender/send.go` and `receiver/receive.go` file, set the correct provider and address and you are ready to run.

Start the receiver first: `$ go run receiver/receive.go`. He checks every 5 seconds if there are new messages, until cancelled.

Then start the sender: `$ go run sender/send.go`.

If you pick up the transaction hash from the Terminal output and paste it into the input field on the site https://thetangle.org you find your transaction.

If the Node is offline try another one, mentioned above.

### TODOs
- [ ] GoDoc
- [ ] Travis
- [ ] More Read options
- [X] Read by TransactionId





