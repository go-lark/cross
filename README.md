# Cross

[![Go Module](https://badge.fury.io/go/github.com%2Fgo-lark%2Fcross.svg)](https://badge.fury.io/go/github.com%2Fgo-lark%2Fcross.svg)

Cross converting [go-lark/lark](https://github.com/go-lark/lark) messages to [chyroc/lark](https://github.com/chyroc/lark) format.

## Getting Started

```sh
go get github.com/go-lark/cross
```

```go
msg := golark.NewMsgBuffer(golark.MsgInteractive)
om := msg.BindEmail(testUserEmail).Card(card.String()).Build()
crossMsg, err := BuildMessage(golark.UIDEmail, om)
if err != nil {
    // error handling ...
}
cli := clark.New(clark.WithAppCredential(testAppID, testAppSecret))
resp, _, err := cli.Message.SendRawMessage(t.Context(), crossMsg)
```

## License

Copyright (c) David Zhang, 2025. Licensed under MIT License.
