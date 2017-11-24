# mamgiota

Small project to implement Masked Authenticated Messaging on the IOTA tangle with Golang.

This project is still under construction (see TODO) so it's not functioning yet.

# Install
It is assumed that you have Golang installed. You also need to install the Go library API for IOTA which you can download at:
```
$ go get -u github.com/iotaledger/giota
```
After which you can download `send-message` from
```
$ go get -u github.com/habpygo/mamgoiota
```

# TODO
Currently there is an error message `Invalid transaction hash` when sending the messasge to the tangle; however, this could be due to a bug in the `iota.lib.go` library.
