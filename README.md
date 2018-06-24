# webmamgiota

<span style="color:red">WARNING: DO NOT USE THE SEED OF YOUR WALLET!</span> 

NOTE: Tab `Check for IoT data` not working yet. Work in progress.
NOTE: You must fill in a number (0 is fine) in the Value IOTAs field, otherwise the app will panic.

This web-app is still under construction and no safety tests have been conducted yet. If you want to, only use small IOTA values, and do not use our wallet seed. Otherwise and even better, use https://nodes.devnet.thetangle.org:443 (Previous testnet) for testing purposes,with a MWM of 9 to 14 as is suggested here https://blog.iota.org/first-of-the-new-testnets-live-f8f41b99e9a3 I've set mwm to 9 which is fast. You can change these settings in the `metadata.go` file.

`webmamgiota` is a small project to implement Masked Authenticated Messaging on the IOTA tangle with Golang via a web UI.

This is work in progress and still under construction (see TODO) with the aim to get IoT sensors and devices to send MAMs. No Merkle Tree authentication is implemented yet, which 
is a priority for now.

## Install

It is assumed that you have Golang installed. You also need to install the Go library API for IOTA which you can download at:

```go
go get -u github.com/iotaledger/giota
```

After that you can download the webmamgiota package.

```go
go get -u github.com/habpygo/webmamgiota
```
It is assumed that in will be installed in your `$GOPATH/src` otherwise you will have to vendor it yourself.
## Sending MAMs to the IOTA tangle with Go

### API

#### Run the web-app

In the root directory enter `go run main.go`. Your browser automatically opens up on the Send message page. 

#### Connection
A new connection is automatically created when the app is started and pointed to port 3000 of your local webbrowser.

```go
func main() {
	...
	open("http://localhost:3000/")
	web.Serve(msgwebpage)
}
```

If you don't have a seed yet, follow the description here: https://iota.readme.io/docs/securely-generating-a-seed

Please keep in mind that you may NEVER lose this seed nor give it to anybody else, because the seed is the connection to your funds!

#### Send a MAM to the IOTA tangle

After the webpage has opened, you can write a message in the input field labeled "Text message". Fill in 0 for Value (not working yet) and press the `Send message` button. This might take a while depending on the traffic.
After sending, you find your transaction by clicking on the `Query address for all messages`.

You can also peruse it here https://thetangle.org giving the TransactionId.

<!-- If you want to transfer value aswell (here 100 IOTA) call the send method like this: ```Send("the receiving address", 100, "your stringified message", c)```. -->



### TODOs

- [ ] Make use of Merkle Tree (binary) to make proper masked authenticated messages, a priority for now
- [ ] Still receiving Travis errors
- [ ] GoDoc
- [ ] Read sensor data, e.g. RuuVi tag
- [ ] More Read options
- [ ] Send Value
- [X] Read by TransactionId





